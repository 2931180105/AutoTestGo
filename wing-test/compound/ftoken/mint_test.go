package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"
	//OntCommon "github.com/ontio/ontology/common"
	"testing"
)

func TestApproveAndMintETHandBTCS(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FBTC)
	OAddr3, _ := utils.AddressFromHexString(cfg.WBTC)

	comptroller.ApproveAndMint(cfg, account, sdk, FAddr, OAddr3, account.Address, Utils.ToIntByPrecise("100", 8))

}

func TestBatchMint(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(40, 10)
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	FAddr2, _ := utils.AddressFromHexString(cfg.FRENBTC)
	FAddr3, _ := utils.AddressFromHexString(cfg.FBTC)
	FAddr4, _ := utils.AddressFromHexString(cfg.FUSDC)
	FAddr5, _ := utils.AddressFromHexString(cfg.FWING)
	OAddr, _ := utils.AddressFromHexString(cfg.ETH)
	OAddr2, _ := utils.AddressFromHexString(cfg.RENBTC)
	OAddr3, _ := utils.AddressFromHexString(cfg.WBTC)
	OAddr4, _ := utils.AddressFromHexString(cfg.USDC)
	OAddr5, _ := utils.AddressFromHexString(cfg.GovToken)
	//comptroller.Mint(cfg, account, sdk, FAddr, account.Address, 1000000)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run*****:%d***********", i)
		otoken.TransferAllTestToken(cfg, account, sdk, accounts[i].Address.ToBase58())
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr, OAddr, accounts[i].Address, Utils.ToIntByPrecise("100", 18))
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr4, OAddr2, accounts[i].Address, Utils.ToIntByPrecise("100", 8))
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr2, OAddr3, accounts[i].Address, Utils.ToIntByPrecise("100", 8))
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr3, OAddr4, accounts[i].Address, Utils.ToIntByPrecise("100", 6))
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr5, OAddr5, accounts[i].Address, Utils.ToIntByPrecise("100", 9))
	}
}

func TestBatchBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	accounts := dbHelper.QueryAccountFromDb(530, 10)
	FAddr4, _ := utils.AddressFromHexString(cfg.FUSDC)
	OAddr4, _ := utils.AddressFromHexString(cfg.USDC)
	//FAddr2, _ := utils.AddressFromHexString(cfg.FRENBTC)
	FAddr3, _ := utils.AddressFromHexString(cfg.FBTC)
	//comptroller.Mint(cfg, account, sdk, FAddr2, account.Address, 100000000000)
	//comptroller.Mint(cfg, account, sdk, FAddr3, account.Address, 100000000000)
	//comptroller.ApproveAndMint(cfg,account, sdk, FAddr4,  OAddr4,account.Address, Utils.ToIntByPrecise("100000",8))

	//otoken.DelegateToProxyAllTestToken(cfg,account,sdk)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run****************:%d", i)
		//otoken.TransferAllTestToken(cfg,account,sdk,accounts[i].Address.ToBase58())
		//_, _ = sdk.Native.Ong.Transfer(cfg.GasPrice, cfg.GasLimit, account, account, accounts[i].Address, 100000000000)
		otoken.OTokenTransfer(cfg, account, sdk, accounts[i].Address.ToBase58(), cfg.USDC, 8)
		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr4, OAddr4, accounts[i].Address, Utils.ToIntByPrecise("10000", 6))
		comptroller.EnterMarkets(cfg, accounts[i], sdk, comptrollerAddr, accounts[i].Address, []interface{}{FAddr3, FAddr4})
		//comptroller.Borrow(cfg, accounts[i], sdk, FAddr2, accounts[i].Address,  Utils.ToIntByPrecise("1",8))
		comptroller.Borrow(cfg, accounts[i], sdk, FAddr3, accounts[i].Address, Utils.ToIntByPrecise("8", 7))
	}
}

func TestBatchBorrow2(t *testing.T) {
	cfg, _, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(100, 10)
	FAddr3, _ := utils.AddressFromHexString(cfg.FRENBTC)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run****************:%d", i)
		//comptroller.Borrow(cfg, accounts[i], sdk, FAddr2, accounts[i].Address,  Utils.ToIntByPrecise("1",8))
		comptroller.Borrow(cfg, accounts[i], sdk, FAddr3, accounts[i].Address, Utils.ToIntByPrecise("1", 4))
	}
}

