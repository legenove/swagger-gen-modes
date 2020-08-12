package gin4grpc_mode

import (
	"fmt"
	"testing"
)

func TestRouterPath(t *testing.T) {
	in := []string{
		"/pet",
		"/pet/{petId}",
		"/pet/{petId}/uploadImage",
		"/pet/findByStatus",
		"/store/{petId}/order/{orderId}",
	}
	out := []string{
		"/pet",
		"/pet/:petId",
		"/pet/:petId/uploadImage",
		"/pet/findByStatus",
		"/store/:petId/order/:orderId",
	}

	for i := range in {
		s := &service{pathName: in[i]}
		res := s.getRouterPath()
		if res == out[i] {
			fmt.Println("true: ", i)
		} else {
			fmt.Println("false: ", res)
		}
	}
}