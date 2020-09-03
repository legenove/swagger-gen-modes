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
	g.SetFilename("base.go")
	g.P("package services")
	g.P()
	g.P("import (")
	p.GenImportPb(g)
	g.P(")")
	g.P()
	g.P("//server is used to implement ", p.swaggerPub.PackageName, ".", strings.Title(p.swaggerPub.PackageName), "Server.")
	g.P("type ", p.swaggerPub.PackageName, "Server struct {")
	g.P("  pb.Unimplemented", strings.Title(p.swaggerPub.PackageName), "Server")
	g.P("}")
	g.P("func NewServer() *",p.swaggerPub.PackageName,"Server{")
	g.P("    return &",p.swaggerPub.PackageName,"Server{}")
	g.P("}")
	g.GenFile()
	for _, s := range p.services {
		g = mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/services", p.swaggerPub.Md5)
		g.SetFilename(s.FuncName + ".go")
		g.P("package services")
		g.P()
		g.P("import (")
		g.P("    \"context\"")
		p.GenImportPb(g)
		g.P("    \"fmt\"")
		g.P(")")
		g.P()
		//func (*V1Server) DeleteTestPetId(ctx context.Context, req *pb.DeleteTestPetIdRequest) (*pb.CommonReply, error) {
		//
		//	return nil, nil
		//}

		g.P("func (*", p.swaggerPub.PackageName, "Server) ", s.FuncName,
			"(ctx context.Context, req *pb.", s.ReqName, ") (*pb.", s.ReplyName, ", error) {")
		g.P("    fmt.Println(\"in\", req)")
		g.P("    return nil, nil")
		g.P("}")
		g.GenFile()
	}

}
