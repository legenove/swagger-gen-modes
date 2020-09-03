package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) DeleteUserUsername(ctx context.Context, req *pb.DeleteUserUsernameRequest) (*pb.DeleteUserUsernameReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
