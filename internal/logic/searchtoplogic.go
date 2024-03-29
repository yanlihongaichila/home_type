package logic

import (
	"context"
	"fmt"
	"github.com/yanlihongaichila/home_type/home"
	"github.com/yanlihongaichila/home_type/internal/pkg"
	"github.com/yanlihongaichila/home_type/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTopLogic {
	return &SearchTopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTopLogic) SearchTop(in *home.SearchTopRequest) (*home.SearchTopResponse, error) {
	// todo: add your logic here and delete this line
	if in.Where != 1 {
		return nil, fmt.Errorf("传入参数有问题")
	}

	topType, err := pkg.SearchTopType()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	topInfos := home.SearchTopResponse{}
	tInfos := topInfos.TopInfos

	for _, val := range topType {
		tIs := home.TopInfo{
			Title: val.Name,
			Url:   val.Url,
		}

		tInfos = append(tInfos, &tIs)
	}
	return &home.SearchTopResponse{TopInfos: tInfos}, nil
}
