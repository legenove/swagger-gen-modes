package proto_mode

import (
	"github.com/legenove/spec4pb"
	"github.com/legenove/swagger-gen-modes/mode_pub"
	"github.com/legenove/utils"
	"sort"
)

type BufGenOpt struct {
	Locations string // reply definations request
	Method    int    // method
	Key       string
	G         mode_pub.BufGenInterface
}

type SortBufGenOpts []*BufGenOpt

func (s SortBufGenOpts) Less(i, j int) bool {
	if s[i].Locations != s[j].Locations {
		return s[i].Locations < s[j].Locations
	} else if s[i].Method != s[j].Method {
		return s[i].Method < s[j].Method
	}
	return s[i].Key < s[j].Key
}
func (s SortBufGenOpts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortBufGenOpts) Len() int      { return len(s) }
func (s SortBufGenOpts) MergeG(g mode_pub.BufGenInterface) {
	sort.Sort(s)
	for _, j := range s {
		if j != nil && len(j.G.GetBytes()) > 0 {
			g.P(string(j.G.GetBytes()))
		}
	}
}


type fieldOpt struct {
	FieldNumber int32
	FieldName   string
	Propertie  spec4pb.Schema
}

func NewFieldOpt(key string, fieldNumber int32, propertie spec4pb.Schema) *fieldOpt {
	if fieldNumber <= 0 {
		fieldNumber = utils.MaxInt32
	}
	return &fieldOpt{
		FieldNumber: fieldNumber,
		FieldName:   key,
		Propertie: propertie,
	}

}

type SortFieldOpts []*fieldOpt

func (s SortFieldOpts) Less(i, j int) bool {
	if s[i].FieldNumber != s[j].FieldNumber {
		return s[i].FieldNumber < s[j].FieldNumber
	}
	return s[i].FieldName < s[j].FieldName
}
func (s SortFieldOpts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortFieldOpts) Len() int      { return len(s) }
