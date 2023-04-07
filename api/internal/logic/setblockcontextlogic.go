package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBlockContextLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetBlockContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBlockContextLogic {
	return &SetBlockContextLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetBlockContextLogic) SetBlockContext(req *types.SetBlockContextRequest) (resp *types.SetBlockContextResponse, err error) {

	res, err := l.svcCtx.EvmRpc.SetBlockContext(l.ctx, &rpc.SetBlockContextRequest{
		BlockCtx: []byte(req.BlockCtx),
	})

	return &types.SetBlockContextResponse{
		Code: res.Code,
	}, err
}
