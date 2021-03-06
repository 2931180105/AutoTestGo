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
func GovTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	//CuurTime := time.Now().Unix() + 100
	CuurTime := 1599814800
	log.Infof("curr time : %d", CuurTime)
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
func WingTokenSetGov(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, newWingGovAddr string) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(newWingGovAddr)
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

//query_pool_count
func GovTokenBalanceOfToWingGov(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	balanceAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{"balanceOf", []interface{}{balanceAddr}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("GovTokenBalanceOf: %s", resut.Result)
	return resut
}

//query_pool_count
func GovTokenBalanceOf(cfg *config.Config, addr string, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	balanceAddr, _ := utils.AddressFromHexString(addr)
	params := []interface{}{"balanceOf", []interface{}{balanceAddr}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("GovTokenBalanceOf: %s", resut.Result)
	return resut
}

//query_pool_count
func WingTokenGetGovAddr(cfg *config.Config, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{"getGovernanceAddress", []interface{}{}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("GovTokenBalanceOf: %s", resut.Result)
	return resut
}

//getGovernanceAddress
func GetGovAddress(cfg *config.Config, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{"getGovernanceAddress", []interface{}{}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("getGovernanceAddress: %s", resut.Result)
	return resut
}

//getGovernanceAddress
func WingTokenTotalSupply(cfg *config.Config, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{"totalSupply", []interface{}{}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("getGovernanceAddress: %s", resut.Result)
	return resut
}
