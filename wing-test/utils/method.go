package utils

import (
	"github.com/mockyz/AutoTestGo/common/log"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"io/ioutil"
)

//
import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
)

//contract  init TODO: add need more invoke
func ContractInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	GovTokenInit(cfg, account, genSdk)
	GovTokenSetGov(cfg, account, genSdk)
	WingProfitInit(cfg, account, genSdk)
	WingGovInit(cfg, account, genSdk)
	//zero pool init
	// init zero pool (global , wing token)
	// invoke gov regsiter pool
	// staking , unstaking ,withdraw_wing ,get amount_wing,
	// global init
	//	ADD Token support
	//	 oracle init : admin address ,setDecimal

}

func DeployContreacProfit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/profit.wasm.str")
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
func DeployContreacGov(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/wing_dao_contracts.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	GovCodeStr := string(bytes)
	GovCodeContract, err := utils.GetContractAddress(GovCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("GovCodeContract address : %s", GovCodeContract.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, GovCodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}
