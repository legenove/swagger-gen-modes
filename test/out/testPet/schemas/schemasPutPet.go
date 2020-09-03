package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
)

func GetPutPetParams(c *gin.Context, in *pb.PutPetRequest) (map[string][]string, error) {
    // body Pet
	var val *pb.DefinitionsPet
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Pet = val

    return c.Request.Header, nil
}
