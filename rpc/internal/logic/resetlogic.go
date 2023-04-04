package logic

import (
	"context"

	"evm-container/common"
	"evm-container/config"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLogic {
	return &ResetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetLogic) Reset(in *rpc.ResetRequest) (*rpc.ResetResponse, error) {
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	if StateDB == nil {
		return nil, ErrMissingStateDBInstance
	}

	txHash := common.BytesToHash(in.TxHash)
	index := in.Index
	StateDB.SetTxContext(txHash, index)

	txContext, err := config.NewTxContext(in.TxCtx)
	if err != nil {
		return &rpc.ResetResponse{
			Code: err.Error(),
		}, err
	}

	Evm.Reset(vm.TxContext(*txContext), StateDB)
	return &rpc.ResetResponse{
		Code: "success",
	}, nil
}
