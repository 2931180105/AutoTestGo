package main

import (
	OToken "github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
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
	configPath := "wing-test/config_testnet.json"
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
	rpcClient.SetAddress(cfg.Rpc[1])
	sdk.SetDefaultClient(rpcClient)
	//WingGov.BatchStaking(cfg, account, sdk,Utils.GetAccounts(cfg))
	//time.Sleep(time.Second*10)
	//WingGov.BatchUnStaking(cfg, account, sdk,Utils.GetAccounts(cfg))
	//
	//WingGov.WingGovMigrate(cfg, account, sdk)
	//WingGov.Get_admin_address(cfg, account, sdk)
	//WingGov.DeployContractWingGov(cfg, account, sdk)
	//WingGov.WingTokenGetGovAddr(cfg, sdk)

	//AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p
	//AJkQo3Fo7JKxtrKZPqYJQuh9cXH38w7rVt
	//OToken.OTokenDelegateToProxy(cfg, account, sdk, cfg.ODAI)
	//OToken.OTokenTransfer(cfg, account, sdk, "ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ", cfg.ODAI)
	OToken.DelegateToProxyAllTestToken(cfg, account, sdk)
	OToken.TransferAllTestToken(cfg, account, sdk, "AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p")
}

//func deployContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
//	//reslut := Utils.DeployContractOracle(cfg, account, sdk)
//	//log.Infof("hash",reslut.ToHexString())
//	reslut1 := DeployContractWingToken(cfg, account, sdk)
//	log.Infof("hash", reslut1.ToHexString())
//	reslut2 := DeployContractWingGov(cfg, account, sdk)
//	log.Infof("hash", reslut2.ToHexString())
//	reslut3 := DeployContractProfit(cfg, account, sdk)
//	log.Infof("hash", reslut3.ToHexString())
//	reslut4 := DeployContractOracle(cfg, account, sdk)
//	log.Infof("hash", reslut4.ToHexString())
//	//Utils.DeployContractFlash(cfg, account, sdk)
//}

func bacthTest(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	exitChan := make(chan int)
	wallet, _ := genSdk.CreateWallet("tmp.dat")
	var txNum = cfg.TxNum * cfg.TxFactor
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	startTestTime := time.Now().UnixNano() / 1e6
	for i := uint(0); i < cfg.RoutineNum; i++ {
		//rand.Int()%len(cfg.Rpc)随机获取一个接口
		//client := NewRpcClient(cfg.Rpc[rand.Int()%len(cfg.Rpc)])
		go func(nonce uint32, routineIndex uint) {
			startTime := time.Now().UnixNano() / 1e6 // ms
			sentNum := int64(0)
			for j := uint(0); j < txNumPerRoutine; j++ {
				if cfg.SendTx {
					sentNum++
					//TODO : change to your batch func
					Utils.NewAccountToDb(wallet)
					now := time.Now().UnixNano() / 1e6 // ms
					diff := sentNum - (now-startTime)/1e3*tpsPerRoutine
					if now > startTime && diff > 0 {
						sleepTime := time.Duration(diff*1000/tpsPerRoutine) * time.Millisecond
						time.Sleep(sleepTime)
						log.Infof("sleep %d ms", sleepTime.Nanoseconds()/1e6)
					}
				}
				nonce++
			}
			exitChan <- 1
		}(uint32(txNumPerRoutine*i)+cfg.StartNonce, i)
	}
	for i := uint(0); i < cfg.RoutineNum; i++ {
		<-exitChan
	}
	endTestTime := time.Now().UnixNano() / 1e6
	log.Infof("send tps is %f", float64(txNum*1000)/float64(endTestTime-startTestTime))
}
