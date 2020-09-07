package gin4grpc_mode

import (
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/utils"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type service struct {
	ServerName string
	Method     string
	PathName   string
	FuncName   string
	SFuncName  string // 首字母小写的funcname
	ReqName    string
	ReplyName  string
	Params     []spec4pb.Parameter
}

func (s *service) getRouterPath() string {
	re := regexp.MustCompile("\\{(.+?)\\}")
	matched := re.FindAllStringSubmatch(s.PathName, -1)
	res := s.PathName
	for _, match := range matched {
		res = strings.Replace(res, match[0], ":"+match[1], 1)
	}
	return res
}

type sortServices []*service

func (s sortServices) Less(i, j int) bool {
	if s[i].ServerName != s[j].ServerName {
		return s[i].ServerName < s[j].ServerName
	}
	return s[i].FuncName < s[j].FuncName
}
func (s sortServices) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortServices) Len() int      { return len(s) }

func (p *Gin4GrpcMode) prepareServices() {
	wg := sync.WaitGroup{}
	for pth, pathItem := range p.swaggerPub.Swagger.Paths.Paths {
		for _, method := range common.Methods {
			operation := common.GetOptionsFromPathItemByMethod(pathItem, method)
			if operation == nil {
				continue
			}
			wg.Add(1)
			go func(pth, method string, operation *spec4pb.Operation) {
				defer wg.Done()
				p.prepareService(pth, method, operation)
			}(pth, method, operation)
		}
	}
	wg.Wait()
	sort.Sort(p.services)
}

func (p *Gin4GrpcMode) prepareService(pth, method string, operation *spec4pb.Operation) {
	serviceName := common.UriPathToName(pth)
	reqName := "EmptyMessage"
	replyName := "CommonReply"
	if p.analyseParams(operation.Parameters) {
		reqName = utils.ConcatenateStrings(method, serviceName, "Request")
	}
	if p.analyseReply(operation.Responses) {
		replyName = utils.ConcatenateStrings(method, serviceName, "Reply")
	}
	p.Lock()
	defer p.Unlock()
	p.services = append(p.services,
		&service{
			ServerName: serviceName,
			Method:     method,
			PathName:   pth,
			FuncName:   utils.ConcatenateStrings(method, serviceName),
			SFuncName:  utils.ConcatenateStrings(strings.ToLower(method), serviceName),
			ReqName:    reqName,
			ReplyName:  replyName,
			Params:     operation.Parameters,
		},
	)
}

func (p *Gin4GrpcMode) analyseParams(params []spec4pb.Parameter) bool {
	for _, param := range params {
		if param.In == "header" {
			continue
		}
		return true
	}
	return false
}

func (p *Gin4GrpcMode) analyseReply(response *spec4pb.Responses) bool {
	if response.Default != nil {
		if response.Default.Schema != nil {
			return true
		}
	}
	for i, r := range response.StatusCodeResponses {
		if i >= 200 && i < 300 {
			if r.Schema != nil {
				return true
			}
		}
	}
	return false
}

func (p *Gin4GrpcMode) genServices() {
	g := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/services", p.swaggerPub.Md5)
	g.SetFilename("a_base.go")
	p.GenDoNotChange(g)
	g.P("package services")
	g.P()
	g.P("import (")
	g.P("    \"context\"")
	g.P()
	p.GenImportPb(g)
	g.P(")")
	g.P()
	g.P("//server is used to implement ", p.swaggerPub.PackageName, ".", strings.Title(p.swaggerPub.PackageName), "Server.")
	g.P("type ", p.swaggerPub.PackageName, "Server struct {")
	g.P("  pb.Unimplemented", strings.Title(p.swaggerPub.PackageName), "Server")
	g.P("}")
	g.P("func NewServer() *", p.swaggerPub.PackageName, "Server{")
	g.P("    return &", p.swaggerPub.PackageName, "Server{}")
	g.P("}")
	g.P()

	gd := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/services", p.swaggerPub.Md5)
	gd.SetFilename("a_decorator.go")
	gd.P(decoratorFile)
	gd.GenFile()
	
	for _, s := range p.services {
		gs := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/services", p.swaggerPub.Md5)
		gs.SetFilename(s.FuncName + ".go")
		var mapper = map[string]interface{}{
			"packageName": p.swaggerPub.PackageName,
			"funcName":    s.FuncName,
			"reqName":     s.ReqName,
			"replyName":   s.ReplyName,
			"sFuncName":   s.SFuncName,
		}
		gs.P("package services")
		gs.P()
		gs.P("import (")
		gs.P("    \"context\"")
		gs.P("    \"fmt\"")
		gs.P()
		p.GenImportPb(gs)
		gs.P("    \"github.com/legenove/server-sdk-go/grpccore\"")
		gs.P(")")
		gs.P()
		
		g.P(FormatByKey(serverDecoratorHandle, mapper))
		gs.P(FormatByKey(serverFuncHandle, mapper))
		gs.GenFile()
	}
	g.GenFile()

}

var serverDecoratorHandle = `func (s *{{.packageName}}Server) {{.funcName}}(ctx context.Context, req *pb.{{.reqName}}) (*pb.{{.replyName}}, error) {
	res, err := {{.sFuncName}}Handler(ctx, req)
	var _r = res.(*pb.{{.replyName}})
	if _r == nil {
		_r = &pb.{{.replyName}}{}
	}
	return _r, err
}

var {{.sFuncName}}Handler = decoratorHandler("{{.sFuncName}}", func(ctx context.Context, req interface{}) (interface{}, error) {
	return {{.sFuncName}}(ctx, req.(*pb.{{.reqName}}))
}, {{.sFuncName}}Decors...)

`

var serverFuncHandle = `
var {{.sFuncName}}Decors = []grpccore.GrpcDecoratorFunc{}

func {{.sFuncName}}(ctx context.Context, req *pb.{{.reqName}}) (*pb.{{.replyName}}, error) {
	fmt.Println("in", req)
	return nil, nil
}
`

var decoratorFile = `package services

import (
	"context"
	"github.com/legenove/server-sdk-go/grpccore"
	"google.golang.org/grpc"
)

// 公共方法
func commonHandler(funcName string, handler grpc.UnaryHandler) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// before
		res, err := handler(ctx, req)
		// after
		return res, err
	}
}

func decoratorHandler(funcName string, handler grpc.UnaryHandler, decors ...grpccore.GrpcDecoratorFunc) grpc.UnaryHandler {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		handler = d(funcName, handler)
	}
	handler = grpccore.LoggerRecoveryHandler(funcName, handler)
	return commonHandler(funcName, handler)
}
`