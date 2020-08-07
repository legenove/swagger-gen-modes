package proto_mode

import "github.com/legenove/swagger-gen-modes/swagger_gen"

const AnyImport = "import \"google/protobuf/any.proto\";"
const AnyProto = "google.protobuf.Any"

func (p *ProtoMode) addImport(s string) {
	p.imports[s] = true
}
func (p *ProtoMode) genImport(g swagger_gen.BufGenInterface) {
	if p.imports != nil && len(p.imports) > 0 {
		for s := range p.imports {
			g.P(s)
		}
	}
	g.P()
}

func (p *ProtoMode) getAnyProto() string {
	p.addImport(AnyImport)
	return AnyProto
}
