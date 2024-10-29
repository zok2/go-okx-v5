package publicdata

import (
	"github.com/zok2/go-okx-v5"
)

type (
	Instrument struct {
		InstID       string               `json:"instId"`                 // 产品ID
		InstType     okex.InstrumentType  `json:"instType"`               // 产品类型
		Uly          string               `json:"uly,omitempty"`          // 标的指数
		InstFamily   string               `json:"instFamily,omitempty"`   // 交易品种
		BaseCcy      string               `json:"baseCcy,omitempty"`      // 交易货币币种
		QuoteCcy     string               `json:"quoteCcy,omitempty"`     // 计价货币币种
		SettleCcy    string               `json:"settleCcy,omitempty"`    // 盈亏结算和保证金币种
		CtValCcy     string               `json:"ctValCcy,omitempty"`     // 合约面值计价币种
		CtVal        okex.JSONFloat64     `json:"ctVal,omitempty"`        // 合约面值
		CtMult       okex.JSONFloat64     `json:"ctMult,omitempty"`       // 合约乘数
		Stk          okex.JSONFloat64     `json:"stk,omitempty"`          // 行权价格
		TickSz       okex.JSONFloat64     `json:"tickSz,omitempty"`       // 下单价格精度
		LotSz        okex.JSONFloat64     `json:"lotSz,omitempty"`        // 下单数量精度
		MinSz        okex.JSONFloat64     `json:"minSz,omitempty"`        // 最小下单数量
		Lever        okex.JSONFloat64     `json:"lever"`                  // 最大杠杆倍数
		OptType      okex.OptionType      `json:"optType,omitempty"`      // 期权类型
		ListTime     okex.JSONTime        `json:"listTime"`               // 上线时间
		ExpTime      okex.JSONTime        `json:"expTime,omitempty"`      // 产品下线时间
		CtType       okex.ContractType    `json:"ctType,omitempty"`       // 合约类型
		Alias        okex.AliasType       `json:"alias,omitempty"`        // 合约日期别名
		State        okex.InstrumentState `json:"state"`                  // 产品状态
		RuleType     string               `json:"ruleType,omitempty"`     // 交易规则类型
		MaxLmtSz     okex.JSONFloat64     `json:"maxLmtSz,omitempty"`     // 限价单的单笔最大委托数量
		MaxMktSz     okex.JSONFloat64     `json:"maxMktSz,omitempty"`     // 市价单的单笔最大委托数量
		MaxLmtAmt    okex.JSONFloat64     `json:"maxLmtAmt,omitempty"`    // 限价单的单笔最大美元价值
		MaxMktAmt    okex.JSONFloat64     `json:"maxMktAmt,omitempty"`    // 市价单的单笔最大美元价值
		MaxTwapSz    okex.JSONFloat64     `json:"maxTwapSz,omitempty"`    // 时间加权单的单笔最大委托数量
		MaxIcebergSz okex.JSONFloat64     `json:"maxIcebergSz,omitempty"` // 冰山委托的单笔最大委托数量
		MaxTriggerSz okex.JSONFloat64     `json:"maxTriggerSz,omitempty"` // 计划委托的单笔最大委托数量
		MaxStopSz    okex.JSONFloat64     `json:"maxStopSz,omitempty"`    // 止盈止损市价委托的单笔最大委托数量
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     okex.JSONFloat64          `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       okex.JSONFloat64    `json:"oi"`
		OiCcy    okex.JSONFloat64    `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     okex.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okex.JSONFloat64    `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   okex.JSONFloat64    `json:"buyLmt"`
		SellLmt  okex.JSONFloat64    `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx okex.JSONFloat64    `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    okex.JSONFloat64    `json:"delta"`
		Gamma    okex.JSONFloat64    `json:"gamma"`
		Vega     okex.JSONFloat64    `json:"vega"`
		Theta    okex.JSONFloat64    `json:"theta"`
		DeltaBS  okex.JSONFloat64    `json:"deltaBS"`
		GammaBS  okex.JSONFloat64    `json:"gammaBS"`
		VegaBS   okex.JSONFloat64    `json:"vegaBS"`
		ThetaBS  okex.JSONFloat64    `json:"thetaBS"`
		Lever    okex.JSONFloat64    `json:"lever"`
		MarkVol  okex.JSONFloat64    `json:"markVol"`
		BidVol   okex.JSONFloat64    `json:"bidVol"`
		AskVol   okex.JSONFloat64    `json:"askVol"`
		RealVol  okex.JSONFloat64    `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string           `json:"ccy"`
		Amt          okex.JSONFloat64 `json:"amt"`
		DiscountLv   okex.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex.JSONInt64 `json:"discountRate"`
		MaxAmt       okex.JSONInt64 `json:"maxAmt"`
		MinAmt       okex.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss okex.JSONFloat64          `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    okex.JSONFloat64  `json:"bkPx"`
		Sz      okex.JSONFloat64  `json:"sz"`
		BkLoss  okex.JSONFloat64  `json:"bkLoss"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   okex.JSONFloat64    `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JSONInt64      `json:"tier"`
		MinSz        okex.JSONFloat64    `json:"minSz"`
		MaxSz        okex.JSONFloat64    `json:"maxSz"`
		Mmr          okex.JSONFloat64    `json:"mmr"`
		Imr          okex.JSONFloat64    `json:"imr"`
		OptMgnFactor okex.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okex.JSONFloat64    `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string           `json:"ccy"`
		Rate  okex.JSONFloat64 `json:"rate"`
		Quota okex.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string           `json:"level"`
		IrDiscount    okex.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int              `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string        `json:"title"`
		State       string        `json:"state"`
		Href        string        `json:"href"`
		ServiceType string        `json:"serviceType"`
		System      string        `json:"system"`
		ScheDesc    string        `json:"scheDesc"`
		Begin       okex.JSONTime `json:"begin"`
		End         okex.JSONTime `json:"end"`
	}
)
