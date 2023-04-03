package logic

import (
	"context"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.EvmRpc.Create(l.ctx, &rpc.CreateRequest{
		Caller: req.Caller,
		Code:   req.Code,
		Gas:    req.Gas,
		Value:  req.Value,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Ret:          res.Ret,
		ContractAddr: res.ContractAddr,
		LeftOverGas:  res.LeftOverGas,
		Error:        res.Error,
	}, nil
}
