package main

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"time"
)

var sdk = goSdk.NewOntologySdk()

func main() {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "wing-test/config.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		return
	}
	wallet, err := sdk.OpenWallet(cfg.Wallet)
	if err != nil {
		log.Errorf("parse wallet err: %s", err)
	}
	account, err := wallet.GetDefaultAccount([]byte(cfg.Password))
	if err != nil {
		log.Errorf("get account err: %s", err)
	}
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sdk.SetDefaultClient(rpcClient)
	tx := Utils.UpdateProfitContract(cfg, account, sdk)

	hash, err := sdk.SendTransaction(tx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
}
