package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/ftoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

func TestReadFile(t *testing.T) {
	fileContent, err := ioutil.ReadFile("./hash.txt")
	if err != nil {
		log.Errorf("NewConfigFromFile: %s", err)
	}
	log.Infof("file :%v", string(fileContent))
}
func TestWingSpeeds(t *testing.T) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	syn := new(sync.WaitGroup)
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
	market,_ := NewMarkets(cfg,account,sdk,cfg.FWBTC)
	ftokenAddressList, err := market.Comptroller.AllMarkets()
	//TODO:markets compare user addr
	for _, ftokenAddress := range ftokenAddressList {
		syn.Add(1)
		market,_ := NewMarkets(cfg,account,sdk,ftokenAddress.ToHexString())
		//time.Sleep(time.Second)
		//go market.WingSpeed4BorrowTestNewByMarketAddr(ftokenAddress, "AXhxR1NDWCABFn6MQmmpVy9BSAgm7bf15D", syn)
		//time.Sleep(time.Second)
		go market.WingSpeed4SuppluyTestNew(ftokenAddress, "ASQmMksvxcC8rbBGbsChUEwD7guXFH3riY", syn)
	}
	syn.Wait()
}
//func TestBorrowRateByTime(t *testing.T) {
//	syn := new(sync.WaitGroup)
//	for i := 0; i < len(MarketNames); i++ {
//		syn.Add(1)
//		go TestBorrowRateNew(MarketNames[i], syn)
//	}
//	syn.Wait()
//}


/*
只有抵押借款才能分wing
根据新的需求测试wing的分配（修改之前的代码）
将测试结果保存到文件中
1.获取所有的市场
ftokenAddressList, err := comptroller.GetAllMarkets(genSdk, cfg.Comptroller)
2.根据市场的名称匹配对应的ftoken

*/
func TestWingSpeedTest(t *testing.T){
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
	market,_ := NewMarkets(cfg,account,sdk,cfg.WBTC)
	//ftokenAddressList, err := market.Comptroller.AllMarkets()
	//TODO:markets compare user addr
	//for _, ftokenAddress := range ftokenAddressList {
	//	market.SetAddr(ftokenAddress)
	//	marketName,err := market.Name()
	//	if err !=nil{
	//		log.Errorf("marketName err :%v ",err)
	//	}
	//	log.Infof("market.Name11: %s",marketName)
	//}

	market.WingSpeed4SuppluyTestByName("ONG","AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p")
}

func TestMintBorrow(t *testing.T){
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
	underlyToken := new(ftoken.FlashToken)
	//market.
	hash0,err :=market.NeoVMApprove(market.GetSigner().Address,market.Comptroller.GetAddr(),utils.ToIntByPrecise("1",10))
	if err != nil {
		log.Errorf("NeoVMApprove: %s", err)
	}
	utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash0)
	ftokenAddressList, err := market.Comptroller.AllMarkets()
	for _, ftokenAddress := range ftokenAddressList {
		market.SetAddr(ftokenAddress)
		mn, err :=market.Name()
		if err !=nil{
			return
		}
		if strings.Contains(mn, "USDT") {
			oAddr :=market.Underlying()
			underlyToken,_ = NewMarkets(cfg,account,sdk,oAddr.ToHexString())
			break
		}
	}
	underlyToken.NeoVMBalanceOf(underlyToken.GetSigner().Address)
	//hash,err :=market.Mint(market.GetSigner().Address,utils.ToIntByPrecise("1",8))
	hash,err :=market.Borrow(market.GetSigner().Address,utils.ToIntByPrecise("1",8),true)
	if err != nil {
		log.Errorf("Mint: %s", err)
	}
	utils.PrintSmartEventByHash_Ont(market.GetGoSdk(),hash)
}
