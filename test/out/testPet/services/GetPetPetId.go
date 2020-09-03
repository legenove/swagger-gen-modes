package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetPetPetId(ctx context.Context, req *pb.GetPetPetIdRequest) (*pb.GetPetPetIdReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
