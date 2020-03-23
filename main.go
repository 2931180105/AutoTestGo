package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mockyz/AutoTestGo/common"
	"github.com/mockyz/AutoTestGo/common/config"
	"github.com/mockyz/AutoTestGo/common/log"
	"io/ioutil"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

//RpcClient for ontology rpc api
type RpcClient struct {
	qid        uint64
	addr       string
	httpClient *http.Client
}

//JsonRpc version
const JSON_RPC_VERSION = "2.0"

//JsonRpcRequest object in rpc
type JsonRpcRequest struct {
	Version string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//JsonRpcResponse object response for JsonRpcRequest
type JsonRpcResponse struct {
	Id     string          `json:"id"`
	Error  int64           `json:"error"`
	Desc   string          `json:"desc"`
	Result json.RawMessage `json:"result"`
}

//sendRpcRequest send Rpc request to ontology
func (this *RpcClient) sendRpcRequest(qid, method string, params []interface{}) ([]byte, error) {
	rpcReq := &JsonRpcRequest{
		Version: JSON_RPC_VERSION,
		Id:      qid,
		Method:  method,
		Params:  params,
	}
	data, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("JsonRpcRequest json.Marsha error:%s", err)
	}
	resp, err := this.httpClient.Post(this.addr, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rpc response body error:%s", err)
	}
	rpcRsp := &JsonRpcResponse{}
	err = json.Unmarshal(body, rpcRsp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal JsonRpcResponse:%s error:%s", body, err)
	}
	if rpcRsp.Error != 0 {
		return nil, fmt.Errorf("JsonRpcResponse error code:%d desc:%s result:%s", rpcRsp.Error, rpcRsp.Desc, rpcRsp.Result)
	}
	return rpcRsp.Result, nil
}

//NewRpcClient return RpcClient instance
func NewRpcClient(addr string) *RpcClient {
	return &RpcClient{
		addr: addr,
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false, //enable keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
			},
			Timeout: time.Second * 300, //timeout for http response
		},
	}
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
				fileObj, err = os.OpenFile(fmt.Sprintf("sendLog/invoke_%d.txt", routineIndex),
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
					//_, err := client.sendRpcRequest(client.GetNextQid(), "batchAdd", addArgs)
					_, err := client.sendRpcRequest(client.GetNextQid(), "verify", addArgs)

					if err != nil {
						fmt.Printf("Add Error: %s\n", err)
						log.Errorf("send tx failed, err: %s", err)
						return
					} else {
						log.Infof("send tx ***%s***", addArgs)
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

func hashLeaf(data []byte) common.Uint256 {
	tmp := append([]byte{0}, data...)
	return sha256.Sum256(tmp)
}
func GenerateLeafv(start uint32, N uint32) []common.Uint256 {
	sink := common.NewZeroCopySink(nil)
	leafs := make([]common.Uint256, 0)
	for i := uint32(start); i < start+N; i++ {
		sink.Reset()
		sink.WriteUint32(i)
		leafs = append(leafs, hashLeaf(sink.Bytes()))
	}

	return leafs
}

func leafvToAddArgs(leafs []common.Uint256) []interface{} {
	addargs := make([]interface{}, 0, len(leafs))

	for i := range leafs {
		addargs = append(addargs, hex.EncodeToString(leafs[i][:]))
	}

	return addargs
}

func (this *RpcClient) GetNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}
