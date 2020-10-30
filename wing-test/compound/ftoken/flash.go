package ftoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"math/big"
)

// TODO: support rest/ws
// TODO: support estimate gas before execute tx

type FlashToken struct {
	sdk    *ontSDK.OntologySdk
	signer *ontSDK.Account
	addr   common.Address
	Comptroller *comptroller.Comptroller
	TestConfig *config.Config
	gasPrice uint64
	gasLimit uint64
}
type AccountSnapshot struct {
	TokenBalance  *big.Int
	BorrowBalance *big.Int
	ExchangeRate  *big.Int
}

func NewFlashToken(sdk *ontSDK.OntologySdk, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64, ) (*FlashToken, error) {
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewFlashToken: cannot access ontology network through addr %s", err)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewFlashToken: invalid contract addr %s", contractAddr)
		}
	}
	return &FlashToken{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}
func NewFlashToken2(sdk *ontSDK.OntologySdk, contractAddr string, signer *ontSDK.Account, cfg *config.Config,comptroller *comptroller.Comptroller) (*FlashToken, error) {
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewFlashToken: cannot access ontology network through addr %s", err)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewFlashToken: invalid contract addr %s", contractAddr)
		}
	}
	return &FlashToken{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: cfg.GasPrice,
		gasLimit: cfg.GasLimit,
		TestConfig:cfg,
		Comptroller: comptroller,
	}, nil
}

func (this *FlashToken) UpdateSigner(newSigner *ontSDK.Account) {
	this.signer = newSigner
}

func (this *FlashToken) GetAddr() common.Address {
	return this.addr
}

func (this *FlashToken) Init(admin, underlying_ common.Address, underlyingName string, comptroller_,
	globalParamContract common.Address, interestRateModel common.Address,
	initialExchangeRateMantissa *big.Int) (string, error) {
	method := "init"
	params := []interface{}{admin, underlying_, underlyingName, comptroller_, globalParamContract, interestRateModel,
		initialExchangeRateMantissa}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetGlobalParam(globalParam common.Address) (string, error) {
	method := "_setGlobalParam"
	params := []interface{}{globalParam}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetGlobalParam: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetPriceOracle(oracle common.Address) (string, error) {
	method := "_setPriceOracle"
	params := []interface{}{oracle}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPriceOracle: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetComptroller(comptroller common.Address) (string, error) {
	method := "_setComptroller"
	params := []interface{}{comptroller}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetComptroller: %s", err)
	}
	return hash, err
}

func (this *FlashToken) InsuranceMintPaused() (bool, error) {
	method := "insuranceMintPaused"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceMintPaused: %s", err)
	}
	return res, err
}

func (this *FlashToken) UpdateInsuranceMintPaused(state bool) (string, error) {
	method := "_updateInsuranceMintPaused"
	params := []interface{}{state}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateInsuranceMintPaused: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetReserveFactor(factor *big.Int) (string, error) {
	method := "_setReserveFactor"
	params := []interface{}{factor}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetReserveFactor: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetInsuranceFactor(factor *big.Int) (string, error) {
	method := "_setInsuranceFactor"
	params := []interface{}{factor}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceFactor: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetInterestRateModel(interestRateModel common.Address) (string, error) {
	method := "_setInterestRateModel"
	params := []interface{}{interestRateModel}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInterestRateModel: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetMarketAddr(market common.Address) (string, error) {
	method := "_setMarketAddr"
	params := []interface{}{market}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMarketAddr: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetInsuranceAddr(insurance common.Address) (string, error) {
	method := "_setInsuranceAddr"
	params := []interface{}{insurance}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceAddr: %s", err)
	}
	return hash, err
}

func (this *FlashToken) Transfer(from, to common.Address, amount *big.Int) (string, error) {
	method := "transfer"
	params := []interface{}{from, to, amount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Transfer: %s", err)
	}
	return hash, err
}

func (this *FlashToken) TransferFrom(from, src, to common.Address, amount *big.Int) (string, error) {
	method := "transferFrom"
	params := []interface{}{from, src, to, amount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferFrom: %s", err)
	}
	return hash, err
}

func (this *FlashToken) Approve(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Approve: %s", err)
	}
	return hash, err
}

func (this *FlashToken) NeoVMApprove(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err := utils.InvokeNeoVMTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr,
		method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMApprove: %s", err)
	}
	return hash, err
}

func (this *FlashToken) AccrueInterest() (string, error) {
	method := "accrueInterest"
	params := []interface{}{}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccrueInterest: %s", err)
	}
	return hash, err
}

func (this *FlashToken) SetPendingAdmin(pendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{pendingAdmin}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *FlashToken) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("_acceptAdmin: %s", err)
	}
	return hash, err
}

func (this *FlashToken) ReduceReserves(reduceAmount *big.Int) (string, error) {
	method := "_reduceReserves"
	params := []interface{}{reduceAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ReduceReserves: %s", err)
	}
	return hash, err
}

func (this *FlashToken) Mint(minter common.Address, mintAmount *big.Int) (string, error) {
	method := "mint"
	params := []interface{}{minter, mintAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Mint: %s", err)
	}
	return hash, err
}

