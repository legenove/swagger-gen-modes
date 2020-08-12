package gin4grpc_mode

import (
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"sync"
)

const ModeName = "gin4grpc_mode"

func RegistMode(gen *swagger_gen.SwaggerGenerator) {
	gen.AddMode(ModeName, &Gin4GrpcMode{})
}

type Gin4GrpcMode struct {
	sync.Mutex
	swaggerPub *swagger_gen.SwaggerPub
	outPath    string
	services   sortServices
}

func (p *Gin4GrpcMode) GenFile(outpath string, swaggerPub *swagger_gen.SwaggerPub) error {
	p.swaggerPub = swaggerPub
	p.outPath = outpath
	p.prepareServices()
	p.genRouters()
	p.genHubs()
	p.genSchemas()
	return nil
}
