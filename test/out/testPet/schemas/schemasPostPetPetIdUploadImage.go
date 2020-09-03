package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
    "errors"
)

func GetPostPetPetIdUploadImageParams(c *gin.Context, in *pb.PostPetPetIdUploadImageRequest) (map[string][]string, error) {
    // formData AdditionalMetadata
	var valAdditionalMetadata string
	if val, ok := c.GetPostFormArray("additionalMetadata"); ok {
		_v, err := setWithKind("string", val[0])
		if err != nil {
			return nil, errors.New("additionalMetadata value not string")
		}
		valAdditionalMetadata, ok = _v.(string)
		if !ok {
			return nil, errors.New("additionalMetadata value not string")
		}
	}
    in.AdditionalMetadata = valAdditionalMetadata

    // formData File
	var valFile []byte
	if val, ok := c.GetPostFormArray("file"); ok {
		_v, err := setWithKind("[]byte", val[0])
		if err != nil {
			return nil, errors.New("file value not []byte")
		}
		valFile, ok = _v.([]byte)
		if !ok {
			return nil, errors.New("file value not []byte")
		}
	}
    in.File = valFile

    // path PetId
	var valPetId int64
	if val, ok := c.Params.Get("petId"); ok {
		_v, err := setWithKind("int64", val)
		if err != nil {
			return nil, errors.New("petId value not int64")
		}
		valPetId, ok = _v.(int64)
		if !ok {
			return nil, errors.New("petId value not int64")
		}
	} else {
		return nil, errors.New("petId required")
	}
    in.PetId = valPetId

    return c.Request.Header, nil
}
