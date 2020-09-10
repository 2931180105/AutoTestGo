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
	Wallet             string
	WingGov            string
	GovToken           string
	ZeroPool           string
	Password           string
	SDRate             int // static and dyamic profit
	ContractCodePath   string
	Oracle             string
	InterestRate       string
	Comptroller        string
	FToken             string
	Token              string
	FTokenInsurance    string
	WingProfit         string // HexString
	GlobalParam        string // HexString
	AuthAddr           string
	OUSDT              string
	OWBTC              string
	OETH               string
	ONT                string
	FBTC               string
	FETH               string
	FUSDT              string
	IBTC               string
	IETH               string
	IUSDT              string
	ExchangeRate       int
	Amount             uint64
	Weight             int
	Eta                int
	Gama               int
	TotalStaticProfit  int
	TotalDynamicProfit int
	LendAmount         []int
	Rpc                []string
	TxNum              uint // whole tx num is *TxFactor
	TxFactor           uint
	RoutineNum         uint // whole tx save to RoutineNum files, and one go-routine per file
	TPS                uint
	StartNonce         uint32
	GasPrice           uint64
	GasLimit           uint64
	SaveTx             bool
	SendTx             bool
	Owner              string
	AccountNum         int
	StakeOnt           int
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
