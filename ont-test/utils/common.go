package utils

import (
	"encoding/hex"
	"fmt"
	config "github.com/mockyz/AutoTestGo/ont-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"os"
)

func GenerateLockParam(cfg *config.Config, account *goSdk.Account) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	contractAddress, err := utils.AddressFromHexString(cfg.LockProxy)
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}
	FromAssetHash, _ := common.HexToBytes(cfg.FromAssetHash)
	toAddress, _ := common.HexToBytes(cfg.ToAddress)
	params := []interface{}{"lock", []interface{}{FromAssetHash, account.Address, cfg.ToChainId, toAddress, cfg.Amount}}
	mutTx, err = genTxSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, contractAddress, params)
	if err != nil {
		fmt.Println("construct tx err", err)
		os.Exit(1)
	}
	return mutTx
}

func GenerateOngParam(cfg *config.Config, account *goSdk.Account) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	toAddr, _ := utils.AddressFromBase58(cfg.To)
	mutTx, err := genTxSdk.Native.Ong.NewTransferTransaction(cfg.GasPrice, cfg.GasLimit, account.Address, toAddr, cfg.Amount*100)
	if err != nil {
		panic(err)
	}
	return mutTx
}

func GenerateBindAsstHashTx(cfg *config.Config, fromAsstHash string, toChainId uint64, toAssetHash string, assetLimt []byte, isTargetChainAsset bool) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	contractAddress, err := utils.AddressFromHexString(cfg.LockProxy)
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}
	FromAssetHash, _ := common.HexToBytes(fromAsstHash)
	ToAssetHash, _ := common.HexToBytes(toAssetHash)

	params := []interface{}{"bindAssetHash", []interface{}{FromAssetHash, toChainId, ToAssetHash, assetLimt, isTargetChainAsset}}
	mutTx, err = genTxSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, contractAddress, params)
	if err != nil {
		fmt.Println("construct tx err", err)
		os.Exit(1)
	}
	return mutTx
}

func GenerateLayer2DepositParam(cfg *config.Config, account *goSdk.Account) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	contractAddress, err := utils.AddressFromHexString(cfg.Contract)
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}
	tokenAddress, _ := hex.DecodeString(cfg.FromAssetHash)
	mutTx, err = genTxSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, contractAddress,
		[]interface{}{"deposit", []interface{}{account.Address, cfg.Amount, tokenAddress}})
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}
	return mutTx
}

func GenerateLayer2WithdrawParam(cfg *config.Config, account *goSdk.Account) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	mutTx, err := genTxSdk.Native.Ong.NewTransferTransaction(cfg.GasPrice, cfg.GasLimit, account.Address, common.ADDRESS_EMPTY, cfg.Amount)
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}
	return mutTx
}
func signTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
	tx.Nonce = nonce
	tx.Sigs = nil
	err := sdk.SignToTransaction(tx, signer)
	if err != nil {
		return fmt.Errorf("sign tx failed, err: %s", err)
	}

	return nil
}
