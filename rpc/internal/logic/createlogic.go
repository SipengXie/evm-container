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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	var value *big.Int
	err := value.UnmarshalJSON(in.Value)
	if err != nil {
		return &rpc.CreateResponse{}, err
	}

	ret, addr, leftOverGas, err := Evm.Create(caller, in.Code, in.Gas, value)
	return &rpc.CreateResponse{
		Ret:          ret,
		ContractAddr: addr.Bytes(),
		LeftOverGas:  leftOverGas,
		Error:        err.Error(),
	}, err
}
