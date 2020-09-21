package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var putUserUsernameDecors = []grpccore.GrpcDecoratorFunc{}

func putUserUsername(ctx context.Context, req *pb.PutUserUsernameRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

