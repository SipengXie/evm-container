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
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	var value *big.Int
	err := value.UnmarshalJSON(in.Value)
	if err != nil {
		return &rpc.Create2Response{}, err
	}

	ret, addr, leftOverGas, err := Evm.Create(caller, in.Code, in.Gas, value)
	return &rpc.Create2Response{
		Ret:          ret,
		ContractAddr: addr.Bytes(),
		LeftOverGas:  leftOverGas,
		Error:        err.Error(),
	}, err

}
