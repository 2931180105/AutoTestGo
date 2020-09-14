package wingGov

import (
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"testing"
	"time"
)

func GetTestConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "../config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("../wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func TestDeployContractOETHToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//fmt.Println(cfg, account, sdk)
	QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
	//DeployContractODAIToken(cfg, account, sdk)
}
func TestGet_support_token(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//fmt.Println(cfg, account, sdk)
	Get_support_token(cfg, account, sdk)
	//DeployContractODAIToken(cfg, account, sdk)
}

//1599814800

//部署后init的所有合约
func TestInitWingToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(GovTokenInit(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(SetGovTokenAddres(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

}
