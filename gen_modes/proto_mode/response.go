package proto_mode

import (
	"fmt"
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/utils"
)

func (p *ProtoMode) analyseReply(name string, method, part string, response *spec4pb.Responses) bool {
	var schema *spec4pb.Schema
	if response.Default != nil {
		if response.Default.Schema != nil {
			schema = response.Default.Schema
		}
	}
	for i, r := range response.StatusCodeResponses {
		if i >= 200 && i < 300 {
			if r.Schema != nil {
				schema = r.Schema
				break
			}
		}
	}
	location := fmt.Sprintf("%d:%s", common.OptLocationMap[part], name)
	name = method + name
	messageName := utils.ConcatenateStrings(name, part)
	g := &mode_pub.BufGenerator{}
	g.P("message ", name, part, " {")
	g.P("  uint32 http_code = 1;")
	g.P("  uint32 err_code = 2;")
	g.P("  string err_msg = 3;")
	g.P("  uint32 comp_id = 4;")
	if schema != nil {
		g.Pl("  ")
		GPSchema(g, schema, method, location, name+part+"Data", "", p)
		GPFieldEnd(g, "data", 15, schema.Description)
	}
	g.P("}")
	p.Lock()
	defer p.Unlock()
	p.messageGenOpt = append(p.messageGenOpt,
		&BufGenOpt{
			location,
			common.OptMethodMap[method],
			messageName,
			g},
	)
	return true
}
