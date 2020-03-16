package sage

type NewJournalEntry struct {
	ID                  int     `json:"ID"`
	Date                Date    `json:"Date"`
	Effect              Effect  `json:"Effect"`
	AccountID           int     `json:"AccountId"`
	Reference           string  `json:"Reference"`
	Description         string  `json:"Description"`
	TaxTypeID           int     `json:"TaxTypeId"`
	TaxPeriodID         int     `json:"TaxPeriodId"`
	Exclusive           float64 `json:"Exclusive"`
	Tax                 float64 `json:"Tax"`
	Total               float64 `json:"Total"`
	ContraAccountID     int     `json:"ContraAccountId"`
	Memo                string  `json:"Memo"`
	HasAttachments      bool    `json:"HasAttachments"`
	AnalysisCategoryID1 int     `json:"AnalysisCategoryId1"`
	AnalysisCategoryID2 int     `json:"AnalysisCategoryId2"`
	AnalysisCategoryID3 int     `json:"AnalysisCategoryId3"`
	Editable            bool    `json:"Editable"`
	Locked              bool    `json:"Locked"`
	TrackingCode        string  `json:"$TrackingCode"`
	Debit               float64 `json:"Debit"`
	Credit              float64 `json:"Credit"`
	BusinessID          int     `json:"BusinessId"`
	PayRunID            int     `json:"PayRunId"`
}
