package logic

import (
	"context"
	"fmt"
	"homeRpc/internal/pkg"

	"homeRpc/home"
	"homeRpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBottomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchBottomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBottomLogic {
	return &SearchBottomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchBottomLogic) SearchBottom(in *home.SearchBottomRequest) (*home.SearchBottomResponse, error) {
	// todo: add your logic here and delete this line
	if in.Where != 4 {
		return nil, fmt.Errorf("传入参数有问题")
	}

	Bottoms, err := pkg.SearchBottom()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	BottomInfos := home.SearchBottomResponse{}
	mBInfos := BottomInfos.Bottoms

	for _, val := range Bottoms {
		mBIs := home.Bottom{
			Title: val.Name,
			Url:   val.Url,
			Image: val.Image,
		}

		mBInfos = append(mBInfos, &mBIs)
	}

	return &home.SearchBottomResponse{Bottoms: mBInfos}, nil
}
