package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetUserLogin(ctx context.Context, req *pb.GetUserLoginRequest) (*pb.GetUserLoginReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
