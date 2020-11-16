package if_borrow

import  (
	"fmt"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	utils2 "github.com/ontio/ontology/core/utils"
	"github.com/ontio/ontology/smartcontract/states"
	"math/big"
)

// TODO: support rest/ws
// TODO: support estimate gas before execute tx

// TODO: adapt to latest borrow pool

type IfBorrowPool struct {
	sdk    *ontSDK.OntologySdk
	signer	*ontSDK.Account
	addr   common.Address
	gasPrice uint64
	gasLimit uint64
}
type OscoreInfo struct {
	Level            byte
	InterestRate     uint64
	CollateralFactor uint64
	MaxValue         uint64
}

func (this *OscoreInfo) Serialize(sink *common.ZeroCopySink) {
	sink.WriteByte(this.Level)
	sink.WriteI128(common.I128FromUint64(this.InterestRate))
	sink.WriteI128(common.I128FromUint64(this.CollateralFactor))
	sink.WriteI128(common.I128FromUint64(this.MaxValue))
}

func (this *OscoreInfo) Deserialize(source *common.ZeroCopySource) error {
	data, eof := source.NextByte()
	if eof {
		return fmt.Errorf("error: %v", eof)
	}
	this.Level = data
	data2, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("error: %v", eof)
	}
	this.InterestRate = data2
	data2, eof = source.NextUint64()
	if eof {
		return fmt.Errorf("error: %v", eof)
	}
	this.CollateralFactor = data2

	data2, eof = source.NextUint64()
	if eof {
		return fmt.Errorf("error: %v", eof)
	}
	this.MaxValue = data2
	return nil
}

func NewIfBorrowPool(nodeRPCAddr string, contractAddr string, signer *ontSDK.Account, gasPrice,
	gasLimit uint64) (*IfBorrowPool, error) {
	sdk := ontSDK.NewOntologySdk()
	client := sdk.NewRpcClient()
	client.SetAddress(nodeRPCAddr)
	sdk.SetDefaultClient(client)
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil, fmt.Errorf("NewBorrowPool: cannot access ontology network through addr %s", nodeRPCAddr)
	}
	addr, err := common.AddressFromHexString(contractAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(contractAddr)
		if err != nil {
			return nil, fmt.Errorf("NewBorrowPool: invalid contract addr %s", contractAddr)
		}
	}
	return &IfBorrowPool{
		sdk:      sdk,
		signer:   signer,
		addr:     addr,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}, nil
}

func (this *IfBorrowPool) UpdateSigner(newSigner *ontSDK.Account) {
	this.signer = newSigner
}

func (this *IfBorrowPool) GetAddr() common.Address {
	return this.addr
}

func (this *IfBorrowPool) Init(admin common.Address, marketName string, oracle, comptroller common.Address, accruedDayNumber,
	formalDays, interimDays,
	interimInterestRate,
	punishInterestRate,
	reservesFactor,
	insuranceInterestFactor uint64,
) (string, error) {
	method := "init"
	params := []interface{}{admin, marketName, oracle, comptroller, accruedDayNumber, formalDays, interimDays,
		interimInterestRate, punishInterestRate, reservesFactor, insuranceInterestFactor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Init: %s", err)
	}
	return hash, err
}

//level: u8, pub interest_rate: U128, pub collateral_factor: U128,
func (this *IfBorrowPool) SetOscoreInfo(param []*OscoreInfo) (string, error) {
	method := "setOscoreInfo"
	sink := common.NewZeroCopySink(nil)
	sink.WriteString(method)
	sink.WriteVarUint(uint64(len(param)))
	for _, p := range param {
		p.Serialize(sink)
	}
	contract := &states.WasmContractParam{}
	contract.Address = this.addr
	contract.Args = sink.Bytes()

	tx, err := utils2.NewWasmSmartContractTransaction(this.gasPrice, this.gasLimit, common.SerializeToBytes(contract))
	if err != nil {
		return "", err
	}
	if err = utils.SignTx(this.sdk, tx,0, this.signer); err != nil {
		return "", err
	}
	txhash, err := this.sdk.SendTransaction(tx)
	if err != nil {
		return "", err
	}
	return txhash.ToHexString(), err
}

