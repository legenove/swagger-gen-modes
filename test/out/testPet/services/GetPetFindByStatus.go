package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetPetFindByStatus(ctx context.Context, req *pb.GetPetFindByStatusRequest) (*pb.GetPetFindByStatusReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
