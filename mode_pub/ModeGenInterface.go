package mode_pub

type ModeGenInterface interface {
	GenFile(outpath string, swaggerPub *SwaggerPub) error
	New() ModeGenInterface
}
