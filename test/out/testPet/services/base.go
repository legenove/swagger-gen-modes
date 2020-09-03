package services

import (
    pb "/testPet/pb"
)

//server is used to implement testPet.TestPetServer.
type testPetServer struct {
  pb.UnimplementedTestPetServer
}
func NewServer() *testPetServer{
    return &testPetServer{}
}
