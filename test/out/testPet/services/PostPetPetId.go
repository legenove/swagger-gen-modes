package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postPetPetIdDecors = []grpccore.GrpcDecoratorFunc{}

func postPetPetId(ctx context.Context, req *pb.PostPetPetIdRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

