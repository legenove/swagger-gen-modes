package main

import (
	"fmt"
	"github.com/legenove/swagger-gen-modes/gen_modes/gin4grpc_mode"
	"github.com/legenove/swagger-gen-modes/gen_modes/proto_mode"
	"github.com/legenove/swagger-gen-modes/swagger_gen"
	"github.com/legenove/utils"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)



func GetAllProtoFileByPath(_fpath string) map[string]string {
	var res = map[string]string{}
	if !utils.PathExists(_fpath) {
		return res
	}
	list, err := ioutil.ReadDir(_fpath)
	if err != nil {
		return res
	}
	for _, n := range list {
		if !n.IsDir() {
			continue
		}
		fname := filepath.Join(_fpath, n.Name(), n.Name()+".proto")
		if !utils.FileExists(fname) {
			continue
		}
		res[n.Name()] = filepath.Join(_fpath, n.Name())
	}
	return res
}

func main() {
	curpath := callerSourcePath() + "/test"
	protoPath := curpath + "/out/proto"
	fileName := curpath + "/testPet.yaml"
	if !utils.FileExists(fileName) {
		fmt.Println("file not exist")
	}
	gen, err := swagger_gen.NewSwaggerGenerator(curpath+"/out", fileName)
	gen.SetGoPackage("github.com/legenove/swagger-gen-modes/test/out")
	if err != nil {
		panic(err.Error())
	}
	gin4grpc_mode.RegistMode(gen)
	proto_mode.RegistMode(gen, protoPath)
	err = gen.Run()
	ps := GetAllProtoFileByPath(protoPath)
	for pn, pp := range ps {
		GinGrpcPb(curpath+"/out", pn, pp)
	}
}

func callerSourcePath() string {
	_, callerPath, _, _ := runtime.Caller(1)
	return path.Dir(callerPath)
}



func GinGrpcPb(curpath string, packageName string, protoPath string) {
	outPath := filepath.Join(curpath, packageName, "pb")
	if !utils.PathExists(outPath) {
		utils.CreateDir(outPath)
	}
	err := exec.Command("protoc",
		fmt.Sprintf("--proto_path=%s", protoPath),
		fmt.Sprintf("--gofast_out=plugins=grpc:%s", outPath),
		fmt.Sprintf("%s.proto", packageName),
	).Run()
	if err != nil {
		panic(err.Error())
	}
}