package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Token uint8

const (
	ONT Token = iota
	ONG
	OEP4
)

type Config struct {
	Wallet           string
	WingGov          string
	GovToken         string
	ZeroPool         string
	Password         string
	SDRate           int // static and dyamic profit
	ContractCodePath string
	Oracle           string
	WingProfit       string // HexString
	GlobalParam      string // HexString
	AuthAddr         string
	Amount           uint64
	Weight           int
	Rpc              []string
	TxNum            uint // whole tx num is *TxFactor
	TxFactor         uint
	RoutineNum       uint // whole tx save to RoutineNum files, and one go-routine per file
	TPS              uint
	StartNonce       uint32
	GasPrice         uint64
	GasLimit         uint64
	SaveTx           bool
	SendTx           bool
	Owner            string
}

func ParseConfig(path string) (*Config, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: failed, err: %s", err)
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: failed, err: %s", err)
	}
	return config, nil
}
