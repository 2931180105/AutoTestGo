package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

func Gov_init(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingProfit, _ := utils.AddressFromHexString(cfg.WingProfit)
	Oracle, _ := utils.AddressFromHexString(cfg.Oracle)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{"init", []interface{}{WingToken, WingProfit, Oracle, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
	//sendTxSdk := goSdk.NewOntologySdk()
	//rpcClient := client.NewRpcClient()
	//rpcClient.SetAddress(cfg.Rpc[0])
	//sendTxSdk.SetDefaultClient(rpcClient)
	//if err := signTx(sendTxSdk, mutTx, cfg.StartNonce, account); err != nil {
	//	log.Error(err)
	//}
	//hash, err := sendTxSdk.SendTransaction(mutTx)
	//if err != nil {
	//	log.Errorf("send tx failed, err: %s********", err)
	//} else {
	//	log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	//}
}

func signTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
	if nonce != 0 {
		tx.Nonce = nonce
	}
	tx.Sigs = nil
	err := sdk.SignToTransaction(tx, signer)
	if err != nil {
		return fmt.Errorf("sign tx failed, err: %s", err)
	}
	return nil
}
