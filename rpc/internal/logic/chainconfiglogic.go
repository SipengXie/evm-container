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
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}

	chainCfg := Evm.ChainConfig()
	data, err := json.Marshal(chainCfg)

	return &rpc.ChainConfigResponse{
		ChainConfig: data,
	}, err
}
