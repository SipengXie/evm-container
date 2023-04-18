package svc

import (
	"evm-container/rpc/internal/config"
	"evm-container/state/rpc/sdbclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SdbRpc sdbclient.Sdb
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SdbRpc: sdbclient.NewSdb(zrpc.MustNewClient(c.SdbRpc)),
	}
}
