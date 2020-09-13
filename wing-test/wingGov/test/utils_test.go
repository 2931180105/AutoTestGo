package test

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
)

func GetTestConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../../config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetPrvConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../../config_prv.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetTestConfigAddAccts() (*config.Config, *goSdk.Account, *goSdk.OntologySdk, []*goSdk.Account) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../../config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	accts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)
	return cfg, account, sdk, accts
}
