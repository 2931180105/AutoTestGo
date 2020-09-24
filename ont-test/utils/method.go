package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/ont-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math"
	"os"
	"time"
)

func BalanceOf(sdk *goSdk.OntologySdk, address common.Address, ContractHash string) {
	contractAddr, err := utils.AddressFromHexString(ContractHash)
	if err != nil {
		log.Errorf("balanceOf: decode contract addr failed, err: %s", err)
		return
	}
	preResult, err := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr, []interface{}{"balanceOf", []interface{}{address}})
	if err != nil {
		log.Errorf("balanceOf: pre-execute failed, err: %s", err)
		return
	}
	balance, err := preResult.Result.ToInteger()
	if err != nil {
		log.Errorf("balanceOf: parse result %v failed, err: %s", preResult, err)
		return
	}
	log.Infof("balanceOf: addr %s is %d", address.ToBase58(), balance)
}

func TestTransfer(cfg *config.Config, account *goSdk.Account) {
	txNum := cfg.TxNum * cfg.TxFactor
	if txNum > math.MaxUint32 {
		txNum = math.MaxUint32
	}
	exitChan := make(chan int)
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	startTestTime := time.Now().UnixNano() / 1e6
	for i := uint(0); i < cfg.RoutineNum; i++ {
		mutTx := GenerateOngParam(cfg, account)
		go func(nonce uint32, routineIndex uint) {
			sendTxSdk := goSdk.NewOntologySdk()
			rpcClient := client.NewRpcClient()
			//rpcClient.SetAddress(cfg.Rpc[nonce%14])

			rpcClient.SetAddress(cfg.Rpc[0])

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
						//time.Sleep(300000*1e6)
						sentNum++
						log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), sentNum)
					}
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
