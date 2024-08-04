package financial

type (
	SavingBalance struct {
		Ccy        string `json:"ccy"`        // 币种
		Earnings   string `json:"earnings"`   // 存款收益
		RedemptAmt string `json:"redemptAmt"` // 赎回金额
		Rate       string `json:"rate"`       // 利率
		Amt        string `json:"amt"`        // 总金额
		LoanAmt    string `json:"loanAmt"`    // 借款金额
		PendingAmt string `json:"pendingAmt"` // 待处理金额
	}
	SavingsPurchaseRedempt struct {
		Ccy  string `json:"ccy"`  // 币种名称
		Amt  string `json:"amt"`  // 申购/赎回数量
		Side string `json:"side"` // 操作类型
		Rate string `json:"rate"` // 申购年利率
	}
	SetLendingRate struct {
		Ccy  string `json:"ccy"`  // 币种名称，如 BTC
		Rate string `json:"rate"` // 贷出年利率，范围在 1% 到 365% 之间
	}
	LendingHistory struct {
		Ccy      string `json:"ccy"`      // 币种，如 BTC
		Amt      string `json:"amt"`      // 出借数量
		Earnings string `json:"earnings"` // 已赚取利息
		Rate     string `json:"rate"`     // 出借年利率
		Ts       string `json:"ts"`       // 出借时间，Unix 时间戳的毫秒数格式，如 1597026383085
	}
	LendingRateSummary struct {
		Ccy       string `json:"ccy"`       // 币种名称，如 BTC
		AvgAmt    string `json:"avgAmt"`    // 过去24小时平均借贷量
		AvgAmtUsd string `json:"avgAmtUsd"` // 过去24小时平均借贷美元价值
		AvgRate   string `json:"avgRate"`   // 过去24小时平均借出利率
		PreRate   string `json:"preRate"`   // 上一次借贷年利率
		EstRate   string `json:"estRate"`   // 下一次预估借贷年利率
	}
	LendingRateHistory struct {
		Ccy  string `json:"ccy"`  // Currency type, e.g., BTC
		Amt  string `json:"amt"`  // Total lending amount in the market
		Rate string `json:"rate"` // Annual lending rate
		Ts   string `json:"ts"`   // Timestamp in milliseconds since Unix epoch
	}
)
