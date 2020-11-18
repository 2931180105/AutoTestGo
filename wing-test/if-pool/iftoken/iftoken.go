package iftoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"math/big"
)

// TODO: support rest/ws
// TODO: support estimate gas before execute tx

// TODO: adapt to latest supply pool

type IFToken struct {
	sdk    *ontSDK.OntologySdk
	signer *ontSDK.Account
	addr   common.Address

	gasPrice uint64
	gasLimit uint64
}

func NewIFToken(nodeRPCAddr string, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64) (*IFToken, error) {
	sdk := ontSDK.NewOntologySdk()
	client := sdk.NewRpcClient()
	client.SetAddress(nodeRPCAddr)
	sdk.SetDefaultClient(client)
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewIFToken: cannot access ontology network through addr %s", nodeRPCAddr)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewIFToken: invalid contract addr %s", contractAddr)
		}
	}
	return &IFToken{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}


func (this *IFToken) UpdateSigner(newSigner *ontSDK.Account) {
	this.signer = newSigner
}

func (this *IFToken) GetAddr() common.Address {
	return this.addr
}

func (this *IFToken) Init(marketName string, initExchangeRate *big.Int, admin,
	comptroller common.Address, debtLimit *big.Int) (string, error) {
	method := "init"
	params := []interface{}{marketName, initExchangeRate, admin, comptroller, debtLimit}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *IFToken) SetPendingAdmin(pendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{pendingAdmin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *IFToken) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("_acceptAdmin: %s", err)
	}
	return hash, err
}

func (this *IFToken) SetComptroller(comptroller common.Address) (string, error) {
	method := "_setComptroller"
	params := []interface{}{comptroller}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetComptroller: %s", err)
	}
	return hash, err
}

func (this *IFToken) Mint(minter common.Address, mintAmount *big.Int) (string, error) {
	method := "mint"
	params := []interface{}{minter, mintAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Mint: %s", err)
	}
	return hash, err
}

func (this *IFToken) Redeem(redeemer common.Address, redeemTokens *big.Int) (string, error) {
	method := "redeem"
	params := []interface{}{redeemer, redeemTokens}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Redeem: %s", err)
	}
	return hash, err
}

func (this *IFToken) RedeemUnderlying(redeemer common.Address, redeemAmount *big.Int) (string, error) {
	method := "redeemUnderlying"
	params := []interface{}{redeemer, redeemAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemUnderlying: %s", err)
	}
	return hash, err
}

func (this *IFToken) Transfer(from, to common.Address, amount *big.Int) (string, error) {
	method := "transfer"
	params := []interface{}{from, to, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Transfer: %s", err)
	}
	return hash, err
}

func (this *IFToken) TransferFrom(from, src, to common.Address, amount *big.Int) (string, error) {
	method := "transferFrom"
	params := []interface{}{from, src, to, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferFrom: %s", err)
	}
	return hash, err
}

func (this *IFToken) Approve(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Approve: %s", err)
	}
	return hash, err
}

func (this *IFToken) NeoVMApprove(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err := utils.InvokeNeoVMTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr,
		method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMApprove: %s", err)
	}
	return hash, err
}

/* pre execute */

func (this *IFToken) BalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOf: %s", err)
	}
	return res, err
}

func (this *IFToken) NeoVMBalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.NeoVMPreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMBalanceOf: %s", err)
	}
	return res, err
}

func (this *IFToken) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	method := "balanceOfUnderlying"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOfUnderlying: %s", err)
	}
	return res, err
}

func (this *IFToken) ExchangeRate() (*big.Int, error) {
	method := "exchangeRate"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ExchangeRate: %s", err)
	}
	return res, err
}

func (this *IFToken) Allowance(owner, spender common.Address) (*big.Int, error) {
	method := "allowance"
	params := []interface{}{owner, spender}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Allowance: %s", err)
	}
	return res, err
}

func (this *IFToken) TotalCash() (*big.Int, error) {
	method := "totalCash"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalCash: %s", err)
	}
	return res, err
}

func (this *IFToken) TotalDebt() (*big.Int, error) {
	method := "totalDebt"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalDebt: %s", err)
	}
	return res, err
}

func (this *IFToken) TotalSupply() (*big.Int, error) {
	method := "totalSupply"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalSupply: %s", err)
	}
	return res, err
}

func (this *IFToken) Name() (string, error) {
	method := "name"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Name: %s", err)
	}
	return res, err
}

func (this *IFToken) Symbol() (string, error) {
	method := "symbol"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Symbol: %s", err)
	}
	return res, err
}

func (this *IFToken) Decimals() (string, error) {
	method := "Decimals"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Decimals: %s", err)
	}
	return res, err
}

func (this *IFToken) Admin() (common.Address, error) {
	method := "admin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *IFToken) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PendingAdmin: %s", err)
	}
	return res, err
}

func (this *IFToken) Comptroller() (common.Address, error) {
	method := "comptroller"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Comptroller: %s", err)
	}
	return res, err
}

func (this *IFToken) IsSupply() (bool, error) {
	method := "isSupply"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsSupply: %s", err)
	}
	return res, err
}
