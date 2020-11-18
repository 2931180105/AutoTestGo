package oscore_oracle

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"math/big"
)

// TODO: support rest/ws
// TODO: support estimate gas before execute tx

type Oracle struct {
	sdk    *ontSDK.OntologySdk
	signer *ontSDK.Account
	addr   common.Address

	gasPrice uint64
	gasLimit uint64
}

func NewOracle(nodeRPCAddr string, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64) (*Oracle, error) {
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
	return &Oracle{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}



func (this *Oracle) UpdateSigner(newSigner *ontSDK.Account) {
	this.signer = newSigner
}

func (this *Oracle) GetAddr() common.Address {
	return this.addr
}

func (this *Oracle) Init(admin common.Address) (string, error) {
	method := "init"
	params := []interface{}{admin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *Oracle) SetPendingAdmin(newPendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{newPendingAdmin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *Oracle) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AcceptAdmin: %s", err)
	}
	return hash, err
}

func (this *Oracle) BindAddressToONTID(addr common.Address, ont_id []byte, index uint64) (string, error) {
	method := "bindAddressToONTID"
	params := []interface{}{addr, ont_id, index}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BindAddressToONTID: %s", err)
	}
	return hash, err
}

func (this *Oracle) PutAddressKYC(addr common.Address, kyc_hash []byte) (string, error) {
	method := "putAddressKYC"
	params := []interface{}{addr, kyc_hash}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutAddressKYC: %s", err)
	}
	return hash, err
}

func (this *Oracle) PutAddressLevel(addr common.Address, level uint64) (string, error) {
	method := "putAddressLevel"
	params := []interface{}{addr, level}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutAddressLevel: %s", err)
	}
	return hash, err
}

func (this *Oracle) GetBindedONTID(addr common.Address) ([]byte, error) {
	method := "getBindedONTID"
	params := []interface{}{addr}
	res, err := utils.PreExecuteByteArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetBindedONTID: %s", err)
	}
	return res, err
}

func (this *Oracle) GetKYC(addr common.Address) ([]byte, error) {
	method := "getKYC"
	params := []interface{}{addr}
	res, err := utils.PreExecuteByteArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetKYC: %s", err)
	}
	return res, err
}

func (this *Oracle) GetLevel(addr common.Address) (*big.Int, error) {
	method := "getLevel"
	params := []interface{}{addr}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetLevel: %s", err)
	}
	return res, err
}
