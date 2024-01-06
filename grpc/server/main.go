package main

import (
	"context"
	"fmt"
	"github.com/ajwlforever/goutils/grpc/server/api"
	"github.com/ajwlforever/goutils/grpc/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

// grpc的客户端
func main() {
	fmt.Println("yes")
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	// 注册服务
	api.RegisterPostServiceServer(grpcServer, &server{&service.PostService{}})

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("start gRPC listen on port ")
	}

}

type server struct {
	svr *service.PostService
}

func (s server) GetPost(ctx context.Context, req *api.PostReq) (*api.PostReply, error) {
	//TODO implement me
	return s.svr.GetPost(ctx, req)
}
