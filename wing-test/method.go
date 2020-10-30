package main

import (
	"crypto/rand"
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/compound/ftoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"math"
	"math/big"
	"time"
)

func BatchSupply(cfg *config.Config, genSdk *goSdk.OntologySdk,acc *goSdk.Account) {
	ftokenAddressList, err := comptroller.GetAllMarkets(genSdk, cfg.Comptroller)
	if err != nil {
		log.Errorf("batchOperate, comptroller.GetAllMarkets error: %s", err)
	}
	for _, ftokenAddress := range ftokenAddressList {
		otokenAddress, err := common.AddressFromHexString(assetMap[ftokenAddress.ToHexString()])
		if err != nil {
			log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
		}
		amount, err := rand.Int(rand.Reader, big.NewInt(50))
		if err != nil {
			log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
		}
		comptroller.ApproveAndMint(cfg, acc, genSdk, ftokenAddress, otokenAddress, acc.Address,
			new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
		itokenAddress, err := ftoken.GetITokenAddress(genSdk, ftokenAddress)
		if err != nil {
			log.Errorf("batchOperate, ftoken.GetITokenAddress error: %s", err)
		}
		comptroller.ApproveAndMint(cfg, acc, genSdk, itokenAddress, otokenAddress, acc.Address,
			new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
	}

}
func BatchBorrow(cfg *config.Config, genSdk *goSdk.OntologySdk,acc *goSdk.Account) {
	flashAddress, err := common.AddressFromHexString(cfg.Comptroller)
	if err != nil {
		log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
	}
	ftokenAddressList, err := comptroller.GetAllMarkets(genSdk, cfg.Comptroller)
	if err != nil {
		log.Errorf("batchOperate, comptroller.GetAllMarkets error: %s", err)
	}
	for _, ftokenAddress := range ftokenAddressList {
		otokenAddress, err := common.AddressFromHexString(assetMap[ftokenAddress.ToHexString()])
		if err != nil {
			log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
		}
		comptroller.ApproveAndMint(cfg, acc, genSdk, ftokenAddress, otokenAddress, acc.Address,
			new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
		comptroller.EnterMarkets(cfg, acc, genSdk, flashAddress, acc.Address, []interface{}{ftokenAddress})
		amount, err := rand.Int(rand.Reader, big.NewInt(50))
		if err != nil {
			log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
		}
		comptroller.Borrow(cfg, acc, genSdk, ftokenAddress, acc.Address,
			new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
	}
}


func batchOperate(cfg *config.Config, genSdk *goSdk.OntologySdk) {
	flashAddress, err := common.AddressFromHexString(cfg.Comptroller)
	if err != nil {
		log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
	}
	ftokenAddressList, err := comptroller.GetAllMarkets(genSdk, cfg.Comptroller)
	if err != nil {
		log.Errorf("batchOperate, comptroller.GetAllMarkets error: %s", err)
	}
	accounts := Utils.GetAccounts(cfg)
	for i := 0; i < cfg.AccountNum-1000; i++ {
		go func() {
			acc := accounts[i]
			for _, ftokenAddress := range ftokenAddressList {
				otokenAddress, err := common.AddressFromHexString(assetMap[ftokenAddress.ToHexString()])
				if err != nil {
					log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
				}
				amount, err := rand.Int(rand.Reader, big.NewInt(50))
				if err != nil {
					log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
				}
				comptroller.ApproveAndMint(cfg, acc, genSdk, ftokenAddress, otokenAddress, acc.Address,
					new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
				itokenAddress, err := ftoken.GetITokenAddress(genSdk, ftokenAddress)
				if err != nil {
					log.Errorf("batchOperate, ftoken.GetITokenAddress error: %s", err)
				}
				comptroller.ApproveAndMint(cfg, acc, genSdk, itokenAddress, otokenAddress, acc.Address,
					new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
			}
		}()
		time.Sleep(1000 * time.Millisecond)
	}
	for i := cfg.AccountNum - 1000; i < cfg.AccountNum; i++ {
		go func() {
			acc := accounts[i]
			for _, ftokenAddress := range ftokenAddressList {
				otokenAddress, err := common.AddressFromHexString(assetMap[ftokenAddress.ToHexString()])
				if err != nil {
					log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
				}
				comptroller.ApproveAndMint(cfg, acc, genSdk, ftokenAddress, otokenAddress, acc.Address,
					new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
				comptroller.EnterMarkets(cfg, acc, genSdk, flashAddress, acc.Address, []interface{}{ftokenAddress})
				amount, err := rand.Int(rand.Reader, big.NewInt(50))
				if err != nil {
					log.Errorf("batchOperate, common.AddressFromHexString error: %s", err)
				}
				comptroller.Borrow(cfg, acc, genSdk, ftokenAddress, acc.Address,
					new(big.Int).Mul(amount, new(big.Int).SetUint64(uint64(math.Pow10(int(decimalMap[otokenAddress.ToHexString()]))))))
			}
		}()
		time.Sleep(1000 * time.Millisecond)
	}
}
