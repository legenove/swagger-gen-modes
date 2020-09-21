/*
### DO NOT CHANGE THIS FILE
### The code is auto generated, your change will be overwritten by
### code generating.
*/
package testPet
import (
    "github.com/legenove/nano-server-sdk/gincore"
    "github.com/legenove/nano-server-sdk/grpccore"
    "google.golang.org/grpc"

    "github.com/legenove/swagger-gen-modes/test/out/testPet/hubs"
    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
    "github.com/legenove/swagger-gen-modes/test/out/testPet/services"
)

const basePath = "/legenove6/test2.0/1.0.0"

func init() {
    grpccore.RegisterToServer("testPet", func(s *grpc.Server) {
        pb.RegisterTestPetServer(s, services.NewServer())
    })
    group := gincore.GetCurrentGroup(basePath)
    group.POST("/pet", decoratorHandler(hubs.PostPet))
    group.PUT("/pet", decoratorHandler(hubs.PutPet))
    group.GET("/pet/findByStatus", decoratorHandler(hubs.GetPetFindByStatus))
    group.GET("/pet/findByTags", decoratorHandler(hubs.GetPetFindByTags))
    group.DELETE("/pet/:petId", decoratorHandler(hubs.DeletePetPetId))
    group.GET("/pet/:petId", decoratorHandler(hubs.GetPetPetId))
    group.POST("/pet/:petId", decoratorHandler(hubs.PostPetPetId))
    group.POST("/pet/:petId/uploadImage", decoratorHandler(hubs.PostPetPetIdUploadImage))
    group.GET("/store/inventory", decoratorHandler(hubs.GetStoreInventory))
    group.POST("/store/order", decoratorHandler(hubs.PostStoreOrder))
    group.DELETE("/store/order/:orderId", decoratorHandler(hubs.DeleteStoreOrderOrderId))
    group.GET("/store/order/:orderId", decoratorHandler(hubs.GetStoreOrderOrderId))
    group.POST("/user", decoratorHandler(hubs.PostUser))
    group.POST("/user/createWithArray", decoratorHandler(hubs.PostUserCreateWithArray))
    group.POST("/user/createWithList", decoratorHandler(hubs.PostUserCreateWithList))
    group.GET("/user/login", decoratorHandler(hubs.GetUserLogin))
    group.GET("/user/logout", decoratorHandler(hubs.GetUserLogout))
    group.DELETE("/user/:username", decoratorHandler(hubs.DeleteUserUsername))
    group.GET("/user/:username", decoratorHandler(hubs.GetUserUsername))
    group.PUT("/user/:username", decoratorHandler(hubs.PutUserUsername))
}
