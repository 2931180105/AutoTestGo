package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"io/ioutil"
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
		go market.TestWingSpeeds4Borrow(ftokenAddress, "AXhxR1NDWCABFn6MQmmpVy9BSAgm7bf15D", syn)
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
	market,_ := NewMarkets(cfg,account,sdk,cfg.FWBTC)
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
	market.WingSpeed4SuppluyTestByName("SUSD","ASQmMksvxcC8rbBGbsChUEwD7guXFH3riY")
}
