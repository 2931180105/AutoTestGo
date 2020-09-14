package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	"testing"
	"time"
)

func TestUpdateComptroller(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	WingGovMethod.MigrateComptroller(cfg, account, sdk, "e58749f563f14488fae547e35ad2120cbf922eba")
}

func TestRegsiterComptroller(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.RegisterPool(cfg, account, sdk, "1f99f0a0bae1c3df3ce6cc1adf975767bdd2dfa7"))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	WingGovMethod.QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
}
