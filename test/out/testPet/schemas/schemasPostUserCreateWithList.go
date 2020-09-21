package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)

func GetPostUserCreateWithListParams(c *gin.Context, in *pb.PostUserCreateWithListRequest) (map[string][]string, error) {
    // body Body
	var val interface{}
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Body = val

    return c.Request.Header, nil
}
