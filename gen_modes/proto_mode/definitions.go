package proto_mode

import (
	"fmt"
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"sort"
	"strings"
)

func (p *ProtoMode) prepareDefinitions(definitions spec4pb.Definitions) {

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
	var maxNum int32 = 0
	var fieldSort = SortFieldOpts{}
	if len(definition.AllOf) > 0 {
		panicErr("not support allOf", locations, name, "nil")
	}
	if len(definition.OneOf) > 0 {
		panicErr("not support oneOf", locations, name, "nil")
	}
	if len(definition.AnyOf) > 0 {
		panicErr("not support anyOf", locations, name, "nil")
	}
	for fieldName, property := range definition.Properties {
		if property.FieldNumber > maxNum {
			maxNum = property.FieldNumber
		}
		fieldSort = append(fieldSort, NewFieldOpt(fieldName, property.FieldNumber))
	}
	sort.Sort(fieldSort)
	g.P("message ", name, " {")
	for _, field := range fieldSort {
		fieldName := field.FieldName
		property := definition.Properties[fieldName]
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
