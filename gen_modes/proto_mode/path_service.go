package proto_mode

import (
	"fmt"
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/utils"
	"strings"
	"sync"
)

func (p *ProtoMode) prepareServices() {
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
}

func (p *ProtoMode) prepareService(pth, method string, operation *spec4pb.Operation) {
	serviceName := common.UriPathToName(pth)
	g := &mode_pub.BufGenerator{}
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
			fmt.Sprintf("%d:%s", common.OptLocationMap["service"], serviceName),
			common.OptMethodMap[method],
			method + serviceName,
			g},
	)
}

func (p *ProtoMode) genServices(g mode_pub.BufGenInterface) {
	g.P("// Services")
	g.P(utils.ConcatenateStrings("service ", strings.Title(p.swaggerPub.PackageName), " {"))
	p.serviceGenOpt.MergeG(g)
	g.P("}")
	g.P()
}
