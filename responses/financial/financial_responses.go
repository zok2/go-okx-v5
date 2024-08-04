package financial

import (
	models "github.com/zok2/go-okx-v5/models/financial"
	"github.com/zok2/go-okx-v5/responses"
)

type (
	GetSavingBalance struct {
		responses.Basic
		SavingBalances []*models.SavingBalance `json:"data"`
	}
	SavingsPurchaseRedempt struct {
		responses.Basic
		SavingsPurchaseRedempt []*models.SavingsPurchaseRedempt
	}
	SetLendingRate struct {
		responses.Basic
		SetLendingRate []*models.SetLendingRate
	}
	GetLendingHistory struct {
		responses.Basic
		LendingHistory []*models.LendingHistory
	}
	GetLendingRateSummary struct {
		responses.Basic
		LendingRateSummary []*models.LendingRateSummary
	}
	GetLendingRateHistory struct {
		responses.Basic
		LendingRateHistory []*models.LendingRateHistory
	}
)
