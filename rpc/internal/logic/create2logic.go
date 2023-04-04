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

type Create2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreate2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create2Logic {
	return &Create2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Create2Logic) Create2(in *rpc.Create2Request) (*rpc.Create2Response, error) {
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	var value *big.Int
	value, _ = value.SetString(string(in.Value), 10)

	ret, addr, leftOverGas, err := Evm.Create(caller, in.Code, in.Gas, value)
	return &rpc.Create2Response{
		Ret:          ret,
		ContractAddr: addr.Bytes(),
		LeftOverGas:  leftOverGas,
	}, err

}
