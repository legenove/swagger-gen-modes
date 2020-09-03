package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetStoreOrderOrderId(ctx context.Context, req *pb.GetStoreOrderOrderIdRequest) (*pb.GetStoreOrderOrderIdReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
