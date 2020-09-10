package compound

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

//mint
func FtokenMint(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
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
func FtokenRedeem(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	params := []interface{}{account.Address, cfg.Amount}}
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
func FtokenRedeemUnderlying(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	params :=  []interface{}{account.Address, cfg.Amount}}
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
func FtokenBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	params :=  []interface{}{account.Address, cfg.Amount}}
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
func FtokenRepayBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	params :=  []interface{}{account.Address, cfg.Amount}}
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
func FtokenRepayBorrowBehalf(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	payer := account.Address
	borrower, _ := utils.AddressFromBase58(cfg.AuthAddr)
	params :=  []interface{}{payer, borrower, cfg.Amount}}
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
func FtokenLiquidateBorrow(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	payer := account.Address
	borrower, _ := utils.AddressFromBase58(cfg.AuthAddr)
	ftokenCollateral, _ := utils.AddressFromHexString(cfg.OUSDT)
	params := []interface{}{payer, borrower, cfg.Amount, ftokenCollateral}}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, "liquidateBorrow", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
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
