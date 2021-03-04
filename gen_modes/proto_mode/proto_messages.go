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
}
