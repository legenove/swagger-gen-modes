package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "errors"
)

func GetPutUserUsernameParams(c *gin.Context, in *pb.PutUserUsernameRequest) (map[string][]string, error) {
    // body Body
	var val *pb.DefinitionsUser
	err := c.ShouldBind(&val)
	if err != nil {
		return nil, err
	}
	in.Body = val

    // path Username
	var valUsername string
	if val, ok := c.Params.Get("username"); ok {
		_v, err := setWithKind("string", val)
		if err != nil {
			return nil, errors.New("username value not string")
		}
		valUsername, ok = _v.(string)
		if !ok {
			return nil, errors.New("username value not string")
		}
	} else {
		return nil, errors.New("username required")
	}
    in.Username = valUsername

    return c.Request.Header, nil
}
