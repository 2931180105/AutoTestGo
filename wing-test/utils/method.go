package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

func GetGovTokenAddres(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	//WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "get_governance_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

func setGovTokenAddres(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{WingToken}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_governance_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//gov init
func GovInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	Oracle, _ := utils.AddressFromHexString(cfg.Oracle)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	GlobalParam, _ := utils.AddressFromHexString(cfg.GlobalParam)
	params := []interface{}{WingToken, WingProfit, Oracle, GlobalParam, 20}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
