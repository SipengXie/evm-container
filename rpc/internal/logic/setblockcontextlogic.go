package logic

import (
	"context"

	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBlockContextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetBlockContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBlockContextLogic {
	return &SetBlockContextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetBlockContextLogic) SetBlockContext(in *rpc.SetBlockContextRequest) (*rpc.SetBlockContextResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.SetBlockContextResponse{}, nil
}
