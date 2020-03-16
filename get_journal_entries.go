package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetJournalEntriesRequest() GetJournalEntriesRequest {
	r := GetJournalEntriesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetJournalEntriesQueryParams()
	r.pathParams = r.NewGetJournalEntriesPathParams()
	r.requestBody = r.NewGetJournalEntriesRequestBody()
	return r
}

type GetJournalEntriesRequest struct {
	client      *Client
	queryParams *GetJournalEntriesQueryParams
	pathParams  *GetJournalEntriesPathParams
	method      string
	headers     http.Header
	requestBody GetJournalEntriesRequestBody
}

func (r GetJournalEntriesRequest) NewGetJournalEntriesQueryParams() *GetJournalEntriesQueryParams {
	return &GetJournalEntriesQueryParams{}
}

type GetJournalEntriesQueryParams struct {
	CompanyID int `schema:"CompanyId"`
}

func (p GetJournalEntriesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetJournalEntriesRequest) QueryParams() *GetJournalEntriesQueryParams {
	return r.queryParams
}

func (r GetJournalEntriesRequest) NewGetJournalEntriesPathParams() *GetJournalEntriesPathParams {
	return &GetJournalEntriesPathParams{}
}

type GetJournalEntriesPathParams struct {
}

func (p *GetJournalEntriesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetJournalEntriesRequest) PathParams() *GetJournalEntriesPathParams {
	return r.pathParams
}

func (r *GetJournalEntriesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetJournalEntriesRequest) Method() string {
	return r.method
}

func (r GetJournalEntriesRequest) NewGetJournalEntriesRequestBody() GetJournalEntriesRequestBody {
	return GetJournalEntriesRequestBody{}
}

type GetJournalEntriesRequestBody struct {
}

func (r *GetJournalEntriesRequest) RequestBody() *GetJournalEntriesRequestBody {
	return &r.requestBody
}

func (r *GetJournalEntriesRequest) SetRequestBody(body GetJournalEntriesRequestBody) {
	r.requestBody = body
}

func (r *GetJournalEntriesRequest) NewResponseBody() *GetJournalEntriesResponseBody {
	return &GetJournalEntriesResponseBody{}
}

type GetJournalEntriesResponseBody struct {
	TotalResults    int
	ReturnedResults int
	Results         JournalEntries
}

func (r *GetJournalEntriesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/JournalEntry/Get", r.PathParams())
}

func (r *GetJournalEntriesRequest) Do() (GetJournalEntriesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

type JournalEntries []JournalEntry

type JournalEntry struct {
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
