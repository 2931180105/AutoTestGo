package if_ctrl

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"math/big"
)

type AccountLiquidity struct {
	Liquidity *big.Int
	Shortfall *big.Int
}

func DeserializeAccountLiquidity(data []byte) (*AccountLiquidity, error) {
	source := common.NewZeroCopySource(data)
	liquidityStr, _, ill, eof := source.NextString()
	if ill {
		return nil, fmt.Errorf("read liquidityStr ill")
	}
	if eof {
		return nil, fmt.Errorf("read liquidityStr eof")
	}
	shortfallStr, _, ill, eof := source.NextString()
	if ill {
		return nil, fmt.Errorf("read shortfallStr ill")
	}
	if eof {
		return nil, fmt.Errorf("read shortfallStr eof")
	}
	liquidity, ok := new(big.Int).SetString(liquidityStr, 10)
	if !ok {
		return nil, fmt.Errorf("parse liquidity %s failed", liquidityStr)
	}
	shortfall, ok := new(big.Int).SetString(shortfallStr, 10)
	if !ok {
		return nil, fmt.Errorf("parse liquidity %s failed", shortfallStr)
	}
	return &AccountLiquidity{
		Liquidity: liquidity,
		Shortfall: shortfall,
	}, nil
}

type MarketInfo struct {
	BorrowPool    common.Address
	SupplyPool    common.Address
	InsurancePool common.Address

	Underlying common.Address

	UnderlyingDecimals uint8
	WingWeight         uint8
}

func DeserializeMarketInfo(data []byte) (*MarketInfo, error) {
	source := common.NewZeroCopySource(data)
	borrow, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read borrow eof")
	}
	supply, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read supply eof")
	}
	insurance, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read insurance eof")
	}
	underlying, eof := source.NextAddress()
	if eof {
		return nil, fmt.Errorf("read underlying eof")
	}
	decimals, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read decimals eof")
	}
	wingWeight, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read wingWeight eof")
	}
	return &MarketInfo{
		BorrowPool:         borrow,
		SupplyPool:         supply,
		InsurancePool:      insurance,
		Underlying:         underlying,
		UnderlyingDecimals: decimals,
		WingWeight:         wingWeight,
	}, nil
}

type WingSBI struct {
	SupplyPortion    uint8
	BorrowPortion    uint8
	InsurancePortion uint8
}

func DeserializeWingSBI(data []byte) (*WingSBI, error) {
	source := common.NewZeroCopySource(data)
	supply, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read supply eof")
	}
	borrow, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read borrow eof")
	}
	insurance, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("read insurance eof")
	}
	return &WingSBI{
		SupplyPortion:    supply,
		BorrowPortion:    borrow,
		InsurancePortion: insurance,
	}, nil
}
