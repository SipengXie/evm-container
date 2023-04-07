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
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}
	caller := vm.AccountRef(common.BytesToAddress(in.Caller))
	var value *big.Int = new(big.Int)
	value, _ = value.SetString(string(in.Value), 10)

	rules := Evm.ChainConfig().Rules(Evm.Context.BlockNumber, Evm.Context.Random != nil, Evm.Context.Time)
	StateDB.Prepare(rules, Evm.TxContext.Origin, Evm.Context.Coinbase, nil, vm.ActivePrecompiles(rules), nil)

	ret, addr, leftOverGas, err := Evm.Create(caller, in.Code, in.Gas, value)

	StateDB.SetCode(addr, ret)

	return &rpc.CreateResponse{
		Ret:          ret,
		ContractAddr: addr.Bytes(),
		LeftOverGas:  leftOverGas,
	}, err
}
