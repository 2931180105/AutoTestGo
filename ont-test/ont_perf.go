package main

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/ont-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"math"
	"os"
	"time"
)

func main() {
	log.InitLog(log.InfoLog, log.PATH, log.Stdout)
	configPath := "ont-test/config.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		return
	}
	sdk := goSdk.NewOntologySdk()
	wallet, err := sdk.OpenWallet(cfg.Wallet)
	if err != nil {
		log.Errorf("parse wallet err: %s", err)
		return
	}
	account, err := wallet.GetDefaultAccount([]byte(cfg.Password))
	if err != nil {
		log.Errorf("get account err: %s", err)
		return
	}
	testTransfer(cfg, account)

}
func GenerateLockParam(cfg *config.Config, account *goSdk.Account) *types.MutableTransaction {
	genTxSdk := goSdk.NewOntologySdk()
	var mutTx *types.MutableTransaction
	contractAddress, err := utils.AddressFromHexString(cfg.Contract)
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

func testTransfer(cfg *config.Config, account *goSdk.Account) {
	txNum := cfg.TxNum * cfg.TxFactor
	if txNum > math.MaxUint32 {
		txNum = math.MaxUint32
	}
	exitChan := make(chan int)
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	startTestTime := time.Now().UnixNano() / 1e6
	for i := uint(0); i < cfg.RoutineNum; i++ {
		mutTx := GenerateLockParam(cfg, account)
		go func(nonce uint32, routineIndex uint) {
			sendTxSdk := goSdk.NewOntologySdk()
			rpcClient := client.NewRpcClient()
			rpcClient.SetAddress(cfg.Rpc[int(routineIndex)%len(cfg.Rpc)])
			sendTxSdk.SetDefaultClient(rpcClient)
			startTime := time.Now().UnixNano() / 1e6 // ms
			sentNum := int64(0)
			var fileObj *os.File
			if cfg.SaveTx {
				fileObj, _ = os.OpenFile(fmt.Sprintf("invoke_%d.txt", routineIndex),
					os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			}
			for j := uint(0); j < txNumPerRoutine; j++ {
				if err := signTx(sendTxSdk, mutTx, nonce, account); err != nil {
					log.Error(err)
					continue
				}
				tx, err := mutTx.IntoImmutable()
				if err != nil {
					log.Error("error****%s", err)
				}
				txHash := tx.Hash()
				sink := common.NewZeroCopySink(nil)
				tx.Serialization(sink)
				txContent := common.ToHexString(sink.Bytes())
				//log.Info(txHash.ToHexString() + "," + txContent + "\n")
				if cfg.SendTx {
					hash, err := sendTxSdk.SendTransaction(mutTx)
					if err != nil {
						log.Errorf("send tx failed, err: %s********", err)
					} else {
						log.Infof("send tx %s", hash.ToHexString())
					}
					sentNum++
					now := time.Now().UnixNano() / 1e6 // ms
					diff := sentNum - (now-startTime)/1e3*tpsPerRoutine
					if now > startTime && diff > 0 {
						sleepTime := time.Duration(diff*1000/tpsPerRoutine) * time.Millisecond
						time.Sleep(sleepTime)
						log.Infof("sleep %d ms", sleepTime.Nanoseconds()/1e6)
					}
				}
				nonce++
				if cfg.SaveTx {
					fileObj.WriteString(txHash.ToHexString() + "," + txContent + "\n")
				}
			}
			exitChan <- 1
		}(uint32(txNumPerRoutine*i)+cfg.StartNonce, i)
	}
	for i := uint(0); i < cfg.RoutineNum; i++ {
		<-exitChan
	}
	endTestTime := time.Now().UnixNano() / 1e6
	log.Infof("send tps is %f", float64(txNum*1000)/float64(endTestTime-startTestTime))
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
