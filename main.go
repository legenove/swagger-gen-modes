package main

import (
	"fmt"
	"github.com/legenove/swagger-gen-modes/gen_modes/proto_mode"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/utils"
	"path"
	"runtime"
)

func main() {
	curpath := callerSourcePath() + "/test"
	fileName := curpath + "/testPet.yaml"
	if !utils.FileExists(fileName) {
		fmt.Println("file not exist")
	}
	gen, err := swagger_gen.NewSwaggerGenerator(curpath+"/out", fileName)
	if err != nil {
		panic(err.Error())
	}
	proto_mode.RegistMode(gen)
	err = gen.Run()
}

func callerSourcePath() string {
	_, callerPath, _, _ := runtime.Caller(1)
	return path.Dir(callerPath)
}
