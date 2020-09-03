package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetUserLogout(ctx context.Context, req *pb.EmptyMessage) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
