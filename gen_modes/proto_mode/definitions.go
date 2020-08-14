package proto_mode

import (
	"fmt"
	"github.com/legenove/swagger-gen-modes/gen_modes/common"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/spec4pb"
	"sort"
	"strings"
)

func (p *ProtoMode) prepareDefinitions(definitions spec4pb.Definitions) {

	for name, definition := range definitions {

		if len(definition.Type) == 0 || definition.Type.Contains("object") {
			locations := fmt.Sprintf("%d%s",common.OptLocationMap[common.DefinitionPreName], strings.Title(name))
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

func GPProperties(p *ProtoMode, definition *spec4pb.Schema, method, locations, name string) {
	g := &swagger_gen.BufGenerator{}
	var maxNum int32 = 0
	var fieldSort = SortFieldOpts{}
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
			reoutSwagger = true
			maxNum = maxNum + 1
			property.FieldNumber = maxNum
			// 覆盖原有数据
			definition.Properties[fieldName] = property
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
