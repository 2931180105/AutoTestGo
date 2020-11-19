package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"testing"
)

func TestApproveAndBorrow(t *testing.T){
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
	market,_ := NewMarkets(cfg,account,sdk,cfg.WING)
	hash0,err :=market.NeoVMApprove(market.GetSigner().Address,market.Comptroller.GetAddr(),utils.ToIntByPrecise("1",9))
	if err != nil {
		log.Errorf("NeoVMApprove: %s", err)
	}
	utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash0)
	ftokenAddressList, err := market.Comptroller.AllMarkets()
	for _, ftokenAddress := range ftokenAddressList {
		//market.SetAddr(ftokenAddress)
		mk,_ := NewMarkets(cfg,account,sdk,ftokenAddress.ToHexString())
		//oAddr :=market.Underlying()
		//underlyToken,_ := NewMarkets(cfg,account,sdk,oAddr.ToHexString())
		//hash1,err := underlyToken.NeoVMApprove(market.GetSigner().Address,market.Comptroller.GetAddr(),utils.ToIntByPrecise("1",2))
		//if err != nil {
		//	log.Errorf("NeoVMApprove: %s", err)
		//}
		//utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash1)
		hash2,err :=mk.Borrow(market.GetSigner().Address,utils.ToIntByPrecise("1",2),true)
		if err != nil {
			log.Errorf("Mint: %s", err)
		}
		utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash2)
	}
}


func TestRepayBorrow(t *testing.T){
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
	market,_ := NewMarkets(cfg,account,sdk,cfg.WING)
	ftokenAddressList, err := market.Comptroller.AllMarkets()
	for _, ftokenAddress := range ftokenAddressList {
		//market.SetAddr(ftokenAddress)
		mk,_ := NewMarkets(cfg,account,sdk,ftokenAddress.ToHexString())
		//oAddr :=market.Underlying()
		//underlyToken,_ := NewMarkets(cfg,account,sdk,oAddr.ToHexString())
		//hash1,err := underlyToken.NeoVMApprove(market.GetSigner().Address,market.Comptroller.GetAddr(),utils.ToIntByPrecise("1",2))
		//if err != nil {
		//	log.Errorf("NeoVMApprove: %s", err)
		//}
		//utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash1)
		hash2,err :=mk.RepayBorrow(market.GetSigner().Address,utils.ToIntByPrecise("1",2))
		if err != nil {
			log.Errorf("Mint: %s", err)
		}
		utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash2)
	}
}
