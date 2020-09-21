package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postPetDecors = []grpccore.GrpcDecoratorFunc{}

func postPet(ctx context.Context, req *pb.PostPetRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

