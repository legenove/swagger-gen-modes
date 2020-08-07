package proto_mode

import (
	"fmt"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/spec4pb"
	"github.com/legenove/utils"
	"strings"
)

//message HelloRequest {
//string name = 1;
//}
func (p *ProtoMode) analyseParams(name string, method, part string, params []spec4pb.Parameter) bool {
	res := false
	var maxNum int32 = 0
	for _, param := range params {
		if param.In == "header" {
			continue
		}
		res = true
		if param.FieldNumber > maxNum {
			maxNum = param.FieldNumber
		}
	}
	if !res {
		return false
	}
	location := fmt.Sprintf("%d:%s", OptLocationMap[part], utils.ConcatenateStrings(name, part))
	messageName := utils.ConcatenateStrings(method, name, part)
	g := &swagger_gen.BufGenerator{}
	g.P("message ", messageName, " {")
	for i, param := range params {
		if param.In == "header" {
			continue
		}
		if param.FieldNumber == 0 {
			reoutSwagger = true
			maxNum = maxNum + 1
			param.FieldNumber = maxNum
			// 覆盖原有的值
			params[i] = param
		}
		GPParam(g, param, method, location, messageName, p)
	}
	g.P("}")
	p.Lock()
	defer p.Unlock()
	p.messageGenOpt = append(p.messageGenOpt,
		&BufGenOpt{
			location,
			OptMethodMap[method],
			messageName,
			g},
	)
	return true
}

func GPParam(g swagger_gen.BufGenInterface, param spec4pb.Parameter, method, locations, ppath string, p *ProtoMode) {
	g.Pl("  ")
	switch param.Type {
	case "string":
		GPString(g)
	case "integer":
		GPInt(g, param.Format, param.Minimum)
	case "number":
		GPNumber(g, param.Format)
	case "boolean":
		GPBoolean(g)
	case "array":
		g.Pl("repeated ")
		GPItem(g, param.Items, ppath)
	case "file":
		GPFile(g)
	default:
		if param.Schema != nil {
			GPSchema(g, param.Schema, method, locations, ppath+strings.Title(param.Name), "", p)
		} else {
			// 否则是any
			g.Pl(p.getAnyProto())
		}
	}
	GPFieldEnd(g, param.Name, param.FieldNumber, param.Description)
}

func GPFieldEnd(g swagger_gen.BufGenInterface, name string, fieldNumber int32, description string) {
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

func GPItem(g swagger_gen.BufGenInterface, items *spec4pb.Items, ppath string) {
	switch items.Type {
	case "string":
		GPString(g)
	case "integer":
		GPInt(g, items.Format, items.Minimum)
	case "number":
		GPNumber(g, items.Format)
	case "boolean":
		GPBoolean(g)
	case "array":
		g.Pl("repeated ")
		GPItem(g, items.Items, ppath)
	default:
		panic("+++-- GPItem, default error" + ppath)
	}
}
