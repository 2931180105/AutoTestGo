package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"time"
)

//wing token init
func GovTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	CuurTime := time.Now().Unix()
	params := []interface{}{"init", []interface{}{WingGovAddr, CuurTime}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, WingToken, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
func GovTokenSetGov(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingProfit)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{"setGovernanceAddress", []interface{}{WingGovAddr}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, WingToken, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
