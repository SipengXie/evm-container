package state

import (
	"context"
	"encoding/json"
	"evm-container/common"
	"evm-container/crypto"
	"evm-container/params"
	"evm-container/state/rpc/sdbclient"
	"evm-container/state/rpc/types/sdb"
	"evm-container/types"
	"fmt"
	"math/big"
)

// StateDB 实现vm的StateDB的接口 用于进行测试
type StateDB struct {
	Accounts   map[common.Address]*accountObject `json:"accounts,omitempty"`
	accessList *accessList
	SdbRpc     sdbclient.Sdb
	// ctx        context.Context
}

func NewStateDB(sdbRpc sdbclient.Sdb, ctx context.Context) *StateDB {
	return &StateDB{
		Accounts: make(map[common.Address]*accountObject),
		SdbRpc:   sdbRpc,
		// ctx:      ctx,
	}
}

func (accSt *StateDB) getAccountObject(addr common.Address) *accountObject {
	obj, ok := accSt.Accounts[addr]
	if ok {
		return obj
	} else {
		return nil
	}
}

func (accSt *StateDB) setAccountObject(obj *accountObject) {
	accSt.Accounts[obj.Address] = obj
}

func (accSt *StateDB) getOrsetAccountObject(addr common.Address) *accountObject {
	get := accSt.getAccountObject(addr)
	if get != nil {
		return get
	}
	set := newAccountObject(addr, accountData{})
	accSt.setAccountObject(set)
	return set
}

// CreateAccount 创建账户接口
func (accSt *StateDB) CreateAccount(addr common.Address) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.CreateAccount(context.Background(), &sdb.CreateAccountRequest{
			Addr: addr.Hex(),
		})
	}

	if accSt.getAccountObject(addr) != nil {
		return
	}
	obj := newAccountObject(addr, accountData{})
	accSt.setAccountObject(obj)
}

// SubBalance 减去某个账户的余额
func (accSt *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SubBalance(context.Background(), &sdb.SubBalanceRequest{
			Addr:   addr.Hex(),
			Amount: amount.String(),
		})
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SubBalance(amount)
	}
}

// AddBalance 增加某个账户的余额
func (accSt *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.AddBalance(context.Background(), &sdb.AddBalanceRequest{
			Addr:   addr.Hex(),
			Amount: amount.String(),
		})
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.AddBalance(amount)
	}
}

// // GetBalance 获取某个账户的余额
func (accSt *StateDB) GetBalance(addr common.Address) *big.Int {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetBalance(context.Background(), &sdb.GetBalanceRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return new(big.Int).SetInt64(0)
		}
		return new(big.Int).SetInt64(res.Amount)
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		return stateObject.Balance()
	}
	return new(big.Int).SetInt64(0)
}

// GetNonce 获取nonce
func (accSt *StateDB) GetNonce(addr common.Address) uint64 {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetNonce(context.Background(), &sdb.GetNonceRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return 0
		}
		return res.Amount
	}
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.Nonce()
	}
	return 0
}

// SetNonce 设置nonce
func (accSt *StateDB) SetNonce(addr common.Address, nonce uint64) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SetNonce(context.Background(), &sdb.SetNonceRequest{
			Addr:   addr.Hex(),
			Amount: nonce,
		})
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SetNonce(nonce)
	}
}

// GetCodeHash 获取代码的hash值
func (accSt *StateDB) GetCodeHash(addr common.Address) common.Hash {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetCodeHash(context.Background(), &sdb.GetCodeHashRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return common.Hash{}
		}
		return common.HexToHash(res.Hash)
	}
	stateObject := accSt.getAccountObject(addr)
	if stateObject == nil {
		return common.Hash{}
	}
	return common.BytesToHash(stateObject.CodeHash())
}

// GetCode 获取智能合约的代码
func (accSt *StateDB) GetCode(addr common.Address) []byte {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetCode(context.Background(), &sdb.GetCodeRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return nil
		}
		return []byte(res.Code)
	}
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.Code()
	}
	return nil
}

// SetCode 设置智能合约的code
func (accSt *StateDB) SetCode(addr common.Address, code []byte) {
	if accSt.SdbRpc != nil {
		_, err := accSt.SdbRpc.SetCode(context.Background(), &sdb.SetCodeRequest{
			Addr: addr.Hex(),
			Code: code,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SetCode(crypto.Keccak256Hash(code).Bytes(), code)
	}
}

// GetCodeSize 获取code的大小
func (accSt *StateDB) GetCodeSize(addr common.Address) int {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetCodeSize(context.Background(), &sdb.GetCodeSizeRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return 0
		}
		return int(res.Size)
	}
	stateObject := accSt.getAccountObject(addr)
	if stateObject == nil {
		return 0
	}
	if stateObject.ByteCode != nil {
		return len(stateObject.ByteCode)
	}
	return 0
}

// AddRefund
func (accSt *StateDB) AddRefund(amount uint64) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.AddRefund(context.Background(), &sdb.AddRefundRequest{
			Amount: amount,
		})
	}
}

