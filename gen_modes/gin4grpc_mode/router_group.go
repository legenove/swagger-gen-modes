package gin4grpc_mode

import (
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"strings"
)

func (p *Gin4GrpcMode) genRouters() {
	g := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName, p.swaggerPub.Md5)
	g.SetFilename("router_group.go")
	p.GenDoNotChange(g)
	g.P("package ", p.swaggerPub.PackageName)
	g.P("import (")
	g.P("    \"github.com/legenove/nano-server-sdk/gincore\"")
	g.P("    \"github.com/legenove/nano-server-sdk/grpccore\"")
	g.P("    \"google.golang.org/grpc\"")
	g.P()
	g.P("    \"", p.swaggerPub.GoPackageName, "/", p.swaggerPub.PackageName, "/hubs\"")
	p.GenImportPb(g)
	p.GenImportServices(g)
	g.P(")")
	g.P()
	g.P("const basePath = \"", p.swaggerPub.Swagger.BasePath, "\"")
	g.P()
	g.P("func init() {")
	g.P("    grpccore.RegisterToServer(\"" + p.swaggerPub.PackageName + "\", func(s *grpc.Server) {")
	g.P("        pb.Register" + strings.Title(p.swaggerPub.PackageName) + "Server(s, services.NewServer())")
	g.P("    })")
	g.P("    group := gincore.GetCurrentGroup(basePath)")
	for _, s := range p.services {
		g.Pl("    ")
		g.Pl("group.", strings.ToUpper(s.Method), "(\"", s.getRouterPath(), "\", ")
		g.Pl("decoratorHandler(hubs.", s.FuncName, "))")
		g.P()
	}
	g.P("}")
	g.GenFile()

	g = mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName, p.swaggerPub.Md5)
	g.SetFilename("router.go")
	//g.Skip()
	g.P("package ", p.swaggerPub.PackageName)
	g.P()
	g.P("import (")
	g.P("    \"github.com/gin-gonic/gin\"")
	g.P("    \"github.com/legenove/nano-server-sdk/gincore\"")
	g.P(")")
	g.P()
	g.P("type ApiBaseHandler func(c *gin.Context) (int, interface{})")
	g.P()
	g.P(`func decoratorHandler(handler ApiBaseHandler, decors ...gincore.HandlerDecorator) gin.HandlerFunc {
	apiFunc := func(c *gin.Context) {
		code, obj := handler(c)
		if obj != nil {
			c.JSON(code, obj)
		} else {
			c.Status(code)
		}
	}
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		apiFunc = d(apiFunc)
	}
	return apiFunc
}
`)
	g.GenFile()

}
