package main

import (
	"encoding/json"
	"fmt"
	"github.com/mockyz/AutoTestGo/common"
	"github.com/mockyz/AutoTestGo/common/config"
	"github.com/mockyz/AutoTestGo/common/log"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type ServerConfig struct {
	Walletname      string `json:"walletname"`
	Walletpassword  string `json:"walletpassword"`
	OntNode         string `json:"ontnode"`
	SignerAddress   string `json:"signeraddress"`
	ServerPort      int    `json:"serverport"`
	ContracthexAddr string `json:"contracthexaddr"`
}
type ClientConfig struct {
	numBatch    int64
	redomNum    int64
	startNumber int64
	batchCount  int64
}

var DefConfig ServerConfig

//var client ClientConfig
func TestVerifyNum(t *testing.T) {
	numbatch := uint32(10000 * 3000 * 2)
	//67296- 72091
	for m := uint32(numbatch / 2); m < numbatch; m++ {
		if m%5 == 0 {
			fmt.Printf("add number %d\n", m-numbatch)

		}
	}

}
func TestLoadFile(t *testing.T) {
	configBuff, err := ioutil.ReadFile("../config.con")
	if err != nil {
		log.Debugf("%v", err)
		return
	}
	err = json.Unmarshal([]byte(configBuff), &DefConfig)
	if err != nil {
		log.Debugf("%v", err)
		return
	}
	log.Debugf("%v", &DefConfig)
	return
}

func TestVerfi(t *testing.T) {
	configPath := "config.json"
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		return
	}
	exitChan := make(chan int)
	var txNum = cfg.TxNum * cfg.TxFactor
	txNumPerRoutine := txNum / cfg.RoutineNum
	tpsPerRoutine := int64(cfg.TPS / cfg.RoutineNum)
	client := NewRpcClient(cfg.Rpc[0])
	startTestTime := time.Now().UnixNano() / 1e6
	for i := uint(0); i < cfg.RoutineNum; i++ {
		//rand.Int()%len(cfg.Rpc)随机获取一个接口
		//client := NewRpcClient(cfg.Rpc[rand.Int()%len(cfg.Rpc)])
		go func(nonce uint32, routineIndex uint) {
			startTime := time.Now().UnixNano() / 1e6 // ms
			sentNum := int64(0)
			var fileObj *os.File
			if cfg.SaveTx {
				fileObj, err = os.OpenFile(fmt.Sprintf("invoke_%d.txt", routineIndex),
					os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Failed to open the file", err.Error())
					os.Exit(2)
				}
			}
			for j := uint(0); j < txNumPerRoutine; j++ {

				var leafs []common.Uint256
				leafs = GenerateLeafv(uint32(0)+cfg.BatchCount*nonce, cfg.BatchCount)
				addArgs := leafvToAddArgs(leafs)
				if cfg.SendTx {
					_, err := client.sendRpcRequest(client.GetNextQid(), "batchAdd", addArgs)
					if err != nil {
						fmt.Printf("Add Error: %s\n", err)
						log.Errorf("send tx failed, err: %s", err)
						return
					} else {
						log.Infof("send tx ***%s*** status ---%s", addArgs)
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
