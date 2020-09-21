package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "errors"
)

func GetGetPetFindByTagsParams(c *gin.Context, in *pb.GetPetFindByTagsRequest) (map[string][]string, error) {
    // query Tags
	var valTags []string
	if val, ok := c.GetQueryArray("tags"); ok {
        valTags = make([]string, len(val))
        for i, v := range val {
            err := setWithKind("string", val[i], &valTags[i])
		    if err != nil {
				return nil, errors.New("tags value not []string")
			}
        }
	} else {
        return nil, errors.New("tags required")
    }
    in.Tags = valTags

    return c.Request.Header, nil
}
