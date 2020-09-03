package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostUserCreateWithArray(ctx context.Context, req *pb.PostUserCreateWithArrayRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
