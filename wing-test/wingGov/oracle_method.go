package wingGov

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

//init
func OracleInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	OracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	//AbNHrhXT3gsDpNrkmAtsagLSuyzX19UQsn
	params := []interface{}{account.Address}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, OracleAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//SetDecimal
func OracleSetDecimal(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	OracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	params := []interface{}{9}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, OracleAddr, "setDecimal", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//SetDecimal
func OracleGetDecimal(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	OracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	params := []interface{}{}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(OracleAddr, "getDecimal", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("OracleGetDecimal:%s", result.Result)
}

//SetDecimal
func OraclePutUnderlyingPrice(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	OracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	tokenNames := []string{"DAI", "ONT"}
	prices := []string{"4e2d223d000000000000000000000000", "ebb3fd25000000000000000000000000"}
	params := []interface{}{tokenNames, prices}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(OracleAddr, "putUnderlyingPrice", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("OracleGetDecimal:%s", result.Result)
}

//SetDecimal
func OracleGetUnderlyingPrice(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, tokenName string) {
	OracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	params := []interface{}{tokenName}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(OracleAddr, "getUnderlyingPrice", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("getUnderlyingPrice:%s", result.Result)
}

//get_gama
func Get_value(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingProfitAddr, _ := utils.AddressFromHexString(cfg.GlobalParam)
	params := []interface{}{"wing-dao-contract"}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingProfitAddr, "get_value", params)
	log.Infof("Get_value: %s", resut.Result)
	return resut
}

// set_gama
func Set_value(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingProfit, _ := utils.AddressFromHexString(cfg.GlobalParam)
	WingGov, _ := utils.AddressFromHexString(cfg.WingGov)

	params := []interface{}{"wing-dao-contract", WingGov[:], account.Address}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingProfit, "set_value", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
