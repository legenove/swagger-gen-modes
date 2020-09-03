package gin4grpc_mode

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"sort"
)

func (p *Gin4GrpcMode) genSchemas() {
	g := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/schemas", p.swaggerPub.Md5)
	g.SetFilename("schemas.go")
	g.P(schemaHeader)
	g.P()
	g.P("var schemaString = `{")
	for i, s := range p.services {
		jb, _ := jsoniter.Marshal(s.Params)
		g.Pl("    \"", s.FuncName, "\" : ", string(jb))
		if i < len(p.services)-1 {
			g.Pl(",")
		}
		g.P()
	}
	g.P("}`")
	g.P(schemaMainBody)
	g.GenFile()

	for _, s := range p.services {
		g = mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/schemas", p.swaggerPub.Md5)
		g.SetFilename("schemas" + s.FuncName + ".go")
		fields := prepareParams(s.Params)
		g.P("package schemas")
		g.P()
		g.P("import (")
		g.P("    \"github.com/gin-gonic/gin\"")
		p.GenImportPb(g)
		fields.ImportsG(g)
		g.P(")")
		g.P()
		g.P("func Get", s.FuncName, "Params(c *gin.Context, in *pb.", s.ReqName, ") (map[string][]string, error) {")
		fields.MergeG(g)
		g.P("    return c.Request.Header, nil")
		g.P("}")
		g.GenFile()
	}
}

type paramField struct {
	fieldName string
	g         mode_pub.BufGenInterface
	imports   map[string]bool
}

type sortParamFields []*paramField

func (s sortParamFields) Less(i, j int) bool {
	return s[i].fieldName < s[j].fieldName
}
func (s sortParamFields) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortParamFields) Len() int      { return len(s) }
func (s sortParamFields) MergeG(g mode_pub.BufGenInterface) {
	sort.Sort(s)
	for _, j := range s {
		if j != nil && len(j.g.GetBytes()) > 0 {
			g.P(string(j.g.GetBytes()))
		}
	}
}
func (s sortParamFields) ImportsG(g mode_pub.BufGenInterface) {
	imports := map[string]bool{}
	for _, j := range s {
		for i := range j.imports {
			if _, ok := imports[i]; ok {
				continue
			}
			imports[i] = true
			g.P("    \"", i, "\"")
		}
	}
}

