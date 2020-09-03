package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostPet(ctx context.Context, req *pb.PostPetRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
