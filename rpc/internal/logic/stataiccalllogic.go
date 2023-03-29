package logic

import (
	"context"

	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StataicCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStataicCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StataicCallLogic {
	return &StataicCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StataicCallLogic) StataicCall(in *rpc.StataicCallRequest) (*rpc.StataicCallResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.StataicCallResponse{}, nil
}
