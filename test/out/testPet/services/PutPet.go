package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PutPet(ctx context.Context, req *pb.PutPetRequest) (*pb.CommonReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
