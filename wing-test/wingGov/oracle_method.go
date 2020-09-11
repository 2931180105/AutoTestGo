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

//init
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
