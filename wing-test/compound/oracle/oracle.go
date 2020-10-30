package oracle

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
	gasLimit uint64, ) (*Oracle, error) {
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

func (this *Oracle) PutUnderlyingPrice(keyList []string, price []*big.Int) (string, error) {
	method := "putUnderlyingPrice"
	keyParam := make([]interface{}, 0)
	for _, key := range keyList {
		keyParam = append(keyParam, key)
	}
	priceParam := make([]interface{}, 0)
	for _, p := range price {
		priceParam = append(priceParam, p)
	}
	params := []interface{}{keyParam, priceParam}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutUnderlyingPrice: %s", err)
	}
	return hash, err
}

func (this *Oracle) SetDecimal(decimals uint8) (string, error) {
	method := "setDecimal"
	params := []interface{}{decimals}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetDecimal: %s", err)
	}
	return hash, err
}

func (this *Oracle) IsPriceOracle() (bool, error) {
	method := "isPriceOracle"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsPriceOracle: %s", err)
	}
	return res, err
}

func (this *Oracle) GetUnderlyingPrice(key string) (*big.Int, error) {
	method := "getUnderlyingPrice"
	params := []interface{}{key}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetUnderlyingPrice: %s", err)
	}
	return res, err
}

func (this *Oracle) GetDecimal() (uint8, error) {
	method := "getDecimal"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetDecimal: %s", err)
	}
	return uint8(res.Uint64()), err
}
