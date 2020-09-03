package mode_pub

import (
	"github.com/legenove/spec4pb"
	"gopkg.in/yaml.v2"
)

var swaggerSortKeys = []string{"swagger", "info", "schemes", "tags", "externalDocs", "host",
	"basePath", "paths", "definitions", "securityDefinitions"}

type SwaggerPub struct {
	Swagger       *spec4pb.Swagger
	PackageName   string
	GoPackageName string
	Md5           string
}

func (p *SwaggerPub) ToYaml() ([]byte, error) {
	jb, err := p.Swagger.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var jsonObj map[string]interface{}
	err = yaml.Unmarshal(jb, &jsonObj)
	if err != nil {
		return nil, err
	}
	g := BufGenerator{}
	for _, k := range swaggerSortKeys {
		if v, ok := jsonObj[k]; ok {
			b, err := yaml.Marshal(map[string]interface{}{k: v})
			if err != nil {
				return nil, err
			}
			g.Pl(string(b))
			delete(jsonObj, k)
		}
	}
	for k, v := range jsonObj {
		b, err := yaml.Marshal(map[string]interface{}{k: v})
		if err != nil {
			return nil, err
		}
		g.Pl(string(b))
	}
	return g.GetBytes(), nil
}
