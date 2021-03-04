package common

const DefinitionPreName = "Definitions"

var OptLocationMap = map[string]int{
	"Service":     0,
	"Request":     1,
	"Reply":       2,
	"Definitions": 3,
}
var OptReLocationMap = map[string]string{
	"0": "Service",
	"1": "Request",
	"2": "Reply",
	"3": "Definitions",
}

var OptMethodMap = map[string]int{
	"Get":     0,
	"Post":    1,
	"Put":     2,
	"Delete":  4,
	"Options": 5,
	"Head":    6,
	"Patch":   7,
}
