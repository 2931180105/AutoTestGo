package comptroller

import (
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"
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
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	OToken, _ := utils.AddressFromHexString(cfg.GovToken)
	ApproveAndMintWing(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
}
func TestApproveAndMintETH(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FETH)
	OToken, _ := utils.AddressFromHexString(cfg.OETH)
	ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
}

func TestApproveAndMintUSDC(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	OToken, _ := utils.AddressFromHexString(cfg.OUSDC)
	ApproveAndMint(cfg, account, sdk, FAddr, OToken, account.Address, uint64(1000))
}

func TestApproveAndReedemWing(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FWING)
	RedeemToken(cfg, account, sdk, FAddr, account.Address, uint64(1000))
}

func TestApproveAndBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	Borrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}

func TestApproveAndRepayBorrow(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	FAddr, _ := utils.AddressFromHexString(cfg.FUSDC)
	RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(990))
	RepayBorrow(cfg, account, sdk, FAddr, account.Address, big.NewInt(10))
}
