package utils

import (
	"fmt"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

//
import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	"github.com/ontio/ontology/common/log"
)

//gov init
func GovInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	Oracle, _ := utils.AddressFromHexString(cfg.Oracle)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	GlobalParam, _ := utils.AddressFromHexString(cfg.GlobalParam)
	params := []interface{}{WingToken, WingProfit, Oracle, GlobalParam, cfg.SDRate}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
