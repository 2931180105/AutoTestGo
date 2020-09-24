package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"
	"io/ioutil"
	"math/big"
	//OntCommon "github.com/ontio/ontology/common"
	"testing"
)

//TODO： init/wing token set gov address
func TestEnterMarkets(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	FETH, _ := utils.AddressFromHexString(cfg.FETH)

	ftokenS := []interface{}{FETH}
	comptroller.EnterMarkets(cfg, account, sdk, comptrollerAddr, account.Address, ftokenS)
}

func TestApproveAndMint(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	//accounts := dbHelper.QueryAccountFromDb(0, 1)
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	OToken, _ := utils.AddressFromHexString(cfg.OETH)
	//Utils.ToIntByPrecise("1", precise)
	comptroller.ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, Utils.ToIntByPrecise("10000", 8))
	//comptroller.ApproveAndMintWing(cfg, accounts[0], sdk, FAddr, OToken, accounts[0].Address, Utils.ToIntByPrecise("1",100000))
	AccountSnapshot(account.Address.ToBase58(), sdk, FAddr.ToHexString())
}

func TestApproveAndMintWing2(t *testing.T) {
	cfg, _, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 3)
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	OToken, _ := utils.AddressFromHexString(cfg.GovToken)
	comptroller.ApproveAndMintWing(cfg, accounts[2], sdk, FAddr, OToken, accounts[2].Address, Utils.ToIntByPrecise("1", 1000))
	comptroller.ApproveAndMintWing(cfg, accounts[2], sdk, FAddr, OToken, accounts[2].Address, Utils.ToIntByPrecise("1", 100))

}
func TestApproveAndMintETH(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	comptroller.Mint2(cfg, account, sdk, FAddr, account.Address, "ffffffffffffffffffffffffffffffff00")

}

func TestApproveAndMintUSDC(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	OToken, _ := utils.AddressFromHexString(cfg.OUSDC)
	comptroller.ApproveAndMintWing(cfg, account, sdk, FAddr, OToken, account.Address, Utils.ToIntByPrecise("1", 1000))
	//u128,_:=common.Uint256FromHexString("ffffffffffffffffffffffffffffffff00")
	comptroller.ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, Utils.ToIntByPrecise("1", 1000))
}

func TestAReedemWing(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 1)
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	comptroller.RedeemToken(cfg, account, sdk, FAddr, account.Address, Utils.ToIntByPrecise("1", 1))
	comptroller.RedeemToken(cfg, accounts[0], sdk, FAddr, accounts[0].Address, Utils.ToIntByPrecise("1", 1))
}

func TestAReedem(t *testing.T) {
	cfg, _, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDbByBase58("AHrrpe6zXduNZNvH1XkJPK9VXdC51b6oqv")
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	//pkey, _ := hex.DecodeString("")
	//pri, _ := keypair.DeserializePrivateKey(pkey)
	//accounts := goSdk.Account{
	//	PrivateKey: pri,
	//	PublicKey:  pri.Public(),
	//	Address:    types.AddressFromPubKey(pri.Public()),
	//	SigScheme:  signature.SHA256withECDSA,
	//}
	comptroller.RedeemToken2(cfg, accounts[0], sdk, FAddr, accounts[0].Address, 22000)
	//comptroller.RedeemToken(cfg, accounts[0], sdk, FAddr, accounts[0].Address, Utils.ToIntByPrecise("1",10000000))
}

func TestBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	comptroller.Borrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}

func TestRepayBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	comptroller.RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}

func TestRepayBorrowEth(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	comptroller.RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(1))
	comptroller.RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}
func TestGetBase58(t *testing.T) {
	bytes, err := ioutil.ReadFile("")
	if err != nil {
		log.Fatal(err)
	}
	contractCodeStr := string(bytes)
	Contract, err := utils.GetContractAddress(contractCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("contract address : %s", Contract.ToBase58())
	log.Infof("contract address : %s", Contract.ToHexString())

	addr, _ := utils.AddressFromHexString("d034792f80deeacd983dc257d29784ea71a1d5ec")
	log.Infof("base58: %s", addr.ToBase58())
}
