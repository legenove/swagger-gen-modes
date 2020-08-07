package proto_mode

import (
	"fmt"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/spec4pb"
	"github.com/legenove/utils"
	"strings"
	"sync"
)

var Methods = []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch"}

func getOptionsFromPathItemByMethod(pathItem spec4pb.PathItem, method string) *spec4pb.Operation {
	switch strings.ToLower(method) {
	case "get":
		return pathItem.Get
	case "post":
		return pathItem.Post
	case "put":
		return pathItem.Put
	case "delete":
		return pathItem.Delete
	case "options":
		return pathItem.Options
	case "head":
		return pathItem.Head
	case "patch":
		return pathItem.Patch
	default:
		return nil
	}
}

func pathToName(pth string) string {
	pth = strings.ReplaceAll(pth, "_", "/")
	pth = strings.ReplaceAll(pth, "{", "/")
	pth = strings.ReplaceAll(pth, "}", "/")
	ps := strings.Split(pth, "/")
	for i := range ps {
		ps[i] = strings.Title(ps[i])
	}
	return utils.ConcatenateStrings(ps...)
}

func (p *ProtoMode) prepareServices() {
	wg := sync.WaitGroup{}
	for pth, pathItem := range p.swaggerPub.Swagger.Paths.Paths {
		for _, method := range Methods {
			operation := getOptionsFromPathItemByMethod(pathItem, method)
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
}

func (p *ProtoMode) prepareService(pth, method string, operation *spec4pb.Operation) {
	serviceName := pathToName(pth)
	g := &swagger_gen.BufGenerator{}
	g.P("  // ", pth, " | ", method)
	g.Pl("  // operationId: ", operation.ID)
	if operation.Summary != "" {
		g.Pl(" | Summary : ", operation.Summary)
	}
	if operation.Description != "" {
		g.Pl(" | Description : ", operation.Description)
	}
	g.P()
	g.Pl("  rpc ", method, serviceName, " (")
	if !p.analyseParams(serviceName, method, "Request", operation.Parameters) {
		p.genEmpty(g)
	} else {
		g.Pl(method, serviceName, "Request")
	}
	g.Pl(") returns (")
	if !p.analyseReply(serviceName, method, "Reply", operation.Responses) {
		p.genReply(g)
	} else {
		g.Pl(method, serviceName, "Reply")
	}
	g.Pl(") {}")
	p.Lock()
	defer p.Unlock()
	p.serviceGenOpt = append(p.serviceGenOpt,
		&BufGenOpt{
			fmt.Sprintf("%d:%s", OptLocationMap["service"], serviceName),
			OptMethodMap[method],
			method + serviceName,
		g},
	)
}

func (p *ProtoMode) genServices(g swagger_gen.BufGenInterface) {
	g.P("// Services")
	g.P(utils.ConcatenateStrings("service ", strings.Title(p.swaggerPub.PackageName), " {"))
	p.serviceGenOpt.MergeG(g)
	g.P("}")
	g.P()
}
