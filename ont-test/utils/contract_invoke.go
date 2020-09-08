package utils

import (
	config "github.com/mockyz/AutoTestGo/ont-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math/big"
)

func LockInvoke(cfg *config.Config, account *goSdk.Account) {
	sendTxSdk := goSdk.NewOntologySdk()
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sendTxSdk.SetDefaultClient(rpcClient)
	mutTx := GenerateLockParam(cfg, account)
	if err := signTx(sendTxSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := sendTxSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func BindAssetHash(cfg *config.Config, account *goSdk.Account) {
	fromAsstHash := cfg.EthX
	toAssetHash := "0000000000000000000000000000000000000000"
	toChainID := uint64(2)
	Factor := int64(1000000000000000000)
	assetLimt := common.BigIntToNeoBytes(big.NewInt(4 * Factor))
	isTargetChainAsset := bool(true)
	mutTx := GenerateBindAsstHashTx(cfg, fromAsstHash, toChainID, toAssetHash, assetLimt, isTargetChainAsset)
	sendTxSdk := goSdk.NewOntologySdk()
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sendTxSdk.SetDefaultClient(rpcClient)
	if err := signTx(sendTxSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := sendTxSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func GetAseetHash(cfg *config.Config, account *goSdk.Account) {
	sendTxSdk := goSdk.NewOntologySdk()
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sendTxSdk.SetDefaultClient(rpcClient)
	contractAddr, err := utils.AddressFromHexString(cfg.LockProxy)
	if err != nil {
		log.Errorf("balanceOf: decode contract addr failed, err: %s", err)
		return
	}
	FromAssetHash, _ := common.HexToBytes(cfg.EthX)
	params := []interface{}{"getAssetHash", []interface{}{FromAssetHash, cfg.ToChainId}}
	preResult, err := sendTxSdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr, params)
	if err != nil {
		log.Errorf("balanceOf: pre-execute failed, err: %s", err)
		return
	}
	log.Infof("result is %x", preResult.Result)
}

func GetCrossLimt(cfg *config.Config, account *goSdk.Account) {
	sendTxSdk := goSdk.NewOntologySdk()
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[0])
	sendTxSdk.SetDefaultClient(rpcClient)
	contractAddr, err := utils.AddressFromHexString(cfg.LockProxy)
	if err != nil {
		log.Errorf("balanceOf: decode contract addr failed, err: %s", err)
		return
	}
	FromAssetHash, _ := common.HexToBytes(cfg.EthX)
	params := []interface{}{"getCrossedLimit", []interface{}{FromAssetHash, cfg.ToChainId}}
	preResult, err := sendTxSdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr, params)
	if err != nil {
		log.Errorf("balanceOf: pre-execute failed, err: %s", err)
		return
	}
	log.Infof("result is %x", preResult.Result)
}
