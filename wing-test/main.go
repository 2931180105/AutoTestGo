package main

import (
	"github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	test_case "github.com/mockyz/AutoTestGo/wing-test/test-case"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"time"
)

var sdk = goSdk.NewOntologySdk()
var assetMap = map[string]string{
	"162654622062ea6762647babab2462dffa712837": "f44504f293aef233f0978f08b50de78ac482b9ba",
	"423dbe4f99c5958d01d4b2993318f28f1baedb02": "c1f7446cf1e20c90699520ff5fe43a92a17d9954",
	"e37e6a5283682b383ecebb712e68cf1d36239d1b": "1a3d6116cf5ebd67f5be0f5216e43d037d509a09",
	"5ae2c02bc24138820e336e18f4aa4eb13d25a84d": "0bbb55ae5bfb2e7a4c6021d1911f68175f785cc7",
	"e5059eac5fb89ae1c8976ea3c3cba7e0786165a9": "86246c2d1cee5ddd71dea2b355a9596ff62fcbb9",
	"789c0aa3277430c10d0dc5b1924ece0060bfa100": "dd9580a013f2c515a846d56ec71f8ae4edce98eb",
	"37c99c10037e000149541d925f61c4e1fa77a60d": "ab6ea82ae231b9140ebbe7d7e61c77f64e5a3979",
	"a1ad38537dd0d9240ff40e56be654434303e858a": "bdb2e7f64ddc5b5cf1d289cec60b5562b74d9410",
	"cb09fdb1745510ee07dfc9d8e6dc69741578513a": "939c0747c3fdc57493df00dd9fabac6f79b491c0",
	"d7edfcd57d784c4a19c991cdb9b391e6b6442fd5": "276bfedfbaae07569e3d77715880645b9434a4af",
	"4eb4dc1179b8192243273017e8729162134c1af4": "9d64665c20363d3204482f95875d8ebbdd39fb56",
	"d3ae6627969363d18517208ba59f540400d7dcaf": "3e8c1a687b7cb9c5d796f5d5e576f804e770dac9",
}
var decimalMap = map[string]uint64{
	"f44504f293aef233f0978f08b50de78ac482b9ba": 8,
	"c1f7446cf1e20c90699520ff5fe43a92a17d9954": 9,
	"1a3d6116cf5ebd67f5be0f5216e43d037d509a09": 8,
	"0bbb55ae5bfb2e7a4c6021d1911f68175f785cc7": 6,
	"86246c2d1cee5ddd71dea2b355a9596ff62fcbb9": 9,
	"dd9580a013f2c515a846d56ec71f8ae4edce98eb": 18,
	"ab6ea82ae231b9140ebbe7d7e61c77f64e5a3979": 18,
	"bdb2e7f64ddc5b5cf1d289cec60b5562b74d9410": 6,
	"939c0747c3fdc57493df00dd9fabac6f79b491c0": 18,
	"276bfedfbaae07569e3d77715880645b9434a4af": 8,
	"9d64665c20363d3204482f95875d8ebbdd39fb56": 18,
	"3e8c1a687b7cb9c5d796f5d5e576f804e770dac9": 18,
}

func main() {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "wing-test/config_prv.json"
	//configPath := "wing-test/config_testnet.json"

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
	otoken.TransferAllTestToken(cfg, account, sdk, "AJExyJ6rk3mdExs3F77jPpZ5xgKSVmYvLA")
	return
	//test wbtc borrow rate ,use defult account and use address
	accounts := Utils.GetAccounts2(3000,1)
	market,_ := test_case.NewMarkets(cfg,accounts[0],sdk,cfg.FWBTC)
	//market.TestBorrowRateByBlock()
	market.TestBorrowRateByTime()

	//market.TestBorrowRateByBlock2Addr("AG4pZwKa9cr8ca7PED7FqzUfcwnrQ2N26w")
	market.TestBorrowRateByTime2Addr("AG4pZwKa9cr8ca7PED7FqzUfcwnrQ2N26w")

	market.WingSpeed4SuppluyTest("AG4pZwKa9cr8ca7PED7FqzUfcwnrQ2N26w")
	//AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p
	//AJkQo3Fo7JKxtrKZPqYJQuh9cXH38w7rVt
	//OToken.DelegateToProxyAllTestToken(cfg, account, sdk)
	//OToken.OTokenTransfer(cfg, account, sdk, "ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ", cfg.ODAI)
	return
	otoken.GenerateAccountsToken(cfg, account, sdk)
	//AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p
	//OToken.BalanceOfAllToken(cfg, sdk, account.Address.ToBase58())
	//bacthTest(cfg, sdk)
}



func bacthTest(cfg *config.Config, genSdk *goSdk.OntologySdk) {
	exitChan := make(chan int)
	//wallet, _ := genSdk.CreateWallet("tmp.dat")
	var txNum = cfg.TxNum * cfg.TxFactor
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	startTestTime := time.Now().UnixNano() / 1e6
	accounts :=Utils.GetAccounts2(0,4000)
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
					//Utils.NewAccountToDb(wallet)
					//BatchBorrow(cfg,genSdk,accounts[nonce])
					BatchSupply(cfg,genSdk,accounts[nonce])
					now := time.Now().UnixNano() / 1e6 // ms
					diff := sentNum - (now-startTime)/1e3*tpsPerRoutine
					if now > startTime && diff > 0 {
						sleepTime := time.Duration(diff*1000/tpsPerRoutine) * time.Millisecond
						time.Sleep(sleepTime)
						log.Infof("sleep %d ms", sleepTime.Nanoseconds()/1e6)
					}
				}
				nonce++
				log.Infof("run nonce :%d",nonce)
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
