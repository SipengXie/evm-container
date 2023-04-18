package logic

import (
	"context"
	super "evm-container"
	"evm-container/config"
	"evm-container/rpc/internal/svc"
	"evm-container/rpc/types/rpc"
	"evm-container/state"
	"evm-container/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

var Evm *vm.EVM = nil
var StateDB *state.StateDB = nil

type NewEnvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNewEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewEnvLogic {
	return &NewEnvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NewEnvLogic) NewEnv(in *rpc.NewEnvRequest) (*rpc.NewEnvResponse, error) {

	cfg, err := config.NewConfig(in.Config)
	if err != nil {
		return &rpc.NewEnvResponse{
			Code: err.Error(),
		}, err
	}

	StateDB = state.NewStateDB(l.svcCtx.SdbRpc, l.ctx)
	cfg.State = StateDB
	Evm = super.NewEnv(cfg)
	if Evm == nil {
		return nil, ErrMissingEvmInstance
	}
	return &rpc.NewEnvResponse{
		Code: "success",
	}, nil
}
