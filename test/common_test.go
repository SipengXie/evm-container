package test

import (
	"evm-container/common"
	"testing"
)

type attrbute interface {
	Add_one()
	Get_num() int
}

type sa struct {
	Num int
}

func (a *sa) Add_one() {
	a.Num += 1
}
func (a *sa) Get_num() int {
	return a.Num
}

type sb struct {
	Att attrbute
}

func TestXxx(t *testing.T) {
	a := sa{Num: 0}
	b := sb{Att: &a}
	a.Add_one()
	t.Log(b.Att.Get_num())
}

func TestCommon(t *testing.T) {
	bytes := []byte("test")
	address := common.BytesToAddress(bytes)
	hex := address.Hex()
	new_addr := common.BytesToAddress([]byte(hex))
	new_hex := new_addr.Hex()
	t.Log(hex)
	t.Log(new_hex)
}
