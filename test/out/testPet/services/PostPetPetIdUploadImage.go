package services

import (
    "context"
    pb "/testPet/pb"
    "fmt"
)

func (*testPetServer) PostPetPetIdUploadImage(ctx context.Context, req *pb.PostPetPetIdUploadImageRequest) (*pb.PostPetPetIdUploadImageReply, error) {
    fmt.Println("in", req)
    return nil, nil
}
