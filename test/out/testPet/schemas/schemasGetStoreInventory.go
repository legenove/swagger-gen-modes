package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)

func GetGetStoreInventoryParams(c *gin.Context, in *pb.EmptyMessage) (map[string][]string, error) {
    return c.Request.Header, nil
}
