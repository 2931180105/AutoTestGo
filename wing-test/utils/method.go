package utils

import (
	goSdk "github.com/ontio/ontology-go-sdk"
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
}
