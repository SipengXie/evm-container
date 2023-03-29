package logic

import (
	"context"

	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelegateCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelegateCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelegateCallLogic {
	return &DelegateCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelegateCallLogic) DelegateCall(in *rpc.DelegateCallRequest) (*rpc.DelegateCallResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.DelegateCallResponse{}, nil
}
