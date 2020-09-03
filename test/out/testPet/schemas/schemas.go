/*
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
)

var schemaString = `{
    "PostPet" : [{"fieldNumber":1,"description":"Pet object that needs to be added to the store","name":"body","in":"body","required":true,"schema":{"$ref":"#/definitions/Pet"}}],
    "PutPet" : [{"fieldNumber":1,"description":"Pet object that needs to be added to the store","name":"pet","in":"body","required":true,"schema":{"$ref":"#/definitions/Pet"}}],
    "GetPetFindByStatus" : [{"type":"array","fieldNumber":1,"items":{"enum":["available","pending","sold"],"type":"string","default":"available"},"collectionFormat":"multi","description":"Status values that need to be considered for filter","name":"status","in":"query","required":true}],
    "GetPetFindByTags" : [{"type":"array","fieldNumber":1,"items":{"type":"string"},"collectionFormat":"multi","description":"Tags to filter by","name":"tags","in":"query","required":true}],
    "DeletePetPetId" : [{"type":"integer","fieldNumber":1,"format":"int64","description":"Pet id to delete","name":"petId","in":"path","required":true}],
    "GetPetPetId" : [{"type":"integer","fieldNumber":1,"format":"int64","description":"ID of pet to return","name":"petId","in":"path","required":true}],
    "PostPetPetId" : [{"type":"integer","fieldNumber":1,"format":"int64","description":"ID of pet that needs to be updated","name":"petId","in":"path","required":true},{"type":"string","fieldNumber":2,"description":"Updated name of the pet","name":"name","in":"formData"},{"type":"string","fieldNumber":3,"description":"Updated status of the pet","name":"status","in":"formData"}],
    "PostPetPetIdUploadImage" : [{"type":"integer","fieldNumber":1,"format":"int64","description":"ID of pet to update","name":"petId","in":"path","required":true},{"type":"string","fieldNumber":2,"description":"Additional data to pass to server","name":"additionalMetadata","in":"formData"},{"type":"file","fieldNumber":3,"description":"file to upload","name":"file","in":"formData"}],
    "GetStoreInventory" : null,
    "PostStoreOrder" : [{"fieldNumber":1,"description":"order placed for purchasing the pet","name":"body","in":"body","required":true,"schema":{"$ref":"#/definitions/Order"}}],
    "DeleteStoreOrderOrderId" : [{"minimum":1,"type":"integer","fieldNumber":1,"format":"int64","description":"ID of the order that needs to be deleted","name":"orderId","in":"path","required":true}],
    "GetStoreOrderOrderId" : [{"maximum":10,"minimum":1,"type":"integer","fieldNumber":1,"format":"int64","description":"ID of pet that needs to be fetched","name":"orderId","in":"path","required":true}],
    "PostUser" : [{"fieldNumber":1,"description":"Created user object","name":"body","in":"body","required":true,"schema":{"$ref":"#/definitions/User"}}],
    "PostUserCreateWithArray" : [{"fieldNumber":1,"description":"List of user object","name":"body","in":"body","required":true,"schema":{"type":"array","items":{"$ref":"#/definitions/User"}}}],
    "PostUserCreateWithList" : [{"fieldNumber":1,"description":"List of user object","name":"body","in":"body","required":true,"schema":{"type":"array","items":{"$ref":"#/definitions/User"}}}],
    "GetUserLogin" : [{"type":"string","fieldNumber":1,"description":"The user name for login","name":"username","in":"query","required":true},{"type":"string","fieldNumber":2,"description":"The password for login in clear text","name":"password","in":"query","required":true}],
    "GetUserLogout" : null,
    "DeleteUserUsername" : [{"type":"string","fieldNumber":1,"description":"The name that needs to be deleted","name":"username","in":"path","required":true}],
    "GetUserUsername" : [{"type":"string","fieldNumber":1,"description":"The name that needs to be fetched. Use user1 for testing.","name":"username","in":"path","required":true}],
    "PutUserUsername" : [{"type":"string","fieldNumber":1,"description":"name that need to be updated","name":"username","in":"path","required":true},{"fieldNumber":2,"description":"Updated user object","name":"body","in":"body","required":true,"schema":{"$ref":"#/definitions/User"}}]
}`


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
}
