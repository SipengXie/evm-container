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

type Create2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create2Logic {
	return &Create2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create2Logic) Create2(req *types.Create2Request) (resp *types.Create2Response, err error) {
	gas, _ := strconv.ParseUint(req.Gas, 10, 64)
	res, err := l.svcCtx.EvmRpc.Create2(l.ctx, &rpc.Create2Request{
		Caller: common.Hex2Bytes(req.Caller),
		Code:   common.Hex2Bytes(req.Code),
		Gas:    gas,
		Value:  req.Value,
	})

	return &types.Create2Response{
		Ret:          res.Ret,
		ContractAddr: res.ContractAddr,
		LeftOverGas:  res.LeftOverGas,
	}, err
}
