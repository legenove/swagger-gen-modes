package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var deleteUserUsernameDecors = []grpccore.GrpcDecoratorFunc{}

func deleteUserUsername(ctx context.Context, req *pb.DeleteUserUsernameRequest) (*pb.DeleteUserUsernameReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

