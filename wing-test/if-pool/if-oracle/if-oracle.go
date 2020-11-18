package if_oracle

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)

type IfOracle struct {
	sdk    *ontSDK.OntologySdk
	signer *ontSDK.Account
	addr   common.Address

	gasPrice uint64
	gasLimit uint64
}

func NewIfOracle(nodeRPCAddr string, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64) (*IfOracle, error) {
	sdk := ontSDK.NewOntologySdk()
	client := sdk.NewRpcClient()
	client.SetAddress(nodeRPCAddr)
	sdk.SetDefaultClient(client)
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewOracle: cannot access ontology network through addr %s", nodeRPCAddr)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewOracle: invalid contract addr %s", contractAddr)
		}
	}
	return &IfOracle{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}


func (this *IfOracle) Init(admin common.Address) (string, error) {
	method := "init"
	params := []interface{}{admin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *IfOracle) SetPendingAdmin(pendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{pendingAdmin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *IfOracle) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("_acceptAdmin: %s", err)
	}
	return hash, err
}

func (this *IfOracle) Admin() (common.Address, error) {
	method := "admin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}
func (this *IfOracle) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *IfOracle) BindAddressToONTID(addr common.Address, ont_id string, index uint64) (string, error) {
	method := "bindAddressToONTID"
	params := []interface{}{addr, ont_id, index}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *IfOracle) PutAddressKYC(addr common.Address, kycHash string) (string, error) {
	method := "putAddressKYC"
	params := []interface{}{addr, kycHash}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *IfOracle) PutAddressLevel(addr common.Address, level uint64) (string, error) {
	method := "putAddressLevel"
	params := []interface{}{addr, level}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutAddressLevel: %s", err)
	}
	return hash, err
}

func (this *IfOracle) GetBindedONTID(addr common.Address) (string, error) {
	method := "getBindedONTID"
	params := []interface{}{addr}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *IfOracle) GetKYC(addr common.Address) (string, error) {
	method := "getKYC"
	params := []interface{}{addr}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *IfOracle) GetLevel(addr common.Address) (*AddressLevelInfo, error) {
	method := "getLevel"
	params := []interface{}{addr}
	res, err := utils.PreExecuteByteArray(this.sdk, this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("Admin: %s", err)
	}
	source := common.NewZeroCopySource(res)
	info := &AddressLevelInfo{}
	err = info.Deserialize(source)
	if err != nil {
		return nil, err
	}
	return info, err
}

type AddressLevelInfo struct {
	OScoreLevel uint8  `json:"o_score_level"` // level
	Timestamp   uint64 `json:"timestamp"`
}

func (this *AddressLevelInfo) Deserialize(source *common.ZeroCopySource) error {
	level, eof := source.NextByte()
	if eof {
		return fmt.Errorf("%v", eof)
	}
	stamp, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("%v", eof)
	}
	this.OScoreLevel = level
	this.Timestamp = stamp
	return nil
}
