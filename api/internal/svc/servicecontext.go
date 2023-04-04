package svc

import (
	"evm-container/api/internal/config"
	"evm-container/rpc/rpcclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	EvmRpc rpcclient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		EvmRpc: rpcclient.NewRpc(zrpc.MustNewClient(c.EvmRpc)),
	}

}
