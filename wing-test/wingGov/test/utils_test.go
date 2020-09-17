package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
)

func GetTestConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../../config_testnet2.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetMainConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/config_main.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("ParseConfig error:%s", err)
	}
	wallet, _ := sdk.OpenWallet("/Volumes/MM/WING_OTHER_OWNER.dat")
	//wallet, err := sdk.OpenWallet("/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/WING_OWNER.dat")
	if err != nil {
		log.Errorf("OpenWallet error:%s", err)
	}
	account, err := wallet.GetDefaultAccount([]byte(cfg.Password))
	if err != nil {
		log.Errorf("GetDefaultAccount error:%s", err)
	}
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetPrvConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
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
