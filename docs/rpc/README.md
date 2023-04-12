## RPC Service
`etcd key: evm.rpc`
```go
EvmRpc: rpcclient.NewRpc(zrpc.MustNewClient(c.EvmRpc))
```
### NewEnv
- Param Content:
```protobuf
	message NewEnvRequest{
		// json byte array of Config structure
		bytes config = 1;
	}
```
- Example: see `evm-container/test/rpc_test.go`
- Returned Value:
```protobuf
	message NewEnvResponse{
		// A string, "success" or err.Error()
		string code = 1;
	}
```

### SetBlockContext
- Param Content:
```protobuf
	message SetBlockContextRequest{
		// json byte array of BlockContext structure
		bytes blockCtx = 1;
	}
```
- Returned Value:
```protobuf
	message SetBlockContextResponse{
		// A string, "success" or err.Error()
		string code = 1;
	}
```

### Reset
- Param Content:
```protobuf
	message ResetRequest{
		// transaction hash in byte array
		bytes txHash = 1;
		// transaction index
		int32 index = 2;
		// json byte array of TransactionContext structure
		bytes txCtx = 3;
	}
```
- Returned Value:
```protobuf
	message ResetResponse{
		// A string, "success" or err.Error()
		string code = 1;
	}
```

### Cancel
- Param Content:
```protobuf
	message CancelRequset{
		// just an empty structure
	}
```
- Returned Value:
```protobuf
	message CancelResponse{
		// A string, "success" or err.Error()
		string code = 1;
	}
```

### Cancelled
- Param Content:
```protobuf
	message CancelledRequest{
		// just an empty structure
	}
```
- Returned Value:
```protobuf
	message CancelledResponse{
		// a bool result, true means Evm has been cancelled
		// false means the opposite way.
		bool result = 1;
	}
```

### ChainConfig
- Param Content:
```protobuf
	message ChainConfigRequest{
		// just an empty structure
	}
```
- Returned Value:
```protobuf
	message CancelledResponse{
		// a bool result, true means Evm has been cancelled
		// false means the opposite way.
		bool result = 1;
	}
```

### Create
- Param Content:
```protobuf
	message CreateRequest{
		// caller address in byte arrary (common.Address.Byte())
		bytes caller = 1;
		// code in byte array(common.Hex2Bytes())
		bytes code = 2;
		// gas limit
		uint64 gas = 3;
		// big.Int in string
		string value =4;
	}
```
- Example: see `evm-container/test/rpc_test.go`
- Returned Value:
```protobuf
	message CreateResponse{
		// returned value in byte array
		bytes ret = 1;
		// contract address in byte array
		bytes contractAddr = 2;
		// left gas
		uint64 leftOverGas = 3;
	}
```

### Create2
```protobuf
	message Create2Request{
		// caller address in byte arrary (common.Address.Byte())
		bytes caller = 1;
		// code in byte array(common.Hex2Bytes())
		bytes code = 2;
		// gas limit
		uint64 gas = 3;
		// big.Int in string
		string value =4;
	}
```
- Returned Value:
```protobuf
	message Create2Response{
		// returned value in byte array
		bytes ret = 1;
		// contract address in byte array
		bytes contractAddr = 2;
		// left gas
		uint64 leftOverGas = 3;
	}
```

### Call
```protobuf
	message CallRequest{
		// caller address in byte arrary (common.Address.Byte())
		bytes caller = 1;
		// contract address in byte arrary (common.Address.Byte())
		bytes addr = 2;
		// input in byte array(common.Hex2Bytes())
		bytes input = 3;
		// gas Limit
		uint64 gas = 4;
		// big.Int in string
		string value = 5;
	}
```
- Example: see `evm-container/test/rpc_test.go`
- Returned Value:
```protobuf
	message CallResponse{
		// returned value in byte array
		bytes ret = 1;
		// left gas
		uint64 leftOverGas = 2;
	}
```

### DelegateCall
```protobuf
	message DelegateCallRequest{
		// caller address in byte arrary (common.Address.Byte())
		bytes caller = 1;
		// contract address in byte arrary (common.Address.Byte())
		bytes addr = 2;
		// input in byte array(common.Hex2Bytes())
		bytes input = 3;
		// gas Limit
		uint64 gas = 4;
		// big.Int in string
		string value = 5;
	}
```
- Returned Value:
```protobuf
	message DelegateCallResponse{
		// returned value in byte array
		bytes ret = 1;
		// left gas
		uint64 leftOverGas = 2;
	}
```

### StaticCall
```protobuf
	message StaticCallRequest{
		// caller address in byte arrary (common.Address.Byte())
		bytes caller = 1;
		// contract address in byte arrary (common.Address.Byte())
		bytes addr = 2;
		// input in byte array(common.Hex2Bytes())
		bytes input = 3;
		// gas Limit
		uint64 gas = 4;
		// big.Int in string
		string value = 5;
	}
```
- Returned Value:
```protobuf
	message StaticCallResponse{
		// returned value in byte array
		bytes ret = 1;
		// left gas
		uint64 leftOverGas = 2;
	}
```