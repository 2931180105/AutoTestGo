package wingGov

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
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

//get_eta
func Get_eta(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingProfitAddr, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingProfitAddr, "get_eta", params)
	log.Infof("get_eta: %s", resut.Result)
	return resut
}

// set_eta
func Set_eta(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{cfg.Eta}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "set_eta", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_gama
func Get_gama(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingProfitAddr, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingProfitAddr, "get_gama", params)
	log.Infof("get_gama: %s", resut.Result)
	return resut
}

// set_gama
func Set_gama(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{cfg.Gama}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "set_gama", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

// static_profit
func Static_profit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{cfg.TotalStaticProfit, cfg.LendAmount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "static_profit", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

// dynamic_profit
func Dynamic_profit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{cfg.TotalDynamicProfit, cfg.LendAmount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "dynamic_profit", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

// dynamic_profit
func DestroyProfit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "destroy", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

// migrate TODO: add param
func MigrateProfit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "migrate", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
