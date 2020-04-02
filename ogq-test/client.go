package main

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/ogq-test/config-ogq"
	log "github.com/ontio/ontology/common/log"
)

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/signature"
	"github.com/ontio/ontology/merkle"
)

//JsonRpc version
const JSON_RPC_VERSION = "2.0"

//JsonRpcRequest object in rpc
type JsonRpcRequest struct {
	Version string   `json:"jsonrpc"`
	Id      string   `json:"id"`
	Method  string   `json:"method"`
	Params  RpcParam `json:"params"`
}

//JsonRpcResponse object response for JsonRpcRequest
type JsonRpcBatchAddResponse struct {
	Id     string `json:"id"`
	Error  int64  `json:"error"`
	Desc   string `json:"desc"`
	Result string `json:"result"`
}

type JsonRpcVerifyResponse struct {
	Id     string       `json:"id"`
	Error  int64        `json:"error"`
	Desc   string       `json:"desc"`
	Result VerifyResult `json:"result"`
}

//RpcClient for ontology rpc api
type RpcClient struct {
	qid        uint64
	addr       string
	httpClient *http.Client
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

//SetAddress set rpc server address. Simple http://localhost:20336
func (this *RpcClient) SetAddress(addr string) *RpcClient {
	this.addr = addr
	return this
}

func (this *RpcClient) GetNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}

//sendRpcRequest send Rpc request to ontology
func (this *RpcClient) sendRpcRequest(qid, method string, params RpcParam) (interface{}, error) {
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
	//fmt.Printf("request: \n%s\n", data)
	resp, err := this.httpClient.Post(this.addr, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rpc response body error:%s", err)
	}

	if method == "batchAdd" {
		//fmt.Printf("response:\n%s", string(body))
		rpcRsp := &JsonRpcBatchAddResponse{}
		err = json.Unmarshal(body, rpcRsp)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal JsonRpcResponse:%s error:%s", body, err)
		}
		if rpcRsp.Error != 0 {
			return nil, fmt.Errorf("JsonRpcResponse error code:%d desc:%s result:%s", rpcRsp.Error, rpcRsp.Desc, rpcRsp.Result)
		}

		return nil, nil
	} else if method == "verify" {
		fmt.Printf("response:\n%s", string(body))
		rpcRsp := &JsonRpcVerifyResponse{}
		err = json.Unmarshal(body, rpcRsp)
		if rpcRsp.Error != 0 {
			return nil, fmt.Errorf("JsonRpcResponse error code:%d desc:%s", rpcRsp.Error, rpcRsp.Desc)
		}
		return &rpcRsp.Result, nil
	}

	return nil, errors.New("error method")
}

var (
	N uint32 = 1
)

