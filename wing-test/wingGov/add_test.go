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
	hash1, err := sdk.SendTransaction(WingTokenSetGov(cfg, account, sdk, cfg.WingGov))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}
func TestDep(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	wasmFile := "../contract/private/zero_pool.wasm.str"
	DeployContractProfit(cfg, account, sdk, wasmFile)
}

//部署后init的所有合约
func TestIint(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	hash1, err := sdk.SendTransaction(WingTokenSetGov(cfg, account, sdk, cfg.WingGov))
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
	BatchUnStaking(cfg, account, sdk, accts)
}

func TestUnStaking(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accounts := dbHelper.QueryAccountFromDb(21, cfg.AccountNum)
	//accts := Utils.GenerateAccounts(cfg, account, sdk )
	//accts := dbHelper.QueryAccountFromDb(0,10)
	BatchUnStaking(cfg, account, sdk, accounts)
}

//query account balance of wing
func TestGetUndlying(t *testing.T) {
	cfg, _, sdk := GetTestConfig()
	//accts := Utils.GenerateAccounts(cfg, account, sdk)
	accts := dbHelper.QueryAccountFromDb(0, 33)
	for i := 0; i < len(accts); i++ {
		wing_balance := ZeroPoolGetUndlying(cfg, accts[i], sdk)
		balance_int, err := wing_balance.ToInteger()
		if err != nil {
			log.Errorf("balance error: %s", err)
		}
		log.Infof("wing_balance: %d", balance_int)
		Utils.UpdateWingBalance(balance_int, accts[i].Address.ToBase58())
		if 1 < balance_int.Uint64() {
			log.Infof("withDraw address: %s", accts[i].Address.ToBase58())
			ZeroPoolWithDraw(cfg, accts[i], sdk)
		}
	}
}

//withDraw
func TestWithDrawWing(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	accts := Utils.GenerateAccounts(cfg, account, sdk)
	//accts := dbHelper.QueryAccountFromDb(0, 33)
	for i := 0; i < len(accts); i++ {
		ZeroPoolWithDraw(cfg, accts[i], sdk)
	}
}

//query account balance of Staking
func TestGetStakingBalance(t *testing.T) {
	cfg, _, sdk := GetTestConfig()
	//accts := Utils.GenerateAccounts(cfg, account, sdk)
	accts := dbHelper.QueryAccountFromDb(0, 33)
	for i := 0; i < len(accts); i++ {
		wing_balance := ZeroPoolGetStakingBalance(cfg, accts[i], sdk)
		balance_int, err := wing_balance.ToInteger()
		if err != nil {
			log.Errorf("balance error: %s", err)
		}
		log.Infof("wing_balance: %d", balance_int)
		Utils.UpdateStakingBalance(balance_int, accts[i].Address.ToBase58())
	}

}

//query account balance of Staking
func TestAddAllSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	//Query_total_pool_bonus(cfg, sdk )
	AddAllSuuportToken(cfg, account, sdk)

}
func TestUpdateAllSuuportToken(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	UpdateAllSuuportToken(cfg, account, sdk)

}
func TestSetOracleAddr(t *testing.T) {
	cfg, account, sdk := GetTestConfig()
	Set_oracle_address(cfg, account, sdk)

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
