package wingGov

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"io/ioutil"
	"time"
)

//
import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
)

//contract  init TODO: add need more invoke
func ContractInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) []*types.MutableTransaction {
	var InitTx []*types.MutableTransaction
	InitTx = append(InitTx, GovTokenInit(cfg, account, genSdk), WingTokenSetGov(cfg, account, genSdk), WingProfitInit(cfg, account, genSdk), WingGovInit(cfg, account, genSdk), OracleInit(cfg, account, genSdk), OracleSetDecimal(cfg, account, genSdk))
	//GovTokenInit(cfg, account, genSdk)
	//GovTokenSetGov(cfg, account, genSdk)
	//WingProfitInit(cfg, account, genSdk)
	//WingGovInit(cfg, account, genSdk)
	////zero pool init
	//// init zero pool (global , wing token)
	//// invoke gov regsiter pool
	//// staking , unstaking ,withdraw_wing ,get amount_wing,
	//// global init
	////	ADD Token support
	//
	////	 oracle init : admin address ,setDecimal
	//OracleInit(cfg, account, genSdk)
	//OracleSetDecimal(cfg, account, genSdk)
	return InitTx

}

func DeployContractProfit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/testnet/profit.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	profitCodeStr := string(bytes)
	profitContract, err := utils.GetContractAddress(profitCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("profitContract address : %s", profitContract.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, profitCodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}
func DeployContractWingGov(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/contract/testnet/wing_dao_contracts.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	GovCodeStr := string(bytes)
	GovCodeContract, err := utils.GetContractAddress(GovCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("wing Gov address : %s", GovCodeContract.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, GovCodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}

func DeployContractOracle(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/oracle.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	CodeStr := string(bytes)
	CodeContractAddr, err := utils.GetContractAddress(CodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Oracle address : %s", CodeContractAddr.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, CodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}

func DeployContractWingToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/testnet/wing.avm")
	if err != nil {
		log.Fatal(err)
	}
	CodeStr := string(bytes)
	CodeContractAddr, err := utils.GetContractAddress(CodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Wing address : %s", CodeContractAddr.ToHexString())
	result, err := genSdk.NeoVM.DeployNeoVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, true, CodeStr, "WING Token", "1.0.1", "Wing Team", "support@wing.finance", "Wing is a credit-based, cross-chain DeFi platform.")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}
func BatchStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, accts []*goSdk.Account) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	for i := 0; i < len(accts); i++ {
		acct := accts[i]
		params := []interface{}{acct.Address, cfg.StakeOnt}
		mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "staking", params)
		if err := signTx(genSdk, mutTx, cfg.StartNonce, acct); err != nil {
			log.Error(err)
		}
		time.Sleep(time.Second)
		txhash, err := genSdk.SendTransaction(mutTx)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send staking tx %s****sentnum:***%d", txhash.ToHexString(), i)
		}
	}
}

func BatchUnStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, accts []*goSdk.Account) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	//accts := GenerateAccounts(cfg, account, genSdk)
	for i := 0; i < len(accts); i++ {
		acct := accts[i]
		params := []interface{}{acct.Address, cfg.StakeOnt}
		mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "unstaking", params)
		if err := signTx(genSdk, mutTx, cfg.StartNonce, acct); err != nil {
			log.Error(err)
		}
		txhash, err := genSdk.SendTransaction(mutTx)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send unstaking tx %s****sentnum:***%d", txhash.ToHexString(), i)
		}
		ZeroPoolWithDraw(cfg, acct, genSdk)
		time.Sleep(time.Second * 3)
		resut2, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ZeroPoolAddr, []interface{}{"balanceOf", []interface{}{acct.Address}})
		log.Infof("Address %s , WIng Token BalanceOf : %s", acct.Address.ToBase58(), resut2)
	}
}
func AddAllSuuportToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk) {
	hash1, err := sdk.SendTransaction(Add_support_token(cfg, account, sdk, "DAI", cfg.ODAI))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "ETH", cfg.OETH))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "BTC", cfg.OWBTC))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "USDT", cfg.OUSDT))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

}
