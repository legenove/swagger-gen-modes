package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var putPetDecors = []grpccore.GrpcDecoratorFunc{}

func putPet(ctx context.Context, req *pb.PutPetRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

