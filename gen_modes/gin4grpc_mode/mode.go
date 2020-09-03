package gin4grpc_mode

import (
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"sync"
)

const ModeName = "gin4grpc_mode"

func RegistMode(gen *swagger_gen.SwaggerGenerator, outPath ...string) {
	op := ""
	if len(outPath) > 0 {
		op = outPath[0]
	}
	gen.AddMode(ModeName, &Gin4GrpcMode{outPath: op})
}

type Gin4GrpcMode struct {
	sync.Mutex
	swaggerPub *mode_pub.SwaggerPub
	outPath    string
	services   sortServices
}

func (p *Gin4GrpcMode) New() mode_pub.ModeGenInterface {
	return &Gin4GrpcMode{outPath: p.outPath}
}

func (p *Gin4GrpcMode) GenFile(outpath string, swaggerPub *mode_pub.SwaggerPub) error {
	p.swaggerPub = swaggerPub
	if p.outPath == "" {
		p.outPath = outpath
	}
	p.prepareServices()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer func() {
			wg.Done()
		}()
		p.genRouters()
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		p.genHubs()
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		p.genSchemas()
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		p.genServices()
	}()
	wg.Wait()
	return nil
}

func (p *Gin4GrpcMode) GenImportPb(g mode_pub.BufGenInterface) {
	g.P("    pb \"", p.swaggerPub.GoPackageName, "/", p.swaggerPub.PackageName, "/pb\"")
}

func (p *Gin4GrpcMode) GenImportSchemas(g mode_pub.BufGenInterface) {
	g.P("    \"", p.swaggerPub.GoPackageName, "/", p.swaggerPub.PackageName, "/schemas\"")
}

func (p *Gin4GrpcMode) GenImportServices(g mode_pub.BufGenInterface) {
	g.P("    \"", p.swaggerPub.GoPackageName, "/", p.swaggerPub.PackageName, "/services\"")
}
