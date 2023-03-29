package logic

import (
	"context"

	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelledLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelledLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelledLogic {
	return &CancelledLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelledLogic) Cancelled(in *rpc.CancelledRequest) (*rpc.CancelledResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.CancelledResponse{}, nil
}
