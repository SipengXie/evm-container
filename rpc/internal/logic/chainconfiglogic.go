package logic

import (
	"context"

	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChainConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChainConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChainConfigLogic {
	return &ChainConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChainConfigLogic) ChainConfig(in *rpc.ChainConfigRequest) (*rpc.ChainConfigResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.ChainConfigResponse{}, nil
}
