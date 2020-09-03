package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostStoreOrder(ctx context.Context, req *pb.PostStoreOrderRequest) (*pb.PostStoreOrderReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
