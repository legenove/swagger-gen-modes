package gin4grpc_mode

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
)

func (p *Gin4GrpcMode) genSchemas() {
	g := swagger_gen.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/schemas", p.swaggerPub.Md5)
	g.SetFilename("schemas.go")
	g.P("package schemas")
	g.P(`import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/legenove/spec4pb"
)`)
	g.P("/*")
	g.P("### DO NOT CHANGE THIS FILE")
	g.P("### The code is auto generated, your change will be overwritten by")
	g.P("### code generating.")
	g.P("*/")
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
	g.P(schemaFormatString)
	g.GenFile()
}

var schemaFormatString = `

var schemaMap map[string][]spec4pb.Parameter
var schemaMapMap map[string]map[string]spec4pb.Parameter

func init() {
	jsoniter.Unmarshal([]byte(schemaString), &schemaMap)
	schemaMapMap = make(map[string]map[string]spec4pb.Parameter, len(schemaMap))
	for k, ps := range schemaMap {
		schemaMapMap[k] = make(map[string]spec4pb.Parameter, len(ps))
		for _, p := range ps {
			schemaMapMap[k][p.Name] = p
		}
	}
}

func GetParams(c *gin.Context, key string, in interface{}) map[string][]string {
	var headers = map[string][]string{}
	scheMap := schemaMapMap[key]
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
		if inputFieldName == "-" {
			continue
		}
		var values []string
		var hasValue bool = true
		p := scheMap[inputFieldName]
		switch p.In {
		case "body":
			// 结构体
			switch c.ContentType() {
			case "application/json":
				err := c.ShouldBind(&structField)
				fmt.Println(err)
			case "application/xml":
				err := c.ShouldBind(&structField)
				fmt.Println(err)
			}
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
				if err := setWithProperType(sliceOf, values[i], slice.Index(i)); err != nil {
					return headers
				}
			}
			val.Field(i).Set(slice)
		} else {
			if _, isTime := structField.Interface().(time.Time); isTime {
				if err := setTimeField(values[0], typeField, structField); err != nil {
					return headers
				}
				continue
			}
			if err := setWithProperType(typeField.Type.Kind(), values[0], structField); err != nil {
				return headers
			}
		}
	}
	return headers
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
}`