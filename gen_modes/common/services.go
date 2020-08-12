package common

import (
	"github.com/legenove/spec4pb"
	"github.com/legenove/utils"
	"strings"
)

var Methods = []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch"}

func GetOptionsFromPathItemByMethod(pathItem spec4pb.PathItem, method string) *spec4pb.Operation {
	switch strings.ToLower(method) {
	case "get":
		return pathItem.Get
	case "post":
		return pathItem.Post
	case "put":
		return pathItem.Put
	case "delete":
		return pathItem.Delete
	case "options":
		return pathItem.Options
	case "head":
		return pathItem.Head
	case "patch":
		return pathItem.Patch
	default:
		return nil
	}
}

func UriPathToName(pth string) string {
	pth = strings.ReplaceAll(pth, "_", "/")
	pth = strings.ReplaceAll(pth, "{", "/")
	pth = strings.ReplaceAll(pth, "}", "/")
	ps := strings.Split(pth, "/")
	for i := range ps {
		ps[i] = strings.Title(ps[i])
	}
	return utils.ConcatenateStrings(ps...)
}