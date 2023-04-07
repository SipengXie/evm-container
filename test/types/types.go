// Code generated by goctl. DO NOT EDIT.
package types

type NewEnvRequest struct {
	Config string `json:"config"`
}

type NewEnvResponse struct {
	Code string `json:"code"`
}

type SetBlockContextRequest struct {
	BlockCtx string `json:"blockCtx"`
}

type SetBlockContextResponse struct {
	Code string `json:"code"`
}

type ResetRequest struct {
	TxHash string `json:"txHash"`
	Index  string `json:"index"`
	TxCtx  string `json:"txCtx"`
}

type ResetResponse struct {
	Code string `json:"code"`
}

type CancelRequest struct {
}

type CancelResponse struct {
	Code string `json:"code"`
}

type CancelledRequest struct {
}

type CancelledResponse struct {
	Result bool `json:"result"`
}

type ChainConfigRequest struct {
}

type ChainConfigResponse struct {
	ChainConfig []byte `json:"chainConfig"`
}

type CreateRequest struct {
	Caller string `json:"caller"`
	Code   string `json:"code"`
	Gas    string `json:"gas"`
	Value  string `json:"value"`
}

type CreateResponse struct {
	Ret          []byte `json:"ret"`
	ContractAddr []byte `json:"contractAddr"`
	LeftOverGas  uint64 `json:"leftOverGas"`
}

type Create2Request struct {
	Caller string `json:"caller"`
	Code   string `json:"code"`
	Gas    string `json:"gas"`
	Value  string `json:"value"`
}

type Create2Response struct {
	Ret          []byte `json:"ret"`
	ContractAddr []byte `json:"contractAddr"`
	LeftOverGas  uint64 `json:"leftOverGas"`
}

type CallRequest struct {
	Caller string `json:"caller"`
	Addr   string `json:"addr"`
	Input  string `json:"input"`
	Gas    string `json:"gas"`
	Value  string `json:"value"`
}

type CallResponse struct {
	Ret         []byte `json:"ret"`
	LeftOverGas uint64 `json:"leftOverGas"`
}

type DelegateCallRequest struct {
	Caller string `json:"caller"`
	Addr   string `json:"addr"`
	Input  string `json:"input"`
	Gas    string `json:"gas"`
}

type DelegateCallResponse struct {
	Ret         []byte `json:"ret"`
	LeftOverGas uint64 `json:"leftOverGas"`
}

type StaticCallRequest struct {
	Caller string `json:"caller"`
	Addr   string `json:"addr"`
	Input  string `json:"input"`
	Gas    string `json:"gas"`
}

type StaticCallResponse struct {
	Ret         []byte `json:"ret"`
	LeftOverGas uint64 `json:"leftOverGas"`
}
