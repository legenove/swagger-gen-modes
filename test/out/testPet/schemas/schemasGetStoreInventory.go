package schemas

import (
    "github.com/gin-gonic/gin"
    pb "/testPet/pb"
)

func GetGetStoreInventoryParams(c *gin.Context, in *pb.EmptyMessage) (map[string][]string, error) {
    return c.Request.Header, nil
}
