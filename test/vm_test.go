package test

import (
	"evm-container/common"
	"evm-container/config"
	"evm-container/vm"
	"math/big"
	"testing"

	"evm-container/state"
)

func TestDefaults(t *testing.T) {
	cfg := new(config.Config)
	config.SetDefaults(cfg)

	if cfg.ChainConfig == nil {
		t.Fatal("incorrect ChainConfig")
	}

	if cfg.BlockCtx == nil {
		t.Fatal("incorrect BlockContext")
	}

	if cfg.TxCtx == nil {
		t.Fatal("incorrect TransactionContext")
	}
}

func TestEVM(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("crashed with: %v", r)
		}
	}()

	Execute([]byte{
		byte(vm.DIFFICULTY),
		byte(vm.TIMESTAMP),
		byte(vm.GASLIMIT),
		byte(vm.PUSH1),
		byte(vm.ORIGIN),
		byte(vm.BLOCKHASH),
		byte(vm.COINBASE),
	}, nil, nil)
}

func TestCall(t *testing.T) {
	state := state.NewStateDB(nil, nil)
	address := common.HexToAddress("0x0a")
	state.SetCode(address, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})

	ret, _, err := Call(address, nil, &config.Config{State: state})
	if err != nil {
		t.Fatal("didn't expect error", err)
	}

	num := new(big.Int).SetBytes(ret)
	if num.Cmp(big.NewInt(10)) != 0 {
		t.Error("Expected 10, got", num)
	}
}

func TestCreate(t *testing.T) {
	state := state.NewStateDB(nil, nil)
	HexTestCode := common.Hex2Bytes("608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033")

	deployer := common.BytesToAddress([]byte("deployer"))
	state.CreateAccount(deployer)
	runtimeCfg := config.Config{
		TxCtx: &config.TransactionContext{
			Origin: deployer,
		},
		State: state,
	}

	code, address, leftGas, err := Create(HexTestCode, &runtimeCfg)
	if err != nil {
		t.Fatal("didn't expect error", err)
	}
	t.Logf("\n\n code:%v \n\naddress:%v\n\n leftGas:%v", code, address, leftGas)
}

func TestCreateAndCall(t *testing.T) {
	state := state.NewStateDB(nil, nil)
	// keccak256
	HexTestCode := common.Hex2Bytes("608060405234801561001057600080fd5b50610211806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063fabeda8c14610030575b600080fd5b61004a600480360381019061004591906100e9565b610060565b6040516100579190610183565b60405180910390f35b6000828260405160200161007592919061016a565b60405160208183030381529060405280519060200120905092915050565b60008083601f8401126100a9576100a86101c7565b5b8235905067ffffffffffffffff8111156100c6576100c56101c2565b5b6020830191508360018202830111156100e2576100e16101cc565b5b9250929050565b60008060208385031215610100576100ff6101d6565b5b600083013567ffffffffffffffff81111561011e5761011d6101d1565b5b61012a85828601610093565b92509250509250929050565b61013f816101a9565b82525050565b6000610151838561019e565b935061015e8385846101b3565b82840190509392505050565b6000610177828486610145565b91508190509392505050565b60006020820190506101986000830184610136565b92915050565b600081905092915050565b6000819050919050565b82818337600083830152505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fdfea264697066735822122056ba8d7c8efa70f3baf94097841dcd99dac6e425cc1d75940e0fa8c16549b0a964736f6c63430008070033")
	TestInput := common.Hex2Bytes("fabeda8c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000033132330000000000000000000000000000000000000000000000000000000000")
	deployer := common.BytesToAddress([]byte("deployer"))
	caller := common.BytesToAddress([]byte("caller"))
	state.CreateAccount(deployer)
	state.CreateAccount(caller)
	state.AddBalance(deployer, big.NewInt(1<<40))
	state.AddBalance(caller, big.NewInt(1<<40))

	runtimeCfg := config.Config{
		TxCtx: &config.TransactionContext{
			Origin: deployer,
		},
		State: state,
	}

	code, ContractAddress, _, err := Create(HexTestCode, &runtimeCfg)
	if err != nil {
		t.Fatal("\ndidn't expect error", err)
	}
	state.SetCode(ContractAddress, code)

	t.Logf("\nCode:%v", code)

	runtimeCfg = config.Config{
		TxCtx: &config.TransactionContext{
			Origin: caller,
		},
		State: state,
	}

	ret, _, err := Call(ContractAddress, TestInput, &runtimeCfg)
	if err != nil {
		t.Fatal("\ndidn't expect error", err)
	}
	t.Logf("\n\nReturn:%v", string(ret))

}
