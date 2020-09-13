package wingGov

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
)

//ZeroPoolWithDraw
func ZeroPoolWithDraw(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "withdraw_wing", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	_, err2 := genSdk.SendTransaction(mutTx)
	if err2 != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
}

//TODO : check method
func ZeroPoolGetUndlying(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.ResultItem {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(ZeroPoolAddr, "get_unbound_wing", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("get_unbound_wing :%s", result.Result)
	return result.Result
}

func ZeroPoolInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	Global, _ := utils.AddressFromHexString(cfg.GlobalParam)
	GovToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{Global, GovToken}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err2 := genSdk.SendTransaction(mutTx)
	if err2 != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	log.Infof("txhash: %s", hash.ToHexString())
}
