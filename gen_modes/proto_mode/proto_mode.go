package proto_mode

import (
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"path"
	"sync"
)

const ModeName = "proto3_mode"

var reoutSwagger bool = false

func RegistMode(gen *swagger_gen.SwaggerGenerator, outPath ...string) {
	op := ""
	if len(outPath) > 0 {
		op = outPath[0]
	}
	gen.AddMode(ModeName, &ProtoMode{outPath: op})
}

type ProtoMode struct {
	sync.Mutex
	outPath       string
	swaggerPub    *mode_pub.SwaggerPub
	g             *mode_pub.FileGenerator
	messageGenOpt SortBufGenOpts
	serviceGenOpt SortBufGenOpts
	imports       map[string]bool
}

func (p *ProtoMode) New() mode_pub.ModeGenInterface {
	return &ProtoMode{outPath: p.outPath}
}

func (p *ProtoMode) GenFile(outpath string, swaggerPub *mode_pub.SwaggerPub) error {
	if p.outPath != "" {
		outpath = p.outPath
	}
	p.swaggerPub = swaggerPub
	p.messageGenOpt = SortBufGenOpts{}
	p.serviceGenOpt = SortBufGenOpts{}
	p.imports = map[string]bool{}
	g := mode_pub.NewFileGen(outpath+ "/" + swaggerPub.PackageName, swaggerPub.Md5)
	p.g = g
	g.SetFilename(swaggerPub.PackageName + ".proto")
	g.P("// Code generated by swagger-gen mode proto. DO NOT EDIT.")
	g.P()
	p.genHeader(g)
	p.prepareServices()
	p.prepareMessages()
	p.genImport(g)
	p.genServices(g)
	p.genMessages(g)
	err := g.GenFile()
	if err != nil {
		return err
	}
	if reoutSwagger {
		yb, err := p.swaggerPub.ToYaml()
		if err != nil {
			return err
		}
		pth := path.Join(outpath, swaggerPub.PackageName+"ForProto.yaml")
		err = g.WriteFile(pth, yb)
	}
	return err
}

func (p *ProtoMode) genHeader(g mode_pub.BufGenInterface) {
	if p.swaggerPub.Swagger.Info != nil {
		g.P("// Title: ", p.swaggerPub.Swagger.Info.Title)
		g.P("// Version: ", p.swaggerPub.Swagger.Info.Version)
		g.P("// Description: ", p.swaggerPub.Swagger.Info.Description)
		g.P()
	}
	g.P("syntax = \"proto3\";")
	g.P("package ", p.swaggerPub.PackageName, ";")
	g.P()
}
