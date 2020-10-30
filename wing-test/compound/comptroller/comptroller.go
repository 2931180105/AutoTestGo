package comptroller

import (
	"encoding/json"
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/payload"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/states"
	"io/ioutil"
	"math/big"
	"time"
)


type Comptroller struct {
	sdk    *ontSDK.OntologySdk
	signer *ontSDK.Account
	addr   common.Address

	gasPrice uint64
	gasLimit uint64
}

func NewComptroller(sdk *ontSDK.OntologySdk, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64, ) (*Comptroller, error) {

	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewComptroller: cannot access ontology network through addr %s", err)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewComptroller: invalid contract addr %s", contractAddr)
		}
	}
	return &Comptroller{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}

func (this *Comptroller) UpdateSigner(newSigner *ontSDK.Account) {
	this.signer = newSigner
}

func (this *Comptroller) GetAddr() common.Address {
	return this.addr
}

func (this *Comptroller) Init(admin, globalParam common.Address) (string, error) {
	method := "init"
	params := []interface{}{admin, globalParam}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetGlobalParam(globalParam common.Address) (string, error) {
	method := "_setGlobalParam"
	params := []interface{}{globalParam}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetGlobalParam: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetPriceOracle(oracle common.Address) (string, error) {
	method := "_setPriceOracle"
	params := []interface{}{oracle}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPriceOracle: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetWingAddr(wing common.Address) (string, error) {
	method := "_setWingAddr"
	params := []interface{}{wing}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetWingAddr: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetCloseFactor(factor *big.Int) (string, error) {
	method := "_setCloseFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetCloseFactor: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetInsuranceRepayFactor(factor *big.Int) (string, error) {
	method := "_setInsuranceRepayFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceRepayFactor: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetMaxAssets(max *big.Int) (string, error) {
	method := "_setMaxAssets"
	params := []interface{}{max}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMaxAssets: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetLiquidationIncentive(incentive *big.Int) (string, error) {
	method := "_setLiquidationIncentive"
	params := []interface{}{incentive}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetLiquidationIncentive: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetWingRate(rate *big.Int) (string, error) {
	method := "_setWingRate"
	params := []interface{}{rate}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetWingRate: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SupportMarket(fToken, insurance common.Address, underlyingDecimals uint64) (string, error) {
	method := "_supportMarket"
	params := []interface{}{fToken, insurance, underlyingDecimals}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SupportMarket: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketUnderlyingDecimals(fToken common.Address, underlyingDecimals uint64) (string, error) {
	method := "_updateMarketUnderlyingDecimals"
	params := []interface{}{fToken, underlyingDecimals}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketUnderlyingDecimals: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetCollateralFactor(market common.Address, factor *big.Int) (string, error) {
	method := "_setCollateralFactor"
	params := []interface{}{market, factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetCollateralFactor: %s", err)
	}
	return hash, err
}

func (this *Comptroller) AddWingMarkets(markets []common.Address, weights []uint64) (string, error) {
	method := "_addWingMarkets"
	sink := common.NewZeroCopySink(nil)
	sink.WriteString(method)
	length := uint64(len(markets))
	sink.WriteVarUint(length)
	for _, v := range markets {
		sink.WriteAddress(v)
	}
	sink.WriteVarUint(length)
	for _, w := range weights {
		sink.WriteI128(common.I128FromUint64(w))
	}

	contract := &states.WasmContractParam{}
	contract.Address = this.addr
	//bf := bytes.NewBuffer(nil)
	argbytes := sink.Bytes()
	contract.Args = argbytes

	invokePayload := &payload.InvokeCode{
		Code: common.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		GasPrice: this.gasPrice,
		GasLimit: this.gasLimit,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	err := utils.SignTx(this.sdk, tx, 0,this.signer)
	if err != nil {
		return "", fmt.Errorf("AddWingMarkets: %s", err)
	}
	hash, err := this.sdk.SendTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("AddWingMarkets: %s", err)
	}
	return hash.ToHexString(), nil
}

func (this *Comptroller) SetPendingAdmin(newPendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{newPendingAdmin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *Comptroller) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AcceptAdmin: %s", err)
	}
	return hash, err
}

func (this *Comptroller) EnterMarkets(from common.Address, markets []common.Address, preExecute bool) ([]string,
	string, error) {
	method := "enterMarkets"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		log.Infof("EnterMarkets: %s", m.ToHexString())
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{from, marketsParam}
	if !preExecute {
		hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
		if err != nil {
			return nil, "", fmt.Errorf("EnterMarkets: %s", err)
		}
		return nil, hash, nil
	} else {
		res, err := utils.PreExecuteStringArray(this.sdk, this.addr, method, params)
		if err != nil {
			return nil, "", fmt.Errorf("EnterMarkets: %s", err)
		}
		return res, "", nil
	}
}

func (this *Comptroller) ExitMarket(from, market common.Address) (string, error) {
	method := "exitMarket"
	params := []interface{}{from, market}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ExitMarket: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetPauseGuardian(newPausedGuardian common.Address) (string, error) {
	method := "_setPauseGuardian"
	params := []interface{}{newPausedGuardian}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPauseGuardian: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetMintPaused(market common.Address, state bool) (string, error) {
	method := "_setMintPaused"
	params := []interface{}{market, state}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMintPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetBorrowPaused(market common.Address, state bool) (string, error) {
	method := "_setBorrowPaused"
	params := []interface{}{market, state}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetBorrowPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetTransferPaused(state bool) (string, error) {
	method := "_setTransferPaused"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetTransferPaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetSeizePaused(state bool) (string, error) {
	method := "_setSeizePaused"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetSeizePaused: %s", err)
	}
	return hash, err
}

func (this *Comptroller) RefreshWingRate() (string, error) {
	method := "refreshWingRate"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RefreshWingRate: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetDistributeWingSwitch(state bool) (string, error) {
	method := "setDistributeWingSwitch"
	params := []interface{}{state}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetDistributeWingSwitch: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetNewMarketId(newMarket, oldMarket common.Address) (string, error) {
	method := "setNewMarketID"
	params := []interface{}{newMarket, oldMarket}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetNewMarketId: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetWingIndexFactor(market common.Address, factor *big.Int) (string, error) {
	method := "_setWingDeltaIndexFactor"
	params := []interface{}{market, factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("MigrateMarket: %s", err)
	}
	return hash, err
}

func (this *Comptroller) SetOldMarketId(oldMarket, newMarket common.Address) (string, error) {
	method := "setOldMarketID"
	params := []interface{}{oldMarket, newMarket}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetOldMarketId: %s", err)
	}
	return hash, err
}

func (this *Comptroller) ClaimWing(holder common.Address, preExecute bool) (hash string,
	remains *big.Int, err error) {
	method := "claimWing"
	params := []interface{}{holder}
	if !preExecute {
		hash, err = utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWing: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaimWing(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWing: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimWingAtMarkets(holder common.Address, markets []common.Address,
	preExecute bool) (hash string, distributed, remains *big.Int, err error) {
	method := "claimWingAtMarkets"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{holder, marketsParam}
	if !preExecute {
		hash, err = utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWingAtMarkets: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaimWing(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimWingAtMarkets: %s", err)
		}
		return
	}
}

func (this *Comptroller) ClaimAllWing(holder common.Address, markets []common.Address, borrows, suppliers,
	insurance bool, preExecute bool) (hash string, distributed, remains *big.Int, err error) {
	method := "claimAllWing"
	marketsParam := make([]interface{}, 0)
	for _, m := range markets {
		marketsParam = append(marketsParam, m)
	}
	params := []interface{}{holder, marketsParam, borrows, suppliers, insurance}
	if !preExecute {
		hash, err = utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllWing: %s", err)
		}
		return
	} else {
		remains, err = this.preExecuteClaimWing(method, params)
		if err != nil {
			err = fmt.Errorf("ClaimAllWing: %s", err)
		}
		return
	}
}

func (this *Comptroller) preExecuteClaimWing(method string, params []interface{}) (remains *big.Int, err error) {
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("preExecuteClaimWing: %s", err)
		return
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		err = fmt.Errorf("preExecuteClaimWing: %s", err)
		return
	}
	source := common.NewZeroCopySource(data)
	remainsNum, eof := source.NextI128()
	if eof {
		err = fmt.Errorf("preExecuteClaimWing: read remains eof")
		return
	}
	return remainsNum.ToBigInt(), nil
}

func (this *Comptroller) SetWingSBIPortion(market common.Address, supplyPortion, borrowPortion, insurancePortion *big.Int) (string, error) {
	method := "_setWingSBIPortion"
	params := []interface{}{market, supplyPortion, borrowPortion, insurancePortion}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetWingSBIPortion: %s", err)
	}
	return hash, err
}

func (this *Comptroller) UpdateMarketWingWeight(market common.Address, weight *big.Int) (string, error) {
	method := "_updateMarketWingWeight"
	params := []interface{}{market, weight}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("UpdateMarketWingWeight: %s", err)
	}
	return hash, err
}

type UserAllowance struct {
	User   common.Address
	Amount *big.Int
}

func UnmarshalWholeUserAllowance(filePath string) ([]*UserAllowance, error) {
	type jsonAllowance struct {
		User   string `json:"user"`
		Amount string `json:"amount"`
	}
	userAllowance := make([]*jsonAllowance, 0)
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalWholeUserAllowance: %s", err)
	}
	if err := json.Unmarshal(fileContent, &userAllowance); err != nil {
		return nil, fmt.Errorf("UnmarshalWholeUserAllowance: %s", err)
	}
	result := make([]*UserAllowance, 0)
	for _, a := range userAllowance {
		addr, err := common.AddressFromBase58(a.User)
		if err != nil {
			return nil, fmt.Errorf("UnmarshalWholeUserAllowance: %s, %s", a.User, err)
		}
		amount, ok := new(big.Int).SetString(a.Amount, 10)
		if !ok {
			return nil, fmt.Errorf("UnmarshalWholeUserAllowance: %s", a.Amount)
		}
		log.Infof("user: %s, allowance: %s, %d", addr.ToBase58(), a.Amount, amount.Uint64())
		result = append(result, &UserAllowance{User: addr, Amount: amount})
	}
	return result, nil
}

func SerializeUserAllowance(sink *common.ZeroCopySink, allowance *UserAllowance) []byte {
	sink.WriteAddress(allowance.User)
	amount, err := common.I128FromBigInt(allowance.Amount)
	if err != nil {
		panic("cannot parse allowance.Amount")
	}
	sink.WriteI128(amount)
	return sink.Bytes()
}

func (this *Comptroller) ApproveWing(allowance []*UserAllowance) (string, error) {
	method := "_approveWing"
	sink := common.NewZeroCopySink(nil)
	sink.WriteString(method)
	contract := &states.WasmContractParam{}
	contract.Address = this.addr
	sink.WriteVarUint(uint64(len(allowance)))
	for _, a := range allowance {
		SerializeUserAllowance(sink, a)
	}
	contract.Args = sink.Bytes()
	invokePayload := &payload.InvokeCode{
		Code: common.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		GasPrice: this.gasPrice,
		GasLimit: this.gasLimit,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	err := utils.SignTx(this.sdk, tx,0, this.signer)
	if err != nil {
		return "", fmt.Errorf("AddWingMarkets: %s", err)
	}
	hash, err := this.sdk.SendTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("AddWingMarkets: %s", err)
	}
	return hash.ToHexString(), nil
}

func (this *Comptroller) DropWingMarket(market common.Address) (string, error) {
	method := "_dropWingMarket"
	params := []interface{}{market}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("DropWingMarket: %s", err)
	}
	return hash, err
}

func (this *Comptroller) RepayByInsurance(borrower common.Address) (string, error) {
	method := "repayByInsurance"
	params := []interface{}{borrower}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayByInsurance: %s", err)
	}
	return hash, err
}

/* pre-execute */

func (this *Comptroller) MintAllowed(market, minter common.Address) (bool, error) {
	method := "mintAllowed"
	params := []interface{}{market, minter}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("MintAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) RedeemAllowed(market, redeemer common.Address, redeemTokens *big.Int) (bool, error) {
	method := "redeemAllowed"
	params := []interface{}{market, redeemer, redeemTokens}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) BorrowAllowed(market, borrower common.Address, borrowAmount *big.Int) (bool, error) {
	method := "borrowAllowed"
	params := []interface{}{market, borrower, borrowAmount}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) RepayBorrowAllowed(market, payer, borrower common.Address, repayAmount *big.Int) (bool, error) {
	method := "repayBorrowAllowed"
	params := []interface{}{market, payer, borrower, repayAmount}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayBorrowAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) LiquidateBorrowAllowed(borrowedMarket, collateralMarket, liquidator, borrower common.Address, repayAmount *big.Int) (bool, error) {
	method := "liquidateBorrowAllowed"
	params := []interface{}{borrowedMarket, collateralMarket, liquidator, borrower, repayAmount}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("LiquidateBorrowAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) SeizeAllowed(collateralMarket, borrowedMarket, liquidator, borrower common.Address, seizeTokens *big.Int) (bool, error) {
	method := "seizeAllowed"
	params := []interface{}{collateralMarket, borrowedMarket, liquidator, borrower, seizeTokens}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SeizeAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) TransferAllowed(collateralMarket, borrowedMarket, liquidator, borrower common.Address, seizeTokens *big.Int) (bool, error) {
	method := "transferAllowed"
	params := []interface{}{collateralMarket, borrowedMarket, liquidator, borrower, seizeTokens}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferAllowed: %s", err)
	}
	return res, err
}

func (this *Comptroller) AssetsIn(account common.Address) ([]common.Address, error) {
	method := "assetsIn"
	params := []interface{}{account}
	res, err := utils.PreExecuteAddrArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AssetsIn: %s", err)
	}
	return res, err
}

func (this *Comptroller) CheckMembership(account, market common.Address) (bool, error) {
	method := "checkMembership"
	params := []interface{}{account, market}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("CheckMembership: %s", err)
	}
	return res, err
}

type AccountLiquidity struct {
	Error     string
	Liquidity common.I128
	Shortfall common.I128
}

func DeserializeAccountLiquidity(data []byte) (*AccountLiquidity, error) {
	source := common.NewZeroCopySource(data)
	errStr, _, ill, eof := source.NextString()
	if ill {
		return nil, fmt.Errorf("read errStr ill")
	}
	if eof {
		return nil, fmt.Errorf("read errStr eof")
	}
	liquidity, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read liquidity eof")
	}
	shortfall, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read shortfall eof")
	}
	return &AccountLiquidity{
		Error:     errStr,
		Liquidity: liquidity,
		Shortfall: shortfall,
	}, nil
}

func (this *Comptroller) GetAccountLiquidity(account common.Address) (*AccountLiquidity, error) {
	method := "getAccountLiquidity"
	params := []interface{}{account}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
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

func (this *Comptroller) GetHypotheticalAccountLiquidity(account, fTokenModify common.Address, redeemTokens,
	borrowAmount *big.Int) (*AccountLiquidity, error) {
	method := "getHypotheticalAccountLiquidity"
	params := []interface{}{account, fTokenModify, redeemTokens, borrowAmount}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("GetHypotheticalAccountLiquidity: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetHypotheticalAccountLiquidity: %s", err)
	}
	result, err := DeserializeAccountLiquidity(data)
	if err != nil {
		return nil, fmt.Errorf("GetHypotheticalAccountLiquidity: %s", err)
	}
	return result, nil
}

func (this *Comptroller) LiquidateCalculateSeizeTokens(account common.Address) (string, *big.Int, error) {
	method := "liquidateCalculateSeizeTokens"
	params := []interface{}{account}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
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
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *Comptroller) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PendingAdmin: %s", err)
	}
	return res, err
}

func (this *Comptroller) GlobalParam() (common.Address, error) {
	method := "globalParam"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GlobalParam: %s", err)
	}
	return res, err
}

func (this *Comptroller) Oracle() (common.Address, error) {
	method := "oracle"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Oracle: %s", err)
	}
	return res, err
}

func (this *Comptroller) CloseFactorMantissa() (*big.Int, error) {
	method := "closeFactorMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("CloseFactorMantissa: %s", err)
	}
	return res, err
}

func (this *Comptroller) InsuranceRepayFactorMantissa() (*big.Int, error) {
	method := "insuranceRepayFactorMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceRepayFactorMantissa: %s", err)
	}
	return res, err
}

