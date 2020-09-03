package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostUser(ctx context.Context, req *pb.PostUserRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
