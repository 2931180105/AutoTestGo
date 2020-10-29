package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"io/ioutil"
	"sync"
	"testing"
)

func TestTestRunner_WingSpeed4BorrowTest(t *testing.T) {

	testRunner, err := NewTestRunner("/home/ubuntu/go/src/github.com/mockyz/AutoTestGo/wing-test/config_prv.json",)
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	testRunner.WingSpeed4BorrowTest("APcNMJBUVEnbdRKMuBuhUGjJZbhpRKgkwY")
}
func TestTestRunner_WingSpeed4SuppluyTestTest(t *testing.T) {
	testRunner, err := NewTestRunner("config/config.json",)
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	testRunner.WingSpeed4SuppluyTest("APcNMJBUVEnbdRKMuBuhUGjJZbhpRKgkwY")
}
func TestTestRunner_WingSpeed4InsuranceTest(t *testing.T) {
	testRunner, err := NewTestRunner("config.json")
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	testRunner.WingSpeed4InsuranceTest("APcNMJBUVEnbdRKMuBuhUGjJZbhpRKgkwY")
}
func TestReadFile(t *testing.T) {
	fileContent, err := ioutil.ReadFile("./hash.txt")
	if err != nil {
		log.Errorf("NewConfigFromFile: %s", err)
	}
	log.Infof("file :%v", string(fileContent))
}
func TestGetTimeByTxhash(t *testing.T) {
	hashTimes := make([]uint32, 0)
	blockHeights := make([]uint32, 0)
	hashs := []string{"95a73fc63eff77eef89fc7d440f168eeb72ae716a6881e816a0f3cfa81e2ad0d",
		"25f9f32a744d3398e2faa238d6ba68a564db542de79130ef2c4bf58eb2c7e2c4",
		"b865eb87967f111af94f84cfe6970a7b99c6b272360c3026b2bbe62afedc22e5",
		"d9e50650f2498df36d20be3a37342f8178a598390fce26a56d48bd929d56ffce",
		"eeeaad44565f35eaad098385354a630efed2763cf0c13f92d4683f528c71b4e1",
		"9f8989f5e4d19c2b29c82ad0f2b6cfb954f2740c3e1a700161f61dcb81622748",
		"2b61e614eb5c9cc12d95262408a12d9e4610a79eea2ac4d2f2d1bff99777013f",
		"465532bc680a5bdc8471f41716e2384ca0d005be2f08fbc9d9142f4fc418bc8d",
		"90a54b3903762f23a4d67efd14ab7c638a4d9dca0e5ee9877226942cddeb4c2e",
		"79b8aa61db9a94ab24b7e96f5d9301ff097cb2592fb90e54a698b386e31fbce8",
		"8e9fb0b19b2336b2d3b41ba60fb6599f1ddbfb923a43dc3c89b90432df863976",
		"3a69d97ac99574049417cbbf5865fbfb54b8cc66690c24e2f8a08261f714316f"}
	cfgFile := "../config.json"
	tr1, err := NewTestRunner(cfgFile)
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	for i := 0; i < len(hashs); i++ {
		blockHeight, err := tr1.OntSDk.GetBlockHeightByTxHash(hashs[i])
		if err != nil {
			log.Errorf("NewTestRunner err : %v", err)
		}
		log.Infof("blockHeight:%v", blockHeight)
		//log.Infof("blockTime:%v",GetTimeByTxhash(tr1.OntSDk,hashs[i]))
		//println(i,blockHeight,GetTimeByTxhash(tr1.OntSDk,hashs[i]))
		blockHeights = append(blockHeights, blockHeight)
		hashTimes = append(hashTimes, GetTimeByTxhash(tr1.OntSDk, hashs[i]))
	}
	log.Infof("hashTimes :%v", hashTimes)
	log.Infof("blockHeights :%v", blockHeights)

}
func TestGetTimeByTxhash2(t *testing.T) {
	//	133598b7595ef17c5910bfa8fd1256cb10c5bf390f9445b7e8e7e413ea5e7e59
	cfgFile := "../config.json"
	hash := "7bd3c03aa47e07cd65dc4b5222a6f37dcd784f38decee1b6638a794769b67fd8"
	tr1, err := NewTestRunner(cfgFile)
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	GetTimeByTxhash(tr1.OntSDk, hash)

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
