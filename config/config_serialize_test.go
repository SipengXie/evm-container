package config_test

import (
	"encoding/json"
	"evm-container/config"
	"testing"
)

func TestSerialization(t *testing.T) {
	cfg := config.Config{}
	config.SetDefaults(&cfg)
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal("Unable to Marshel cfg", err)
	}

	cfg2, err := config.NewConfig(data)
	if err != nil {
		t.Fatal("Unable to unmarshel data", err)
	}

	if cfg2.Value.Cmp(cfg.Value) != 0 {
		t.Fatal("Incorrect unmarshel", cfg.Value, cfg2.Value)
	}
}
