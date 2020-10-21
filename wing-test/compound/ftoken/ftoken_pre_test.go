package ftoken

import (
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"

	//OntCommon "github.com/ontio/ontology/common"
	"testing"
)

//TODO： init/wing token set gov address
func TestBalanceOfUnderlying(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	BalanceOfUnderlying(cfg, account.Address.ToBase58(), sdk, cfg.FBTC)
}
func TestAccountSnapshot(t *testing.T) {
	cfg, account, sdk := Utils.GetPrvConfig()
	AccountSnapshot(account.Address.ToBase58(), sdk, cfg.FBTC)
}
func TestAllmarkets(t *testing.T) {
	cfg, _, sdk := Utils.GetPrvConfig()
	addr, _ := utils.AddressFromHexString(cfg.Comptroller)
	AllMarkets(sdk, addr)
	//AccountSnapshot(account.Address.ToBase58(),sdk,cfg.FBTC)
}
