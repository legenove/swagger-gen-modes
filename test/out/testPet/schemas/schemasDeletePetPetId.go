package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
    "errors"
)

func GetDeletePetPetIdParams(c *gin.Context, in *pb.DeletePetPetIdRequest) (map[string][]string, error) {
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
