package proto_mode

import (
	"fmt"
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/utils"
	"strings"
)

//message HelloRequest {
//string name = 1;
//}
func (p *ProtoMode) analyseParams(name string, method, part string, params []spec4pb.Parameter) bool {
	var maxNum int32 = 0
	for _, param := range params {
		if param.In == "header" {
			continue
		}
		if param.FieldNumber > maxNum {
			maxNum = param.FieldNumber
		}
	}
	location := fmt.Sprintf("%d:%s", common.OptLocationMap[part], utils.ConcatenateStrings(name, part))
	messageName := utils.ConcatenateStrings(method, name, part)
	g := &mode_pub.BufGenerator{}
	g.P("message ", messageName, " {")
	for _, param := range params {
		if param.In == "header" {
			continue
		}
		if param.FieldNumber == 0 {
			panic(fmt.Sprintf("fildNumber not define: loaction: paths, routerName: %s, method: %s, paramsName: %s",
				name, method, param.Name))
		}
		GPParam(g, param, method, location, messageName, p)
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

func GPParam(g mode_pub.BufGenInterface, param spec4pb.Parameter, method, locations, ppath string, p *ProtoMode) {
	g.Pl("  ")
	_type := common.GetPBType(param)
	switch _type {
	case "array":
		g.Pl("repeated ")
		GPItem(g, param.Items, ppath)
	case "object":
		if param.Schema != nil {
			GPSchema(g, param.Schema, method, locations, ppath+strings.Title(param.Name), "", p)
		} else {
			// 否则是any
			g.Pl(p.getAnyProto())
		}
	default:
		g.Pl(_type)
	}
	GPFieldEnd(g, param.Name, param.FieldNumber, param.Description)
}

func GPFieldEnd(g mode_pub.BufGenInterface, name string, fieldNumber int32, description string) {
	g.Pl(" ", name, " = ", fieldNumber, ";")
	if description != "" {
		g.Pl("  // ", description)
	}
	g.P()
}

func FormatRefUrl(s string) string {
	ss := strings.Split(s, "/")
	res := ""
	for i := 1; i < len(ss); i++ {
		res += strings.Title(ss[i])
	}
	return res
}

func GPItem(g mode_pub.BufGenInterface, items *spec4pb.Items, ppath string) {
	_type := common.GetPBType(items)
	switch _type {
	case "array":
		g.Pl("repeated ")
		GPItem(g, items.Items, ppath)
	case "object":
		panic("+++-- GPItem, default error" + ppath)
	default:
		g.Pl(_type)
	}
}
