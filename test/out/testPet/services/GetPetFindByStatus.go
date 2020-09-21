package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getPetFindByStatusDecors = []grpccore.GrpcDecoratorFunc{}

func getPetFindByStatus(ctx context.Context, req *pb.GetPetFindByStatusRequest) (*pb.GetPetFindByStatusReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

