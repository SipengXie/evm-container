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

var Evm *vm.EVM = nil

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

	Evm = super.NewEnv(cfg)
	return &rpc.NewEnvResponse{
		Code: "success",
	}, nil
}