func (this *IfBorrowPool) SetPriceOracle(oracle common.Address) (string, error) {
	method := "_setPriceOracle"
	params := []interface{}{oracle}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPriceOracle: %s", err)
	}
	log.Infof("borrow pool SetPriceOracle txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) SetComptroller(comptroller common.Address) (string, error) {
	method := "_setComptroller"
	params := []interface{}{comptroller}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetComptroller: %s", err)
	}
	return hash, err
}
func (this *IfBorrowPool) SetOracle(oracle common.Address) (string, error) {
	method := "_setOracle"
	params := []interface{}{oracle}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetOracle: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) SetReserveFactor(factor *big.Int) (string, error) {
	method := "_setReserveFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetReserveFactor: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetInterimInterestRate(factor *big.Int) (string, error) {
	method := "_setInterimInterestRate"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInterimInterestRate: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetPunishInterestRate(factor *big.Int) (string, error) {
	method := "_setPunishInterestRate"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPunishInterestRate: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetFormalBorrowDays(days uint64) (string, error) {
	method := "_setFormalBorrowDays"
	params := []interface{}{days}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetFormalBorrowDays: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetInterimBorrowDays(days uint64) (string, error) {
	method := "_setInterimBorrowDays"
	params := []interface{}{days}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInterimBorrowDays: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetInsuranceInterestFactor(factor *big.Int) (string, error) {
	method := "setInsuranceInterestFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceInterestFactor: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) SetInsuranceFactor(factor *big.Int) (string, error) {
	method := "_setInsuranceFactor"
	params := []interface{}{factor}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceFactor: %s", err)
		return "", err
	}
	log.Infof("borrow pool SetInsuranceFactor txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) SetMarketAddr(market common.Address) (string, error) {
	method := "_setMarketAddr"
	params := []interface{}{market}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMarketAddr: %s", err)
		return "", err
	}
	log.Infof("borrow pool SetMarketAddr txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) SetInsuranceAddr(insurance common.Address) (string, error) {
	method := "_setInsuranceAddr"
	params := []interface{}{insurance}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetInsuranceAddr: %s", err)
		return "", err
	}
	log.Infof("borrow pool SetInsuranceAddr txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) SetMarketName(marketName string) (string, error) {
	method := "putMarketName"
	params := []interface{}{marketName}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetMarketName: %s", err)
		return "", err
	}
	return hash, err
}

func (this *IfBorrowPool) Transfer(from, to common.Address, amount *big.Int) (string, error) {
	method := "transfer"
	params := []interface{}{from, to, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Transfer: %s", err)
		return "", err
	}
	log.Infof("borrow pool Transfer txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) TransferFrom(from, src, to common.Address, amount *big.Int) (string, error) {
	method := "transferFrom"
	params := []interface{}{from, src, to, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TransferFrom: %s", err)
		return "", err
	}
	log.Infof("borrow pool TransferFrom txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) Approve(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Approve: %s", err)
		return "", err
	}
	log.Infof("borrow pool Approve txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) NeoVMApprove(owner, spender common.Address, amount *big.Int) (string, error) {
	method := "approve"
	params := []interface{}{owner, spender, amount}
	hash, err := utils.InvokeNeoVMTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr,
		method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMApprove: %s", err)
	}
	log.Infof("borrow pool NeoVMApprove txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) AccrueInterest() (string, error) {
	method := "accrueInterest"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccrueInterest: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) AccrueAccountInterest(addr common.Address) (string, error) {
	method := "accrueAccountInterest"
	params := []interface{}{addr}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccrueAccountInterest: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) SetPendingAdmin(pendingAdmin common.Address) (string, error) {
	method := "_setPendingAdmin"
	params := []interface{}{pendingAdmin}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SetPendingAdmin: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) AcceptAdmin() (string, error) {
	method := "_acceptAdmin"
	params := []interface{}{}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("_acceptAdmin: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) ReduceReserves(reduceAmount *big.Int) (string, error) {
	method := "_reduceReserves"
	params := []interface{}{reduceAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ReduceReserves: %s", err)
	}
	log.Infof("borrow pool ReduceReserves txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) Mint(minter common.Address, mintAmount *big.Int) (string, error) {
	method := "mint"
	params := []interface{}{minter, mintAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Mint: %s", err)
	}
	log.Infof("borrow pool Mint txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) Redeem(redeemer common.Address, redeemTokens *big.Int) (string, error) {
	method := "redeem"
	params := []interface{}{redeemer, redeemTokens}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Redeem: %s", err)
	}
	log.Infof("borrow pool Redeem txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) RedeemUnderlying(redeemer common.Address, redeemAmount *big.Int) (string, error) {
	method := "redeemUnderlying"
	params := []interface{}{redeemer, redeemAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RedeemUnderlying: %s", err)
	}
	log.Infof("borrow pool RedeemUnderlying txHash: %s", hash)
	return hash, err
}

func (this *IfBorrowPool) Borrow(borrower common.Address, borrowAmount *big.Int) (string, error) {
	method := "borrow"
	params := []interface{}{borrower, borrowAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Borrow: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) RepayBorrow(borrower common.Address, repayAmount *big.Int) (string, error) {
	method := "repayBorrow"
	params := []interface{}{borrower, repayAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayBorrow: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) RepayBorrowBehalf(payer, borrower common.Address, repayAmount *big.Int) (string, error) {
	method := "repayBorrowBehalf"
	params := []interface{}{payer, borrower, repayAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("RepayBorrowBehalf: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) Liquidate(liquidator, borrower common.Address) (string, error) {
	method := "liquidate"
	params := []interface{}{liquidator, borrower}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Liquidate: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) ReduceCollateral(from common.Address, addAmount *big.Int) (string, error) {
	method := "reduceCollateral"
	params := []interface{}{from, addAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ReduceCollateral: %s", err)
	}
	return hash, err
}

func (this *IfBorrowPool) IncreaseCollateral(from common.Address, addAmount *big.Int) (string, error) {
	method := "increaseCollateral"
	params := []interface{}{from, addAmount}
	hash, err := utils.InvokeTx(this.sdk, this.signer, this.gasPrice, this.gasLimit, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IncreaseCollateral: %s", err)
	}
	return hash, err
}

/* pre execute */
func (this *IfBorrowPool) GetOscoreInfoByLevel(level byte) (*OscoreInfo, error) {
	method := "getOscoreInfoByLevel"
	params := []interface{}{level}
	res, err := utils.PreExecuteByteArray(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetOscoreInfoByLevel: %s", err)
	}
	source := common.NewZeroCopySource(res)
	info := &OscoreInfo{}
	err = info.Deserialize(source)
	return info, err
}

func (this *IfBorrowPool) Allowance(owner, spender common.Address) (*big.Int, error) {
	method := "allowance"
	params := []interface{}{owner, spender}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Allowance: %s", err)
		return nil, err
	}
	log.Infof("borrow pool Allowance: %d", res.Uint64())
	return res, err
}

func (this *IfBorrowPool) BalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOf: %s", err)
		return nil, err
	}
	log.Infof("borrow pool BalanceOf: %d", res.Uint64())
	return res, err
}

func (this *IfBorrowPool) NeoVMBalanceOf(owner common.Address) (*big.Int, error) {
	method := "balanceOf"
	params := []interface{}{owner}
	res, err := utils.NeoVMPreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("NeoVMBalanceOf: %s", err)
		return nil, err
	}
	log.Infof("borrow pool NeoVMBalanceOf: %d", res.Uint64())
	return res, err
}

func (this *IfBorrowPool) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	method := "balanceOfUnderlying"
	params := []interface{}{owner}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BalanceOfUnderlying: %s", err)
		return nil, err
	}
	log.Infof("borrow pool BalanceOfUnderlying: %d", res.Uint64())
	return res, err
}

type AccountSnapshot struct {
	Collateral *big.Int
	// 借款
	Principal *big.Int
	// 用户当前未还的总利息
	Interest *big.Int
	// 用户负债时的正常利息
	FormalInterest *big.Int
	OScoreLevel    uint8
	BorrowIndex    *big.Int // uint256
	BorrowTime     uint64
}

func DeserializeAccountSnapshot(data []byte) (*AccountSnapshot, error) {
	source := common.NewZeroCopySource(data)
	collateral, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read Collateral eof")
	}
	principal, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read Principal eof")
	}
	interest, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read Interest eof")
	}
	formalInterest, eof := source.NextI128()
	if eof {
		return nil, fmt.Errorf("read FormalInterest eof")
	}
	oScoreLevel, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read FormalInterest eof")
	}
	borrowIndexBytes, eof := source.NextBytes(32)
	if eof {
		return nil, fmt.Errorf("read FormalInterest eof")
	}
	borrowIndex := common.BigIntFromNeoBytes(borrowIndexBytes)
	borrowTime, eof := source.NextUint64()
	if eof {
		return nil, fmt.Errorf("read FormalInterest eof")
	}
	return &AccountSnapshot{
		Collateral:     collateral.ToBigInt(),
		Principal:      principal.ToBigInt(),
		Interest:       interest.ToBigInt(),
		FormalInterest: formalInterest.ToBigInt(),
		OScoreLevel:    oScoreLevel,
		BorrowIndex:    borrowIndex,
		BorrowTime:     borrowTime,
	}, nil
}

func (this *IfBorrowPool) AccountSnapshot(owner common.Address) (*AccountSnapshot, error) {
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

func (this *IfBorrowPool) BorrowRatePerBlock() (*big.Int, error) {
	method := "borrowRatePerBlock"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowRatePerBlock: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) SupplyRatePerBlock() (*big.Int, error) {
	method := "supplyRatePerBlock"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("SupplyRatePerBlock: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) GetCash() (*big.Int, error) {
	method := "getCash"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GetCash: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) ExchangeRateStored() (*big.Int, error) {
	method := "exchangeRateStored"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ExchangeRateStored: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	method := "borrowBalanceStored"
	params := []interface{}{account}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("BorrowBalanceStored: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) GlobalParam() (common.Address, error) {
	method := "globalParam"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("GlobalParam: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) MarketName() (string, error) {
	method := "marketName"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Name: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) Symbol() (string, error) {
	method := "symbol"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Symbol: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) Decimals() (string, error) {
	method := "Decimals"
	params := []interface{}{}
	res, err := utils.PreExecuteString(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Decimals: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) Admin() (common.Address, error) {
	method := "admin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Admin: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) PendingAdmin() (common.Address, error) {
	method := "pendingAdmin"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PendingAdmin: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) Comptroller() (common.Address, error) {
	method := "comptroller"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("Comptroller: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) ReserveFactor() (*big.Int, error) {
	method := "reservesFactor"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("ReservesFactor: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) InsuranceInterestFactor() (*big.Int, error) {
	method := "insuranceInterestFactor"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InsuranceInterestFactor: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) TotalReserves() (*big.Int, error) {
	method := "totalReserves"
	params := []interface{}{}
	res, err := utils.PreExecuteBigInt(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalReserves: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) IsBorrow() (bool, error) {
	method := "isBorrow"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("IsBorrow: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) OScoreOracle() (common.Address, error) {
	method := "oScoreOracle"
	params := []interface{}{}
	res, err := utils.PreExecuteAddress(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("OScoreOracle: %s", err)
	}
	return res, err
}
func (this *IfBorrowPool) FormalBorrowDays() (uint64, error) {
	method := "formalBorrowDays"
	params := []interface{}{}
	res, err := utils.PreExecuteUint64(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("FormalBorrowDays: %s", err)
	}
	return res, err
}
func (this *IfBorrowPool) InterimBorrowDays() (uint64, error) {
	method := "interimBorrowDays"
	params := []interface{}{}
	res, err := utils.PreExecuteUint64(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InterimBorrowDays: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) AccruedDayNumber() (uint64, error) {
	method := "accruedDayNumber"
	params := []interface{}{}
	res, err := utils.PreExecuteUint64(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("AccruedDayNumber: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) PunishInterestRate() (bool, error) {
	method := "punishInterestRate"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("PunishInterestRate: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) InterimInterestRate() (bool, error) {
	method := "interimInterestRate"
	params := []interface{}{}
	res, err := utils.PreExecuteBool(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("InterimInterestRate: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) FormalInterest() (common.I128, error) {
	method := "formalInterest"
	params := []interface{}{}
	res, err := utils.PreExecuteUint128(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("FormalInterest: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) TotalInterest() (common.I128, error) {
	method := "totalInterest"
	params := []interface{}{}
	res, err := utils.PreExecuteUint128(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalInterest: %s", err)
	}
	return res, err
}

func (this *IfBorrowPool) TotalInsuranceInterest() (common.I128, error) {
	method := "totalInsuranceInterest"
	params := []interface{}{}
	res, err := utils.PreExecuteUint128(this.sdk, this.addr, method, params)
	if err != nil {
		err = fmt.Errorf("TotalInsuranceInterest: %s", err)
	}
	return res, err
}
