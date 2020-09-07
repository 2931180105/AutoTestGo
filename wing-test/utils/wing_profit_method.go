package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

//wing token init
func WingProfitInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{cfg.Eta, cfg.Gama}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
