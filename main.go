package main

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/config"
	"github.com/mockyz/AutoTestGo/common/log"
	sdk "github.com/ontio/ontology-go-sdk"
	"os"
	"time"
)

//Lock:
//- fromAssetHash：0000000000000000000000000000000000000001
//toChainId：2 ETH
//- toAddress：ETH地址去掉 F41089700D6d950C8c379772f8a12b12955dB886
//- amount： 10
type LockParam struct {
	FromAssetHash string
	ToChainId     int
	ToAddress     string
	Amount        int
}

func main() {
	configPath := "client/config.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		return
	}
	exitChan := make(chan int)
	var txNum = cfg.TxNum * cfg.TxFactor
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	startTestTime := time.Now().UnixNano() / 1e6
	for i := uint(0); i < cfg.RoutineNum; i++ {
		//rand.Int()%len(cfg.Rpc)随机获取一个接口
		//client := NewRpcClient(cfg.Rpc[rand.Int()%len(cfg.Rpc)])
		go func(nonce uint32, routineIndex uint) {
			startTime := time.Now().UnixNano() / 1e6 // ms
			sentNum := int64(0)
			var fileObj *os.File
			if cfg.SaveTx {
				fileObj, err = os.OpenFile(fmt.Sprintf("sendLog/invoke_%d.txt", routineIndex),
					os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Failed to open the file", err.Error())
					os.Exit(2)
				}
			}
			for j := uint(0); j < txNumPerRoutine; j++ {

				if cfg.SendTx {
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
					fileObj.WriteString("1" + "\n")
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

var LockSigner sdk.Signer

func InitLockSigner() error {
	LockSdk := sdk.NewOntologySdk()
	wallet, err := LockSdk.OpenWallet("client/wallet.dat")
	if err != nil {
		return fmt.Errorf("error in OpenWallet:%s\n", err)
	}
	LockSigner, err = wallet.GetAccountByIndex(1, []byte("123456"))

	if err != nil {
		return fmt.Errorf("error in GetDefaultAccount:%s\n", err)
	}

	return nil
}
