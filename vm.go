package evmcontainer

import (
	"evm-container/common"
	"evm-container/config"
	"evm-container/crypto"
	"evm-container/logger"
	"evm-container/vm"
	"math/big"
	"os"
)

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, amount *big.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

func GetHashFn(n uint64) common.Hash {
	return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
}

func NewEnv(cfg *config.Config) *vm.EVM {

	blockContext := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn,
		Coinbase:    cfg.BlockCtx.Coinbase,
		BlockNumber: cfg.BlockCtx.BlockNumber,
		Time:        cfg.BlockCtx.Time,
		Difficulty:  cfg.BlockCtx.Difficulty,
		GasLimit:    cfg.BlockCtx.GasLimit,
		BaseFee:     cfg.BlockCtx.BaseFee,
	}

	txContext := vm.TxContext{
		Origin:   cfg.TxCtx.Origin,
		GasPrice: cfg.TxCtx.GasPrice,
	}

	EVMConfig := vm.Config{
		Debug:                   cfg.EVMConfig.Debug,
		NoBaseFee:               cfg.EVMConfig.NoBaseFee,
		EnablePreimageRecording: cfg.EVMConfig.EnablePreimageRecording,
		ExtraEips:               cfg.EVMConfig.ExtraEips,
	}
	if cfg.LogCfg != nil {
		EVMConfig.Tracer = logger.NewJSONLogger(cfg.LogCfg, os.Stdout)
	}

	// wait for State
	return vm.NewEVM(blockContext, txContext, cfg.State, cfg.ChainConfig, EVMConfig)
}
