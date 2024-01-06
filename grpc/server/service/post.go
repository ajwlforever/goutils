package service

import (
	context "context"
	"github.com/ajwlforever/goutils/grpc/server/api"
)

type PostService struct {
}

func (s *PostService) mustEmbedUnimplementedPostServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) GetPost(ctx context.Context, postReq *api.PostReq) (postReply *api.PostReply, err error) {
	// 业务逻辑
	postReply = &api.PostReply{}
	postReply.Id = postReq.Id + 50
	return
}