func (this *Comptroller) CouldRepayByInsurance(account common.Address) (bool, error) {
	method := "couldRepayByInsurance"
	params := []interface{}{account}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("CouldRepayByInsurance: %s", err)
	}
	return res, err
}

func (this *Comptroller) LiquidationIncentiveMantissa() (*big.Int, error) {
	method := "liquidationIncentiveMantissa"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("LiquidationIncentiveMantissa: %s", err)
	}
	return res, err
}

func (this *Comptroller) MaxAssets() (*big.Int, error) {
	method := "maxAssets"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("MaxAssets: %s", err)
	}
	return res, err
}

func (this *Comptroller) AccountAssets(account common.Address) ([]common.Address, error) {
	method := "accountAssets"
	params := []interface{}{account}
	res, err := utils.PreExecuteAddrArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccountAssets: %s", err)
	}
	return res, err
}

type MarketMeta struct {
	Addr          common.Address
	InsuranceAddr common.Address

	IsListed    bool
	ReceiveWing bool

	WingWeight               *big.Int
	CollateralFactorMantissa *big.Int

	UnderlyingDecimals uint32
}

func DeserializeMarketMeta(data []byte) (*MarketMeta, error) {
	source := common.NewZeroCopySource(data)
	addr, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read addr eof")
	}
	insurance, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read insurance eof")
	}
	isListed, irr, eof := source.NextBool()
	if irr {
		return nil, fmt.Errorf("read isListed irr")
	}
	if eof {
		return nil, fmt.Errorf("read isListed eof")
	}
	receiveWing, irr, eof := source.NextBool()
	if irr {
		return nil, fmt.Errorf("read receiveWing irr")
	}
	if eof {
		return nil, fmt.Errorf("read receiveWing eof")
	}
	wingWeight, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read wingWeight eof")
	}
	collateralFactor, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read collateralFactor eof")
	}
	decimals, eof := source.NextUint32()
	if eof {
		return nil, fmt.Errorf("read collateralFactor eof")
	}
	return &MarketMeta{
		Addr:                     addr,
		InsuranceAddr:            insurance,
		IsListed:                 isListed,
		ReceiveWing:              receiveWing,
		WingWeight:               wingWeight.ToBigInt(),
		CollateralFactorMantissa: collateralFactor.ToBigInt(),
		UnderlyingDecimals:       decimals,
	}, nil
}

