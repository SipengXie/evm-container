package test

import (
	super "evm-container"
	"evm-container/common"
	"evm-container/config"
	"evm-container/state"
	"evm-container/vm"
)

// Execute executes the code using the input as call data during the execution.
// It returns the EVM's return value, the new state and an error if it failed.
//
// Execute sets up an in-memory, temporary, environment for the execution of
// the given code. It makes sure that it's restored to its original state afterwards.
func Execute(code, input []byte, cfg *config.Config) ([]byte, *state.AccountState, error) {
	if cfg == nil {
		cfg = new(config.Config)
	}
	config.SetDefaults(cfg)

	if cfg.State == nil {
		cfg.State = state.NewAccountStateDb()
	}
	var (
		address = common.BytesToAddress([]byte("contract"))
		vmenv   = super.NewEnv(cfg)
		sender  = vm.AccountRef(cfg.Origin)
	)

	cfg.State.CreateAccount(address)
	// set the receiver's (the executing contract) code for execution.
	cfg.State.SetCode(address, code)
	// Call the code with the given configuration.
	ret, _, err := vmenv.Call(
		sender,
		common.BytesToAddress([]byte("contract")),
		input,
		cfg.GasLimit,
		cfg.Value,
	)
	return ret, cfg.State, err
}

// Create executes the code using the EVM create method
func Create(input []byte, cfg *config.Config) ([]byte, common.Address, uint64, error) {
	if cfg == nil {
		cfg = new(config.Config)
	}
	config.SetDefaults(cfg)

	if cfg.State == nil {
		cfg.State = state.NewAccountStateDb()
	}
	var (
		vmenv  = super.NewEnv(cfg)
		sender = vm.AccountRef(cfg.Origin)
		rules  = cfg.ChainConfig.Rules(vmenv.Context.BlockNumber, vmenv.Context.Random != nil, vmenv.Context.Time)
	)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, nil, vm.ActivePrecompiles(rules), nil)

	code, address, leftOverGas, err := vmenv.Create(
		sender,
		input,
		cfg.GasLimit,
		cfg.Value,
	)
	return code, address, leftOverGas, err
}

// Call executes the code given by the contract's address. It will return the
// EVM's return value or an error if it failed.
//
// Call, unlike Execute, requires a config and also requires the State field to
// be set.
func Call(address common.Address, input []byte, cfg *config.Config) ([]byte, uint64, error) {
	config.SetDefaults(cfg)

	var (
		vmenv  = super.NewEnv(cfg)
		caller = vm.AccountRef(cfg.Origin)
	)

	// Call the code with the given configuration.
	ret, leftOverGas, err := vmenv.Call(
		caller,
		address,
		input,
		cfg.GasLimit,
		cfg.Value,
	)
	return ret, leftOverGas, err
}
