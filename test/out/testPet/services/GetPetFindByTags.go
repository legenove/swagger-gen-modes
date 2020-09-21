package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getPetFindByTagsDecors = []grpccore.GrpcDecoratorFunc{}

func getPetFindByTags(ctx context.Context, req *pb.GetPetFindByTagsRequest) (*pb.GetPetFindByTagsReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

