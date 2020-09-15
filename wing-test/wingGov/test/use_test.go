package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"testing"
	"time"
)

//TODO： init/wing token set gov address
func TestWingGovInit(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.WingGovInit(cfg, account, sdk, cfg.WingGov))
	if err != nil {
		log.Errorf("send WingGovInit tx failed, err: %s********", err)
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(WingGovMethod.WingTokenSetGov(cfg, account, sdk, cfg.WingGov))
	if err != nil {
		log.Errorf("send WingTokenSetGov tx failed, err: %s********", err)
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

//升级
func TestStepGovContractUpgrade(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	contractPath := "/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/contract/testnet/wing_dao_contracts_new2.wasm.str" //todo: 放合约地址
	newGovString := WingGovMethod.WingGovMigrate(cfg, account, sdk, contractPath)                                                   //todo: 不同钱包
	hash1, err := sdk.SendTransaction(WingGovMethod.WingTokenSetGov(cfg, account, sdk, newGovString))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("WingTokenSetGov hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	result := WingGovMethod.GetGovAddress(cfg, sdk)
	log.Infof("gov address:%s", result.Result)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "ONT", 0)
	//	TODO:update config file
}

//0.add support token
func TestStep_AddSupportToken(t *testing.T) {
	//todo：add support token
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.AddAllSupportToken(cfg, account, sdk)
}

//set decimals
func TestSet_token_decimals(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "ONTd", 9)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "WBTC", 8)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "renBTC", 8)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "USDC", 6)
}

//step6： 调整两个池子权重： 现设1再设0
func TestStep006(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 1))
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash2, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.ZeroPool, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash2.ToHexString())
}

//03d830df793d80137343731f799d5be40637a310
func TestUpdateComptroller(t *testing.T) {
	cfg, account, sdk := GetMainConfig()
	WingGovMethod.MigrateComptroller(cfg, account, sdk, "efd78c612b66c690a59721b7bdd1c0e090c52ec4")
}

//Add support token
func TestAddAllSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.AddAllSupportToken(cfg, account, sdk)
	WingGovMethod.Get_support_token(cfg, account, sdk)
}

func TestRegsiterComptroller(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, cfg.Comptroller))
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, "d034792f80deeacd983dc257d29784ea71a1d5ec"))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, "d034792f80deeacd983dc257d29784ea71a1d5ec")

}
func TestTmp(t *testing.T) {
	cfg, account, sdk := GetTestConfig()

	//设置comptroller权重值为0
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestUpdatePoolWeight(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 1))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	hash1, err = sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.ZeroPool, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
