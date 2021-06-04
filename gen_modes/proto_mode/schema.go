package proto_mode

import (
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
)

func GPSchema(g mode_pub.BufGenInterface, schema *spec4pb.Schema, method, locations ,ppath string, inMap string, p *ProtoMode) {
	var _inMap = ""
	// map
	if schema.AdditionalProperties != nil && schema.AdditionalProperties.Schema != nil {
		if inMap != "" {
			// 在map里面则不能为再有key-value
			panic("to much map in one struct " + ppath)
		}
		_inMap = "value"
		g.Pl("map<string,")
		GPSchema(g, schema.AdditionalProperties.Schema, method, locations ,ppath, "key", p)
		g.Pl(">")
		return
	}
	// allof anyof oneof 禁止使用
	_type := common.GetPBType(schema)
	switch _type {
	case "array":
		g.Pl("repeated ")
		if schema.Items.Schema != nil {
			GPSchema(g, schema.Items.Schema, method, locations ,ppath+"Arr", _inMap, p)
		} else {
			panic("schema: array: new items " + ppath)
			//GPNewArraySchema()
		}
		//GPItem(g, name, fieldNumber, schema.Items, ppath+strings.Title(name))
	case "object":
		if schema.Ref.GetURL() == nil {
			if len(schema.Properties) != 0 {
				g.Pl(ppath)
				GPProperties(p, schema, method, locations ,ppath)
			} else {
				g.Pl(p.getAnyProto())
			}
		} else {
			g.Pl(common.FormatRefUrl(schema.Ref.GetURL().String()))
		}
	default:
		g.Pl(_type)
	}
}

func getIntegerType(format string, minimum *float64) string {
	preFix := "s"
	if minimum != nil {
		if *minimum >= 256 {
			preFix = "f"
		} else if *minimum >= 0 {
			preFix = "u"
		}
	}
	switch format {
	case "int32":
		return preFix + "int32"
	case "int64":
		return preFix + "int64"
	default:
		return preFix + "int32"
	}
}

func getNumberType(format string) string {
	switch format {
	case "float":
		return "float"
	case "double":
		return "double"
	default:
		return "double"
	}
}

func GPInt(g mode_pub.BufGenInterface, format string, minimum *float64) {
	g.Pl(getIntegerType(format, minimum))
}

func GPNumber(g mode_pub.BufGenInterface, format string) {
	g.Pl(getNumberType(format))
}

func GPString(g mode_pub.BufGenInterface) {
	g.Pl("string")
}

func GPBoolean(g mode_pub.BufGenInterface) {
	g.Pl("bool")
}

// 文件类型使用字节形式传输
func GPFile(g mode_pub.BufGenInterface) {
	g.Pl("bytes")
}
