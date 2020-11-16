package test_case

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigTools struct {
	GlobalParamAddr string `json:"global_param_addr"`
	GasPrice int `json:"gas_price"`
	GasLimit int64 `json:"gas_limit"`
	NodeRPCAddr string `json:"node_rpc_addr"`
	Governance struct {
		FilePath string `json:"file_path"`
		Address string `json:"address"`
		WingAddr string `json:"wing_addr"`
		ProfitAddr string `json:"profit_addr"`
		OracleAddr string `json:"oracle_addr"`
		GlobalParamAddr string `json:"global_param_addr"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
	} `json:"governance"`
	ApproveContract struct {
		FilePath string `json:"file_path"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
		Address string `json:"address"`
		Admin string `json:"admin"`
	} `json:"approve_contract"`
	Comptroller struct {
		FilePath string `json:"file_path"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
		Address string `json:"address"`
		Admin string `json:"admin"`
		OracleAddr string `json:"oracle_addr"`
		WingAddr string `json:"wing_addr"`
		CloseFactor string `json:"close_factor"`
		InsuranceRepayFactor string `json:"insurance_repay_factor"`
		MaxAssets string `json:"max_assets"`
		LiquidationIncentive string `json:"liquidation_incentive"`
		WingRate string `json:"wing_rate"`
	} `json:"comptroller"`
	Oracle struct {
		FilePath string `json:"file_path"`
		Decimals int `json:"decimals"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
		Address string `json:"address"`
		Admin string `json:"admin"`
	} `json:"oracle"`
	InterestRateModel struct {
		FilePath string `json:"file_path"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
		Address string `json:"address"`
		BaseRatePerYear string `json:"base_rate_per_year"`
		MultiplierPerYear string `json:"multiplier_per_year"`
	} `json:"interest_rate_model"`
	Tokens []struct {
		Name string `json:"name"`
		UnderlyingAddr string `json:"underlying_addr"`
		UnderlyingDecimals int `json:"underlying_decimals"`
		OracleAddr string `json:"oracle_addr"`
		ComptrollerAddr string `json:"comptroller_addr"`
		InterestRateModelAddr string `json:"interest_rate_model_addr"`
		InsuranceAddr string `json:"insurance_addr"`
		OriginalAddr string `json:"original_addr"`
		UnderlyingName string `json:"underlying_name"`
		InitialExchangeRate int `json:"initial_exchange_rate"`
		ReserveFactor string `json:"reserve_factor"`
		InsuranceFactor string `json:"insurance_factor"`
		WingWeight int `json:"wing_weight"`
		CollateralFactor string `json:"collateral_factor"`
		FilePath string `json:"file_path"`
		InsuranceFilePath string `json:"insurance_file_path"`
		DeployArgs struct {
			Name string `json:"name"`
			Version string `json:"version"`
			Author string `json:"author"`
			Email string `json:"email"`
			Desc string `json:"desc"`
		} `json:"deploy_args"`
		Address string `json:"address"`
		Admin string `json:"admin"`
	} `json:"tokens"`
}

func ParseConfig(path string) (*ConfigTools, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: failed, err: %s", err)
	}
	config := &ConfigTools{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: failed, err: %s", err)
	}
	return config, nil
}
