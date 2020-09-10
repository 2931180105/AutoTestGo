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
	params := []interface{}{"mint", []interface{}{account.Address, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//redeem
func FtokenRedeem(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	FTokenAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	params := []interface{}{"redeem", []interface{}{account.Address, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
	params := []interface{}{"redeemUnderlying", []interface{}{account.Address, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
	params := []interface{}{"borrow", []interface{}{account.Address, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
	params := []interface{}{"repayBorrow", []interface{}{account.Address, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
	params := []interface{}{"repayBorrowBehalf", []interface{}{payer, borrower, cfg.Amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
	params := []interface{}{"liquidateBorrow", []interface{}{payer, borrower, cfg.Amount, ftokenCollateral}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, FTokenAddr, params)
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
