package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getUserUsernameDecors = []grpccore.GrpcDecoratorFunc{}

func getUserUsername(ctx context.Context, req *pb.GetUserUsernameRequest) (*pb.GetUserUsernameReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

