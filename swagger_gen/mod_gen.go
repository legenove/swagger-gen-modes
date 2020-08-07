package swagger_gen

type ModeGenInterface interface {
	GenFile(outpath string, swaggerPub *SwaggerPub) error
}
