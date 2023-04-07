package logic

import "errors"

var (
	ErrMissingEvmInstance     = errors.New("missing Evm Instance")
	ErrMissingStateDBInstance = errors.New("missing StateDB Instance")
)
