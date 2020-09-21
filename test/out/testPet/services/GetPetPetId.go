package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getPetPetIdDecors = []grpccore.GrpcDecoratorFunc{}

func getPetPetId(ctx context.Context, req *pb.GetPetPetIdRequest) (*pb.GetPetPetIdReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

