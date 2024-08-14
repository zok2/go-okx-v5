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
		SavingsPurchaseRedempt []*models.SavingsPurchaseRedempt `json:"data"`
	}
	SetLendingRate struct {
		responses.Basic
		SetLendingRate []*models.SetLendingRate `json:"data"`
	}
	GetLendingHistory struct {
		responses.Basic
		LendingHistory []*models.LendingHistory `json:"data"`
	}
	GetLendingRateSummary struct {
		responses.Basic
		LendingRateSummary []*models.LendingRateSummary `json:"data"`
	}
	GetLendingRateHistory struct {
		responses.Basic
		LendingRateHistory []*models.LendingRateHistory `json:"data"`
	}
)
