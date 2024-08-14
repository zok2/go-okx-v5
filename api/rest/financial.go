package rest

import (
	"encoding/json"
	okex "github.com/zok2/go-okx-v5"
	requests "github.com/zok2/go-okx-v5/requests/rest/financial"
	responses "github.com/zok2/go-okx-v5/responses/financial"
	"net/http"
)

// Financial
//
// https://www.okx.com/docs-v5/zh/#financial-product
type Financial struct {
	client *ClientRest
}

// NewFinancial returns a pointer to a fresh Financial
func NewFinancial(c *ClientRest) *Financial {
	return &Financial{c}
}

func (c Financial) GetSavingBalance(req requests.GetSavingBalance) (response responses.GetSavingBalance, err error) {
	p := "/api/v5/finance/savings/balance"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

func (c Financial) SetLendingRate(req requests.SetLendingRate) (response responses.SetLendingRate, err error) {
	p := "/api/v5/finance/savings/set-lending-rate"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
func (c Financial) GetLendingHistory(req requests.GetLendingHistory) (response responses.GetLendingHistory, err error) {
	p := "/api/v5/finance/savings/lending-history"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

func (c Financial) GetLendingRateSummary(req requests.GetLendingRateSummary) (response responses.GetLendingRateSummary, err error) {
	p := "/api/v5/finance/savings/lending-rate-summary"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

func (c Financial) GetLendingRateHistory(req requests.GetLendingRateHistory) (response responses.GetLendingRateHistory, err error) {
	p := "/api/v5/finance/savings/lending-rate-history"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