func (this *Comptroller) MarketMeta(market common.Address) (*MarketMeta, error) {
	method := "marketMeta"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	result, err := DeserializeMarketMeta(data)
	if err != nil {
		return nil, fmt.Errorf("MarketMeta: %s", err)
	}
	return result, nil
}

func (this *Comptroller) PauseGuardian() (common.Address, error) {
	method := "pauseGuardian"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PauseGuardian: %s", err)
	}
	return res, err
}

func (this *Comptroller) TransferGuardianPaused() (bool, error) {
	method := "transferGuardianPaused"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) SeizeGuardianPaused() (bool, error) {
	method := "seizeGuardianPaused"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SeizeGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) MintGuardianPaused(market common.Address) (bool, error) {
	method := "mintGuardianPaused"
	params := []interface{}{market}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("MintGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) BorrowGuardianPaused(market common.Address) (bool, error) {
	method := "borrowGuardianPaused"
	params := []interface{}{market}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowGuardianPaused: %s", err)
	}
	return res, err
}

func (this *Comptroller) AllMarkets() ([]common.Address, error) {
	method := "allMarkets"
	params := []interface{}{}
	res, err := utils.PreExecuteAddrArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AllMarkets: %s", err)
	}
	return res, err
}

func (this *Comptroller) AllRawMarkets() ([]common.Address, error) {
	method := "allRawMarkets"
	params := []interface{}{}
	res, err := utils.PreExecuteAddrArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AllMarkets: %s", err)
	}
	return res, err
}

