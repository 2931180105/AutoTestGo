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
	configPath := "wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)

	//Utils.GetUnderlyingName(cfg, account, sdk)

	//Utils.DeployContractFlash(cfg, account, sdk)

	// run when first token
	//Utils.InterestRateInit(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.ComptrollerInit(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetPriceOracle(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetWingAddr(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetCloseFactor(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetInsuranceRepayFactor(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetMaxAssets(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetLiquidationIncentive(cfg, account, sdk)
	//time.Sleep(time.Second * 3)
	//Utils.SetWingRate(cfg, account, sdk)
	//time.Sleep(time.Second * 3)

	Utils.FTokenInit(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SetReserveFactor(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SetInsuranceFactor(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.FTokenInsuranceInit(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SetMarketAddr(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SetInsuranceAddr(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SupportMarket(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.SetCollateralFactor(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	Utils.AddWingMarkets(cfg, account, sdk)
	time.Sleep(time.Second * 3)

	//Utils.Get_support_token(cfg, account, sdk)
	//Utils.Get_exchange_rate(cfg, account, sdk)
	//Utils.WingGovMigrate(cfg, account, sdk)
	//Utils.GovTokenBalanceOf(cfg,"AaLXuZ7rCME3QyL3aXWst5socmj4jw1vjG",sdk)
	//accts := Utils.GenerateAccounts(cfg, account, sdk)
	//Utils.BatchStaking(cfg, account, sdk, accts)
	//Utils.Get_global_address(cfg, account, sdk)
	//reslut2 := Utils.DeployContractWingGov(cfg, account, sdk)
	//log.Infof("hash", reslut2.ToHexString())
	//time.Sleep(time.Second * 3)
	if false {
		hash1, err := sdk.SendTransaction(Utils.SetFFactor(cfg, account, sdk))
		if err != nil {
			log.Errorf("send  tx failed, err: %s********", err)
		}
		time.Sleep(time.Second * 3)
		Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	}
	//Utils.
	//Utils.BatchUnStakeing(cfg, account, sdk, accts)
}
func deployContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	//reslut := Utils.DeployContractOracle(cfg, account, sdk)
	//log.Infof("hash",reslut.ToHexString())
	reslut1 := Utils.DeployContractWingToken(cfg, account, sdk)
	log.Infof("hash", reslut1.ToHexString())
	reslut2 := Utils.DeployContractWingGov(cfg, account, sdk)
	log.Infof("hash", reslut2.ToHexString())
	reslut3 := Utils.DeployContractProfit(cfg, account, sdk)
	log.Infof("hash", reslut3.ToHexString())
	reslut4 := Utils.DeployContractOracle(cfg, account, sdk)
	log.Infof("hash", reslut4.ToHexString())
	Utils.DeployContractFlash(cfg, account, sdk)
}
