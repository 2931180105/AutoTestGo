package wingGov

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
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
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

//
//"a22fde0b0284d4e850558b4a5f95f2e2d36654d2",
//"5de3ac0c8863db961f2c3d3a832d5063dda51777",
//"d69e51230c0f096888c259c6f29a1f5c7cf3542a",
//"386bc44661fef273e6627261736ebc2944273c62",

func TestWingGovUpdatePoolWeight(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, "1f99f0a0bae1c3df3ce6cc1adf975767bdd2dfa7", 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	hash1, err = sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, "5de3ac0c8863db961f2c3d3a832d5063dda51777", 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	hash1, err = sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, "d69e51230c0f096888c259c6f29a1f5c7cf3542a", 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	hash1, err = sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, "386bc44661fef273e6627261736ebc2944273c62", 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}

	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovRegisterPool(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(RegisterPool(cfg, account, sdk, "cfg.Comptroller"))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	QueryPoolByAddress(cfg, account, sdk, cfg.Comptroller)
}

func TestWingGovRegisterPoolToAddress(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	acct := Utils.GetAccounts(cfg)
	hash1, err := sdk.SendTransaction(RegisterPoolToAddress(cfg, account, sdk, acct[0].Address))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	QueryPoolByAddress(cfg, account, sdk, acct[0].Address.ToHexString())
}
func TestStakingMore(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)
	BatchStaking(cfg, account, sdk, accounts)

}
func TestAddSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(Add_support_token(cfg, account, sdk, "renBTC", cfg.RENBTC))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestUpdateSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(Update_support_token(cfg, account, sdk, "renBTC", cfg.RENBTC))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestOracle(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	OraclePutUnderlyingPrice(cfg, account, sdk)
	OracleGetUnderlyingPrice(cfg, account, sdk, "ONT")
}

//GetProductPools

func TestGetProductPools(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	GetProductPools(cfg, account, sdk)
}

func TestUpdatePoolWeight(t *testing.T) {
	cfg, account, sdk := GetTestConfig()

	hash1, err := sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, cfg.Comptroller, 0))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	hash1, err = sdk.SendTransaction(UpdatePoolWeight(cfg, account, sdk, cfg.ZeroPool, 1))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
