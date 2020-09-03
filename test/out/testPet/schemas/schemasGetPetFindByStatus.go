package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
    "errors"
)

func GetGetPetFindByStatusParams(c *gin.Context, in *pb.GetPetFindByStatusRequest) (map[string][]string, error) {
    // query Status
	var valStatus []string
	if val, ok := c.GetQueryArray("status"); ok {
        valStatus = make([]string, len(val))
        for i, v := range val {
            err := setWithKind("string", val[i], &valStatus[i])
		    if err != nil {
				return nil, errors.New("status value not []string")
			}
        }
	} else {
        return nil, errors.New("status required")
    }
    in.Status = valStatus

    return c.Request.Header, nil
}
