package test

import (
	"bytes"
	"encoding/json"
	"evm-container/common"
	"evm-container/config"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"testing"
)

type NewEnvRequest struct {
	Config string `json:"config"`
}

type CreateRequest struct {
	Caller string `json:"caller"`
	Code   string `json:"code"`
	Gas    string `json:"gas"`
	Value  string `json:"value"`
}

func TestApiNewEnv(t *testing.T) {

	cfg := new(config.Config)
	config.SetDefaults(cfg)

	urlNewEnv := "http://127.0.0.1:8888/api/evm/newEnv"

	// urlCall := "127.0.0.1/api/evm/call"

	cfg_bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}

	json_str := string(cfg_bytes)

	payload := NewEnvRequest{
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
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response body:", string(body))
}

func TestApiCreate(t *testing.T) {
	urlCreate := "http://127.0.0.1:8888/api/evm/create"
	caller := common.BytesToAddress([]byte("deployer"))
	code := "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063771602f714610030575b600080fd5b61004a6004803603810190610045919061008b565b610060565b60405161005791906100da565b60405180910390f35b6000818361006e91906100f5565b905092915050565b60008135905061008581610189565b92915050565b600080604083850312156100a2576100a1610184565b5b60006100b085828601610076565b92505060206100c185828601610076565b9150509250929050565b6100d48161014b565b82525050565b60006020820190506100ef60008301846100cb565b92915050565b60006101008261014b565b915061010b8361014b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101405761013f610155565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6101928161014b565b811461019d57600080fd5b5056fea264697066735822122081d5aab419bede5a1093121d0a0ded9cba91d5d72688b3518a2e323b9865936f64736f6c63430008070033"
	gas := strconv.FormatUint(math.MaxUint64, 10)
	value := new(big.Int).String()

	payload := CreateRequest{
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error:", err)
		return
	}
	t.Log(string(body))
}
