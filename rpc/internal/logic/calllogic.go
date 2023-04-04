package logic

import (
	"context"
	"math/big"

	"evm-container/common"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallLogic {
	return &CallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallLogic) Call(in *rpc.CallRequest) (*rpc.CallResponse, error) {

	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	addr := common.BytesToAddress(in.Addr)
	var value *big.Int
	value, _ = value.SetString(string(in.Value), 10)

	ret, leftOverGas, err := Evm.Call(caller, addr, in.Input, in.Gas, value)
	return &rpc.CallResponse{
		Ret:         ret,
		LeftOverGas: leftOverGas,
	}, err

}
