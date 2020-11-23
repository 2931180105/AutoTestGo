package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	if_borrow "github.com/mockyz/AutoTestGo/wing-test/if-pool/if-borrow"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"math/big"
	"testing"
)

func TestIfBorrow(t *testing.T){
	var sdk = goSdk.NewOntologySdk()
	//configPath := ".../config_prv.json"
	configPath := "../config_testnet.json"

	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet(cfg.Wallet)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sdk.SetDefaultClient(rpcClient)

	IB,err :=if_borrow.NewIfBorrowPool(cfg.Rpc[0],"03f5da55634c916a01107719ea5de17d725e9d54",account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil{
		log.Errorf("NewIfBorrowPool err:%v",err)
		return
	}
	syb,err :=IB.MarketName()
	if err !=nil{
		log.Errorf("Symbol err:%v",err)
		return
	}
	log.Infof("Decimals is :%s",syb)
	hash,err :=IB.IncreaseCollateral(account.Address,big.NewInt(10000000))
	if err !=nil{
		log.Errorf("IncreaseCollateral err:%v",err)
		return
	}
	log.Infof("hash is :%s",hash)

}