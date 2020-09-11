package wingGov

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"testing"
	"time"
)

var sdk = goSdk.NewOntologySdk()
var account *goSdk.Account
var cfg *config.Config

//
func TestMain(m *testing.M) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "../config_testnet.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ = wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[3])
	sdk.SetDefaultClient(rpcClient)
	m.Run()
}

func TestWingGovUnboundToken(t *testing.T) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "../config_testnet.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ = wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[3])
	sdk.SetDefaultClient(rpcClient)
	hash1, err := sdk.SendTransaction(UnboundToken(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovUnboundTokenToPool(t *testing.T) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "../config_testnet.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ = wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[3])
	sdk.SetDefaultClient(rpcClient)
	hash1, err := sdk.SendTransaction(UnboundToPool(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovUpdatePoolWeight(t *testing.T) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "../config_testnet.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ = wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[3])
	sdk.SetDefaultClient(rpcClient)
	hash1, err := sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestQuery_unbound_to_pool(t *testing.T) {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "../config_testnet.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ = wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[3])
	sdk.SetDefaultClient(rpcClient)
	Query_unbound_to_pool_count(cfg, account, sdk)
	Query_unbound_to_pool(cfg, account, sdk)
}
