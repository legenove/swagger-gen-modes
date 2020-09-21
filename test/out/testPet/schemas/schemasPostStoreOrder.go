package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)

func GetPostStoreOrderParams(c *gin.Context, in *pb.PostStoreOrderRequest) (map[string][]string, error) {
    // body Body
	var val *pb.DefinitionsOrder
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Body = val

    return c.Request.Header, nil
}
