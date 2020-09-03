package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetStoreInventory(ctx context.Context, req *pb.EmptyMessage) (*pb.GetStoreInventoryReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
