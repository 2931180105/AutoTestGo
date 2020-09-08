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
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	//deployContract(cfg,account,sdk)
	//InitContract(cfg,account,sdk)
	//Utils.GovTokenBalanceOf(cfg,account,sdk)
	Utils.DeployContractOracle(cfg, account, sdk)
	//Utils.Query_unbound_to_pool(cfg,account,sdk)
	//Utils.GovTokenSetGov(cfg,account,sdk)
	// deposit , withDraw
	tx := Utils.OracleInit(cfg, account, sdk)
	hash, err := sdk.SendTransaction(tx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
}
func deployContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	//reslut := Utils.DeployContractOracle(cfg, account, sdk)
	//log.Infof("hash",reslut.ToHexString())
	reslut2 := Utils.DeployContractGov(cfg, account, sdk)
	log.Infof("hash", reslut2.ToHexString())
	reslut3 := Utils.DeployContractProfit(cfg, account, sdk)
	log.Infof("hash", reslut3.ToHexString())
	reslut4 := Utils.DeployContractOracle(cfg, account, sdk)
	log.Infof("hash", reslut4.ToHexString())

}
func InitContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	txs := Utils.ContractInit(cfg, account, sdk)
	//tx := Utils.Set_exchange_rate(cfg, account, sdk)
	for tx := 0; tx < len(txs); tx++ {
		hash, err := genSdk.SendTransaction(txs[tx])
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
		}
		time.Sleep(time.Second * 3)
		Utils.PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
	}

}
