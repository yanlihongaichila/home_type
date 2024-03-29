package logic

import (
	"context"
	"fmt"
	"homeRpc/internal/pkg"

	"homeRpc/home"
	"homeRpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMainBodyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMainBodyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMainBodyLogic {
	return &SearchMainBodyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMainBodyLogic) SearchMainBody(in *home.SearchMainBodyRequest) (*home.SearchMainBodyResponse, error) {
	// todo: add your logic here and delete this line
	if in.Where != 3 {
		return nil, fmt.Errorf("传入参数有问题")
	}

	mainBodys, err := pkg.SearchMainBody()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mainBodyInfos := home.SearchMainBodyResponse{}
	mBInfos := mainBodyInfos.MainBodys

	for _, val := range mainBodys {
		mBIs := home.MainBody{
			Title: val.Name,
			Url:   val.Url,
			Image: val.Image,
		}

		mBInfos = append(mBInfos, &mBIs)
	}

	return &home.SearchMainBodyResponse{MainBodys: mBInfos}, nil
}
