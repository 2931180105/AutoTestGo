package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"io/ioutil"
	"sync"
	"testing"
)

func TestReadFile(t *testing.T) {
	fileContent, err := ioutil.ReadFile("./hash.txt")
	if err != nil {
		log.Errorf("NewConfigFromFile: %s", err)
	}
	log.Infof("file :%v", string(fileContent))
}
func TestWingSpeeds2(t *testing.T) {
	syn := new(sync.WaitGroup)
	for i := 0; i < 1; i++ {
		syn.Add(1)
		go TestWingSpeeds("WBTC", "AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p", syn)
	}
	syn.Wait()
}
//func TestBorrowRateByTime(t *testing.T) {
//	syn := new(sync.WaitGroup)
//	for i := 0; i < len(MarketNames); i++ {
//		syn.Add(1)
//		go TestBorrowRateNew(MarketNames[i], syn)
//	}
//	syn.Wait()
//}


/*
只有抵押借款才能分wing
根据新的需求测试wing的分配（修改之前的代码）
将测试结果保存到文件中
1.获取所有的市场
ftokenAddressList, err := comptroller.GetAllMarkets(genSdk, cfg.Comptroller)
2.根据市场的名称匹配对应的ftoken

 */