// SubRefund
func (accSt *StateDB) SubRefund(amount uint64) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SubRefund(context.Background(), &sdb.SubRefundRequest{
			Amount: amount,
		})
	}
}

// GetRefund ...
func (accSt *StateDB) GetRefund() uint64 {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetRefund(context.Background(), &sdb.GetRefundRequest{})
		if err != nil {
			return 0
		}
		return res.Amount
	}
	return 0
}

// GetState 和SetState 是用于保存合约执行时 存储的变量是否发生变化 evm对变量存储的改变消耗的gas是有区别的
func (accSt *StateDB) GetState(addr common.Address, key common.Hash) common.Hash {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetState(context.Background(), &sdb.GetStateRequest{
			Addr: addr.Hex(),
			Hash: key.Hex(),
		})
		if err != nil {
			return common.Hash{}
		}
		return common.HexToHash(res.Hash)
	}
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.GetStorageState(key)
	}
	return common.Hash{}
}

// SetState 设置变量的状态
func (accSt *StateDB) SetState(addr common.Address, key common.Hash, value common.Hash) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SetState(context.Background(), &sdb.SetStateRequest{
			Addr:  addr.Hex(),
			Key:   key.Hex(),
			Value: value.Hex(),
		})
	}
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		fmt.Printf("SetState key: %x value: %s", key, new(big.Int).SetBytes(value[:]).String())
		stateObject.SetStorageState(key, value)
	}
}

// Suicide
func (accSt *StateDB) Suicide(addr common.Address) bool {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.Suicide(context.Background(), &sdb.SuicideRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return false
		}
		return res.IsSuicide
	}
	return false
}

// HasSuicided ...
func (accSt *StateDB) HasSuicided(addr common.Address) bool {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.HasSuicided(context.Background(), &sdb.HasSuicidedRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return false
		}
		return res.IsSuicide
	}
	return false
}

// Exist 检查账户是否存在
func (accSt *StateDB) Exist(addr common.Address) bool {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.Exist(context.Background(), &sdb.ExistRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return false
		}
		return res.Is_Exist
	}
	return accSt.getAccountObject(addr) != nil
}

// Empty 是否是空账户
func (accSt *StateDB) Empty(addr common.Address) bool {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.Empty(context.Background(), &sdb.EmptyRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return true
		}
		return res.Is_Empty
	}
	so := accSt.getAccountObject(addr)
	return so == nil || so.Empty()
}

// RevertToSnapshot ...
func (accSt *StateDB) RevertToSnapshot(id int) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.RevertToSnapshot(context.Background(), &sdb.RevertToSnapshotRequest{
			Revid: int32(id),
		})
	}
}

// Snapshot ...
func (accSt *StateDB) Snapshot() int {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.Snapshot(context.Background(), &sdb.SnapshotRequest{})
		if err != nil {
			return 0
		}
		return int(res.Revid)
	}
	return 0
}

// AddLog
func (accSt *StateDB) AddLog(log *types.Log) {
	if accSt.SdbRpc != nil {
		bytes, _ := json.Marshal(log)
		accSt.SdbRpc.AddLogJson(context.Background(), &sdb.AddLogJsonRequest{
			Json: string(bytes),
		})
	}
	fmt.Printf("log: %v", log)
}

// AddPreimage
func (accSt *StateDB) AddPreimage(hash common.Hash, preimage []byte) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.AddPreimage(context.Background(), &sdb.AddPreimageRequest{
			Hash:     hash.Hex(),
			Preimage: preimage,
		})
	}
}

// GetCommittedState retrieves a value from the given account's committed storage trie.
func (accSt *StateDB) GetCommittedState(addr common.Address, hash common.Hash) common.Hash {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetCommittedState(context.Background(), &sdb.GetCommittedStateRequest{
			Addr: addr.Hex(),
			Hash: hash.Hex(),
		})
		if err != nil {
			return common.Hash{}
		}
		return common.HexToHash(res.Hash)
	}
	return accSt.GetState(addr, hash)
}

// GetTransientState gets transient storage for a given account.

func (accSt *StateDB) GetTransientState(addr common.Address, hash common.Hash) common.Hash {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.GetTransientState(context.Background(), &sdb.GetTransientStateRequest{
			Addr: addr.Hex(),
			Key:  hash.Hex(),
		})
		if err != nil {
			return common.Hash{}
		}
		return common.HexToHash(res.Value)
	}
	return common.Hash{}
	// return accSt.GetState(addr, hash)
}

// SetTransientState sets transient storage for a given account. It
// adds the change to the journal so that it can be rolled back
// to its previous value if there is a revert.
func (accSt *StateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SetTransientState(context.Background(), &sdb.SetTransientStateRequest{
			Addr:  addr.Hex(),
			Key:   key.Hex(),
			Value: value.Hex(),
		})
	}
	accSt.SetState(addr, key, value)
}

