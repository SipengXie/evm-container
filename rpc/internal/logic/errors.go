package logic

import "errors"

var (
	ErrMissingEvmInstance     = errors.New("Missing Evm Instance")
	ErrMissingStateDBInstance = errors.New("Missing StateDB Instance")
)
