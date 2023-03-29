package config

import (
	"encoding/json"
	"evm-container/common"
	"evm-container/logger"
	"evm-container/params"
	"evm-container/state"
	"math"
	"math/big"
)

// Config is a basic type specifying certain configuration flags for running
// the EVM.

type BlockContext struct {
	Coinbase    common.Address
	BlockNumber *big.Int
	Time        uint64
	Difficulty  *big.Int
	GasLimit    uint64
	BaseFee     *big.Int
}

type TransactionContext struct {
	Origin   common.Address
	GasPrice *big.Int
}

type VmConfig struct {
	Debug                   bool  // Enables debugging
	NoBaseFee               bool  // Forces the EIP-1559 baseFee to 0 (needed for 0 price calls)
	EnablePreimageRecording bool  // Enables recording of SHA3/keccak preimages
	ExtraEips               []int // Additional EIPS that are to be enabled
}

type Config struct {
	ChainConfig *params.ChainConfig
	BlockCtx    *BlockContext
	TxCtx       *TransactionContext
	LogCfg      *logger.Config
	EVMConfig   VmConfig

	Value *big.Int
	State *state.StateDB // TODO: use our State Interface
}

func NewConfig(data []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SetDefaults(cfg *Config) {
	if cfg.ChainConfig == nil {
		cfg.ChainConfig = &params.ChainConfig{
			ChainID:             big.NewInt(1),
			HomesteadBlock:      new(big.Int),
			DAOForkBlock:        new(big.Int),
			DAOForkSupport:      false,
			EIP150Block:         new(big.Int),
			EIP150Hash:          common.Hash{},
			EIP155Block:         new(big.Int),
			EIP158Block:         new(big.Int),
			ByzantiumBlock:      new(big.Int),
			ConstantinopleBlock: new(big.Int),
			PetersburgBlock:     new(big.Int),
			IstanbulBlock:       new(big.Int),
			MuirGlacierBlock:    new(big.Int),
			BerlinBlock:         new(big.Int),
			LondonBlock:         new(big.Int),
		}
	}

	if cfg.BlockCtx == nil {
		cfg.BlockCtx = &BlockContext{
			Difficulty:  new(big.Int),
			GasLimit:    math.MaxUint64,
			BlockNumber: new(big.Int),
			BaseFee:     big.NewInt(params.InitialBaseFee),
		}
	}

	if cfg.TxCtx == nil {
		cfg.TxCtx = &TransactionContext{
			GasPrice: new(big.Int),
		}
	}

	if cfg.Value == nil {
		cfg.Value = new(big.Int)
	}

}
