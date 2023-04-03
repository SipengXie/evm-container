package logic

import (
	"context"

	super "evm-container"
	"evm-container/config"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBlockContextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetBlockContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBlockContextLogic {
	return &SetBlockContextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetBlockContextLogic) SetBlockContext(in *rpc.SetBlockContextRequest) (*rpc.SetBlockContextResponse, error) {

	blockCtx, err := config.NewBlockContext(in.BlockCtx)
	if err != nil {
		return &rpc.SetBlockContextResponse{
			Code: err.Error(),
		}, err
	}

	blockCtx_vm := vm.BlockContext{
		CanTransfer: super.CanTransfer,
		Transfer:    super.Transfer,
		GetHash:     super.GetHashFn,
		Coinbase:    blockCtx.Coinbase,
		BlockNumber: blockCtx.BlockNumber,
		Time:        blockCtx.Time,
		Difficulty:  blockCtx.Difficulty,
		GasLimit:    blockCtx.GasLimit,
		BaseFee:     blockCtx.BaseFee,
	}

	Evm.SetBlockContext(blockCtx_vm)
	return &rpc.SetBlockContextResponse{
		Code: "success",
	}, nil
}
