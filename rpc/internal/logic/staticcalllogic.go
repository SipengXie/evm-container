package logic

import (
	"context"

	"evm-container/common"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

type StaticCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStaticCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StaticCallLogic {
	return &StaticCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StaticCallLogic) StaticCall(in *rpc.StaticCallRequest) (*rpc.StaticCallResponse, error) {
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	addr := common.BytesToAddress(in.Addr)

	ret, leftOverGas, err := Evm.StaticCall(caller, addr, in.Input, in.Gas)

	return &rpc.StaticCallResponse{
		Ret:         ret,
		LeftOverGas: leftOverGas,
	}, err

}
