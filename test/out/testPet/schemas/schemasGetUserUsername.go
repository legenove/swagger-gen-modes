package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
    "errors"
)

func GetGetUserUsernameParams(c *gin.Context, in *pb.GetUserUsernameRequest) (map[string][]string, error) {
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
