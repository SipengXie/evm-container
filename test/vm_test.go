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

	if cfg.Difficulty == nil {
		t.Error("expected difficulty to be non nil")
	}

	if cfg.GasLimit == 0 {
		t.Error("didn't expect gaslimit to be zero")
	}
	if cfg.GasPrice == nil {
		t.Error("expected time to be non nil")
	}
	if cfg.Value == nil {
		t.Error("expected time to be non nil")
	}
	if cfg.GetHashFn == nil {
		t.Error("expected time to be non nil")
	}
	if cfg.BlockNumber == nil {
		t.Error("expected block number to be non nil")
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
	state := state.NewAccountStateDb()
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
	state := state.NewAccountStateDb()
	HexTestCode := common.Hex2Bytes("608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033")

	deployer := common.BytesToAddress([]byte("deployer"))
	state.CreateAccount(deployer)
	runtimeCfg := config.Config{
		Origin: deployer,
		State:  state,
	}

	code, address, leftGas, err := Create(HexTestCode, &runtimeCfg)
	if err != nil {
		t.Fatal("didn't expect error", err)
	}
	t.Logf("\n\n code:%v \n\naddress:%v\n\n leftGas:%v", code, address, leftGas)
}

func TestCreateAndCall(t *testing.T) {
	state := state.NewAccountStateDb()
	// A + B contract
	HexTestCode := common.Hex2Bytes("608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122024ef1eaa50221cd8ca75d909a80743d456152d97e26343abb3471bd207793e3b64736f6c63430008070033")
	TestInput := common.Hex2Bytes("771602f700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002")
	deployer := common.BytesToAddress([]byte("deployer"))
	caller := common.BytesToAddress([]byte("caller"))
	state.CreateAccount(deployer)
	state.CreateAccount(caller)
	state.AddBalance(deployer, big.NewInt(1<<40))
	state.AddBalance(caller, big.NewInt(1<<40))

	runtimeCfg := config.Config{
		Origin: deployer,
		State:  state,
	}

	code, ContractAddress, _, err := Create(HexTestCode, &runtimeCfg)
	if err != nil {
		t.Fatal("\ndidn't expect error", err)
	}
	state.SetCode(ContractAddress, code)

	t.Logf("\nCode:%v", code)

	runtimeCfg = config.Config{
		Origin: caller,
		State:  state,
	}

	ret, _, err := Call(ContractAddress, TestInput, &runtimeCfg)
	if err != nil {
		t.Fatal("\ndidn't expect error", err)
	}
	t.Logf("\n\nReturn:%v", ret)

}
