package ftoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	go_sdk_utils "github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

//mint
func FtokenMint(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{account.Address, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "mint", params)
	if err != nil {
		fmt.Println("construct tx mint err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//redeem
func FtokenRedeem(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{account.Address, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "redeem", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//redeemUnderlying
func FtokenRedeemUnderlying(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{account.Address, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "redeemUnderlying", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//borrow
func FtokenBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{account.Address, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "borrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//repay borrow
func FtokenRepayBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{account.Address, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "repayBorrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//repay borrow behalf
func FtokenRepayBorrowBehalf(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	payer := account.Address
	borrower, _ := go_sdk_utils.AddressFromBase58(cfg.AuthAddr)
	params := []interface{}{payer, borrower, cfg.Amount}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "repayBorrowBehalf", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//liquidate borrow
func FtokenLiquidateBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *types.MutableTransaction {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	payer := account.Address
	borrower, _ := go_sdk_utils.AddressFromBase58(cfg.AuthAddr)
	ftokenCollateral, _ := go_sdk_utils.AddressFromHexString(cfg.USDT)
	params := []interface{}{payer, borrower, cfg.Amount, ftokenCollateral}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "liquidateBorrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get accountSnapshoto
func AccountSnapshot(address string, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	accountAddr, _ := go_sdk_utils.AddressFromBase58(address)
	params := []interface{}{accountAddr}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "accountSnapshot", params)
	if err != nil {
		log.Infof("accountSnapshot: %s", err)
	}
	if resut.Result != nil {
		log.Infof("accountSnapshot: %s", resut.Result)
	} else {
		log.Info("accountSnapshot is nil")
	}
	return resut
}

//get balanceOfUnderlying
func BalanceOfUnderlying(cfg *config.Config, address string, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	accountAddr, _ := go_sdk_utils.AddressFromBase58(address)
	params := []interface{}{accountAddr}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "balanceOfUnderlying", params)
	if err != nil {
		log.Infof("balanceOfUnderlying: %s", err)
	}
	if resut.Result != nil {
		log.Infof("balanceOfUnderlying: %s", resut.Result)
	} else {
		log.Info("balanceOfUnderlying result is nil")
	}
	return resut
}

//get borrowRatePerBlock
func BorrowRatePerBlock(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "borrowRatePerBlock", params)
	if err != nil {
		log.Infof("borrowRatePerBlock: %s", err)
	}
	if resut.Result != nil {
		log.Infof("borrowRatePerBlock: %s", resut.Result)
	} else {
		log.Info("borrowRatePerBlock result is nil")
	}
	return resut
}

//get supplyRatePerBlock
func SupplyRatePerBlock(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "supplyRatePerBlock", params)
	if err != nil {
		log.Infof("supplyRatePerBlock: %s", err)
	}
	if resut.Result != nil {
		log.Infof("supplyRatePerBlock: %s", resut.Result)
	} else {
		log.Info("supplyRatePerBlock result is nil")
	}
	return resut
}

//getCash
func GetCash(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "getCash", params)
	if err != nil {
		log.Infof("getCash: %s", err)
	}
	if resut.Result != nil {
		log.Infof("getCash: %s", resut.Result)
	} else {
		log.Info("getCash result is nil")
	}
	return resut
}

//get exchangeRateStored
func ExchangeRateStored(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "exchangeRateStored", params)
	if err != nil {
		log.Infof("exchangeRateStored: %s", err)
	}
	if resut.Result != nil {
		log.Infof("exchangeRateStored: %s", resut.Result)
	} else {
		log.Info("exchangeRateStored result is nil")
	}
	return resut
}

//get borrowBalanceStored
func BorrowBalanceStored(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "borrowBalanceStored", params)
	if err != nil {
		log.Infof("borrowBalanceStored: %s", err)
	}
	if resut.Result != nil {
		log.Infof("borrowBalanceStored: %s", resut.Result)
	} else {
		log.Info("borrowBalanceStored result is nil")
	}
	return resut
}

//get totalBorrows
func TotalBorrows(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "totalBorrows", params)
	if err != nil {
		log.Infof("totalBorrows: %s", err)
	}
	if resut.Result != nil {
		log.Infof("totalBorrows: %s", resut.Result)
	} else {
		log.Info("totalBorrows result is nil")
	}
	return resut
}

//get globalParam
func GlobalParam(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "globalParam", params)
	if err != nil {
		log.Infof("globalParam: %s", err)
	}
	if resut.Result != nil {
		log.Infof("globalParam: %s", resut.Result)
	} else {
		log.Info("globalParam result is nil")
	}
	return resut
}

//get admin Address
func Admin(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "admin", params)
	if err != nil {
		log.Infof("admin: %s", err)
	}
	if resut.Result != nil {
		log.Infof("admin: %s", resut.Result)
	} else {
		log.Info("admin result is nil")
	}
	return resut
}

//get pendingAdmin Address
func PendingAdmin(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "pendingAdmin", params)
	if err != nil {
		log.Infof("pendingAdmin: %s", err)
	}
	if resut.Result != nil {
		log.Infof("pendingAdmin: %s", resut.Result)
	} else {
		log.Info("pendingAdmin result is nil")
	}
	return resut
}

//get comptroller Address
func Comptroller(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "comptroller", params)
	if err != nil {
		log.Infof("comptroller: %s", err)
	}
	if resut.Result != nil {
		log.Infof("comptroller: %s", resut.Result)
	} else {
		log.Info("comptroller result is nil")
	}
	return resut
}

