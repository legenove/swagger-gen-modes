package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostPetPetId(ctx context.Context, req *pb.PostPetPetIdRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
