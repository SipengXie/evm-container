package state

import (
	"bytes"
	"evm-container/common"
	"evm-container/crypto"
	"math/big"
)

var emptyCodeHash = crypto.Keccak256(nil)

type accountData struct {
	Nonce    uint64      `json:"nonce,omitempty"`
	Balance  *big.Int    `json:"balance,omitempty"`
	Root     common.Hash `json:"root,omitempty"` // merkle root of the storage trie
	CodeHash []byte      `json:"code_hash,omitempty"`
}

type accountObject struct {
	Address      common.Address              `json:"address,omitempty"`
	AddrHash     common.Hash                 `json:"addr_hash,omitempty"` // hash of ethereum address of the account
	ByteCode     []byte                      `json:"byte_code,omitempty"`
	Data         accountData                 `json:"data,omitempty"`
	CacheStorage map[common.Hash]common.Hash `json:"cache_storage,omitempty"` // 用于缓存存储的变量
}

func newAccountObject(address common.Address, data accountData) *accountObject {
	if data.Balance == nil {
		data.Balance = new(big.Int)
	}
	if data.CodeHash == nil {
		data.CodeHash = emptyCodeHash
	}
	return &accountObject{
		Address:      address,
		AddrHash:     common.BytesToHash(crypto.Keccak256(address[:])),
		Data:         data,
		CacheStorage: make(map[common.Hash]common.Hash),
	}
}

func (object *accountObject) Balance() *big.Int {
	return object.Data.Balance
}

func (object *accountObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	object.Data.Balance = new(big.Int).Sub(object.Balance(), amount)
}

func (object *accountObject) AddBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	object.Data.Balance = new(big.Int).Add(object.Balance(), amount)
}

func (object *accountObject) Nonce() uint64 {
	return object.Data.Nonce
}

func (object *accountObject) SetNonce(nonce uint64) {
	object.Data.Nonce = nonce
}

func (object *accountObject) CodeHash() []byte {
	return object.Data.CodeHash
}

func (object *accountObject) Code() []byte {
	return object.ByteCode
}

func (object *accountObject) SetCode(codeHash []byte, code []byte) {
	object.Data.CodeHash = codeHash
	object.ByteCode = code
}

func (object *accountObject) GetStorageState(key common.Hash) common.Hash {
	value, exist := object.CacheStorage[key]
	if exist {
		// fmt.Println("exist cache ", " key: ", key, " value: ", value)
		return value
	}
	return common.Hash{}
}

func (object *accountObject) SetStorageState(key, value common.Hash) {
	object.CacheStorage[key] = value
}

func (object *accountObject) Empty() bool {
	return object.Data.Nonce == 0 && object.Data.Balance.Sign() == 0 && bytes.Equal(object.Data.CodeHash, emptyCodeHash)
}
