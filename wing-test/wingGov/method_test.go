package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"testing"
)

func GetContext() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "/home/ubuntu/go/src/github.com/mockyz/AutoTestGo/wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func TestDeployContractOETHToken(t *testing.T) {
	cfg, account, sdk := GetContext()
	fmt.Println(cfg, account, sdk)
	DeployContractODAIToken(cfg, account, sdk)
}
