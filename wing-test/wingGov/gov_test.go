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
	Query_unbound_to_pool_count(cfg, sdk)
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
func TestStakingMore(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)

	BatchStaking(cfg, account, sdk, accounts)

}
func TestAddSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)

	hash1, err := sdk.SendTransaction(Add_support_token(cfg, account, sdk, "DAI", cfg.ODAI))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestMigrateGov(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)

	//wing-test/contract/private/wing_dao_contracts_new.wasm.str
	WingGovMigrate(cfg, account, sdk, "/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/contract/private/wing_dao_contracts_new.wasm.str")

}
func TestOracle(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	OraclePutUnderlyingPrice(cfg, account, sdk)
	OracleGetUnderlyingPrice(cfg, account, sdk, "ONT")

}
