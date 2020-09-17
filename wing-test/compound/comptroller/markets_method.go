package comptroller

import (
	"fmt"
	Otoken "github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math"
	"math/big"
)

// EnterMarkets  ftoken address
func EnterMarkets(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, comptroller OntCommon.Address, from OntCommon.Address, markets []OntCommon.Address) {
	params := []interface{}{from, markets}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptroller, "enterMarkets", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// ExitMarkets  ftoken address
func ExitMarkets(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, comptroller OntCommon.Address, from OntCommon.Address, market OntCommon.Address) {
	params := []interface{}{from, market}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptroller, "exitMarkets", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// Mint  ftoken  (supply)
func ApproveAndMint(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, otoken OntCommon.Address, from OntCommon.Address, amount uint64) {
	Otoken.ApproveOToken(cfg, account, sdk, ftoken, otoken, big.NewInt(math.MaxInt64))
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "mint", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// Mint  ftoken   MintFtoken
func MintFtoken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, from OntCommon.Address, amount uint64) {
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "mint", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// RedeemToken   ftoken
func RedeemToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, from OntCommon.Address, amount uint64) {
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "redeem", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

func RedeemUnderlying(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, redeemer OntCommon.Address, redeemAmount *big.Int) {
	params := []interface{}{redeemer, redeemAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "redeemUnderlying", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

func Borrow(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, borrower OntCommon.Address, borrowAmount *big.Int) {
	params := []interface{}{borrower, borrowAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "borrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}
func RepayBorrow(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, borrower OntCommon.Address, repayAmount *big.Int) {
	params := []interface{}{borrower, repayAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "repayBorrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}
func RepayBorrowBehalf(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, payer, ftoken OntCommon.Address, borrower OntCommon.Address, repayAmount *big.Int) {
	params := []interface{}{payer, borrower, repayAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "repayBorrowBehalf", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

func LiquidateBorrow(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, liquidator, ftoken OntCommon.Address, borrower OntCommon.Address, repayAmount *big.Int) {
	params := []interface{}{liquidator, borrower, repayAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "liquidateBorrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}
func AddReserves(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken, from OntCommon.Address, addAmount *big.Int) {
	params := []interface{}{from, addAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "_addReserves", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

func AddInsurance(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken, from OntCommon.Address, addAmount *big.Int) {
	params := []interface{}{from, addAmount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "_addInsurance", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}
