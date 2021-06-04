package proto_mode

import (
    "errors"
    "fmt"
    "sort"
    "strings"
    "sync"

    "github.com/legenove/spec4pb"
    "github.com/legenove/swagger-gen-modes/gen_modes/common"
    "github.com/legenove/swagger-gen-modes/mode_pub"
)

var definitionsLock = sync.Mutex{}

func (p *ProtoMode) prepareDefinitions(definitions spec4pb.Definitions) {
    definitionsLock.Lock()
    definitionsLock.Unlock()
    for name, definition := range definitions {

        if len(definition.Type) == 0 || definition.Type.Contains("object") {
            locations := fmt.Sprintf("%d:%s", common.OptLocationMap[common.DefinitionPreName], strings.Title(name))
            GPProperties(p, &definition, "Get", locations, common.DefinitionPreName+strings.Title(name))
        } else {
            panic(" definitions " + name + " must object")
        }
        p.swaggerPub.Swagger.Definitions[name] = definition
    }
}

func (p *ProtoMode) prepareSecurityDefinitions(definitions spec4pb.SecurityDefinitions) {
    //for k, v := range definitions {
    //
    //}
}

func panicErr(errstr, locations, name, fieldName string) {
    ls := strings.Split(locations, ":")
    if len(ls) > 1 {
        panic(fmt.Sprintf("%s: location: %s, objName: %s, name: %s, field: %s",
            errstr, common.OptReLocationMap[ls[0]], ls[1], name, fieldName))
    }
    panic(fmt.Sprintf("%s: loaction: %s,name: %s, field: %s",
        errstr, locations, name, fieldName))
}

func GPProperties(p *ProtoMode, definition *spec4pb.Schema, method, locations, name string) {
    g := &mode_pub.BufGenerator{}
    fieldSort := CheckSchema(definition, locations, name)
    g.P("message ", name, " {")
    for _, field := range fieldSort {
        fieldName := field.FieldName
        property := field.Propertie
        if property.FieldNumber == 0 {
            panicErr("fildNumber not define", locations, name, fieldName)
        }
        g.Pl("  ")
        GPSchema(g, &property, method, locations, name+strings.Title(fieldName), "", p)
        GPFieldEnd(g, fieldName, property.FieldNumber, property.Description)
    }
    g.P("}")
    p.Lock()
    defer p.Unlock()
    p.messageGenOpt = append(p.messageGenOpt,
        &BufGenOpt{
            locations,
            common.OptMethodMap[method],
            name,
            g},
    )
}

func CheckSchema(definition *spec4pb.Schema, locations, name string) SortFieldOpts {
    var fieldSort = SortFieldOpts{}
    var fieldNumMapper = map[int32]bool{}
    var fieldNameMapper = map[string]bool{}
    if len(definition.OneOf) > 0 {
        panicErr("not support oneOf", locations, name, "nil")
    }
    if len(definition.AnyOf) > 0 {
        panicErr("not support anyOf", locations, name, "nil")
    }
    var err error
    if len(definition.AllOf) > 0 {
        fieldSort, err = CheckAllOf(fieldSort, fieldNumMapper, fieldNameMapper, definition.AllOf, p.swaggerPub.Swagger)
        if err != nil {
            panicErr(err.Error(), locations, name, "allOf")
        }
    }
    for fieldName, property := range definition.Properties {
        if _, ok := fieldNumMapper[property.FieldNumber]; ok {
            panicErr(fmt.Sprintf("fieldNumber duplicate: fieldName: %s  ; fieldNumber : %d",
                fieldName, property.FieldNumber), locations, name, fieldName)
        }
        if _, ok := fieldNameMapper[fieldName]; ok {
            panicErr(fmt.Sprintf("fieldName duplicate: fieldName: %s;",
                fieldName), locations, name, fieldName)
        }
        fieldNumMapper[property.FieldNumber] = true
        fieldNameMapper[fieldName] = true
        fieldSort = append(fieldSort, NewFieldOpt(fieldName, property.FieldNumber, definition.Properties[fieldName]))
    }
    sort.Sort(fieldSort)
    return fieldSort
}

func CheckAllOf(fieldSort SortFieldOpts, fieldNumMapper map[int32]bool, fieldNameMapper map[string]bool, allOf []spec4pb.Schema, swagger *spec4pb.Swagger) (SortFieldOpts, error) {
    var err error
    for i := range allOf {
        name := allOf[i].Ref.Ref.GetURL().String()
        if !strings.HasPrefix(name, "#/definitions/") {
            return fieldSort, errors.New("allOf must in definitions")
        }
        _schema, ok := swagger.Definitions[strings.ReplaceAll(name, "#/definitions/", "")]
        if !ok {
            return fieldSort, errors.New("allOf definiions name not found:" + name)
        }
        fieldSort, err = CheckDefinition(fieldSort, fieldNumMapper, fieldNameMapper, _schema)
        if err != nil {
            return fieldSort, err
        }
        if len(_schema.AllOf) > 0 {
            fieldSort, err = CheckAllOf(fieldSort, fieldNumMapper, fieldNameMapper, _schema.AllOf, swagger)
            if err != nil {
                return fieldSort, err
            }
        }
    }
    return fieldSort, nil
}

func CheckDefinition(fieldSort SortFieldOpts, fieldNumMapper map[int32]bool, fieldNameMapper map[string]bool, definition spec4pb.Schema) (SortFieldOpts, error) {
    for fieldName, property := range definition.Properties {
        if _, ok := fieldNumMapper[property.FieldNumber]; ok {
            return nil, errors.New(fmt.Sprintf("fieldNumber duplicate: fieldName: %s  ; fieldNumber : %d",
                fieldName, property.FieldNumber))
        }
        if _, ok := fieldNameMapper[fieldName]; ok {
            return nil, errors.New(fmt.Sprintf("fieldName duplicate: fieldName: %s;",
                fieldName))
        }
        fieldNumMapper[property.FieldNumber] = true
        fieldNameMapper[fieldName] = true
        fieldSort = append(fieldSort, NewFieldOpt(fieldName, property.FieldNumber, definition.Properties[fieldName]))
    }

    return fieldSort, nil
}
