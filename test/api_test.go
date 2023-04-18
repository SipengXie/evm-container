package test

import (
	"bytes"
	"encoding/json"
	"evm-container/common"
	"evm-container/config"
	"evm-container/test/types"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"testing"
)

var urlNewEnv string = "http://127.0.0.1:8000/api/evm/newEnv"
var urlCreate string = "http://127.0.0.1:8000/api/evm/create"
var urlCall string = "http://127.0.0.1:8000/api/evm/call"

func TestApiNewEnv(t *testing.T) {

	cfg := new(config.Config)
	config.SetDefaults(cfg)

	cfg_bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}

	json_str := string(cfg_bytes)
	payload := types.NewEnvRequest{
		Config: json_str,
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", urlNewEnv, reader)
	if err != nil {
		t.Fatal(err)
	}
	defer request.Body.Close()

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	t.Log("status", resp.Status)
	t.Log("response:", resp.Header)
	body, _ := io.ReadAll(resp.Body)

	res := types.NewEnvResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("response body:", res)
}

func TestApiCreate(t *testing.T) {

	caller := common.BytesToAddress([]byte("deployer"))
	code := "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033"
	gas := strconv.FormatUint(math.MaxUint64, 10)
	value := new(big.Int).String()

	payload := types.CreateRequest{
		Caller: caller.Hex(),
		Code:   code,
		Gas:    gas,
		Value:  value,
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	buffer := bytes.NewBuffer(bytesData)

	resp, err := http.Post(urlCreate, "application/json", buffer)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	res := types.CreateResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("response body:", res)
}

func TestApiCreateAndCall(t *testing.T) {
	caller := common.BytesToAddress([]byte("deployer"))
	code := "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033"
	gas := strconv.FormatUint(math.MaxUint64, 10)
	value := new(big.Int).String()

	payload := types.CreateRequest{
		Caller: caller.Hex(),
		Code:   code,
		Gas:    gas,
		Value:  value,
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	buffer := bytes.NewBuffer(bytesData)

	resp, err := http.Post(urlCreate, "application/json", buffer)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	t.Log("response body:", string(body))

	res := types.CreateResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatal(err)
	}

	//////////////////////////////////////////////////////////////////

	TestInput := "771602f700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"
	t.Log("Address is:", common.BytesToAddress(res.ContractAddr))
	payload_2 := types.CallRequest{
		Caller: caller.Hex(),
		Addr:   common.Bytes2Hex(res.ContractAddr),
		Input:  TestInput,
		Gas:    gas,
		Value:  value,
	}
	bytesData, err = json.Marshal(payload_2)
	if err != nil {
		t.Fatal(err)
	}

	buffer = bytes.NewBuffer(bytesData)

	resp, err = http.Post(urlCall, "application/json", buffer)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}

	t.Log("response body:", string(body))
	res_2 := types.CallResponse{}
	err = json.Unmarshal(body, &res_2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result:", res_2)

}
