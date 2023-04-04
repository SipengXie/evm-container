package logic

import (
	"context"
	"strconv"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/common"
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
	gas, _ := strconv.ParseUint(req.Gas, 10, 64)
	res, err := l.svcCtx.EvmRpc.Create(l.ctx, &rpc.CreateRequest{
		Caller: common.Hex2Bytes(req.Caller),
		Code:   common.Hex2Bytes(req.Code),
		Gas:    gas,
		Value:  req.Value,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Ret:          res.Ret,
		ContractAddr: res.ContractAddr,
		LeftOverGas:  res.LeftOverGas,
	}, nil
}
