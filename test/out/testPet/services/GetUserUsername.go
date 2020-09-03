package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetUserUsername(ctx context.Context, req *pb.GetUserUsernameRequest) (*pb.GetUserUsernameReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
