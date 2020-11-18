package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/compound/ftoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"math/big"
)
func NewMarkets(cfg *config.Config, account *ontSDK.Account, sdk *ontSDK.OntologySdk, makretAddr string) (*ftoken.FlashToken, error) {
	comp,err := comptroller.NewComptroller(sdk,cfg.Comptroller,account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil {
		log.Errorf("NewComptroller  err:%v",err)
	}
	market ,err := ftoken.NewFlashToken2(sdk,makretAddr,account,cfg,comp)
	if err !=nil {
		log.Errorf("NewFlashToken  err:%v",err)
	}
	market.TestConfig = cfg

	return market, nil
}
func ExpTestRuslt(wingSpeed, users, total *big.Int, start, end, percetage uint32) *big.Int {

	//	expResult =wingSpeed*percetage*userBorrow/totalBorrow * (start - end time)
	//x =wingSpeed*userBorrow/totalBorrow
	x := big.NewInt(0).Div(big.NewInt(0).Mul(wingSpeed, users), total)
	//y = (start -end) * percetage
	y := (end - start) * percetage
	reslut := big.NewInt(0).Div(big.NewInt(0).Mul(x, big.NewInt(int64(y))), big.NewInt(100))
	log.Infof("wingSpeed: %v,users: %v,total: %v,start: %v,end: %v,percetage: %v", wingSpeed, users, total, start, end, percetage)

	return reslut
}
func GetTimeByTxhash(sdk *ontSDK.OntologySdk, txHash string) uint32 {
	blockHeight, err := sdk.GetBlockHeightByTxHash(txHash)
	if err != nil {
		log.Errorf("  GetBlockHeightByTxHash err : %v", err)
	}
	blockInfo, err := sdk.GetBlockByHeight(blockHeight)
	if err != nil {
		log.Errorf("  GetBlockByHeight err : %v", err)
	}
	log.Infof("hash time is : %v", blockInfo.Header.Timestamp)
	log.Infof("blockHeight is : %v", blockHeight)

	return blockInfo.Header.Timestamp
}
func CmpTestRuslt(expRsult, relRsult *big.Int) *big.Float {
	if relRsult.Cmp(expRsult) > 0 {
		return new(big.Float).Sub(new(big.Float).SetInt64(1), new(big.Float).Quo(new(big.Float).SetInt(relRsult), new(big.Float).SetInt(expRsult)))
	} else {
		return new(big.Float).Sub(new(big.Float).SetInt64(1), new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	}
	return nil
}
func ExpInterestAdd(totalBorrow, delayBlockNum, borrowRatePerBlock *big.Int) *big.Int {
	x := big.NewInt(0).Mul(big.NewInt(0).Mul(totalBorrow, delayBlockNum), borrowRatePerBlock)
	y := big.NewInt(0).Quo(x, big.NewInt(10^9))
	return y
}
