package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) DeleteStoreOrderOrderId(ctx context.Context, req *pb.DeleteStoreOrderOrderIdRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
