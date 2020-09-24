package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"testing"
)

//TODO： init/wing token set gov address
func TestWingGovInit(t *testing.T) {
	cfg, account, sdk := Utils.GetTestConfig()
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
	cfg, account, sdk := Utils.GetTestConfig()
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
	cfg, account, sdk := Utils.GetPrvConfig()

	hash1, err := sdk.SendTransaction(WingGovMethod.Update_support_token(cfg, account, sdk, "ETH", cfg.OETH))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	//WingGovMethod.AddSupportTokenAndSend(cfg, account, sdk, "WING", cfg.GovToken)
	//WingGovMethod.AddSupportTokenAndSend(cfg, account, sdk,"ETH",cfg.OETH)
	//WingGovMethod.AddSupportTokenAndSend(cfg, account, sdk,"DAI",cfg.ODAI)\

}

//set decimals
func TestSet_token_decimals(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	//
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "ONT", 0)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "ONTd", 9)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "WBTC", 8)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "renBTC", 8)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "USDC", 6)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "ETH", 18)
	//WingGovMethod.Set_token_decimals(cfg, account, sdk, "ETH9", 9)
	//WingGovMethod.Set_token_decimals(cfg, account, sdk, "DAI", 18)
	WingGovMethod.Set_token_decimals(cfg, account, sdk, "WING", 9)
}

//set decimals
func TestGet_token_decimals(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	//WingGovMethod.Set_token_decimals(cfg, account, sdk, "ONTd", 9)        b"get_token_decimals" => {
	WingGovMethod.Get_token_decimals(cfg, account, sdk)
}

//step6： 调整两个池子权重： 现设1再设0
func TestStep006(t *testing.T) {
	cfg, account, sdk := Utils.GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 1))
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash2, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.ZeroPool, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash2.ToHexString())
}

//ba7729f31c1c4b0043df8e9189ea0b26dd1653fe
func TestUpdateComptroller(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	//oldAddr := "ea7660cb741caf5e311dfce1e57aa849c4b03069"
	//newAddr := "f1e89adb8053af4fa96f588a30eceb7bdb7a71b8"
	WingGovMethod.MigrateComptroller(cfg, account, sdk, cfg.Comptroller, "fb320bbd7cc747eb57e1bbee64a7600d26dd0667")
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, "fb320bbd7cc747eb57e1bbee64a7600d26dd0667")
}

//Add support token
func TestAddAllSuuportToken(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	WingGovMethod.AddAllSupportToken(cfg, account, sdk)
	WingGovMethod.Get_support_token(cfg, account, sdk)
}

func TestRegsiterComptroller(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, cfg.Comptroller))
	//hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, "8c729377e714ff5013d74e309dad25fbdf1bf889"))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)

}
func TestTmp(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
	//
	////设置comptroller权重值为1
	//hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 1))
	//if err != nil {
	//	log.Errorf("send  tx failed, err: %s********", err)
	//	return
	//}
	//Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestUpdatePoolWeight(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, "a7af1dfc7539db719fc2215577c6c33474016820", 1))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}