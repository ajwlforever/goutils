package main

import (
	"context"
	"fmt"
	"github.com/ajwlforever/goutils/grpc/client/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	address := "127.0.0.1:8001"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := api.NewPostServiceClient(conn)
	ctx := context.Background()

	postReply, err := client.GetPost(ctx, &api.PostReq{Id: 1})
	if err != nil {
		log.Println("err:", err)
	}
	fmt.Println("post reply: ", postReply)
}