func (this *Comptroller) IsMarketExisted(market common.Address) (bool, error) {
	method := "isMarketExisted"
	params := []interface{}{market}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsMarketExisted: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingDistributedNum(market common.Address) (*big.Int, error) {
	method := "wingDistributedNum"
	params := []interface{}{market}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingDistributedNum: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingAddr() (common.Address, error) {
	method := "wingAddr"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingAddr: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingRate() (*big.Int, error) {
	method := "wingRate"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingRate: %s", err)
	}
	return res, err
}

func (this *Comptroller) WingSpeeds(market common.Address) (*big.Int, error) {
	method := "wingSpeeds"
	params := []interface{}{market}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingSpeeds: %s", err)
	}
	return res, err
}

type WingSBI struct {
	SupplyPortion    uint64
	BorrowPortion    uint64
	InsurancePortion uint64
}

func DeserializeWingSBI(data []byte) (*WingSBI, error) {
	source := common.NewZeroCopySource(data)
	supply, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read supply eof")
	}
	borrow, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read borrow eof")
	}
	insurance, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read insurance eof")
	}
	return &WingSBI{
		SupplyPortion:    supply.ToBigInt().Uint64(),
		BorrowPortion:    borrow.ToBigInt().Uint64(),
		InsurancePortion: insurance.ToBigInt().Uint64(),
	}, nil
}

