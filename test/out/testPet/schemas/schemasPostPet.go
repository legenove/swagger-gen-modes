package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
)

func GetPostPetParams(c *gin.Context, in *pb.PostPetRequest) (map[string][]string, error) {
    // body Body
	var val *pb.DefinitionsPet
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Body = val

    return c.Request.Header, nil
}
