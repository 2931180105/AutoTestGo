package test

import (
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"github.com/ontio/ontology/common/log"
	"time"

	"testing"
)

//Migrate gov
func TestMigrateGov(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)
	newGovString := WingGovMethod.WingGovMigrate(cfg, account, sdk, "/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/contract/testnet/wing_dao_contracts.wasm_new.str")
	//	set wing token gov
	hash1, err := sdk.SendTransaction(WingGovMethod.WingTokenSetGov(cfg, account, sdk, newGovString))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("WingTokenSetGov hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	//	TODO:update config file
}

// upgrade zero pool

//Migrate Zero pool
func TestMigrateZeroPool(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	zeroPoolPath := "../../contract/testnet/zero_pool_new2.wasm.str"
	newContractString := WingGovMethod.ContractMigrate(cfg, account, sdk, cfg.ZeroPool, zeroPoolPath)
	time.Sleep(time.Second * 3)
	log.Infof("new zero pool2 : %s", newContractString)
	//	update pool address
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolAddress(cfg, account, sdk, newContractString, cfg.ZeroPool))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("UpdatePoolAddress hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, newContractString)

	//	TODO:update config file
}

//	WingGovMethod.ZeroPoolInit(cfg, account, sdk)

//Deploy zero pool Todo : finish deploy
func TestDeployZeroPool(t *testing.T) {
	cfg, account, sdk := GetPrvConfig()
	wasmFile := "../../contract/testnet/zero_pool.wasm.str"
	zeroPoolAddr := WingGovMethod.DeployContractt(cfg, account, sdk, wasmFile)
	//WingGovMethod.QueryPoolByAddress(cfg, account, sdk, zeroPoolAddr)
	WingGovMethod.ZeroPoolInit(cfg, account, sdk, zeroPoolAddr)
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, zeroPoolAddr))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("SendTransaction RegisterPool hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, zeroPoolAddr)
	//	TODO:update config file
}

//Deploy zero pool Todo : finish deploy
func TestDeployGov(t *testing.T) {
	cfg, account, sdk := GetPrvConfig()
	wasmFile := "../../contract/testnet/wing_dao_contracts_old.wasm.str"
	WingGovAddr := WingGovMethod.DeployContractt(cfg, account, sdk, wasmFile)
	time.Sleep(time.Second * 3)
	//WingGovMethod.QueryPoolByAddress(cfg, account, sdk, zeroPoolAddr)
	hash1, err := sdk.SendTransaction(WingGovMethod.WingGovInit(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("SendTransaction  hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	_, err = sdk.SendTransaction(WingGovMethod.WingTokenSetGov(cfg, account, sdk, WingGovAddr))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	//WingGovMethod.QueryPoolByAddress(cfg, accouWingGovMethod.WingGovInit(cfg, account, sdk)
	//	TODO:update config file
}

//Deploy zero pool Todo : finish deploy
func TestDeployProfit(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	wasmFile := "../../contract/testnet/profit.wasm.str"
	profitAddr := WingGovMethod.DeployContractt(cfg, account, sdk, wasmFile)
	//WingGovMethod.QueryPoolByAddress(cfg, account, sdk, zeroPoolAddr)

	hash1, err := sdk.SendTransaction(WingGovMethod.WingProfitInit2(cfg, account, sdk, profitAddr))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("SendTransaction WingProfitInit2 hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	//	TODO:update config file
}