//get interestRateModel Address
func InterestRateModel(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "interestRateModel", params)
	if err != nil {
		log.Infof("interestRateModel: %s", err)
	}
	if resut.Result != nil {
		log.Infof("interestRateModel: %s", resut.Result)
	} else {
		log.Info("interestRateModel result is nil")
	}
	return resut
}

//get initialExchangeRate
func InitialExchangeRateMantissa(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "initialExchangeRateMantissa", params)
	if err != nil {
		log.Infof("initialExchangeRateMantissa: %s", err)
	}
	if resut.Result != nil {
		log.Infof("initialExchangeRateMantissa: %s", resut.Result)
	} else {
		log.Info("initialExchangeRateMantissa result is nil")
	}
	return resut
}

//get reserveFactor
func ReserveFactorMantissa(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "reserveFactorMantissa", params)
	if err != nil {
		log.Infof("reserveFactorMantissa: %s", err)
	}
	if resut.Result != nil {
		log.Infof("reserveFactorMantissa: %s", resut.Result)
	} else {
		log.Info("reserveFactorMantissa result is nil")
	}
	return resut
}

//get insuranceFactor
func InsuranceFactorMantissa(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "insuranceFactorMantissa", params)
	if err != nil {
		log.Infof("insuranceFactorMantissa: %s", err)
	}
	if resut.Result != nil {
		log.Infof("insuranceFactorMantissa: %s", resut.Result)
	} else {
		log.Info("insuranceFactorMantissa result is nil")
	}
	return resut
}

//get accrualBlockNumber
func AccrualBlockNumber(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "accrualBlockNumber", params)
	if err != nil {
		log.Infof("accrualBlockNumber: %s", err)
	}
	if resut.Result != nil {
		log.Infof("accrualBlockNumber: %s", resut.Result)
	} else {
		log.Info("accrualBlockNumber result is nil")
	}
	return resut
}

//get borrowIndex
func BorrowIndex(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "borrowIndex", params)
	if err != nil {
		log.Infof("borrowIndex: %s", err)
	}
	if resut.Result != nil {
		log.Infof("borrowIndex: %s", resut.Result)
	} else {
		log.Info("borrowIndex result is nil")
	}
	return resut
}

//get totalReserves
func TotalReserves(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "totalReserves", params)
	if err != nil {
		log.Infof("totalReserves: %s", err)
	}
	if resut.Result != nil {
		log.Infof("totalReserves: %s", resut.Result)
	} else {
		log.Info("totalReserves result is nil")
	}
	return resut
}

//get totalInsurance
func TotalInsurance(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "totalInsurance", params)
	if err != nil {
		log.Infof("totalInsurance: %s", err)
	}
	if resut.Result != nil {
		log.Infof("totalInsurance: %s", resut.Result)
	} else {
		log.Info("totalInsurance result is nil")
	}
	return resut
}

//get underlying asset contract address
func Underlying(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(cfg.FBTC)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "underlying", params)
	if err != nil {
		log.Infof("underlying: %s", err)
	}
	if resut.Result != nil {
		log.Infof("underlying: %s", resut.Result)
	} else {
		log.Info("underlying result is nil")
	}
	return resut
}

//get underlying asset name
func UnderlyingName(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "underlyingName", params)
	if err != nil {
		log.Infof("underlyingName: %s", err)
	}
	if resut.Result != nil {
		log.Infof("underlyingName: %s", resut.Result)
	} else {
		log.Info("underlyingName result is nil")
	}
	return resut
}

//get insuranceAddr
func InsuranceAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "insuranceAddr", params)
	if err != nil {
		log.Infof("insuranceAddr: %s", err)
	}
	if resut.Result != nil {
		log.Infof("insuranceAddr: %s", resut.Result)
	} else {
		log.Info("insuranceAddr result is nil")
	}
	return resut
}

//get isFToken
func IsFToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "isFToken", params)
	if err != nil {
		log.Infof("isFToken: %s", err)
	}
	if resut.Result != nil {
		log.Infof("isFToken: %s", resut.Result)
	} else {
		log.Info("isFToken result is nil")
	}
	return resut
}

//get marketAddr(需要在保险池合约上调用，而不是在market合约上调用)
func MarketAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "marketAddr", params)
	if err != nil {
		log.Infof("marketAddr: %s", err)
	}
	if resut.Result != nil {
		log.Infof("marketAddr: %s", resut.Result)
	} else {
		log.Info("marketAddr result is nil")
	}
	return resut
}

//get isInsurance(需要在保险池合约上调用，而不是在market合约上调用)
func IsInsurance(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, fToken string) *common.PreExecResult {
	FTokenAddr, _ := go_sdk_utils.AddressFromHexString(fToken)
	params := []interface{}{}
	resut, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(FTokenAddr, "isInsurance", params)
	if err != nil {
		log.Infof("isInsurance: %s", err)
	}
	if resut.Result != nil {
		log.Infof("isInsurance: %s", resut.Result)
	} else {
		log.Info("isInsurance result is nil")
	}
	return resut
}

func signTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
	if nonce != 0 {
		tx.Nonce = nonce
	}
	tx.Sigs = nil
	err := sdk.SignToTransaction(tx, signer)
	if err != nil {
		return fmt.Errorf("sign tx failed, err: %s", err)
	}
	return nil
}
