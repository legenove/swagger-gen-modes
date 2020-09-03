package proto_mode

import (
	"github.com/legenove/swagger-gen-modes/mode_pub"
)

func (p *ProtoMode) prepareMessages() {
	p.prepareDefinitions(p.swaggerPub.Swagger.Definitions)
	p.prepareSecurityDefinitions(p.swaggerPub.Swagger.SecurityDefinitions)
}

func (p *ProtoMode) genMessages(g mode_pub.BufGenInterface) {
	p.messageGenOpt.MergeG(g)
	if p.hasReply {
		g.P("// Common Reply")
		g.P("message CommonReply {")
		g.P("  uint32 httpCode=1;")
		g.P("  uint32 errorCode=2;")
		g.P("  string errorMsg=3;")
		g.P("}")
		g.P()
	}
	if p.hasEmpty {
		g.P("// Empty message")
		g.P("message EmptyMessage {}")
		g.P()
	}
}



func (p *ProtoMode) genEmpty(g mode_pub.BufGenInterface) {
	p.hasEmpty = true
	g.Pl("EmptyMessage")
}



func (p *ProtoMode) genReply(g mode_pub.BufGenInterface)  {
	p.hasReply = true
	g.Pl("CommonReply")
}
