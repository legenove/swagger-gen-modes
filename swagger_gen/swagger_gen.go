package swagger_gen

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/legenove/spec4pb"
	"github.com/legenove/utils"
	"path/filepath"
	"strings"
	"sync"
)

// swagger 生成器
type SwaggerGenerator struct {
	sync.RWMutex
	outPath        string                      // 输出路径
	goPackageName  string                      // go package 包名
	SourceFilePath []string                    // 原始swagger文件地址
	swaggers       map[string]*spec4pb.Swagger // swagger对象
	genMode        map[string]ModeGenInterface // 需要生成模版的方法
	Errors         []error
}

// 初始化swagger 生成器
func NewSwaggerGenerator(outPath string, filePath ...string) (*SwaggerGenerator, error) {
	o := new(SwaggerGenerator)
	o.outPath = outPath
	o.SourceFilePath = make([]string, 0, 16)
	o.swaggers = make(map[string]*spec4pb.Swagger, 16)
	o.genMode = make(map[string]ModeGenInterface, 16)
	err := o.LoadSource(filePath...)
	return o, err
}

func (s *SwaggerGenerator) SetGoPackage(packagePath string) *SwaggerGenerator {
	s.goPackageName = packagePath
	return s
}

func (s *SwaggerGenerator) LoadSource(filePath ...string) error {
	if len(filePath) == 0 {
		return nil
	}
	for _, fpath := range filePath {
		if _, ok := s.swaggers[fpath]; ok {
			continue
		}
		swagger, err := s.LoadFileToSwagger(fpath)
		if err != nil {
			return errors.New(utils.ConcatenateStrings(fpath, " load error :", err.Error()))
		}
		s.Lock()
		s.SourceFilePath = append(s.SourceFilePath, fpath)
		s.swaggers[fpath] = swagger
		s.Unlock()
	}
	return nil
}

func (s *SwaggerGenerator) AddMode(mode string, i ModeGenInterface) error {
	s.Lock()
	if _, ok := s.genMode[mode]; ok {
		return errors.New("mode already register")
	}
	s.genMode[mode] = i
	s.Unlock()
	return nil
}

func (s *SwaggerGenerator) GetError() error {
	if len(s.Errors) != 0 {
		errorStrs := make([]string, len(s.Errors)*2)
		for i, e := range s.Errors {
			errorStrs[i*2] = e.Error()
			errorStrs[i*2+1] = " ;\n"
		}
		return errors.New(utils.ConcatenateStrings(errorStrs...))
	}
	return nil
}

func (s *SwaggerGenerator) AddError(err error) {
	s.Lock()
	defer s.Unlock()
	s.Errors = append(s.Errors, err)
}

func (s *SwaggerGenerator) Run() error {
	s.Errors = make([]error, 0, 16)
	ch := make(chan struct{}, 5)
	wg := sync.WaitGroup{}
	for mod, modFunc := range s.genMode {
		for fpath, swagger := range s.swaggers {
			wg.Add(1)
			ch <- struct{}{}
			go func(mod, fpath string, modI ModeGenInterface, swagger *spec4pb.Swagger) {
				key := utils.ConcatenateStrings("GenMode: ", mod, "\nsource path: ", fpath)
				m := utils.GetMD5Hash(key)
				fmt.Println("|", m, ": start gen : "+key)
				var err error
				defer func() {
					//if err := recover(); err != nil {
					//	switch err.(type) {
					//	case error:
					//		s.AddError(err.(error))
					//	default:
					//	}
					//}
					fmt.Println("|", m, ": end gen : "+key)
					wg.Done()
					<-ch
				}()
				fname := filepath.Base(fpath)
				pub := &SwaggerPub{
					Swagger:       swagger,
					PackageName:   strings.Split(fname, ".")[0],
					Md5:           m,
					GoPackageName: s.goPackageName,
				}
				err = modI.GenFile(s.outPath, pub)
				if err != nil {
					s.AddError(err)
				}
			}(mod, fpath, modFunc, swagger)
		}
	}
	wg.Wait()
	return s.GetError()
}

func (s *SwaggerGenerator) LoadFileToSwagger(filePath string) (*spec4pb.Swagger, error) {
	ext := filepath.Ext(filePath)
	var b []byte
	var err error
	b, err = swag.LoadFromFileOrHTTP(filePath)
	if err != nil {
		return nil, err
	}
	var jb json.RawMessage
	switch ext {
	case ".yaml", "yml":
		var yb interface{}
		yb, err = swag.BytesToYAMLDoc(b)
		if err != nil {
			return nil, err
		}
		jb, err = swag.YAMLToJSON(yb)
		if err != nil {
			return nil, err
		}
	case ".json":
		jb = b
	default:
		return nil, errors.New("only support json and yaml")
	}
	swagger := new(spec4pb.Swagger)
	err = swagger.UnmarshalJSON(jb)
	return swagger, err
}
