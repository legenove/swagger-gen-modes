package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) GetPetFindByTags(ctx context.Context, req *pb.GetPetFindByTagsRequest) (*pb.GetPetFindByTagsReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
