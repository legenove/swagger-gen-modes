package gin4grpc_mode

import (
	"bytes"
	"github.com/legenove/spec4pb"
	"text/template"
)

func addParamsValidate(field *paramField, goName, valType, _type string, param spec4pb.Parameter){
	switch _type {
	case "string":
		if param.Pattern != "" {
			field.g.P(FormatByKey(paramStringPattern, map[string]interface{}{"goName": goName, "pattern": param.Pattern}))
			field.imports["errors"] = true
			field.imports["regexp"] = true
		}
		if param.MaxLength != nil {
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "len(val" + goName + ")",
					"symbol": "<=",
					"number": *param.MaxLength,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
		if param.MinLength != nil {
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "len(val" + goName + ")",
					"symbol": ">=",
					"number": *param.MinLength,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
	case "int32", "uint32", "uint64", "int64":
		if param.Minimum != nil {
			symbolType := ">="
			if param.ExclusiveMinimum {
				symbolType = ">"
			}
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "in." + goName,
					"symbol": symbolType,
					"number": int(*param.Minimum),
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
		if param.Maximum != nil {
			symbolType := "<="
			if param.ExclusiveMaximum {
				symbolType = "<"
			}
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "in." + goName,
					"symbol": symbolType,
					"number": int(*param.Maximum),
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
	case "float32", "float64":
		if param.Minimum != nil {
			symbolType := ">="
			if param.ExclusiveMinimum {
				symbolType = ">"
			}
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "in." + goName,
					"symbol": symbolType,
					"number": *param.Minimum,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
		if param.Maximum != nil {
			symbolType := "<="
			if param.ExclusiveMaximum {
				symbolType = "<"
			}
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "in." + goName,
					"symbol": symbolType,
					"number": *param.Maximum,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
	case "bool":
	case "[]byte":
	case "struct":
	case "array":
		if param.MaxItems != nil {
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "len(in." + goName + ")",
					"symbol": "<=",
					"number": *param.MaxItems,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}
		if param.MinItems != nil {
			field.g.P(FormatByKey(paramNumberOrStringCmp,
				map[string]interface{}{"condition": "len(in." + goName + ")",
					"symbol": ">=",
					"number": *param.MinItems,
					"goName": goName,
				}))
			field.imports["errors"] = true
		}

	}
}

func FormatByKey(f string, m map[string]interface{}) string {
	var tpl bytes.Buffer
	t := template.Must(template.New("").Parse(f))
	if err := t.Execute(&tpl, m); err != nil {
		return ""
	}
	return tpl.String()
}
var paramStringPattern = `    re := regexp.MustCompile("{{.pattern}}")
	if !re.MatchString(in.{{.goName}}) {
		return nil, errors.New("{{.goName}} not match with '{{.pattern}}'")
	}`

var paramNumberOrStringCmp = `    if !({{.condition}} {{.symbol}} {{.number}}) {
	    return nil, errors.New("{{.goName}} len must {{.symbol}} {{.number}}")
    }`
