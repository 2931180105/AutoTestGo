package wingGov

import (
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common/log"
	"testing"
	"time"
)

func TestWingTokenSetGov(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingTokenSetGov(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestDep(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	wasmFile := "../contract/private/profit.wasm.str"
	DeployContractProfit(cfg, account, sdk, wasmFile)
}

//部署后init的所有合约
func TestIint(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingTokenSetGov(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

	hash1, err = sdk.SendTransaction(WingGovInit(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(RegisterPool(cfg, account, sdk, cfg.ZeroPool))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	QueryPoolByAddress(cfg, account, sdk, cfg.ZeroPool)
	hash1, err = sdk.SendTransaction(WingProfitInit(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

}

func TestStaking(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accts := Utils.GenerateAccounts(cfg, account, sdk)
	//accts := dbHelper.QueryAccountFromDb(0,10)
	BatchStaking(cfg, account, sdk, accts)
}

func TestUnStaking(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accounts := dbHelper.QueryAccountFromDb(21, cfg.AccountNum)
	//accts := Utils.GenerateAccounts(cfg, account, sdk )
	//accts := dbHelper.QueryAccountFromDb(0,10)
	BatchUnStaking(cfg, account, sdk, accounts)
}

func TestGetUndlying(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accts := Utils.GenerateAccounts(cfg, account, sdk)
	//accts := dbHelper.QueryAccountFromDb(0,10)
	ZeroPoolGetUndlying(cfg, accts[0], sdk)
}

//分润合约更新：先部署再init，然后更新治理合约里的分润地址
func TestOracleSetValue(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//accts := Utils.GenerateAccounts(cfg, account, sdk )
	//accts := dbHelper.QueryAccountFromDb(0,10)
	hash1, err := sdk.SendTransaction(WingProfitInit(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(UpdateProfitContract(cfg, account, sdk))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
