package financial

import okex "github.com/zok2/go-okx-v5"

type (
	GetSavingBalance struct {
		Ccy string `json:"ccy"`
	}
	SavingsPurchaseRedempt struct {
		Ccy  string           `json:"ccy"`  // 币种名称，如 BTC
		Amt  string           `json:"amt"`  // 申购/赎回 数量
		Side string           `json:"side"` // 操作类型，"purchase" 表示申购，"redempt" 表示赎回
		Rate okex.JSONFloat64 `json:"rate"` // 申购年利率，适用于申购操作
	}
	SetLendingRate struct {
		Ccy  string           `json:"ccy"`  // 币种名称，如 BTC
		Rate okex.JSONFloat64 `json:"rate"` // 贷出年利率，范围在 1% 到 365% 之间
	}
	GetLendingHistory struct {
		Ccy    string `json:"ccy,omitempty"` // 币种名称，如 BTC
		After  int64  `json:"after,omitempty,string"`
		Before int64  `json:"before,omitempty,string"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
	GetLendingRateSummary struct {
		Ccy string `json:"ccy,omitempty"`
	}
	GetLendingRateHistory struct {
		Ccy    string `json:"ccy,omitempty"` // Currency type, e.g., BTC. Optional.
		After  int64  `json:"after,omitempty,string"`
		Before int64  `json:"before,omitempty,string"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
)
