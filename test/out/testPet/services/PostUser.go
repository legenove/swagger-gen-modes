package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postUserDecors = []grpccore.GrpcDecoratorFunc{}

func postUser(ctx context.Context, req *pb.PostUserRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