func (this *FlashToken) Redeem(redeemer common.Address, redeemTokens *big.Int) (string, error) {
	method := "redeem"
	params := []interface{}{redeemer, redeemTokens}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Redeem: %s", err)
	}
	return hash, err
}

func (this *FlashToken) RedeemUnderlying(redeemer common.Address, redeemAmount *big.Int) (string, error) {
	method := "redeemUnderlying"
	params := []interface{}{redeemer, redeemAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemUnderlying: %s", err)
	}
	return hash, err
}

func (this *FlashToken) Borrow(borrower common.Address, borrowAmount *big.Int) (string, error) {
	method := "borrow"
	params := []interface{}{borrower, borrowAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Borrow: %s", err)
	}
	return hash, err
}

func (this *FlashToken) RepayBorrow(borrower common.Address, repayAmount *big.Int) (string, error) {
	method := "repayBorrow"
	params := []interface{}{borrower, repayAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayBorrow: %s", err)
	}
	return hash, err
}

func (this *FlashToken) RepayBorrowBehalf(payer, borrower common.Address, repayAmount *big.Int) (string, error) {
	method := "repayBorrowBehalf"
	params := []interface{}{payer, borrower, repayAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayBorrowBehalf: %s", err)
	}
	return hash, err
}

func (this *FlashToken) LiquidateBorrow(liquidator, borrower common.Address, repayAmount *big.Int,
	tokenCollateral common.Address) (string, error) {
	method := "liquidateBorrow"
	params := []interface{}{liquidator, borrower, repayAmount, tokenCollateral}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("LiquidateBorrow: %s", err)
	}
	return hash, err
}

func (this *FlashToken) AddReserves(from common.Address, addAmount *big.Int, ) (string, error) {
	method := "_addReserves"
	params := []interface{}{from, addAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AddReserves: %s", err)
	}
	return hash, err
}

func (this *FlashToken) AddInsurance(from common.Address, addAmount *big.Int, ) (string, error) {
	method := "_addInsurance"
	params := []interface{}{from, addAmount}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AddInsurance: %s", err)
	}
	return hash, err
}

/* pre execute */

func (this *FlashToken) Allowance(owner, spender common.Address) (*big.Int, error) {
	method := "allowance"
	params := []interface{}{owner, spender}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Allowance: %s", err)
	}
	return res, err
}

func (this *FlashToken) BalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOf: %s", err)
	}
	return res, err
}

func (this *FlashToken) NeoVMBalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.NeoVMPreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMBalanceOf: %s", err)
	}
	return res, err
}

func (this *FlashToken) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	method := "balanceOfUnderlying"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOfUnderlying: %s", err)
	}
	return res, err
}



func DeserializeAccountSnapshot(data []byte) (*AccountSnapshot, error) {
	source := common.NewZeroCopySource(data)
	t, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read TokenBalance eof")
	}
	b, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read BorrowBalance eof")
	}
	e, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read ExchangeRate eof")
	}
	return &AccountSnapshot{
		TokenBalance:  t.ToBigInt(),
		BorrowBalance: b.ToBigInt(),
		ExchangeRate:  e.ToBigInt(),
	}, nil
}

func (this *FlashToken) AccountSnapshot(owner common.Address) (*AccountSnapshot, error) {
	method := "accountSnapshot"
	params := []interface{}{owner}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("AccountSnapshot: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("AccountSnapshot: %s", err)
	}
	result, err := DeserializeAccountSnapshot(data)
	if err != nil {
		return nil, fmt.Errorf("AccountSnapshot: %s", err)
	}
	return result, nil
}

func (this *FlashToken) BorrowRatePerBlock() (*big.Int, error) {
	method := "borrowRatePerBlock"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowRatePerBlock: %s", err)
	}
	return res, err
}

func (this *FlashToken) SupplyRatePerBlock() (*big.Int, error) {
	method := "supplyRatePerBlock"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SupplyRatePerBlock: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetCash() (*big.Int, error) {
	method := "getCash"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetCash: %s", err)
	}
	return res, err
}

func (this *FlashToken) ExchangeRateStored() (*big.Int, error) {
	method := "exchangeRateStored"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ExchangeRateStored: %s", err)
	}
	return res, err
}

func (this *FlashToken) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	method := "borrowBalanceStored"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowBalanceStored: %s", err)
	}
	return res, err
}

func (this *FlashToken) TotalBorrows() (*big.Int, error) {
	method := "totalBorrows"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowBalanceStored: %s", err)
	}
	return res, err
}

func (this *FlashToken) GlobalParam() (common.Address, error) {
	method := "globalParam"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GlobalParam: %s", err)
	}
	return res, err
}

func (this *FlashToken) Name() (string, error) {
	method := "name"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Name: %s", err)
	}
	return res, err
}

func (this *FlashToken) Symbol() (string, error) {
	method := "symbol"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Symbol: %s", err)
	}
	return res, err
}

func (this *FlashToken) Decimals() (string, error) {
	method := "Decimals"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Decimals: %s", err)
	}
	return res, err
}

