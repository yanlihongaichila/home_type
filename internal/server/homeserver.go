// Code generated by goctl. DO NOT EDIT.
// Source: home.proto

package server

import (
	"context"
	"github.com/yanlihongaichila/home_type/home"
	"github.com/yanlihongaichila/home_type/internal/logic"
	"github.com/yanlihongaichila/home_type/internal/svc"
)

type HomeServer struct {
	svcCtx *svc.ServiceContext
	home.UnimplementedHomeServer
}

func NewHomeServer(svcCtx *svc.ServiceContext) *HomeServer {
	return &HomeServer{
		svcCtx: svcCtx,
	}
}

func (s *HomeServer) Ping(ctx context.Context, in *home.Request) (*home.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *HomeServer) SearchTop(ctx context.Context, in *home.SearchTopRequest) (*home.SearchTopResponse, error) {
	l := logic.NewSearchTopLogic(ctx, s.svcCtx)
	return l.SearchTop(in)
}

func (s *HomeServer) SearchSlideshow(ctx context.Context, in *home.SearchSlideshowRequest) (*home.SearchSlideshowResponse, error) {
	l := logic.NewSearchSlideshowLogic(ctx, s.svcCtx)
	return l.SearchSlideshow(in)
}

func (s *HomeServer) SearchMainBody(ctx context.Context, in *home.SearchMainBodyRequest) (*home.SearchMainBodyResponse, error) {
	l := logic.NewSearchMainBodyLogic(ctx, s.svcCtx)
	return l.SearchMainBody(in)
}

func (s *HomeServer) SearchBottom(ctx context.Context, in *home.SearchBottomRequest) (*home.SearchBottomResponse, error) {
	l := logic.NewSearchBottomLogic(ctx, s.svcCtx)
	return l.SearchBottom(in)
}
