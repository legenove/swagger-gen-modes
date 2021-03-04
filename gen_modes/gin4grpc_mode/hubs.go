package gin4grpc_mode

import (
	"github.com/legenove/swagger-gen-modes/mode_pub"
)

func (p *Gin4GrpcMode) genHubs() {
	g := mode_pub.NewFileGen(p.outPath+"/"+p.swaggerPub.PackageName+"/hubs", p.swaggerPub.Md5)
	g.SetFilename("gin_hubs.go")
	g.P("package hubs")
	g.P("/*")
	g.P("### DO NOT CHANGE THIS FILE")
	g.P("### The code is auto generated, your change will be overwritten by")
	g.P("### code generating.")
	g.P("*/")
	g.P("import (")
	//g.P("    \"", p.swaggerPub.GoPackageName, "/core\"")
	g.P("    \"github.com/gin-gonic/gin\"")
	g.P("    \"github.com/legenove/nano-server-sdk/servers\"")
	g.P("    \"google.golang.org/grpc/metadata\"")
	g.P()
	p.GenImportPb(g)
	p.GenImportSchemas(g)
	p.GenImportServices(g)
	g.P(")")
	g.P()
	g.P("var server = services.NewServer()")
	g.P()
	for _, s := range p.services {
		g.P("func ", s.FuncName, "(c *gin.Context) (int, interface{}) {")
		g.P("    in := new(pb.", s.ReqName, ")")
		g.P("    headers, err := schemas.Get", s.FuncName, "Params(c, in)")
		g.P(`    // header设置
    md := make(metadata.MD, len(headers))
    for k, v := range headers {
        md.Set(k, v...)
    }`)
		g.P("    ctx := metadata.NewIncomingContext(servers.GetRestRequestCtx(), md)")
		g.P("    res, err := server.", s.FuncName, "(ctx, in)")
		g.P("    if err != nil {")
		g.P("        panic(err)")
		g.P("    }")
		g.P(`    if res == nil {
        return 200, nil
    }`)
		if s.ReplyData == "CommonReply" {
			g.P("    return int(res.HttpCode), nil")
		} else {
			g.P("    return int(res.HttpCode), res.Data")
		}
		g.P("}")
		g.P()
	}
	g.GenFile()
}
