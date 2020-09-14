package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"testing"
	"time"
)

//03d830df793d80137343731f799d5be40637a310
func TestUpdateComptroller(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.MigrateComptroller(cfg, account, sdk, "a5a9e4131aa80531acdfed4e24962cb091cae3a7")
}

//Add support token
func TestAddAllSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.AddAllSuuportToken(cfg, account, sdk)
	WingGovMethod.Get_support_token(cfg, account, sdk)

}

func TestRegsiterComptroller(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, cfg.Comptroller))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
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
