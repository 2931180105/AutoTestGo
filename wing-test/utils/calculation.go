package utils


import (
	"github.com/mockyz/AutoTestGo/common/log"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"math/big"
	"time"
)

type ClaimWingAtMarket struct {
	DistributedBorrowerWing  *big.Int
	DistributedSupplierWing  *big.Int
	DistributedGuaranteeWing *big.Int
	Timestamp                uint32
}

func GetClaimWingEventByHash(sdk *ontSDK.OntologySdk, txHash string) *ClaimWingAtMarket {
	claimWingAtMarket := new(ClaimWingAtMarket)
	//wait hash
	for j := 0; j < 50; j++ {
		time.Sleep(time.Second * 3)
		evts, err := sdk.GetSmartContractEvent(txHash)
		if err != nil || evts == nil {
			continue
		} else {
			log.Infof("evts = %v\n", evts)
			log.Infof("TxHash:%v\n", txHash)
			log.Infof("State:%d\n", evts.State)
			for _, notify := range evts.Notify {
				log.Infof("ContractAddress:%v\n", notify.ContractAddress)
				log.Infof("States:%+v\n", notify.States)
				value := notify.States.([]interface{})
				if value[0].(string) == "DistributedSupplierWing" {
					claimWingAtMarket.DistributedSupplierWing = ToIntByPrecise(value[3].(string), 0)
				}
				if value[0].(string) == "DistributedBorrowerWing" {
					claimWingAtMarket.DistributedBorrowerWing = ToIntByPrecise(value[3].(string), 0)
				}
				if value[0].(string) == "DistributedGuaranteeWing" {
					claimWingAtMarket.DistributedGuaranteeWing = ToIntByPrecise(value[3].(string), 0)
				}
			}
			break
		}
	}
	claimWingAtMarket.Timestamp = GetTimeByTxhash(sdk, txHash)
	return claimWingAtMarket
}
func ExpTestRuslt(wingSpeed, users, total *big.Int, start, end, percetage uint32) *big.Int {

	//	expResult =wingSpeed*percetage*userBorrow/totalBorrow * (start - end time)
	//x =wingSpeed*userBorrow/totalBorrow
	x := big.NewInt(0).Div(big.NewInt(0).Mul(wingSpeed, users), total)
	//y = (start -end) * percetage
	y := (end - start) * percetage
	reslut := big.NewInt(0).Div(big.NewInt(0).Mul(x, big.NewInt(int64(y))), big.NewInt(100))
	log.Infof("wingSpeed: %v,users: %v,total: %v,start: %v,end: %v,percetage: %v", wingSpeed, users, total, start, end, percetage)
	if reslut.Cmp(big.NewInt(0))==0{
		reslut =big.NewInt(1)
	}
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
	if relRsult.Cmp(big.NewInt(0))==0 || expRsult.Cmp(big.NewInt(0))==0{
		return big.NewFloat(1)
	}
	if relRsult.Cmp(expRsult) < 0 {
		return new(big.Float).Sub(new(big.Float).SetInt64(1), new(big.Float).Quo(new(big.Float).SetInt(relRsult), new(big.Float).SetInt(expRsult)))
	} else {
		return new(big.Float).Sub(new(big.Float).SetInt64(1), new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	}
}
func ExpInterestAdd(totalBorrow, delayBlockNum, borrowRatePerBlock *big.Int) *big.Int {
	x := big.NewInt(0).Mul(big.NewInt(0).Mul(totalBorrow, delayBlockNum), borrowRatePerBlock)
	//DIV 10^20
	y := big.NewInt(0).Div(x, ToIntByPrecise("1",18))
	return y
}
