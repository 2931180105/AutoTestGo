package if_ctrl

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"math/big"
)

// TODO: support rest/ws
// TODO: support estimate gas before execute tx

type Comptroller struct {
	Sdk    *ontSDK.OntologySdk
	Signer *ontSDK.Account
	Addr   common.Address

	GasPrice uint64
	GasLimit uint64
}

func NewComptroller(nodeRPCAddr string, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64) (*Comptroller, error) {
	sdk := ontSDK.NewOntologySdk()
	client := sdk.NewRpcClient()
	client.SetAddress(nodeRPCAddr)
	sdk.SetDefaultClient(client)
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewComptroller: cannot access ontology network through addr %s", nodeRPCAddr)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewComptroller: invalid contract addr:%s", contractAddr)
		}
	}
	return &Comptroller{
		Sdk:      sdk,
		Signer:   signer,
		Addr:     addr,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}, nil
}


func (this *Comptroller) UpdateSigner(newSigner *ontSDK.Account) {
	this.Signer = newSigner
}

func (this *Comptroller) GetAddr() common.Address {
	return this.Addr
}

func (this *Comptroller) Init(admin, globalParam, wingAddr, priceOracle, oscoreOracle common.Address, maxSupply,
	insuranceRepayFactor, liquidateIncentiveFactor *big.Int) (string, error) {
	method := "init"
	params := []interface{}{admin, globalParam, wingAddr, priceOracle, oscoreOracle,
		maxSupply, insuranceRepayFactor, liquidateIncentiveFactor}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetPendingAdmin(newPendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{newPendingAdmin}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *Comptroller) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("AcceptAdmin: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetPriceOracle(oracle common.Address) (string, error) {
	method := "_setPriceOracle"
	params := []interface{}{oracle}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPriceOracle: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetOScoreOracle(oracle common.Address) (string, error) {
	method := "_setOScoreOracle"
	params := []interface{}{oracle}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetOScoreOracle: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetGlobalParam(globalParam common.Address) (string, error) {
	method := "_setGlobalParam"
	params := []interface{}{globalParam}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetGlobalParam: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetWingAddr(wing common.Address) (string, error) {
	method := "_setWingAddr"
	params := []interface{}{wing}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetWingAddr: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetInsuranceRepayFactor(factor *big.Int) (string, error) {
	method := "_setInsuranceRepayFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceRepayFactor: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetMaxSupplyValue(value *big.Int) (string, error) {
	method := "_setMaxSupplyValue"
	params := []interface{}{value}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMaxSupplyValue: %s", err)
	}
	return hash, err
}

func (this *Comptroller) MaxSupplyValue() (*big.Int, error) {
	method := "maxSupplyValue"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("MaxSupplyValue: %s", err)
	}
	return res, err
}

func (this *Comptroller) TotalSupplyValue() (*big.Int, error) {
	method := "totalSupplyValue"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalSupplyValue: %s", err)
	}
	return res, err
}

func (this *Comptroller) SetLiquidationIncentive(incentive *big.Int) (string, error) {
	method := "_setLiquidationIncentive"
	params := []interface{}{incentive}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetLiquidationIncentive: %s", err)
	}
	return hash, err
}

func (this *Comptroller) RegisterMarket(marketName string, supplyPool common.Address, borrowPool common.Address,
	insurancePool common.Address, underlying common.Address, underlyingDecimals uint8,
	wingWeight uint8) (string, error) {
	method := "registerMarket"
	params := []interface{}{marketName, supplyPool, borrowPool, insurancePool, underlying,
		underlyingDecimals, wingWeight}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("RegisterMarket: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketUnderlying(market string, underlying common.Address) (string, error) {
	method := "updateMarketUnderlying"
	params := []interface{}{market, underlying}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketUnderlying: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketSupplyPool(market string, supplyPool common.Address) (string, error) {
	method := "updateMarketSupplyPool"
	params := []interface{}{market, supplyPool}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketSupplyPool: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketBorrowPool(market string, borrowPool common.Address) (string, error) {
	method := "updateMarketBorrowPool"
	params := []interface{}{market, borrowPool}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketBorrowPool: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketInsurancePool(market string, insurancePool common.Address) (string, error) {
	method := "updateMarketInsurancePool"
	params := []interface{}{market, insurancePool}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketInsurancePool: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketUnderlyingDecimals(market string, underlyingDecimals uint64) (string, error) {
	method := "updateMarketUnderlyingDecimals"
	params := []interface{}{market, underlyingDecimals}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketUnderlyingDecimals: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketWingWeight(market string, weight uint64) (string, error) {
	method := "updateMarketWingWeight"
	params := []interface{}{market, weight}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketWingWeight: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateWingSBI(market common.Address, supplyPortion, borrowPortion, insurancePortion uint8) (string, error) {
	method := "updateWingSBI"
	params := []interface{}{market, supplyPortion, borrowPortion, insurancePortion}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateWingSBI: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetPauseGuardian(newPausedGuardian common.Address) (string, error) {
	method := "_setPauseGuardian"
	params := []interface{}{newPausedGuardian}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPauseGuardian: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetMintPaused(market common.Address, state bool) (string, error) {
	method := "_setMintPaused"
	params := []interface{}{market, state}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMintPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetBorrowPaused(market common.Address, state bool) (string, error) {
	method := "_setBorrowPaused"
	params := []interface{}{market, state}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetBorrowPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetTransferPaused(state bool) (string, error) {
	method := "_setTransferPaused"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetTransferPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetSeizePaused(state bool) (string, error) {
	method := "_setSeizePaused"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetSeizePaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetDistributeWingSwitch(state bool) (string, error) {
	method := "setDistributeWingSwitch"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetDistributeWingSwitch: %s", err)
	}
	return hash, err
}

func (this *Comptroller) RepayInsurance(borrower common.Address) (string, error) {
	method := "repayInsurance"
	params := []interface{}{borrower}
	hash, err := utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayInsurance: %s", err)
	}
	return hash, err
}

func (this *Comptroller) ClaimWing(holder common.Address, preExecute bool) (hash string,
	remains *big.Int, err error) {
	method := "claimWing"
	params := []interface{}{holder}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWing: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWing: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimWingAtMarkets(holder common.Address, markets []string,
	preExecute bool) (hash string, remains *big.Int, err error) {
	method := "claimWingAtMarkets"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{holder, marketsParam}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWingAtMarkets: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWingAtMarkets: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimAllWing(holders []common.Address, markets []string, borrows, suppliers,
	insurance bool, preExecute bool) (hash string, remains *big.Int, err error) {
	method := "claimAllWing"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	holdersParam := make([]interface{}, 0)
	for _, h := range holders {
		holdersParam = append(holdersParam, h)
	}
	params := []interface{}{holdersParam, marketsParam, borrows, suppliers, insurance}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllWing: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllWing: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimInterest(holder common.Address, preExecute bool) (hash string,
	remains *big.Int, err error) {
	method := "claimInterest"
	params := []interface{}{holder}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimInterest: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimInterest: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimInterestAtMarkets(holder common.Address, markets []common.Address,
	preExecute bool) (hash string, distributed, remains *big.Int, err error) {
	method := "claimInterestAtMarkets"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{holder, marketsParam}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimInterestAtMarkets: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimInterestAtMarkets: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimAllInterest(holder common.Address, markets []common.Address, borrows, suppliers,
	insurance bool, preExecute bool) (hash string, distributed, remains *big.Int, err error) {
	method := "claimAllInterest"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{holder, marketsParam, borrows, suppliers, insurance}
	if !preExecute {
		hash, err = utils.InvokeTx(this.Sdk, this.Signer, this.GasPrice, this.GasLimit, this.Addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllInterest: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaim(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllInterest: %s", err)
		}
		return
	}
}

func (this *Comptroller) preExecuteClaim(method string, params []interface{}) (remains *big.Int, err error) {
	res, err := this.Sdk.WasmVM.PreExecInvokeWasmVMContract(this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("preExecuteClaim: %s", err)
		return
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		err = fmt.Errorf("preExecuteClaim: %s", err)
		return
	}
	source := common.NewZeroCopySource(data)
	remainsNum, eof := source.NextI128()
	if eof {
		err = fmt.Errorf("preExecuteClaim: read remains eof")
		return
	}
	return remainsNum.ToBigInt(), nil
}

/* pre-execute */

func (this *Comptroller) GetAccountLiquidity(account common.Address) (*AccountLiquidity, error) {
	method := "getAccountLiquidity"
	params := []interface{}{account}
	res, err := this.Sdk.WasmVM.PreExecInvokeWasmVMContract(this.Addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("GetAccountLiquidity: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetAccountLiquidity: %s", err)
	}
	result, err := DeserializeAccountLiquidity(data)
	if err != nil {
		return nil, fmt.Errorf("GetAccountLiquidity: %s", err)
	}
	return result, nil
}

func (this *Comptroller) LiquidateCalculate(account common.Address) (string, *big.Int, error) {
	method := "liquidateCalculate"
	params := []interface{}{account}
	res, err := this.Sdk.WasmVM.PreExecInvokeWasmVMContract(this.Addr, method, params)
	if err != nil {
		return "", nil, fmt.Errorf("LiquidateCalculateSeizeTokens: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return "", nil, fmt.Errorf("LiquidateCalculateSeizeTokens: %s", err)
	}
	source := common.NewZeroCopySource(data)
	errStr, _, ill, eof := source.NextString()
	if ill {
		return "", nil, fmt.Errorf("read errStr ill")
	}
	if eof {
		return "", nil, fmt.Errorf("read errStr eof")
	}
	num, eof := source.NextI128()
	if eof {
		return "", nil, fmt.Errorf("read num eof")
	}
	return errStr, num.ToBigInt(), nil
}

func (this *Comptroller) Admin() (common.Address, error) {
	method := "admin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *Comptroller) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("PendingAdmin: %s", err)
	}
	return res, err
}

func (this *Comptroller) GlobalParam() (common.Address, error) {
	method := "globalParam"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("GlobalParam: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingAddr() (common.Address, error) {
	method := "wingAddr"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingAddr: %s", err)
	}
	return res, err
}

func (this *Comptroller) PriceOracle() (common.Address, error) {
	method := "priceOracle"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("PriceOracle: %s", err)
	}
	return res, err
}

func (this *Comptroller) OScoreOracle() (common.Address, error) {
	method := "oScoreOracle"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("OScoreOracle: %s", err)
	}
	return res, err
}

func (this *Comptroller) MarketInfo(market string) (*MarketInfo, error) {
	method := "marketInfo"
	params := []interface{}{market}
	res, err := this.Sdk.WasmVM.PreExecInvokeWasmVMContract(this.Addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("MarketInfo: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("MarketInfo: %s", err)
	}
	result, err := DeserializeMarketInfo(data)
	if err != nil {
		return nil, fmt.Errorf("MarketInfo: %s", err)
	}
	return result, nil
}

func (this *Comptroller) LiquidationIncentiveFactor() (*big.Int, error) {
	method := "liquidationIncentiveFactor"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("LiquidationIncentiveFactor: %s", err)
	}
	return res, err
}

func (this *Comptroller) InsuranceRepayFactor() (*big.Int, error) {
	method := "insuranceRepayFactor"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceRepayFactor: %s", err)
	}
	return res, err
}

func (this *Comptroller) PauseGuardian() (common.Address, error) {
	method := "pauseGuardian"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("PauseGuardian: %s", err)
	}
	return res, err
}

func (this *Comptroller) TransferGuardianPaused() (bool, error) {
	method := "transferGuardianPaused"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) SeizeGuardianPaused() (bool, error) {
	method := "seizeGuardianPaused"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("SeizeGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) MintGuardianPaused(market common.Address) (bool, error) {
	method := "mintGuardianPaused"
	params := []interface{}{market}
	res, err := utils.PreExecuteBool(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("MintGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) BorrowGuardianPaused(market common.Address) (bool, error) {
	method := "borrowGuardianPaused"
	params := []interface{}{market}
	res, err := utils.PreExecuteBool(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) AllMarkets() ([]string, error) {
	method := "allMarkets"
	params := []interface{}{}
	res, err := utils.PreExecuteStringArray(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("AllMarkets: %s", err)
	}
	return res, err
}

func (this *Comptroller) AccountBorrowMarket(account common.Address) (string, error) {
	method := "accountBorrowMarket"
	params := []interface{}{account}
	res, err := utils.PreExecuteString(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccountBorrowMarket: %s", err)
	}
	return res, err
}

func (this *Comptroller) AccountBorrowDay(account common.Address) (uint64, error) {
	method := "accountBorrowDay"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccountBorrowDay: %s", err)
	}
	return res.Uint64(), err
}

func (this *Comptroller) AccountCollateralPool(account common.Address) ([]string, error) {
	method := "accountCollateralPool"
	params := []interface{}{account}
	res, err := utils.PreExecuteStringArray(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccountBorrowDay: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingSupplyIndex(market string) (*big.Int, error) {
	method := "wingSupplyIndex"
	params := []interface{}{market}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingSupplyIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingBorrowIndex(market string) (*big.Int, error) {
	method := "wingBorrowIndex"
	params := []interface{}{market}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingSupplyIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingInsuranceIndex(market string) (*big.Int, error) {
	method := "wingInsuranceIndex"
	params := []interface{}{market}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingInsuranceIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingBorrowIndexSnapshot(market string, dayNumber uint64) (*big.Int, error) {
	method := "wingBorrowIndexSnapshot"
	params := []interface{}{market, dayNumber}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingBorrowIndexSnapshot: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingSupplierIndex(market string, account common.Address) (*big.Int, error) {
	method := "wingSupplierIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingSupplierIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingBorrowerIndex(market string, account common.Address) (*big.Int, error) {
	method := "wingBorrowerIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("wingBorrowerIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingGuarantorIndex(market string, account common.Address) (*big.Int, error) {
	method := "wingGuarantorIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingGuarantorIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingUserAccrued(account common.Address) (*big.Int, error) {
	method := "wingUserAccrued"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingUserAccrued: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingUserPunished(account common.Address) (*big.Int, error) {
	method := "wingUserPunished"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingUserPunished: %s", err)
	}
	return res, err
}

func (this *Comptroller) LastWingBalance() (*big.Int, error) {
	method := "lastWingBalance"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("LastWingBalance: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingMarketAccrued(market string) (*big.Int, error) {
	method := "wingMarketAccrued"
	params := []interface{}{market}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingMarketAccrued: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingSBIPortion(market common.Address) (*WingSBI, error) {
	method := "wingSBIPortion"
	params := []interface{}{market}
	res, err := this.Sdk.WasmVM.PreExecInvokeWasmVMContract(this.Addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	result, err := DeserializeWingSBI(data)
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	return result, nil
}

func (this *Comptroller) WingDistributedNum(market string) (*big.Int, error) {
	method := "wingDistributedNum"
	params := []interface{}{market}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingDistributedNum: %s", err)
	}
	return res, err
}

func (this *Comptroller) InterestIndex(market string) (*big.Int, error) {
	method := "interestIndex"
	params := []interface{}{market}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("InterestIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) InterestUserIndex(market string, account common.Address) (*big.Int, error) {
	method := "interestUserIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteU256(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("InterestUserIndex: %s", err)
	}
	return res, err
}

func (this *Comptroller) UserCompletedBorrowNum(account common.Address) (*big.Int, error) {
	method := "userCompletedBorrowNum"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("UserCompletedBorrowNum: %s", err)
	}
	return res, err
}

func (this *Comptroller) IsComptroller() (bool, error) {
	method := "isComptroller"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.Sdk, this.Addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsComptroller: %s", err)
	}
	return res, err
}