func prepareParams(params []spec4pb.Parameter) sortParamFields {
	var fields = make(sortParamFields, 0, len(params))
	for _, p := range params {
		field := prepareParam(p)
		if field == nil {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

func prepareParam(param spec4pb.Parameter) *paramField {
	field := &paramField{param.Name, &mode_pub.BufGenerator{}, map[string]bool{}}
	_type := common.GetGoType(param)
	goName := common.GetGoName(param.Name)
	var valType string
	switch _type {
	case "array":
		valType = "[]" + GetItemType(param.Items)
	case "struct":
		if param.Schema != nil && param.Schema.Ref.GetURL() != nil {
			valType = "*pb." + common.GetSchemaName(param.Schema)
		} else {
			valType = "interface{}"
		}
	default:
		valType = _type
	}
	switch param.In {
	case "path":
		getPathVal(field, goName, valType, _type, param)
	case "query":
		getQueryVal(field, goName, valType, _type, param)
	case "formData":
		getFormDataVal(field, goName, valType, _type, param)
	case "body":
		getBodyVal(field, goName, valType, _type, param)
	case "header":
		return nil
	default:
		return nil
	}
	return field
}

func getBodyVal(field *paramField, goName, valType, _type string, param spec4pb.Parameter) {
	fieldMaps := map[string]interface{}{"goName": goName, "name": param.Name, "valType": valType}
	field.g.P(FormatByKey(`    // body {{.goName}}
	var val {{.valType}}
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.{{.goName}} = val`, fieldMaps))
}

func getPathVal(field *paramField, goName, valType, _type string, param spec4pb.Parameter) {
	fieldMaps := map[string]interface{}{"goName": goName, "name": param.Name, "valType": valType}
	field.g.P(FormatByKey(`    // path {{.goName}}
	var val{{.goName}} {{.valType}}
	if val, ok := c.Params.Get("{{.name}}"); ok {
		_v, err := setWithKind("{{.valType}}", val)
		if err != nil {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
		val{{.goName}}, ok = _v.({{.valType}})
		if !ok {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
	} else {
		return nil, errors.New("{{.name}} required")
	}`, fieldMaps))
	addParamsValidate(field, goName, valType, _type, param)
	field.g.P(FormatByKey("    in.{{.goName}} = val{{.goName}}", fieldMaps))
	field.imports["errors"] = true
}

func getQueryVal(field *paramField, goName, valType, _type string, param spec4pb.Parameter) {
	fieldMaps := map[string]interface{}{"goName": goName, "name": param.Name, "valType": valType, "subType": valType[2:]}
	field.g.P(FormatByKey(`    // query {{.goName}}
	var val{{.goName}} {{.valType}}
	if val, ok := c.GetQueryArray("{{.name}}"); ok {`, fieldMaps))
	if _type == "array" {
		field.g.Pl(FormatByKey(`        val{{.goName}} = make({{.valType}}, len(val))
        for i, v := range val {
            err := setWithKind("{{.subType}}", val[i], &val{{.goName}}[i])
		    if err != nil {
				return nil, errors.New("{{.name}} value not {{.valType}}")
			}
        }
	}`, fieldMaps))
	} else {
		field.g.Pl(FormatByKey(`		_v, err := setWithKind("{{.valType}}", val[0])
		if err != nil {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
		val{{.goName}}, ok = _v.({{.valType}})
		if !ok {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
	}`, fieldMaps))
	}
	if param.Default == nil && param.Required == true {
		field.g.P(" else {")
		field.g.P("        return nil, errors.New(\"" + param.Name + " required\")")
		field.g.P("    }")
	} else {
		field.g.P()
	}
	addParamsValidate(field, goName, valType, _type, param)
	field.g.P(FormatByKey("    in.{{.goName}} = val{{.goName}}", fieldMaps))
	field.imports["errors"] = true
}

func getFormDataVal(field *paramField, goName, valType, _type string, param spec4pb.Parameter) {
	fieldMaps := map[string]interface{}{"goName": goName, "name": param.Name, "valType": valType, "subType": valType[2:]}
	field.g.P(FormatByKey(`    // formData {{.goName}}
	var val{{.goName}} {{.valType}}
	if val, ok := c.GetPostFormArray("{{.name}}"); ok {`, fieldMaps))
	if _type == "array" {
		field.g.Pl(FormatByKey(`        val{{.goName}} = make({{.valType}}, len(val))
        for i := range val {
            _v, err := setWithKind("{{.subType}}", val[i])
		    if err != nil {
				return nil, errors.New("{{.name}} value not {{.valType}}")
			}
			&val{{.goName}}[i] = _v.({{.valType}})
        }
	}`, fieldMaps))
	} else {
		field.g.Pl(FormatByKey(`		_v, err := setWithKind("{{.valType}}", val[0])
		if err != nil {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
		val{{.goName}}, ok = _v.({{.valType}})
		if !ok {
			return nil, errors.New("{{.name}} value not {{.valType}}")
		}
	}`, fieldMaps))
	}
	if param.Default == nil && param.Required == true {
		field.g.P(" else {")
		field.g.P("        return nil, errors.New(\"" + param.Name + " required\")")
		field.g.P("    }")
	} else {
		field.g.P()
	}
	addParamsValidate(field, goName, valType, _type, param)
	field.g.P(FormatByKey("    in.{{.goName}} = val{{.goName}}", fieldMaps))
	field.imports["errors"] = true
}

func GetItemType(items *spec4pb.Items) string {
	_type := common.GetPBType(items)
	switch _type {
	case "array":
		panic("+++-- GPItem, default error")
	case "object":
		panic("+++-- GPItem, default error")
	default:
		return _type
	}
}

var schemaHeader = `/*
### DO NOT CHANGE THIS FILE
### The code is auto generated, your change will be overwritten by
### code generating.
*/
package schemas

import (
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/legenove/spec4pb"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"strings"
	"time"
)`

var schemaMainBody = `

var schemaMap map[string][]spec4pb.Parameter
var schemaMapMap map[string]map[string]spec4pb.Parameter
var schemaHasBody map[string]string

func init() {
	jsoniter.Unmarshal([]byte(schemaString), &schemaMap)
	schemaMapMap = make(map[string]map[string]spec4pb.Parameter, len(schemaMap))
	schemaHasBody = make(map[string]string, len(schemaMap))
	for k, ps := range schemaMap {
		schemaMapMap[k] = make(map[string]spec4pb.Parameter, len(ps))
		schemaHasBody[k] = ""
		for _, p := range ps {
			schemaMapMap[k][p.Name] = p
			if p.In == "body" {
				schemaHasBody[k] = p.Name
			}
		}
	}
}

func prepareBody(c *gin.Context, name string, in interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:  "json",
		Metadata: nil,
		Result:   &in,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	var val interface{}
	err = c.ShouldBind(&val)
	if err != nil {
		return err
	}
	return decoder.Decode(map[string]interface{}{name: val})
}

func GetParams(c *gin.Context, key string, in interface{}) (headers map[string][]string, err error) {
	scheMap := schemaMapMap[key]
	if schemaHasBody[key] != "" {
		err = prepareBody(c, schemaHasBody[key], in)
		if err != nil {
			return
		}
	}
	typ := reflect.TypeOf(in).Elem()
	val := reflect.ValueOf(in).Elem()
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}
		inputFieldName := typeField.Tag.Get("json")
		if inputFieldName == "-" {
			continue
		}
		inputFieldNameList := strings.Split(inputFieldName, ",")
		inputFieldName = inputFieldNameList[0]
		var values []string
		var hasValue bool = true
		p := scheMap[inputFieldName]
		switch p.In {
		case "body":
			continue
		case "formData":
			values, hasValue = c.GetPostFormArray(inputFieldName)
		case "path":
			_, hasValue = c.Params.Get(inputFieldName)
			if hasValue {
				values = []string{c.Params.ByName(inputFieldName)}
			}
		case "query":
			values, hasValue = c.GetQueryArray(inputFieldName)
		case "header":
			// TODO 调试header值
			continue
		}
		if !hasValue || len(values) == 0 {
			continue
		}
		numElems := len(values)
		structFieldKind := structField.Kind()
		if structFieldKind == reflect.Slice && numElems > 0 {
			sliceOf := structField.Type().Elem().Kind()
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
			for i := 0; i < numElems; i++ {
				if err = setWithProperType(sliceOf, values[i], slice.Index(i)); err != nil {
					return
				}
			}
			val.Field(i).Set(slice)
		} else {
			if _, isTime := structField.Interface().(time.Time); isTime {
				if err = setTimeField(values[0], typeField, structField); err != nil {
					return
				}
				continue
			}
			if err = setWithProperType(typeField.Type.Kind(), values[0], structField); err != nil {
				return
			}
		}
	}
	return
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	case reflect.Ptr:
		if !structField.Elem().IsValid() {
			structField.Set(reflect.New(structField.Type().Elem()))
		}
		structFieldElem := structField.Elem()
		return setWithProperType(structFieldElem.Kind(), val, structFieldElem)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		return errors.New("Blank time format")
	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func setWithKind(valueKind string, val string) (interface{}, error) {
	switch valueKind {
	case "int":
		_v, err := setInt(val, 0)
		return int(_v), err
	case "int8":
		_v, err := setInt(val, 8)
		return int8(_v), err
	case "int16":
		_v, err := setInt(val, 16)
		return int16(_v), err
	case "int32":
		_v, err := setInt(val, 32)
		return int32(_v), err
	case "int64":
		_v, err := setInt(val, 64)
		return int64(_v), err
	case "uint":
		_v, err := setUint(val, 0)
		return uint(_v), err
	case "uint8":
		_v, err := setUint(val, 8)
		return uint8(_v), err
	case "uint16":
		_v, err := setUint(val, 16)
		return uint16(_v), err
	case "uint32":
		_v, err := setUint(val, 32)
		return uint32(_v), err
	case "uint64":
		_v, err := setUint(val, 64)
		return _v, err
	case "bool":
		_v, err := setBool(val)
		return _v, err
	case "float32":
		_v, err := setFloat(val, 32)
		return float32(_v), err
	case "float64":
		_v, err := setFloat(val, 64)
		return _v, err
	case "string":
		return val, nil
	case "[]byte":
		return []byte(val), nil
	default:
		return nil, errors.New("Unknown type")
	}
}

func setInt(val string, bitSize int) (int64, error) {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		return intVal, nil
	}
	return 0, err
}

func setUint(val string, bitSize int) (uint64, error) {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		return uintVal, nil
	}
	return 0, err
}

func setBool(val string) (bool, error) {
	switch val {
	case "", "0", "false":
		return false, nil
	case "1", "true":
		return true, nil
	default:
		return true, nil
	}
}

func setFloat(val string, bitSize int) (float64, error) {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		return floatVal, nil
	}
	return 0, err
}`
