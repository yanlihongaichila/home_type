package logic

import (
	"context"
	"fmt"
	"homeRpc/home"
	"homeRpc/internal/pkg"
	"homeRpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSlideshowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSlideshowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSlideshowLogic {
	return &SearchSlideshowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSlideshowLogic) SearchSlideshow(in *home.SearchSlideshowRequest) (*home.SearchSlideshowResponse, error) {
	// todo: add your logic here and delete this line

	if in.Where != 2 {
		return nil, fmt.Errorf("传入参数有问题")
	}

	slideshows, err := pkg.SearchSlideshow()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	slideshowInfos := home.SearchSlideshowResponse{}
	sInfos := slideshowInfos.Slideshows

	for _, val := range slideshows {
		tIs := home.Slideshow{
			Title: val.Name,
			Image: val.Img,
			Url:   val.Url,
		}

		sInfos = append(sInfos, &tIs)
	}

	return &home.SearchSlideshowResponse{Slideshows: sInfos}, nil
}