func verifyleaf(client *RpcClient, leafs []common.Uint256, v bool) {
	for i := uint32(0); i < uint32(len(leafs)); i++ {
		fmt.Printf("enter verify")
		vargs := getVerifyArgs(leafs[i])
		res, err := client.sendRpcRequest(client.GetNextQid(), "verify", vargs)
		if err != nil {
			fmt.Printf("Verify Failed %s\n", err)
			panic("xxx")
			return
		}

		vres, ok := res.(*VerifyResult)
		if !ok {
			panic("error type")
		}
		if v {
			err = Verify(vres, leafs[i])
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("success")
	}
}

func main() {
	configPath := "ogq-test/config.json"
	err := InitSigner()
	if err != nil {
		log.Error(err)
		return
	}
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
			if cfg.VerTx {
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
					_, err := client.sendRpcRequest(client.GetNextQid(), "batchAdd", addArgs)
					if err != nil {
						log.Errorf("send tx failed, err: %s", err)
					} else {
						log.Infof(" *****sentNum***%d****", sentNum)

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
				if cfg.VerTx && !cfg.SendTx {
					//log.Infof("send tx ***%s***", addArgs[0])
					fileObj.WriteString(addArgs.Hashes[0] + "****" + addArgs.PubKey + "*****" + addArgs.Sigature + "\n")
					verifyleaf(client, leafs, true)
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

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {
			fmt.Printf("OGQ server received exit signal: %v.", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
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

func MerkleInit() *merkle.CompactMerkleTree {
	//store, _ := merkle.NewFileHashStore("merkletree.db", 0)
	tree := merkle.NewTree(0, nil, nil)
	return tree
}

func getleafvroot(leafs []common.Uint256, tree *merkle.CompactMerkleTree, needroot bool) []common.Uint256 {
	root := make([]common.Uint256, 0)
	for i := range leafs {
		tree.AppendHash(leafs[i])
		if needroot {
			root = append(root, tree.Root())
		}
	}

	return root
}

type RpcParam struct {
	PubKey   string   `json:"pubKey"`
	Sigature string   `json:"signature"`
	Hashes   []string `json:"hashes"`
}

func leafvToAddArgs(leafs []common.Uint256) RpcParam {
	leafargs := make([]string, 0, len(leafs))

	for i := range leafs {
		leafargs = append(leafargs, hex.EncodeToString(leafs[i][:]))
	}

	sigData, err := DefSigner.Sign(leafs[0][:])
	if err != nil {
		panic(err)
	}

	addargs := RpcParam{
		PubKey:   hex.EncodeToString(keypair.SerializePublicKey(DefSigner.GetPublicKey())),
		Sigature: hex.EncodeToString(sigData),
		Hashes:   leafargs,
	}

	err = signature.Verify(DefSigner.GetPublicKey(), leafs[0][:], sigData)
	if err != nil {
		panic(err)
	}

	return addargs
}

type VerifyResult struct {
	Root        common.Uint256   `json:"root"`
	TreeSize    uint32           `json:"size"`
	BlockHeight uint32           `json:"blockheight"`
	Index       uint32           `json:"index"`
	Proof       []common.Uint256 `json:"proof"`
}

func (self VerifyResult) MarshalJSON() ([]byte, error) {
	root := hex.EncodeToString(self.Root[:])
	proof := make([]string, 0, len(self.Proof))
	for i := range self.Proof {
		proof = append(proof, hex.EncodeToString(self.Proof[i][:]))
	}

	res := struct {
		Root        string   `json:"root"`
		TreeSize    uint32   `json:"size"`
		BlockHeight uint32   `json:"blockheight"`
		Index       uint32   `json:"index"`
		Proof       []string `json:"proof"`
	}{
		Root:        root,
		TreeSize:    self.TreeSize,
		BlockHeight: self.BlockHeight,
		Index:       self.Index,
		Proof:       proof,
	}

	return json.Marshal(res)
}

func (self *VerifyResult) UnmarshalJSON(buf []byte) error {
	res := struct {
		Root        string   `json:"root"`
		TreeSize    uint32   `json:"size"`
		BlockHeight uint32   `json:"blockheight"`
		Index       uint32   `json:"index"`
		Proof       []string `json:"proof"`
	}{}

	if len(buf) == 0 {
		return nil
	}

	json.Unmarshal(buf, &res)

	root, err := HashFromHexString(res.Root)
	if err != nil {
		return err
	}
	proof, err := convertParamsToLeafs(res.Proof)
	if err != nil {
		return err
	}

	self.Root = root
	self.TreeSize = res.TreeSize
	self.BlockHeight = res.BlockHeight
	self.Index = res.Index
	self.Proof = proof

	return nil
}

func convertParamsToLeafs(params []string) ([]common.Uint256, error) {
	leafs := make([]common.Uint256, len(params), len(params))

	for i := uint32(0); i < uint32(len(params)); i++ {
		s := params[i]
		leaf, err := HashFromHexString(s)
		if err != nil {
			return nil, err
		}
		leafs[i] = leaf
	}

	return leafs, nil
}

func getVerifyArgs(leaf common.Uint256) RpcParam {
	leafs := make([]string, 1, 1)
	leafs[0] = hex.EncodeToString(leaf[:])

	//sigData, err := DefSigner.Sign(leaf[:])
	//if err != nil {
	//	panic(err)
	//}

	vargs := RpcParam{
		PubKey: hex.EncodeToString(keypair.SerializePublicKey(DefSigner.GetPublicKey())),
		Hashes: leafs,
	}

	return vargs
}

func clean() {
	os.RemoveAll("merkletree.db")
}

func printLeafs(prefix string, leafs []common.Uint256) {
	for i := range leafs {
		fmt.Printf("%s[%d]: %x\n", prefix, i, leafs[i])
	}
}

func HashFromHexString(s string) (common.Uint256, error) {
	hx, err := common.HexToBytes(s)
	if err != nil {
		return merkle.EMPTY_HASH, err
	}
	res, err := common.Uint256ParseFromBytes(hx)
	if err != nil {
		return merkle.EMPTY_HASH, err
	}
	return res, nil
}

var DefSigner sdk.Signer

func InitSigner() error {
	DefSdk := sdk.NewOntologySdk()
	wallet, err := DefSdk.OpenWallet("ogq-test/wallet.dat")
	if err != nil {
		return fmt.Errorf("error in OpenWallet:%s\n", err)
	}
	DefSigner, err = wallet.GetAccountByIndex(1, []byte("123456"))

	if err != nil {
		return fmt.Errorf("error in GetDefaultAccount:%s\n", err)
	}

	return nil
}

func Verify(vres *VerifyResult, leaf common.Uint256) error {
	verify := merkle.NewMerkleVerifier()
	err := verify.VerifyLeafHashInclusion(leaf, vres.Index, vres.Proof, vres.Root, vres.TreeSize)
	if err != nil {
		return err
	}

	return nil
}
