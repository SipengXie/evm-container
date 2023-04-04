package logic

import (
	"context"

	"evm-container/common"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/vm"

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
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	addr := common.BytesToAddress(in.Addr)

	ret, leftOverGas, err := Evm.DelegateCall(caller, addr, in.Input, in.Gas)

	return &rpc.DelegateCallResponse{
		Ret:         ret,
		LeftOverGas: leftOverGas,
	}, err
}
