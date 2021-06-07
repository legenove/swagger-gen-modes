package mode_pub

import (
    "bytes"
    "fmt"
    "os"
    "path"
    "path/filepath"

    "github.com/legenove/utils"
)

type BufGenInterface interface {
    P(v ...interface{})
    Pl(v ...interface{})
    GetBytes() []byte
}

type BufGenerator struct {
    buf bytes.Buffer
}

// P prints a line to the generated output. It converts each parameter to a
// string following the same rules as fmt.Print. It never inserts spaces
// between parameters.
func (g *BufGenerator) P(v ...interface{}) {
    g.Pl(v...)
    fmt.Fprintln(&g.buf)
}

func (g *BufGenerator) Pl(v ...interface{}) {
    for _, x := range v {
        switch x := x.(type) {
        //case GoIdent:
        //	fmt.Fprint(&g.buf, g.QualifiedGoIdent(x))
        default:
            fmt.Fprint(&g.buf, x)
        }
    }
}

func (g *BufGenerator) GetBytes() []byte {
    return g.buf.Bytes()
}

type FileGenerator struct {
    BufGenerator
    skip     bool
    outPath  string
    filename string
    params   map[string]interface{}
    md5      string
}

func NewFileGen(outPath string, md5 string) *FileGenerator {
    return &FileGenerator{
        outPath: outPath,
        params:  map[string]interface{}{},
        md5:     md5,
    }
}

func (g *FileGenerator) SetFilename(fileName string) {
    g.filename = fileName
}

func (g *FileGenerator) AddParam(key string, value interface{}) {
    g.params[key] = value
}

func (g *FileGenerator) GetFullPath() string {
    return path.Join(g.outPath, g.filename)
}

func (g *FileGenerator) GenTmplFile() error {
    // todo use buf as tmpl gen file
    return nil
}

func (g *FileGenerator) GenFile() error {
    paths, _ := filepath.Split(g.GetFullPath())
    if !utils.PathExists(paths) {
        utils.CreateDir(paths)
    }
    if utils.FileExists(g.GetFullPath()) && g.skip {
        return nil
    }
    err := g.WriteFile(g.GetFullPath(), g.buf.Bytes())
    return err
}

func (g *FileGenerator) WriteFile(pth string, b []byte) error {
    f, err := os.Create(pth)
    fmt.Println("|", g.md5, ": create file: ", pth)
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = f.Write(b)
    return err
}

// Write implements io.Writer.
func (g *FileGenerator) Write(p []byte) (n int, err error) {
    return g.buf.Write(p)
}

// Skip removes the generated file from the plugin output.
func (g *FileGenerator) Skip() {
    g.skip = true
}

// Unskip reverts a previous call to Skip, re-including the generated file in
// the plugin output.
func (g *FileGenerator) Unskip() {
    g.skip = false
}
