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

type CallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallLogic {
	return &CallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallLogic) Call(req *types.CallRequest) (resp *types.CallResponse, err error) {
	//	logx.Info("Address is:", common.HexToAddress(req.Addr))
	gas, _ := strconv.ParseUint(req.Gas, 10, 64)
	res, err := l.svcCtx.EvmRpc.Call(l.ctx, &rpc.CallRequest{
		Caller: common.Hex2Bytes(req.Caller),
		Addr:   common.Hex2Bytes(req.Addr),
		Input:  common.Hex2Bytes(req.Input),
		Gas:    gas,
		Value:  req.Value,
	})

	return &types.CallResponse{
		Ret:         res.Ret,
		LeftOverGas: res.LeftOverGas,
	}, err
}
