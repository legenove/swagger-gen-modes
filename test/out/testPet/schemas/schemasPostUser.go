package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
)

func GetPostUserParams(c *gin.Context, in *pb.PostUserRequest) (map[string][]string, error) {
    // body Body
	var val *pb.DefinitionsUser
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Body = val

    return c.Request.Header, nil
}
