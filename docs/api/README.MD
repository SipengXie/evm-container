## API Service

### NewEnv
- Method : post 
- Path : [Service_Host]/api/evm/newEnv
- Param Type : Json
- Param Content:
```go
	NewEnvRequest {
        // A json string of Config (BlockContext, TxContext, ChainConfig and EVMConfig)
		Config string `json:"config"`
	}
```
- example: `see evm-container/test/api_test.go`
- Returned Value:
```go
	NewEnvResponse {
        // A string, "success" or err.Error()
		Code string `json:"code"`
	}
```

### SetBlockContext
- Method : post 
- Path : [Service_Host]/api/evm/setBlockContext
- Param Type : Json
- Param Content:
```go
    SetBlockContextRequest {
        // A json string of BlockContext
		BlockCtx string `json:"blockCtx"`
	}
```
- Returned Value:
```go
	SetBlockContextResponse {
        // A string, "success" or err.Error()
		Code string `json:"code"`
	}
```

### Reset
- Method : post 
- Path : [Service_Host]/api/evm/reset
- Param Type : Json
- Param Content:
```go
    ResetRequest {
        // Given Transaction Hash, Transaction index of the Block, and the Json string of the TransactionContext
		TxHash string `json:"txHash"`
		Index  string `json:"index"`
		TxCtx  string `json:"txCtx"`
	}
```
- Returned Value:
```go
	ResetResponse {
        // A string, "success" or err.Error()
		Code string `json:"code"`
	}
```

### Cancel
- Method : post 
- Path : [Service_Host]/api/evm/cancel
- Param Type : Json
- Param Content:
```go
    // just an empty request struct
	CancelRequest {
	}
```
- Returned Value:
```go
	CancelResponse {
        // A string, "success" or err.Error()
		Code string `json:"code"`
	}
```

### Cancelled
- Method : post 
- Path : [Service_Host]/api/evm/cancelled
- Param Type : Json
- Param Content:
```go
    // just an empty request struct
	CancelledRequest {
	}
```
- Returned Value:
```go
	CancelledResponse {
        // True means successfully cancelled
        // Flase means failed
		Result bool `json:"result"`
	}
```

### ChainConfig
- Method : post 
- Path : [Service_Host]/api/evm/chainConfig
- Param Type : Json
- Param Content:
```go
    // just an empty request struct
	ChainConfigRequest {
	}
```
- Returned Value:
```go
	ChainConfigResponse {
        // byte array that can be unmarshal to ChainConfig
		ChainConfig []byte `json:"chainConfig"`
	}
```

### Create
- Method : post 
- Path : [Service_Host]/api/evm/create
- Param Type : Json
- Param Content:
```go
	CreateRequest {
        // Hex string caller
		Caller string `json:"caller"`
        // Hex string code
		Code   string `json:"code"`
        // Decimal uint64 string
		Gas    string `json:"gas"`
        // Decimal Big.Int string
		Value  string `json:"value"`
	}
```
- example: `see evm-container/test/api_test.go`
- Returned Value:
```go
	CreateResponse {
        // Returned code in byte array
		Ret          []byte `json:"ret"`
        // Returned Contract Address in byte array
        // note that we have done the "StateDB.SetCode(addr, ret)" command.
		ContractAddr []byte `json:"contractAddr"`
        // Left Gas
		LeftOverGas  uint64 `json:"leftOverGas"`
	}
```

### Create2
- Method : post 
- Path : [Service_Host]/api/evm/create2
- Param Type : Json
- Param Content:
```go
	Create2Request {
        // Hex string caller
		Caller string `json:"caller"`
        // Hex string code
		Code   string `json:"code"`
        // Decimal uint64 string
		Gas    string `json:"gas"`
        // Decimal Big.Int string
		Value  string `json:"value"`
	}
```
- Returned Value:
```go
	Create2Response {
        // Returned code in byte array
		Ret          []byte `json:"ret"`
        // Returned Contract Address in byte array
        // note that we have done the "StateDB.SetCode(addr, ret)" command.
		ContractAddr []byte `json:"contractAddr"`
        // Left Gas
		LeftOverGas  uint64 `json:"leftOverGas"`
	}
```

### Call
- Method : post 
- Path : [Service_Host]/api/evm/call
- Param Type : Json
- Param Content:
```go
	CallRequest {
        // Hex string
		Caller string `json:"caller"`
        // Hex string
		Addr   string `json:"addr"`
        // Hex string
		Input  string `json:"input"`
        // Decimal uint64 string
		Gas    string `json:"gas"`
        // Decimal big.Int string
		Value  string `json:"value"`
	}
```
- example: `see evm-container/test/api_test.go`
- Returned Value:
```go
	CallResponse {
        // Returned value in byte array
		Ret         []byte `json:"ret"`
        // Left Gas
		LeftOverGas uint64 `json:"leftOverGas"`
	}
```

### DelegateCall
- Method : post 
- Path : [Service_Host]/api/evm/delegateCall
- Param Type : Json
- Param Content:
```go
	DelegateCallRequest {
        // Hex string
		Caller string `json:"caller"`
        // Hex string
		Addr   string `json:"addr"`
        // Hex string
		Input  string `json:"input"`
        // Decimal uint64 string
		Gas    string `json:"gas"`
	}
```
- Returned Value:
```go
	DelegateCallResponse {
        // Returned value in byte array
		Ret         []byte `json:"ret"`
        // Left Gas
		LeftOverGas uint64 `json:"leftOverGas"`
	}
```

### StaticCall
- Method : post 
- Path : [Service_Host]/api/evm/staticCall
- Param Type : Json
- Param Content:
```go
	StaticCallRequest {
        // Hex string
		Caller string `json:"caller"`
        // Hex string
		Addr   string `json:"addr"`
        // Hex string
		Input  string `json:"input"`
        // Decimal uint64 string
		Gas    string `json:"gas"`
        // Decimal big.Int string
		Value  string `json:"value"`
	}
```

- Returned Value:
```go
	StaticCallResponse {
        // Returned value in byte array
		Ret         []byte `json:"ret"`
        // Left Gas
		LeftOverGas uint64 `json:"leftOverGas"`
	}
```
