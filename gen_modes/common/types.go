package common

import (
    "github.com/legenove/spec4pb"
)

var TypeGoMap = map[string]string{
    "string": "string",
    "fint32": "int32",
    "sint32": "int32",
    "uint32": "uint32",
    "fint64": "int64",
    "sint64": "int64",
    "uint64": "uint64",
    "float":  "float32",
    "double": "float64",
    "bool":   "bool",
    "bytes":  "[]byte",
    "object": "struct",
    "array":  "array",
}

func GetGoType(o interface{}) string {
    var t string
    t = GetPBType(o)
    return TypeGoMap[t]
}

func GetPBType(o interface{}) string {
    s, ok := o.(*spec4pb.Schema)
    if ok {
        return GetPBTypeBySchema(s)
    }
    p, ok := o.(spec4pb.Parameter)
    if ok {
        return GetPBTypeByParam(p)
    }
    i, ok := o.(*spec4pb.Items)
    if ok {
        return GetPBTypeByItem(i)
    }
    panic("obj error type")
}

func GetPBTypeBySchema(schema *spec4pb.Schema) string {
    var _type string
    if len(schema.Type) > 0 {
        _type = schema.Type[0]
    }
    return getPbType(_type, schema.Format, schema.Minimum)
}

func GetPBTypeByParam(param spec4pb.Parameter) string {
    return getPbType(param.Type, param.Format, param.Minimum)
}

func GetPBTypeByItem(items *spec4pb.Items) string {
    return getPbType(items.Type, items.Format, items.Minimum)
}

func getPbType(_type string, format string, minimum *float64) string {
    switch _type {
    case "string":
        return getPbStringType(format)
    case "integer":
        return getBpIntegerType(format, minimum)
    case "number":
        return getNumberType(format)
    case "boolean":
        return "bool"
    case "array":
        return "array"
    case "file":
        return "bytes"
    default:
        return "object"
    }
}

func getPbStringType(format string) string {
    switch format {
    case "byte", "binary":
        return "bytes"
    default:
        return "string"
    }
}

func getBpIntegerType(format string, minimum *float64) string {
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
    case "uint64":
        return "uint64"
    case "uint32":
        return "uint32"
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
