package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var deletePetPetIdDecors = []grpccore.GrpcDecoratorFunc{}

func deletePetPetId(ctx context.Context, req *pb.DeletePetPetIdRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

