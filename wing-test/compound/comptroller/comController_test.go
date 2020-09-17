package comptroller

import (
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"testing"
)

//TODO： init/wing token set gov address
func TestEnterMarkets(t *testing.T) {
	cfg, account, sdk := Utils.GetTestConfig()
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	FETH, _ := utils.AddressFromHexString(cfg.FETH)
	ftokenS := make([]OntCommon.Address, 0)
	ftokenS = append(ftokenS, FETH)
	EnterMarkets(cfg, account, sdk, comptrollerAddr, account.Address, ftokenS)
}
