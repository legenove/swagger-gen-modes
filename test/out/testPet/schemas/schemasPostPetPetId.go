package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "errors"
)

func GetPostPetPetIdParams(c *gin.Context, in *pb.PostPetPetIdRequest) (map[string][]string, error) {
    // formData Name
	var valName string
	if val, ok := c.GetPostFormArray("name"); ok {
		_v, err := setWithKind("string", val[0])
		if err != nil {
			return nil, errors.New("name value not string")
		}
		valName, ok = _v.(string)
		if !ok {
			return nil, errors.New("name value not string")
		}
	}
    in.Name = valName

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

    // formData Status
	var valStatus string
	if val, ok := c.GetPostFormArray("status"); ok {
		_v, err := setWithKind("string", val[0])
		if err != nil {
			return nil, errors.New("status value not string")
		}
		valStatus, ok = _v.(string)
		if !ok {
			return nil, errors.New("status value not string")
		}
	}
    in.Status = valStatus

    return c.Request.Header, nil
}
