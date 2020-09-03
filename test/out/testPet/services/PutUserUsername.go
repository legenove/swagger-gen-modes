package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PutUserUsername(ctx context.Context, req *pb.PutUserUsernameRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
