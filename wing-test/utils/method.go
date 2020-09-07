package utils

import (
	goSdk "github.com/ontio/ontology-go-sdk"
)

//
import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
)

//gov init
func GovInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	GovTokenInit(cfg, account, genSdk)
	GovTokenSetGov(cfg, account, genSdk)
	WingProfitInit(cfg, account, genSdk)
	WingGovInit(cfg, account, genSdk)
}
