package wingGov

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common/log"
	"testing"
	"time"
)

var sdk = goSdk.NewOntologySdk()
var account *goSdk.Account
var cfg *config.Config

func TestWingGovUnboundToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(UnboundToken(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovUnboundTokenToPool(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(UnboundToPool(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovUpdatePoolWeight(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestQuery_unbound_to_pool(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	Query_unbound_to_pool_count(cfg, account, sdk)
	Query_unbound_to_pool(cfg, account, sdk)
}

func TestWingGovRegisterPool(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(RegisterPool(cfg, account, sdk, cfg.Comptroller))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
}
