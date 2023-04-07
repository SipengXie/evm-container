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

type DelegateCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelegateCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelegateCallLogic {
	return &DelegateCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelegateCallLogic) DelegateCall(req *types.DelegateCallRequest) (resp *types.DelegateCallResponse, err error) {
	gas, _ := strconv.ParseUint(req.Gas, 10, 64)
	res, err := l.svcCtx.EvmRpc.DelegateCall(l.ctx, &rpc.DelegateCallRequest{
		Caller: common.Hex2Bytes(req.Caller),
		Addr:   common.Hex2Bytes(req.Addr),
		Input:  common.Hex2Bytes(req.Input),
		Gas:    gas,
	})

	return &types.DelegateCallResponse{
		Ret:         res.Ret,
		LeftOverGas: res.LeftOverGas,
	}, err
}
