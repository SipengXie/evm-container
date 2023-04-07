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

type StaticCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStaticCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StaticCallLogic {
	return &StaticCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StaticCallLogic) StaticCall(req *types.StaticCallRequest) (resp *types.StaticCallResponse, err error) {
	gas, _ := strconv.ParseUint(req.Gas, 10, 64)
	res, err := l.svcCtx.EvmRpc.StaticCall(l.ctx, &rpc.StaticCallRequest{
		Caller: common.Hex2Bytes(req.Caller),
		Addr:   common.Hex2Bytes(req.Addr),
		Input:  common.Hex2Bytes(req.Input),
		Gas:    gas,
	})

	return &types.StaticCallResponse{
		Ret:         res.Ret,
		LeftOverGas: res.LeftOverGas,
	}, err
}
