package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostUserCreateWithList(ctx context.Context, req *pb.PostUserCreateWithListRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
