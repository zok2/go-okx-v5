package funding

import "github.com/zok2/go-okx-v5"

type (
	Currency struct {
		Ccy                  string `json:"ccy"`                  // 币种名称，如 BTC
		Name                 string `json:"name"`                 // 币种名称，不显示则无对应名称
		LogoLink             string `json:"logoLink"`             // 币种Logo链接
		Chain                string `json:"chain"`                // 币种链信息，如 USDT-ERC20，USDT-TRC20
		CanDep               bool   `json:"canDep"`               // 当前是否可充值
		CanWd                bool   `json:"canWd"`                // 当前是否可提币
		CanInternal          bool   `json:"canInternal"`          // 当前是否可内部转账
		MinDep               string `json:"minDep"`               // 币种单笔最小充值量
		MinWd                string `json:"minWd"`                // 币种单笔最小链上提币量
		MaxWd                string `json:"maxWd"`                // 币种单笔最大链上提币量
		WdTickSz             string `json:"wdTickSz"`             // 提币精度,表示小数点后的位数
		WdQuota              string `json:"wdQuota"`              // 过去24小时内提币额度，单位为USD
		UsedWdQuota          string `json:"usedWdQuota"`          // 过去24小时内已用提币额度，单位为USD
		MinFee               string `json:"minFee"`               // 普通地址最小提币手续费数量
		MaxFee               string `json:"maxFee"`               // 普通地址最大提币手续费数量
		MinFeeForCtAddr      string `json:"minFeeForCtAddr"`      // 合约地址最小提币手续费数量
		MaxFeeForCtAddr      string `json:"maxFeeForCtAddr"`      // 合约地址最大提币手续费数量
		MainNet              bool   `json:"mainNet"`              // 当前链是否为主链
		NeedTag              bool   `json:"needTag"`              // 当前链是否需要标签（tag/memo）信息
		MinDepArrivalConfirm string `json:"minDepArrivalConfirm"` // 充值到账最小网络确认数
		MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`   // 提现解锁最小网络确认数
		DepQuotaFixed        string `json:"depQuotaFixed"`        // 充币固定限额，单位为USD
		UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`    // 已用充币固定额度，单位为USD
		DepQuoteDailyLayer2  string `json:"depQuoteDailyLayer2"`  // Layer2网络每日充值上限
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     okex.JSONFloat64 `json:"amt"`
		From    okex.AccountType `json:"from,string"`
		To      okex.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string           `json:"billId"`
		Ccy    string           `json:"ccy"`
		Bal    okex.JSONFloat64 `json:"bal"`
		BalChg okex.JSONFloat64 `json:"balChg"`
		Type   okex.BillType    `json:"type,string"`
		TS     okex.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtID    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to,string"`
		TS       okex.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxID  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   okex.JSONFloat64  `json:"amt"`
		State okex.DepositState `json:"state,string"`
		TS    okex.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		WdID  okex.JSONInt64   `json:"wdId"`
		Amt   okex.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		TxID  string               `json:"txId"`
		From  string               `json:"from"`
		To    string               `json:"to"`
		Tag   string               `json:"tag,omitempty"`
		PmtID string               `json:"pmtId,omitempty"`
		Memo  string               `json:"memo,omitempty"`
		Amt   okex.JSONFloat64     `json:"amt"`
		Fee   okex.JSONFloat64     `json:"fee"`
		WdID  okex.JSONInt64       `json:"wdId"`
		State okex.WithdrawalState `json:"state,string"`
		TS    okex.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string           `json:"ccy"`
		Amt  okex.JSONFloat64 `json:"amt"`
		Side okex.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy"`
		Amt      okex.JSONFloat64 `json:"amt"`
		Earnings okex.JSONFloat64 `json:"earnings"`
	}
)
