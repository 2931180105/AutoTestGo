package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Wallet             string   `json:"Wallet"`
	Password           string   `json:"Password"`
	SDRate             uint64   `json:"SDRate"`
	WingStartTime      uint64   `json:"WingStartTime"`
	Oracle             string   `json:"Oracle"`
	GovToken           string   `json:"GovToken"`
	ZeroPool           string   `json:"ZeroPool"`
	WingGov            string   `json:"WingGov"`
	Comptroller        string   `json:"Comptroller"`
	OldZeroPool        string   `json:"OldZeroPool"`
	WingProfit         string   `json:"WingProfit"`
	GlobalParam        string   `json:"GlobalParam"`
	Pool2              string   `json:"Pool2"`
	uint64erestRate    string   `json:"uint64erestRate"`
	FBTC               string   `json:"FBTC"`
	FETH               string   `json:"FETH"`
	FUSDT              string   `json:"FUSDT"`
	IBTC               string   `json:"IBTC"`
	IETH               string   `json:"IETH"`
	IUSDT              string   `json:"IUSDT"`
	AuthAddr           string   `json:"AuthAddr"`
	WBTC               string   `json:"WBTC"`
	ONTd               string   `json:"ONTd"`
	RenBTC             string   `json:"RenBTC"`
	USDC               string   `json:"USDC"`
	WING               string   `json:"WING"`
	ETH                string   `json:"ETH"`
	DAI                string   `json:"DAI"`
	USDT               string   `json:"USDT"`
	SUSD               string   `json:"SUSD"`
	NEO                string   `json:"NEO"`
	OKB                string   `json:"OKB"`
	UNI                string   `json:"UNI"`
	Weight             uint64   `json:"Weight"`
	Amount             uint64   `json:"Amount"`
	Eta                uint64   `json:"Eta"`
	Gama               uint64   `json:"Gama"`
	ExchangeRate       uint64   `json:"ExchangeRate"`
	GasPrice           uint64   `json:"GasPrice"`
	GasLimit           uint64   `json:"GasLimit"`
	RPC                []string `json:"Rpc"`
	TotalStaticProfit  uint64   `json:"TotalStaticProfit"`
	LendAmount         []uint64 `json:"LendAmount"`
	TotalDynamicProfit uint64   `json:"TotalDynamicProfit"`
	Wallets            []string `json:"Wallets"`
	Wifs               []string `json:"Wifs"`
	TxNum              uint64   `json:"TxNum"`
	TxFactor           uint64   `json:"TxFactor"`
	RoutineNum         uint64   `json:"RoutineNum"`
	TPS                uint64   `json:"TPS"`
	AccountNum         uint64   `json:"AccountNum"`
	StakeOnt           uint64   `json:"StakeOnt"`
	StartNonce         uint64   `json:"StartNonce"`
	SaveTx             bool     `json:"SaveTx"`
	SendTx             bool     `json:"SendTx"`
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
