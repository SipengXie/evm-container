package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelledLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelledLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelledLogic {
	return &CancelledLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelledLogic) Cancelled(req *types.CancelledRequest) (resp *types.CancelledResponse, err error) {

	res, err := l.svcCtx.EvmRpc.Cancelled(l.ctx, &rpc.CancelledRequest{})

	return &types.CancelledResponse{
		Result: res.Result,
	}, err
}