func TestBorrWBTC(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	accounts := dbHelper.QueryAccountFromDb(60, 10)
	FAddr4, _ := utils.AddressFromHexString(cfg.FRENBTC)
	OAddr4, _ := utils.AddressFromHexString(cfg.RENBTC)
	//FAddr2, _ := utils.AddressFromHexString(cfg.FRENBTC)
	FAddr3, _ := utils.AddressFromHexString(cfg.FBTC)
	//comptroller.Mint(cfg, account, sdk, FAddr2, account.Address, 100000000000)
	//comptroller.Mint(cfg, account, sdk, FAddr3, account.Address, 100000000000)
	//comptroller.ApproveAndMint(cfg,account, sdk, FAddr3,  OAddr4,account.Address, Utils.ToIntByPrecise("100",8))

	//otoken.DelegateToProxyAllTestToken(cfg,account,sdk)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run****************:%d", i)
		//otoken.TransferAllTestToken(cfg,account,sdk,accounts[i].Address.ToBase58())
		otoken.OTokenTransfer(cfg, account, sdk, accounts[i].Address.ToBase58(), cfg.RENBTC, 8)

		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr4, OAddr4, accounts[i].Address, Utils.ToIntByPrecise("1", 8))
		comptroller.EnterMarkets(cfg, accounts[i], sdk, comptrollerAddr, accounts[i].Address, []interface{}{FAddr3, FAddr4})
		//comptroller.Borrow(cfg, accounts[i], sdk, FAddr2, accounts[i].Address,  Utils.ToIntByPrecise("1",8))
		comptroller.Borrow(cfg, accounts[i], sdk, FAddr3, accounts[i].Address, Utils.ToIntByPrecise("8", 7))
	}
}

func TestStakeETHBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	accounts := dbHelper.QueryAccountFromDb(0, 10)
	FAddr4, _ := utils.AddressFromHexString(cfg.FETH)
	OAddr4, _ := utils.AddressFromHexString(cfg.ETH)
	//FAddr2, _ := utils.AddressFromHexString(cfg.FRENBTC)
	FAddr3, _ := utils.AddressFromHexString(cfg.FRENBTC)
	//comptroller.Mint(cfg, account, sdk, FAddr2, account.Address, 100000000000)
	//comptroller.Mint(cfg, account, sdk, FAddr3, account.Address, 100000000000)
	//comptroller.ApproveAndMint(cfg,account, sdk, FAddr3,  OAddr4,account.Address, Utils.ToIntByPrecise("100",8))

	//otoken.DelegateToProxyAllTestToken(cfg,account,sdk)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run****************:%d", i)
		//otoken.TransferAllTestToken(cfg,account,sdk,accounts[i].Address.ToBase58())
		otoken.OTokenTransfer(cfg, account, sdk, accounts[i].Address.ToBase58(), cfg.ETH, 18)

		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr4, OAddr4, accounts[i].Address, Utils.ToIntByPrecise("1", 18))
		comptroller.EnterMarkets(cfg, accounts[i], sdk, comptrollerAddr, accounts[i].Address, []interface{}{FAddr3, FAddr4})
		//comptroller.Borrow(cfg, accounts[i], sdk, FAddr2, accounts[i].Address,  Utils.ToIntByPrecise("1",8))
		comptroller.Borrow(cfg, accounts[i], sdk, FAddr3, accounts[i].Address, Utils.ToIntByPrecise("8", 7))
	}
}

func TestStakeETHBorrow2(t *testing.T) {

	////accounts := dbHelper.QueryAccountFromDb(0, 1)
	//FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	//OToken, _ := utils.AddressFromHexString(cfg.ETH)
	////Utils.ToIntByPrecise("1", precise)
	//comptroller.ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, Utils.ToIntByPrecise("10000",8))
	////comptroller.ApproveAndMintWing(cfg, accounts[0], sdk, FAddr, OToken, accounts[0].Address, Utils.ToIntByPrecise("1",100000))
	//AccountSnapshot(account.Address.ToBase58(),sdk,FAddr.ToHexString())
}

func TestStakeBTCBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	accounts := dbHelper.QueryAccountFromDb(520, 10)
	FAddr4, _ := utils.AddressFromHexString(cfg.FBTC)
	OAddr4, _ := utils.AddressFromHexString(cfg.WBTC)
	//FAddr2, _ := utils.AddressFromHexString(cfg.FRENBTC)
	FAddr3, _ := utils.AddressFromHexString(cfg.FETH)
	//OAddr3, _ := utils.AddressFromHexString(cfg.ETH)

	//comptroller.Mint(cfg, account, sdk, FAddr2, account.Address, 100000000000)
	//comptroller.Mint(cfg, account, sdk, FAddr3, account.Address, 100000000000)
	//comptroller.ApproveAndMint(cfg,account, sdk, FAddr3,  OAddr3,account.Address, Utils.ToIntByPrecise("10",18))

	//otoken.DelegateToProxyAllTestToken(cfg,account,sdk)
	for i := 0; i < len(accounts); i++ {
		log.Infof("run****************:%d", i)
		//otoken.TransferAllTestToken(cfg,account,sdk,accounts[i].Address.ToBase58())
		otoken.OTokenTransfer(cfg, account, sdk, accounts[i].Address.ToBase58(), cfg.WBTC, 8)

		comptroller.ApproveAndMint(cfg, accounts[i], sdk, FAddr4, OAddr4, accounts[i].Address, Utils.ToIntByPrecise("1", 8))
		comptroller.EnterMarkets(cfg, accounts[i], sdk, comptrollerAddr, accounts[i].Address, []interface{}{FAddr3, FAddr4})
		//comptroller.Borrow(cfg, accounts[i], sdk, FAddr2, accounts[i].Address,  Utils.ToIntByPrecise("1",8))
		comptroller.Borrow(cfg, accounts[i], sdk, FAddr3, accounts[i].Address, Utils.ToIntByPrecise("1", 18))
	}
}
