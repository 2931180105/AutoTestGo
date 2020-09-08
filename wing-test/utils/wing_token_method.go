package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"time"
)

//wing token init
func GovTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	TokenBytes, _ := OntCommon.HexToBytes(cfg.GovToken)
	WingToken, _ := utils.AddressParseFromBytes(OntCommon.ToArrayReverse(TokenBytes))
	CuurTime := uint32(time.Now().Unix()) + 100
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

//query_pool_count
func GovTokenBalanceOf(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	ContractAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	balanceAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{"balanceOf", []interface{}{balanceAddr}}
	resut, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ContractAddr, params)
	log.Infof("GovTokenBalanceOf: %s", resut.Result)
	return resut
}