func (this *Comptroller) WingSBIPortion(market common.Address) (*WingSBI, error) {
	method := "wingSBIPortion"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
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

type WingMarketState struct {
	Index          *big.Int
	BlockTimestamp uint64
}

func DeserializeWingMarketState(data []byte) (*WingMarketState, error) {
	source := common.NewZeroCopySource(data)
	index, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read index eof")
	}
	timestamp, eof := source.NextUint64()
	if eof {
		return nil, fmt.Errorf("read timestamp eof")
	}
	return &WingMarketState{
		Index:          index.ToBigInt(),
		BlockTimestamp: timestamp,
	}, nil
}

func (this *Comptroller) WingSupplyState(market common.Address) (*WingMarketState, error) {
	method := "wingSupplyState"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("WingSupplyState: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("WingSupplyState: %s", err)
	}
	result, err := DeserializeWingMarketState(data)
	if err != nil {
		return nil, fmt.Errorf("WingSupplyState: %s", err)
	}
	return result, nil
}

func (this *Comptroller) WingBorrowState(market common.Address) (*WingMarketState, error) {
	method := "wingBorrowState"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("WingBorrowState: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("WingBorrowState: %s", err)
	}
	result, err := DeserializeWingMarketState(data)
	if err != nil {
		return nil, fmt.Errorf("WingBorrowState: %s", err)
	}
	return result, nil
}

