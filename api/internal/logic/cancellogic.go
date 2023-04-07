package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req *types.CancelRequest) (resp *types.CancelResponse, err error) {

	res, err := l.svcCtx.EvmRpc.Cancel(l.ctx, &rpc.CancelRequset{})

	return &types.CancelResponse{
		Code: res.Code,
	}, err
}
