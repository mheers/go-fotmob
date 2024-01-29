package transfers

import "time"

type Transfers struct {
	Transfers []Transfer `json:"transfers"`
	Hits      int        `json:"hits"`
}

type Transfer struct {
	Name              string       `json:"name"`
	PlayerID          int          `json:"playerId"`
	Position          *Position    `json:"position"`
	TransferDate      time.Time    `json:"transferDate"`
	TransferText      []string     `json:"transferText"`
	FromClub          string       `json:"fromClub"`
	FromClubID        int          `json:"fromClubId"`
	ToClub            string       `json:"toClub"`
	ToClubID          int          `json:"toClubId"`
	Fee               *Fee         `json:"fee"`
	TransferType      TransferType `json:"transferType"`
	ContractExtension bool         `json:"contractExtension"`
	OnLoan            bool         `json:"onLoan"`
	FromDate          time.Time    `json:"fromDate"`
	ToDate            *time.Time   `json:"toDate,omitempty"`
	MarketValue       *string      `json:"marketValue,omitempty"`
}

type Fee struct {
	FeeText          string  `json:"feeText"`
	LocalizedFeeText string  `json:"localizedFeeText"`
	Value            *string `json:"value,omitempty"`
}

type Position struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type TransferType struct {
	Text            string `json:"text"`
	LocalizationKey string `json:"localizationKey"`
}
