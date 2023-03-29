package test

import (
	super "evm-container"
	"evm-container/common"
	"evm-container/config"
	"evm-container/state"
	"evm-container/vm"

	"github.com/holiman/uint256"
)

// Execute executes the code using the input as call data during the execution.
// It returns the EVM's return value, the new state and an error if it failed.
//
// Execute sets up an in-memory, temporary, environment for the execution of
// the given code. It makes sure that it's restored to its original state afterwards.
// Execute is basicly used for testing
func Execute(code, input []byte, cfg *config.Config) ([]byte, *state.StateDB, error) {
	if cfg == nil {
		cfg = new(config.Config)
	}
	config.SetDefaults(cfg)

	if cfg.State == nil {
		cfg.State = state.NewStateDB()
	}
	var (
		address = common.BytesToAddress([]byte("contract"))
		vmenv   = super.NewEnv(cfg)
		sender  = vm.AccountRef(cfg.TxCtx.Origin)
	)

	cfg.State.CreateAccount(address)
	// set the receiver's (the executing contract) code for execution.
	cfg.State.SetCode(address, code)
	// Call the code with the given configuration.
	ret, _, err := vmenv.Call(
		sender,
		common.BytesToAddress([]byte("contract")),
		input,
		cfg.BlockCtx.GasLimit,
		cfg.Value,
	)
	return ret, cfg.State, err
}

func Create(input []byte, cfg *config.Config) ([]byte, common.Address, uint64, error) {
	if cfg == nil {
		cfg = new(config.Config)
	}
	config.SetDefaults(cfg)

	if cfg.State == nil {
		cfg.State = state.NewStateDB()
	}

	// content above is for testing
	var (
		vmenv  = super.NewEnv(cfg)
		sender = vm.AccountRef(cfg.TxCtx.Origin)
		rules  = cfg.ChainConfig.Rules(vmenv.Context.BlockNumber, vmenv.Context.Random != nil, vmenv.Context.Time)
	)
	cfg.State.Prepare(rules, cfg.TxCtx.Origin, cfg.BlockCtx.Coinbase, nil, vm.ActivePrecompiles(rules), nil)

	code, address, leftOverGas, err := vmenv.Create(
		sender,
		input,
		cfg.BlockCtx.GasLimit,
		cfg.Value,
	)

	return code, address, leftOverGas, err
}

func Create2(input []byte, salt *uint256.Int, cfg *config.Config) ([]byte, common.Address, uint64, error) {
	if cfg == nil {
		cfg = new(config.Config)
	}
	config.SetDefaults(cfg)

	if cfg.State == nil {
		cfg.State = state.NewStateDB()
	}

	// content above is for testing
	var (
		vmenv  = super.NewEnv(cfg)
		sender = vm.AccountRef(cfg.TxCtx.Origin)
		rules  = cfg.ChainConfig.Rules(vmenv.Context.BlockNumber, vmenv.Context.Random != nil, vmenv.Context.Time)
	)
	cfg.State.Prepare(rules, cfg.TxCtx.Origin, cfg.BlockCtx.Coinbase, nil, vm.ActivePrecompiles(rules), nil)

	code, address, leftOverGas, err := vmenv.Create2(
		sender,
		input,
		cfg.BlockCtx.GasLimit,
		cfg.Value,
		salt,
	)

	return code, address, leftOverGas, err
}

func Call(address common.Address, input []byte, cfg *config.Config) ([]byte, uint64, error) {
	config.SetDefaults(cfg)

	// content above is for testing
	var (
		vmenv  = super.NewEnv(cfg)
		caller = vm.AccountRef(cfg.TxCtx.Origin)
	)

	ret, leftOverGas, err := vmenv.Call(
		caller,
		address,
		input,
		cfg.BlockCtx.GasLimit,
		cfg.Value,
	)
	return ret, leftOverGas, err
}

func DelegateCall(address common.Address, input []byte, cfg *config.Config) ([]byte, uint64, error) {
	config.SetDefaults(cfg)

	// content above is for testing
	var (
		vmenv  = super.NewEnv(cfg)
		caller = vm.AccountRef(cfg.TxCtx.Origin)
	)

	ret, leftOverGas, err := vmenv.DelegateCall(
		caller,
		address,
		input,
		cfg.BlockCtx.GasLimit,
	)
	return ret, leftOverGas, err
}

func StaticCall(address common.Address, input []byte, cfg *config.Config) ([]byte, uint64, error) {
	config.SetDefaults(cfg)

	// content above is for testing
	var (
		vmenv  = super.NewEnv(cfg)
		caller = vm.AccountRef(cfg.TxCtx.Origin)
	)

	ret, leftOverGas, err := vmenv.StaticCall(
		caller,
		address,
		input,
		cfg.BlockCtx.GasLimit,
	)
	return ret, leftOverGas, err
}
