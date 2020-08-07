package proto_mode

import (
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/spec4pb"
)

func GPSchema(g swagger_gen.BufGenInterface, schema *spec4pb.Schema, method, locations ,ppath string, inMap string, p *ProtoMode) {
	var isMap = false
	var _inMap = ""
	// map
	if schema.AdditionalProperties != nil && schema.AdditionalProperties.Schema != nil {
		if inMap != "" {
			// 在map里面则不能为再有key-value
			panic("to much map in one struct " + ppath)
		}
		isMap = true
		_inMap = "value"
		g.Pl("map<")
		GPSchema(g, schema.AdditionalProperties.Schema, method, locations ,ppath, "key", p)
		g.Pl(",")
	}
	// allof anyof oneof 禁止使用
	var _type string
	if len(schema.Type) > 0 {
		_type = schema.Type[0]
	}
	switch _type {
	case "string":
		GPString(g)
	case "integer":
		GPInt(g, schema.Format, schema.Minimum)
	case "number":
		GPNumber(g, schema.Format)
	case "boolean":
		GPBoolean(g)
	case "array":
		if inMap == "key" {
			panic("map key cant be array" + ppath)
		}
		g.Pl("repeated ")
		if schema.Items.Schema != nil {
			GPSchema(g, schema.Items.Schema, method, locations ,ppath+"Arr", _inMap, p)
		} else {
			panic("schema: array: new items " + ppath)
			//GPNewArraySchema()
		}
		//GPItem(g, name, fieldNumber, schema.Items, ppath+strings.Title(name))
	default:
		if inMap == "key" {
			panic("map key cant be struct" + ppath)
		}
		if schema.Ref.GetURL() == nil {
			// todo 有可能是 pro
			if len(schema.Properties) != 0 {
				g.Pl(ppath)
				GPProperties(p, schema, method, locations ,ppath)
			} else {
				g.Pl(p.getAnyProto())
			}
		} else {
			g.Pl(FormatRefUrl(schema.Ref.GetURL().String()))
		}
	}
	if isMap {
		g.Pl(">")
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

func GPInt(g swagger_gen.BufGenInterface, format string, minimum *float64) {
	g.Pl(getIntegerType(format, minimum))
}

func GPNumber(g swagger_gen.BufGenInterface, format string) {
	g.Pl(getNumberType(format))
}

func GPString(g swagger_gen.BufGenInterface) {
	g.Pl("string")
}

func GPBoolean(g swagger_gen.BufGenInterface) {
	g.Pl("bool")
}

// 文件类型使用字节形式传输
func GPFile(g swagger_gen.BufGenInterface) {
	g.Pl("bytes")
}
