package comptroller

import (
	"fmt"
	Otoken "github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/payload"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/states"
	"math"
	"math/big"
	"time"
)

// EnterMarkets  ftoken address
func EnterMarkets(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, comptroller OntCommon.Address, from OntCommon.Address, markets []interface{}) {
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
func ApproveAndMint(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, otoken OntCommon.Address, from OntCommon.Address, amount *big.Int) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	Otoken.ApproveOToken(cfg, account, sdk, WingGovAddr, otoken, big.NewInt(math.MaxInt64))
	//Otoken.ApproveOToken(cfg, account, sdk, ftoken, otoken, big.NewInt(100000000000000))
	log.Infof("ftoken：%s", ftoken.ToHexString())
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "mint", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Errorf("Mint Token err:%s, Ftoken address: %s", err, ftoken.ToHexString())
	}
}

// Mint  ftoken  (supply)
func Mint(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, from OntCommon.Address, amount uint64) {
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "mint", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Errorf("Mint Token err:%s, Ftoken address: %s", err, ftoken.ToHexString())
	}
}

// Mint  ftoken  (supply)
func Mint2(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken, from OntCommon.Address, amount string) {
	sink := OntCommon.NewZeroCopySink(nil)
	sink.WriteString("mint")
	sink.WriteAddress(from)
	sink.WriteHash(Utils.Uint256FromhexString(amount))
	contract := &states.WasmContractParam{}
	contract.Address = ftoken
	argbytes := sink.Bytes()
	contract.Args = argbytes
	invokePayload := &payload.InvokeCode{
		Code: OntCommon.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		Payer:    account.Address,
		GasPrice: 2500,
		GasLimit: 300000,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	if err := Utils.SignTxAndSendTx(sdk, tx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// Mint  ftoken  (supply)
func ApproveAndMintWing(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, otoken OntCommon.Address, from OntCommon.Address, amount *big.Int) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	Otoken.ApproveOToken(cfg, account, sdk, WingGovAddr, otoken, big.NewInt(1000000))

	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "mint", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Errorf("Mint Token err:%s, Ftoken address: %s", err, ftoken.ToHexString())
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
func RedeemToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, from OntCommon.Address, amount *big.Int) {
	params := []interface{}{from, amount}
	mutTx, err := sdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ftoken, "redeem", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := Utils.SignTxAndSendTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
}

// RedeemToken   ftoken
func RedeemToken2(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, ftoken OntCommon.Address, from OntCommon.Address, amount uint64) {
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
		log.Errorf("borrow : %s", err)
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

func TotalBorrowCurrent(sdk *goSdk.OntologySdk, ftoken OntCommon.Address) {
	params := []interface{}{}
	result, err := sdk.WasmVM.PreExecInvokeWasmVMContract(ftoken, "totalBorrowCurrent", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("totalBorrowCurrent:%s", result.Result)
}

func BorrowBalanceCurrent(sdk *goSdk.OntologySdk, ftoken, account OntCommon.Address) {
	params := []interface{}{account}
	result, err := sdk.WasmVM.PreExecInvokeWasmVMContract(ftoken, "borrowBalanceCurrent", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("totalBorrowCurrent:%s", result.Result)
}

func ExchangeRateCurrent(sdk *goSdk.OntologySdk, ftoken OntCommon.Address) {
	params := []interface{}{}
	result, err := sdk.WasmVM.PreExecInvokeWasmVMContract(ftoken, "totalBorrowCurrent", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("totalBorrowCurrent:%s", result.Result)
}

//allMarkets

func AllMarkets(sdk *goSdk.OntologySdk, comptrooller OntCommon.Address) {
	params := []interface{}{}
	result, err := sdk.WasmVM.PreExecInvokeWasmVMContract(comptrooller, "allMarkets", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("totalBorrowCurrent:%s", result.Result)
}
