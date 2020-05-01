package main

import (
	config "github.com/mockyz/AutoTestGo/ont-test/config_ont"
	OntTools "github.com/mockyz/AutoTestGo/ont-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
)

var sdk = goSdk.NewOntologySdk()

func main() {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "config.json"
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
	//OntTools.BindAssetHash(cfg, account)
	//OntTools.GetCrossLimt(cfg, account)
	//OntTools.GetAseetHash(cfg, account)

	OntTools.TestTransfer(cfg, account)
	//address,_ :=common.AddressFromBase58("AUeKhaRr9xwy114zwsVarYvVG13C2T3C9o")
	//OntTools.BalanceOf(sdk,address,cfg.EthX)
}
