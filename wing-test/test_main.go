package main

import (
	OToken "github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGov "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"time"
)

var sdk = goSdk.NewOntologySdk()

func main() {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	//WingGov.WingGovMigrate(cfg, account, sdk)
	//Compound.BalanceOfOToken(cfg,account, sdk, "BalanceOfOToken")
	OToken.TransferAllTestToken(cfg, account, sdk, "ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ")
	//accts := Utils.GenerateAccounts(cfg, account, sdk)
	//Utils.BatchStaking(cfg, account, sdk, accts)
	//reslut2 := Utils.DeployContractWingToken(cfg, account, sdk)
	//log.Infof("hash", reslut2.ToHexString())
	//WingGov.QueryPoolByAddress(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	if false {
		hash1, err := sdk.SendTransaction(WingGov.Add_support_token(cfg, account, sdk))
		if err != nil {
			log.Errorf("send  tx failed, err: %s********", err)
			return
		}
		time.Sleep(time.Second * 3)
		Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
		//
		//hash1, err = sdk.SendTransaction(Utils.Set_pool_operator(cfg, account, sdk))
		//if err != nil {
		//	log.Errorf("send  tx failed, err: %s********", err)
		//	return
		//}
		//time.Sleep(time.Second * 3)
		//Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	}
	//Compound.ApproveOToken(cfg, account, sdk, cfg.FETH)
	//if false {
	//	hash1, err := sdk.SendTransaction(Compound.OUSDTTokenTransfer(cfg, account, sdk))
	//	if err != nil {
	//		log.Errorf("send  tx failed, err: %s********", err)
	//		return
	//	}
	//	time.Sleep(time.Second * 3)
	//	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	//}
	//Utils.QueryPoolCount(cfg, account, sdk)

	//Utils.BatchUnStakeing(cfg, account, sdk, accts)
}
func deployContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	//reslut := Utils.DeployContractOracle(cfg, account, sdk)
	//log.Infof("hash",reslut.ToHexString())
	reslut1 := WingGov.DeployContractWingToken(cfg, account, sdk)
	log.Infof("hash", reslut1.ToHexString())
	reslut2 := WingGov.DeployContractWingGov(cfg, account, sdk)
	log.Infof("hash", reslut2.ToHexString())
	reslut3 := WingGov.DeployContractProfit(cfg, account, sdk)
	log.Infof("hash", reslut3.ToHexString())
	reslut4 := WingGov.DeployContractOracle(cfg, account, sdk)
	log.Infof("hash", reslut4.ToHexString())
	//Utils.DeployContractFlash(cfg, account, sdk)
}

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
