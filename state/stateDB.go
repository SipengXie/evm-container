package state

import (
	"encoding/json"
	"evm-container/common"
	"evm-container/crypto"
	"evm-container/params"
	"evm-container/types"
	"fmt"
	"math/big"
	"os"
)

// AccountState 实现vm的StateDB的接口 用于进行测试
type AccountState struct {
	Accounts   map[common.Address]*accountObject `json:"accounts,omitempty"`
	accessList *accessList
}

func NewAccountStateDb() *AccountState {
	return &AccountState{
		Accounts: make(map[common.Address]*accountObject),
	}
}

func (accSt *AccountState) getAccountObject(addr common.Address) *accountObject {
	obj, ok := accSt.Accounts[addr]
	if ok {
		return obj
	} else {
		return nil
	}
}

func (accSt *AccountState) setAccountObject(obj *accountObject) {
	accSt.Accounts[obj.Address] = obj
}

func (accSt *AccountState) getOrsetAccountObject(addr common.Address) *accountObject {
	get := accSt.getAccountObject(addr)
	if get != nil {
		return get
	}
	set := newAccountObject(addr, accountData{})
	accSt.setAccountObject(set)
	return set
}

// CreateAccount 创建账户接口
func (accSt *AccountState) CreateAccount(addr common.Address) {
	if accSt.getAccountObject(addr) != nil {
		return
	}
	obj := newAccountObject(addr, accountData{})
	accSt.setAccountObject(obj)
}

// SubBalance 减去某个账户的余额
func (accSt *AccountState) SubBalance(addr common.Address, amount *big.Int) {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SubBalance(amount)
	}
}

// AddBalance 增加某个账户的余额
func (accSt *AccountState) AddBalance(addr common.Address, amount *big.Int) {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.AddBalance(amount)
	}
}

// // GetBalance 获取某个账户的余额
func (accSt *AccountState) GetBalance(addr common.Address) *big.Int {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		return stateObject.Balance()
	}
	return new(big.Int).SetInt64(0)
}

// GetNonce 获取nonce
func (accSt *AccountState) GetNonce(addr common.Address) uint64 {
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.Nonce()
	}
	return 0
}

// SetNonce 设置nonce
func (accSt *AccountState) SetNonce(addr common.Address, nonce uint64) {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SetNonce(nonce)
	}
}

// GetCodeHash 获取代码的hash值
func (accSt *AccountState) GetCodeHash(addr common.Address) common.Hash {
	stateObject := accSt.getAccountObject(addr)
	if stateObject == nil {
		return common.Hash{}
	}
	return common.BytesToHash(stateObject.CodeHash())
}

// GetCode 获取智能合约的代码
func (accSt *AccountState) GetCode(addr common.Address) []byte {
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.Code()
	}
	return nil
}

// SetCode 设置智能合约的code
func (accSt *AccountState) SetCode(addr common.Address, code []byte) {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		stateObject.SetCode(crypto.Keccak256Hash(code).Bytes(), code)
	}
}

// GetCodeSize 获取code的大小
func (accSt *AccountState) GetCodeSize(addr common.Address) int {
	stateObject := accSt.getAccountObject(addr)
	if stateObject == nil {
		return 0
	}
	if stateObject.ByteCode != nil {
		return len(stateObject.ByteCode)
	}
	return 0
}

// AddRefund 暂时先忽略补偿
func (accSt *AccountState) AddRefund(uint64) {
	return
}

// SubRefund 暂时先忽略补偿
func (accSt *AccountState) SubRefund(uint64) {
	return
}

// GetRefund ...
func (accSt *AccountState) GetRefund() uint64 {
	return 0
}

// GetState 和SetState 是用于保存合约执行时 存储的变量是否发生变化 evm对变量存储的改变消耗的gas是有区别的
func (accSt *AccountState) GetState(addr common.Address, key common.Hash) common.Hash {
	stateObject := accSt.getAccountObject(addr)
	if stateObject != nil {
		return stateObject.GetStorageState(key)
	}
	return common.Hash{}
}