func (this *FlashToken) TotalSupply() (*big.Int, error) {
	method := "totalSupply"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalSupply: %s", err)
	}
	return res, err
}

func (this *FlashToken) Admin() (common.Address, error) {
	method := "admin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *FlashToken) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PendingAdmin: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetComptroller() (common.Address, error) {
	method := "comptroller"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Comptroller: %s", err)
	}
	return res, err
}

func (this *FlashToken) InterestRateModel() (common.Address, error) {
	method := "interestRateModel"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InterestRateModel: %s", err)
	}
	return res, err
}

func (this *FlashToken) InitialExchangeRateMantissa() (*big.Int, error) {
	method := "initialExchangeRateMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InitialExchangeRateMantissa: %s", err)
	}
	return res, err
}

func (this *FlashToken) ReserveFactorMantissa() (*big.Int, error) {
	method := "reserveFactorMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ReserveFactorMantissa: %s", err)
	}
	return res, err
}

func (this *FlashToken) InsuranceFactorMantissa() (*big.Int, error) {
	method := "insuranceFactorMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceFactorMantissa: %s", err)
	}
	return res, err
}

func (this *FlashToken) AccrualBlockNumber() (uint32, error) {
	method := "accrualBlockNumber"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccrualBlockNumber: %s", err)
		return 0, err
	}
	return uint32(res.Uint64()), nil
}

func (this *FlashToken) BorrowIndex() (*big.Int, error) {
	method := "borrowIndex"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowIndex: %s", err)
	}
	return res, err
}

func (this *FlashToken) TotalReserves() (*big.Int, error) {
	method := "totalReserves"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalReserves: %s", err)
	}
	return res, err
}

func (this *FlashToken) TotalInsurance() (*big.Int, error) {
	method := "totalInsurance"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalInsurance: %s", err)
	}
	return res, err
}

func (this *FlashToken) Underlying() (common.Address, error) {
	method := "underlying"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Underlying: %s", err)
	}
	return res, err
}

func (this *FlashToken) UnderlyingName() (string, error) {
	method := "underlyingName"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("UnderlyingName: %s", err)
	}
	return res, err
}

func (this *FlashToken) MarketAddr() (common.Address, error) {
	method := "marketAddr"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("MarketAddr: %s", err)
	}
	return res, err
}

func (this *FlashToken) InsuranceAddr() (common.Address, error) {
	method := "insuranceAddr"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceAddr: %s", err)
	}
	return res, err
}

func (this *FlashToken) IsFToken() (bool, error) {
	method := "isFToken"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsFToken: %s", err)
	}
	return res, err
}

func (this *FlashToken) IsInsurance() (bool, error) {
	method := "isInsurance"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsInsurance: %s", err)
	}
	return res, err
}

/* insurance redeem */

func (this *FlashToken) RedeemApply(redeemer common.Address, redeemTokens *big.Int) (string, error) {
	method := "redeemApply"
	params := []interface{}{redeemer, redeemTokens}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemApply: %s", err)
	}
	return hash, err
}

func (this *FlashToken) RedeemApplySettlement() (string, error) {
	method := "redeemApplySettlement"
	params := []interface{}{}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemApplySettlement: %s", err)
	}
	return hash, err
}

func (this *FlashToken) PutSettleInterval(interval uint64) (string, error) {
	method := "putSettleInterval"
	params := []interface{}{interval}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutSettleInterval: %s", err)
	}
	return hash, err
}

func (this *FlashToken) PutSettlementAmtMax(amt *big.Int) (string, error) {
	method := "putSettlementAmtMax"
	params := []interface{}{amt}
	hash, err:= utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PutSettlementAmtMax: %s", err)
	}
	return hash, err
}

func (this *FlashToken) GetInsuranceRedeemApply(account common.Address) (*big.Int, error) {
	method := "getInsuranceRedeemApply"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetInsuranceRedeemApply: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetRedeemApplyTotal() (*big.Int, error) {
	method := "getRedeemApplyTotal"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetRedeemApplyTotal: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetSettleInterval() (*big.Int, error) {
	method := "getSettleInterval"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetSettleInterval: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetSettlementAmtMax() (*big.Int, error) {
	method := "getSettlementAmtMax"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetSettlementAmtMax: %s", err)
	}
	return res, err
}

func (this *FlashToken) GetLastSettleTime() (*big.Int, error) {
	method := "getLastSettleTime"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetLastSettleTime: %s", err)
	}
	return res, err
}


func GetITokenAddress(genSdk *ontSDK.OntologySdk, ftokenAddress common.Address) (common.Address, error) {
	preExecResult, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(ftokenAddress,
		"insuranceAddr", []interface{}{})
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("GetITokenAddress, this.sdk.WasmVM.PreExecInvokeWasmVMContract error: %s", err)
	}
	r, err := preExecResult.Result.ToByteArray()
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("GetITokenAddress, preExecResult.Result.ToByteArray error: %s", err)
	}
	insuranceAddress, err := common.AddressParseFromBytes(r)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("GetITokenAddress, common.AddressParseFromBytes error: %s", err)
	}
	return insuranceAddress, nil
}
