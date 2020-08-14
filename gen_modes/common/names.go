package common

import (
	"github.com/legenove/spec4pb"
	"strings"
)

func GetGoName(n string) string {
	ss := strings.Split(n, "_")
	res := ""
	for _, s := range ss {
		res += strings.Title(s)
	}
	return res
}

func GetSchemaName(schema *spec4pb.Schema) string {
	return FormatRefUrl(schema.Ref.GetURL().String())
}

func FormatRefUrl(s string) string {
	ss := strings.Split(s, "/")
	res := ""
	for i := 1; i < len(ss); i++ {
		res += strings.Title(ss[i])
	}
	return res
}
