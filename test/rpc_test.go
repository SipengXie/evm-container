package test

import (
	"context"
	"encoding/json"
	"evm-container/common"
	"evm-container/config"
	"evm-container/rpc/rpcclient"
	"evm-container/rpc/types/rpc"
	"math"
	"math/big"
	"testing"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestRpcNewEnv(t *testing.T) {

	rpcConf := zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "evm.rpc",
		},
	}
	EvmRpc := rpcclient.NewRpc(zrpc.MustNewClient(rpcConf))

	cfg := new(config.Config)
	config.SetDefaults(cfg)

	cfg_bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res, err := EvmRpc.NewEnv(context.Background(), &rpc.NewEnvRequest{
		Config: cfg_bytes,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.Code)
}

func TestRpcCreate(t *testing.T) {

	rpcConf := zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "evm.rpc",
		},
	}
	EvmRpc := rpcclient.NewRpc(zrpc.MustNewClient(rpcConf))

	caller := common.BytesToAddress([]byte("deployer"))
	code := "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033"
	value := new(big.Int).String()

	res, err := EvmRpc.Create(context.Background(), &rpc.CreateRequest{
		Caller: caller.Bytes(),
		Code:   common.Hex2Bytes(code),
		Gas:    math.MaxUint64,
		Value:  value,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("contract_address:", res.ContractAddr)
	t.Log("left Gas:", res.LeftOverGas)
	t.Log("returned code", res.Ret)
}

func TestRpcCreateAndCall(t *testing.T) {

	rpcConf := zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "evm.rpc",
		},
	}
	EvmRpc := rpcclient.NewRpc(zrpc.MustNewClient(rpcConf))

	caller := common.BytesToAddress([]byte("deployer"))
	code := "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033"
	value := new(big.Int).String()

	res, err := EvmRpc.Create(context.Background(), &rpc.CreateRequest{
		Caller: caller.Bytes(),
		Code:   common.Hex2Bytes(code),
		Gas:    math.MaxUint64,
		Value:  value,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("contract_address:", res.ContractAddr)
	t.Log("left Gas:", res.LeftOverGas)
	t.Log("returned code", res.Ret)

	TestInput := "771602f700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"
	newres, err := EvmRpc.Call(context.Background(), &rpc.CallRequest{
		Caller: caller.Bytes(),
		Addr:   res.ContractAddr,
		Input:  common.Hex2Bytes(TestInput),
		Gas:    math.MaxUint64,
		Value:  value,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("return:", newres.Ret)
	t.Log("left gas:", newres.LeftOverGas)

}
