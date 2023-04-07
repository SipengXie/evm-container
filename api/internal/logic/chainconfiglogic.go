package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChainConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChainConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChainConfigLogic {
	return &ChainConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChainConfigLogic) ChainConfig(req *types.ChainConfigRequest) (resp *types.ChainConfigResponse, err error) {

	res, err := l.svcCtx.EvmRpc.ChainConfig(l.ctx, &rpc.ChainConfigRequest{})

	return &types.ChainConfigResponse{
		ChainConfig: res.ChainConfig,
	}, err
}
