package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postPetPetIdUploadImageDecors = []grpccore.GrpcDecoratorFunc{}

func postPetPetIdUploadImage(ctx context.Context, req *pb.PostPetPetIdUploadImageRequest) (*pb.PostPetPetIdUploadImageReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

