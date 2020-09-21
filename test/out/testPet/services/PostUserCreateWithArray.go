package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postUserCreateWithArrayDecors = []grpccore.GrpcDecoratorFunc{}

func postUserCreateWithArray(ctx context.Context, req *pb.PostUserCreateWithArrayRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