func (this *Comptroller) WingInsuranceState(market common.Address) (*WingMarketState, error) {
	method := "wingInsuranceState"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("WingInsuranceState: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("WingInsuranceState: %s", err)
	}
	result, err := DeserializeWingMarketState(data)
	if err != nil {
		return nil, fmt.Errorf("WingInsuranceState: %s", err)
	}
	return result, nil
}

func (this *Comptroller) WingSupplierIndex(market, account common.Address) (*big.Int, error) {
	method := "wingSupplierIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingSupplierIndex: %s", err)
	}
	return res, nil
}

func (this *Comptroller) WingBorrowerIndex(market, account common.Address) (*big.Int, error) {
	method := "wingBorrowerIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("wingBorrowerIndex: %s", err)
	}
	return res, nil
}

func (this *Comptroller) WingInsuranceIndex(market, account common.Address) (*big.Int, error) {
	method := "wingInsuranceIndex"
	params := []interface{}{market, account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingInsuranceIndex: %s", err)
	}
	return res, nil
}

func (this *Comptroller) WingAccrued(account common.Address) (*big.Int, error) {
	method := "wingAccrued"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingAccrued: %s", err)
	}
	return res, nil
}

func (this *Comptroller) GetNewMarketId(newMarket common.Address) (common.Address, error) {
	method := "getNewMarketID"
	params := []interface{}{newMarket}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetNewMarketId: %s", err)
	}
	return res, nil
}

func (this *Comptroller) GetOldMarketId(oldMarket common.Address) (common.Address, error) {
	method := "getOldMarketID"
	params := []interface{}{oldMarket}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetOldMarketId: %s", err)
	}
	return res, nil
}

type WingIndexFactorInfo struct {
	Factor         *big.Int
	SupplyIndex    *big.Int
	InsuranceIndex *big.Int
}

func DeserializeWingIndexFactorInfo(data []byte) (*WingIndexFactorInfo, error) {
	source := common.NewZeroCopySource(data)
	factor, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read factor eof")
	}
	supplyIndex, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read supplyIndex eof")
	}
	insuranceIndex, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read insuranceIndex eof")
	}
	return &WingIndexFactorInfo{
		Factor:         factor.ToBigInt(),
		SupplyIndex:    supplyIndex.ToBigInt(),
		InsuranceIndex: insuranceIndex.ToBigInt(),
	}, nil
}

func (this *Comptroller) GetWingIndexFactorInfo(market common.Address) (*WingIndexFactorInfo, error) {
	method := "getWingDeltaIndexInfo"
	params := []interface{}{market}
	res, err := this.sdk.WasmVM.PreExecInvokeWasmVMContract(this.addr, method, params)
	if err != nil {
		return nil, fmt.Errorf("GetWingIndexFactorInfo: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetWingIndexFactorInfo: %s", err)
	}
	result, err := DeserializeWingIndexFactorInfo(data)
	if err != nil {
		return nil, fmt.Errorf("GetWingIndexFactorInfo: %s", err)
	}
	return result, nil
}

func (this *Comptroller) isComptroller() (bool, error) {
	method := "isComptroller"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("WingAccrued: %s", err)
	}
	return res, nil
}


