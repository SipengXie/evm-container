package logic

import (
	"context"
	"encoding/json"

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

	chainCfg := Evm.ChainConfig()
	data, err := json.Marshal(chainCfg)
	if err != nil {
		return &rpc.ChainConfigResponse{}, err
	}

	return &rpc.ChainConfigResponse{
		ChainConfig: data,
	}, nil
}