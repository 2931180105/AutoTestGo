package test

import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	WingGovMethod "github.com/mockyz/AutoTestGo/wing-test/wingGov"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common/log"
	"testing"
	"time"
)

var sdk = goSdk.NewOntologySdk()
var account *goSdk.Account
var cfg *config.Config
var SleepTime = time.Duration(1)

func TestWingGovUnboundTokenAndToPool(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	hash1, err := sdk.SendTransaction(WingGovMethod.UnboundToken(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(WingGovMethod.UnboundToPool(cfg, account, sdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func TestWingGovQueryUnboundToken(t *testing.T) {
	cfg, _, sdk := Utils.GetTestConfig()
	//WingGovMethod.Query_unbound_to_pool_count(cfg, sdk)
	WingGovMethod.GetGovAddress(cfg, sdk)
}

func TestStakingCaseStep01(t *testing.T) {
	cfg, account, sdk, accts := Utils.GetTestConfigAddAccts()
	accts = Utils.GenerateAccounts(cfg, account, sdk)
	//old staking A 1000; B 100, D : staking 300 unstaking -300
	log.Infof("fist time staking")
	WingGovMethod.ZeroPoolStaking(cfg, accts[1], sdk, 1000)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[2], sdk, 100)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[4], sdk, 300)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[4], sdk, 300)
	time.Sleep(time.Second * SleepTime)
	log.Infof("fist time invoke unbound to pool")
	//invoke unbound to pool before 10 min
	WingGovMethod.UnboundTokenSend(cfg, account, sdk)
	WingGovMethod.UnboundToPoolSend(cfg, account, sdk)
}
func TestStakingCaseStep02(t *testing.T) {
	cfg, _, sdk, accts := Utils.GetTestConfigAddAccts()
	//	C: +500;A -500;B:-100;E  +100
	WingGovMethod.ZeroPoolStaking(cfg, accts[3], sdk, 500)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[1], sdk, 500)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[2], sdk, 100)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[5], sdk, 100)
	//wait time to unbound wing
	time.Sleep(time.Second * 60 * 5)
}
func TestStakingCaseStep03(t *testing.T) {
	cfg, account, sdk, accts := Utils.GetTestConfigAddAccts()
	// F : +100; G:+1000;H:+500;F:-100;G:-500
	WingGovMethod.ZeroPoolStaking(cfg, accts[6], sdk, 100)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[7], sdk, 1000)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[8], sdk, 500)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[6], sdk, 100)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[7], sdk, 500)
	time.Sleep(time.Second * SleepTime)
	// 1  invoke unbound wing
	WingGovMethod.UnboundTokenSend(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	WingGovMethod.UnboundToPoolSend(cfg, account, sdk)
	// TODO:start check address get unbound wing
}

//
func TestStakingCaseStep04(t *testing.T) {
	cfg, account, sdk, accts := Utils.GetTestConfigAddAccts()
	//J : +100 ; L:+200;G:-200
	WingGovMethod.ZeroPoolStaking(cfg, accts[10], sdk, 100)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolStaking(cfg, accts[11], sdk, 200)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[7], sdk, 500)
	time.Sleep(time.Second * SleepTime)
	//	2 invoke unbound to pool
	WingGovMethod.UnboundTokenSend(cfg, account, sdk)
	time.Sleep(time.Second * 3)
	WingGovMethod.UnboundToPoolSend(cfg, account, sdk)
	// TODO:start check address get unbound wing
	WingGovMethod.ZeroPoolStaking(cfg, accts[12], sdk, 200)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStaking(cfg, accts[7], sdk, 100)
	time.Sleep(time.Second * SleepTime)
}

func TestStakingCaseStep05(t *testing.T) {
	cfg, admin, sdk, accts := Utils.GetTestConfigAddAccts()
	//	upgrade pool to new pool
	zeroPoolPath := "../../contract/testnet/wing_dao_contracts_new.wasm.str"
	newZeroPoolAddr := WingGovMethod.MigrateZeroPool(cfg, admin, sdk, zeroPoolPath)
	WingGovMethod.ZeroPoolStakingByAddr(cfg, accts[13], sdk, 200, newZeroPoolAddr)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStakingByAddr(cfg, accts[12], sdk, 100, newZeroPoolAddr)
	time.Sleep(time.Second * SleepTime)
	// TODO:start check address get unbound wing
}
func TestStakingCaseStep06(t *testing.T) {
	cfg, admin, sdk, accts := Utils.GetTestConfigAddAccts()
	//	upgrade pool to new pool
	zeroPoolPath := "../../contract/private/zero_pool4.wasm.str"
	newZeroPoolAddr := WingGovMethod.MigrateZeroPool(cfg, admin, sdk, zeroPoolPath)
	WingGovMethod.ZeroPoolStakingByAddr(cfg, accts[13], sdk, 200, newZeroPoolAddr)
	time.Sleep(time.Second * SleepTime)
	WingGovMethod.ZeroPoolUnStakingByAddr(cfg, accts[12], sdk, 100, newZeroPoolAddr)
	time.Sleep(time.Second * SleepTime)
	// TODO:start check address get unbound wing
}
