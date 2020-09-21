package schemas

import (
    "github.com/gin-gonic/gin"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "errors"
)

func GetDeleteStoreOrderOrderIdParams(c *gin.Context, in *pb.DeleteStoreOrderOrderIdRequest) (map[string][]string, error) {
    // path OrderId
	var valOrderId uint64
	if val, ok := c.Params.Get("orderId"); ok {
		_v, err := setWithKind("uint64", val)
		if err != nil {
			return nil, errors.New("orderId value not uint64")
		}
		valOrderId, ok = _v.(uint64)
		if !ok {
			return nil, errors.New("orderId value not uint64")
		}
	} else {
		return nil, errors.New("orderId required")
	}
    if !(in.OrderId >= 1) {
	    return nil, errors.New("OrderId len must >= 1")
    }
    in.OrderId = valOrderId

    return c.Request.Header, nil
}
