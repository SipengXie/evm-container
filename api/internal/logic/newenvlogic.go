package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewEnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewEnvLogic {
	return &NewEnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewEnvLogic) NewEnv(req *types.NewEnvRequest) (resp *types.NewEnvResponse, err error) {
	res, err := l.svcCtx.EvmRpc.NewEnv(l.ctx, &rpc.NewEnvRequest{
		Config: []byte(req.Config),
	})

	return &types.NewEnvResponse{
		Code: res.Code,
	}, err
}
