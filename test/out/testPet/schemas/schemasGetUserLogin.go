package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "errors"
)

func GetGetUserLoginParams(c *gin.Context, in *pb.GetUserLoginRequest) (map[string][]string, error) {
    // query Password
	var valPassword string
	if val, ok := c.GetQueryArray("password"); ok {
		_v, err := setWithKind("string", val[0])
		if err != nil {
			return nil, errors.New("password value not string")
		}
		valPassword, ok = _v.(string)
		if !ok {
			return nil, errors.New("password value not string")
		}
	} else {
        return nil, errors.New("password required")
    }
    in.Password = valPassword

    // query Username
	var valUsername string
	if val, ok := c.GetQueryArray("username"); ok {
		_v, err := setWithKind("string", val[0])
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