// SetState 设置变量的状态
func (accSt *AccountState) SetState(addr common.Address, key common.Hash, value common.Hash) {
	stateObject := accSt.getOrsetAccountObject(addr)
	if stateObject != nil {
		fmt.Printf("SetState key: %x value: %s", key, new(big.Int).SetBytes(value[:]).String())
		stateObject.SetStorageState(key, value)
	}
}

// Suicide 暂时禁止自杀
func (accSt *AccountState) Suicide(common.Address) bool {
	return false
}

// HasSuicided ...
func (accSt *AccountState) HasSuicided(common.Address) bool {
	return false
}

// Exist 检查账户是否存在
func (accSt *AccountState) Exist(addr common.Address) bool {
	return accSt.getAccountObject(addr) != nil
}

// Empty 是否是空账户
func (accSt *AccountState) Empty(addr common.Address) bool {
	so := accSt.getAccountObject(addr)
	return so == nil || so.Empty()
}

// RevertToSnapshot ...
func (accSt *AccountState) RevertToSnapshot(int) {

}

// Snapshot ...
func (accSt *AccountState) Snapshot() int {
	return 0
}

// AddLog 添加事件触发日志
func (accSt *AccountState) AddLog(log *types.Log) {
	fmt.Printf("log: %v", log)
}

// AddPreimage
func (accSt *AccountState) AddPreimage(common.Hash, []byte) {

}

// ForEachStorage  暂时没发现vm调用这个接口
func (accSt *AccountState) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) {

}

// Commit 进行持久存储 这里我们只将其简单的json话之后保存到本地磁盘中。
func (accSt *AccountState) Commit() error {
	// 将bincode写入文件
	file, err := os.Create("./account_sate.db")
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(accSt)
	//fmt.Println("len(binCode): ", len(binCode), " code: ", binCode)
	// bufW := bufio.NewWriter(file)
	// bufW.Write(binCode)
	// // bufW.WriteByte('\n')
	// bufW.Flush()
	file.Close()
	return err
}

// GetCommittedState retrieves a value from the given account's committed storage trie.
// TODO: need change persist storage
func (accSt *AccountState) GetCommittedState(addr common.Address, hash common.Hash) common.Hash {
	return accSt.GetState(addr, hash)
}

// GetTransientState gets transient storage for a given account.
// TODO: need change transient storage
func (accSt *AccountState) GetTransientState(addr common.Address, hash common.Hash) common.Hash {
	return accSt.GetState(addr, hash)
}

// SetTransientState sets transient storage for a given account. It
// adds the change to the journal so that it can be rolled back
// to its previous value if there is a revert.
// TODO: need change transient storage
func (accSt *AccountState) SetTransientState(addr common.Address, key, value common.Hash) {
	accSt.SetState(addr, key, value)
}

// TryLoadFromDisk  尝试从磁盘加载AccountState
func TryLoadFromDisk() (*AccountState, error) {
	file, err := os.Open("./account_sate.db")
	if err != nil && os.IsNotExist(err) {
		return NewAccountStateDb(), nil
	}
	if err != nil {
		return nil, err
	}

	// stat, _ := file.Stat()
	// // buf := stat.Size()
	var accStat AccountState

	err = json.NewDecoder(file).Decode(&accStat)
	return &accStat, err
}

func (accSt *AccountState) Prepare(rules params.Rules, sender, coinbase common.Address, dst *common.Address, precompiles []common.Address, list types.AccessList) {
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
func (accSt *AccountState) AddAddressToAccessList(addr common.Address) {
	accSt.accessList.AddAddress(addr)
}

// AddSlotToAccessList adds the given (address, slot)-tuple to the access list
func (accSt *AccountState) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	accSt.accessList.AddSlot(addr, slot)
}

// AddressInAccessList returns true if the given address is in the access list.
func (accSt *AccountState) AddressInAccessList(addr common.Address) bool {
	return accSt.accessList.ContainsAddress(addr)
}

// SlotInAccessList returns true if the given (address, slot)-tuple is in the access list.
func (accSt *AccountState) SlotInAccessList(addr common.Address, slot common.Hash) (addressPresent bool, slotPresent bool) {
	return accSt.accessList.Contains(addr, slot)
}
