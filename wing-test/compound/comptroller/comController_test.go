package comptroller

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"
	"io/ioutil"
	"math"
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
	EnterMarkets(cfg, account, sdk, comptrollerAddr, account.Address, ftokenS)
}

func TestApproveAndMintWing(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 3)
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	OToken, _ := utils.AddressFromHexString(cfg.GovToken)
	ApproveAndMintWing(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
	ApproveAndMintWing(cfg, accounts[2], sdk, FAddr, OToken, accounts[0].Address, uint64(100000))

}

func TestApproveAndMintWing2(t *testing.T) {
	cfg, _, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 3)
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	OToken, _ := utils.AddressFromHexString(cfg.GovToken)
	ApproveAndMintWing(cfg, accounts[2], sdk, FAddr, OToken, accounts[2].Address, uint64(1000))
	ApproveAndMintWing(cfg, accounts[2], sdk, FAddr, OToken, accounts[2].Address, uint64(100))

}
func TestApproveAndMintETH(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(1, 3)
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	OToken, _ := utils.AddressFromHexString(cfg.OETH)
	ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1))
	ApproveAndMint(cfg, accounts[0], sdk, FAddr, OToken, accounts[0].Address, uint64(1000))
}

func TestApproveAndMintUSDC(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	OToken, _ := utils.AddressFromHexString(cfg.OUSDC)
	ApproveAndMintWing(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
	ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
}

func TestAReedemWing(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 1)
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	RedeemToken(cfg, account, sdk, FAddr, account.Address, uint64(1))
	RedeemToken(cfg, accounts[0], sdk, FAddr, accounts[0].Address, math.MaxInt64)
}

func TestAReedem(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	accounts := dbHelper.QueryAccountFromDb(0, 1)
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	RedeemToken(cfg, account, sdk, FAddr, account.Address, uint64(100))
	RedeemToken(cfg, accounts[0], sdk, FAddr, accounts[0].Address, uint64(10000000))
}

func TestBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	Borrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}

func TestRepayBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(-1))
}

func TestRepayBorrowEth(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(1))
	RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
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
