package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"testing"
)

/**
1.更新治理合约
2. 需要合约
*/
func TestStep00_GovContractUpgrade(t *testing.T) {
	cfg, account, sdk := GetMainConfig()
	contractPath := "" //todo: 放合约地址
	newGovString := WingGovMethod.WingGovMigrate(cfg, account, sdk, contractPath)
	hash1, err := sdk.SendTransaction(WingGovMethod.WingTokenSetGov(cfg, account, sdk, newGovString))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	log.Infof("WingTokenSetGov hash : %s", hash1.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	reslut := WingGovMethod.GetGovAddress(cfg, sdk)
	log.Infof("gov address:%s", reslut.Result)
	//	TODO:update config file
}

//0.add support token
func TestStep01_AddSupportToken(t *testing.T) {
	//todo：add support token
	cfg, account, sdk := GetMainConfig()
	WingGovMethod.AddAllSupportToken(cfg, account, sdk)
}

//1.register comptroller set weight =0
func TestStep02_RegisterComptroller(t *testing.T) {
	cfg, account, sdk := GetMainConfig()
	//todo: 修改flash合约地址
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, cfg.Comptroller))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
	//设置comptroller权重值为0
	hash1, err = sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

//2.合约设置借贷宝比例： 暂时不知道怎么做 ---
func TestStep02(t *testing.T) {
	//todo： 王成处理或者提供脚本
}

//Step3： flash pool网页上线
//step4： zero页面更新
//step5： 结算
func TestStep05(t *testing.T) {
	cfg, account, sdk := GetMainConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UnboundToken(cfg, account, sdk))
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash2, err := sdk.SendTransaction(WingGovMethod.UnboundToPool(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash2.ToHexString())

}

//step6： 调整两个池子权重： 现设1再设0
func TestStep06(t *testing.T) {
	cfg, account, sdk := GetMainConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 1))
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash2, err := sdk.SendTransaction(WingGovMethod.UpdatePoolWeight(cfg, account, sdk, cfg.ZeroPool, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash2.ToHexString())
}

//step7:web端按钮跳转-To-Supply 生效/To Supply测试： TODO： TO Supply 测试
//Step8: 启动结算机器人 # 9/15结束

//#9/16
