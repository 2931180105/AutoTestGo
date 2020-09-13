package wingGov

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"testing"
)

func GetTestConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../config_prv.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func TestDeployContractOETHToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//fmt.Println(cfg, account, sdk)
	QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
	//DeployContractODAIToken(cfg, account, sdk)
}