func (accSt *StateDB) Prepare(rules params.Rules, sender, coinbase common.Address, dst *common.Address, precompiles []common.Address, list types.AccessList) {
	if accSt.SdbRpc != nil {
		var precompiles_str []string = make([]string, 0)
		for i := 0; i < len(precompiles); i++ {
			precompiles_str = append(precompiles_str, precompiles[i].Hex())
		}

		var access_list []*sdb.AccessTuple = make([]*sdb.AccessTuple, 0)
		var storage_key_str []string

		for i := 0; i < len(list); i++ {
			storage_key_str = make([]string, 0)
			for j := 0; j < len(list[i].StorageKeys); j++ {
				storage_key_str = append(storage_key_str, list[i].StorageKeys[j].Hex())
			}

			access_list = append(access_list, &sdb.AccessTuple{
				Addr:        list[i].Address.Hex(),
				StorageKeys: storage_key_str,
			})
		}
		var dstAddr *common.Address = &common.Address{}
		if dst != nil {
			dstAddr = dst
		}

		accSt.SdbRpc.Prepare(context.Background(), &sdb.PrepareRequest{
			Rule: &sdb.Rules{
				ChainID:          rules.ChainID.String(),
				IsHomestead:      rules.IsHomestead,
				IsEIP150:         rules.IsEIP150,
				IsEIP155:         rules.IsEIP155,
				IsEIP158:         rules.IsEIP158,
				IsByzantium:      rules.IsByzantium,
				IsConstantinople: rules.IsConstantinople,
				IsPetersburg:     rules.IsPetersburg,
				IsIstanbul:       rules.IsIstanbul,
				IsBerlin:         rules.IsBerlin,
				IsLondon:         rules.IsLondon,
				IsMerge:          rules.IsMerge,
				IsShanghai:       rules.IsShanghai,
				IsCancun:         rules.IsCancun,
				IsPrague:         rules.IsPrague,
			},
			SenderAddr:   sender.Hex(),
			CoinbaseAddr: coinbase.Hex(),
			DestAddr:     dstAddr.Hex(),
			PreCompiles:  precompiles_str,
			List:         access_list,
		})
	}

	if rules.IsBerlin {
		// Clear out any leftover from previous executions
		al := newAccessList()
		accSt.accessList = al

		al.AddAddress(sender)
		if dst != nil {
			al.AddAddress(*dst)
			// If it's a create-tx, the destination will be added inside evm.create
		}
		for _, addr := range precompiles {
			al.AddAddress(addr)
		}
		for _, el := range list {
			al.AddAddress(el.Address)
			for _, key := range el.StorageKeys {
				al.AddSlot(el.Address, key)
			}
		}
		if rules.IsShanghai { // EIP-3651: warm coinbase
			al.AddAddress(coinbase)
		}
	}
}

// AddAddressToAccessList adds the given address to the access list
func (accSt *StateDB) AddAddressToAccessList(addr common.Address) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.AddAddressToAccessList(context.Background(), &sdb.AddAddressToAccessListRequest{
			Addr: addr.Hex(),
		})
	}
	accSt.accessList.AddAddress(addr)
}

// AddSlotToAccessList adds the given (address, slot)-tuple to the access list
func (accSt *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.AddSlotToAccessList(context.Background(), &sdb.AddSlotToAccessListRequest{
			Addr: addr.Hex(),
			Slot: slot.Hex(),
		})
	}
	accSt.accessList.AddSlot(addr, slot)
}

// AddressInAccessList returns true if the given address is in the access list.
func (accSt *StateDB) AddressInAccessList(addr common.Address) bool {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.AddressInAccessList(context.Background(), &sdb.AddressInAccessListRequest{
			Addr: addr.Hex(),
		})
		if err != nil {
			return false
		}
		return res.IsIn
	}
	return accSt.accessList.ContainsAddress(addr)
}

// SlotInAccessList returns true if the given (address, slot)-tuple is in the access list.
func (accSt *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressPresent bool, slotPresent bool) {
	if accSt.SdbRpc != nil {
		res, err := accSt.SdbRpc.SlotInAccessList(context.Background(), &sdb.SlotInAccessListRequest{
			Addr: addr.Hex(),
			Hash: slot.Hex(),
		})
		if err != nil {
			return false, false
		}
		return res.AddrOk, res.SlotOk
	}
	return accSt.accessList.Contains(addr, slot)
}

func (accSt *StateDB) SetTxContext(txHash common.Hash, index int32) {
	if accSt.SdbRpc != nil {
		accSt.SdbRpc.SetTxContext(context.Background(), &sdb.SetTxContextRequest{
			TxHash: txHash.Hex(),
			Index:  index,
		})
	}
}
