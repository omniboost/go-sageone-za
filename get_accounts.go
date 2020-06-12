package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/odata"
	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetAccountsRequest() GetAccountsRequest {
	r := GetAccountsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetAccountsQueryParams()
	r.pathParams = r.NewGetAccountsPathParams()
	r.requestBody = r.NewGetAccountsRequestBody()
	return r
}

type GetAccountsRequest struct {
	client      *Client
	queryParams *GetAccountsQueryParams
	pathParams  *GetAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsRequestBody
}

func (r GetAccountsRequest) NewGetAccountsQueryParams() *GetAccountsQueryParams {
	return &GetAccountsQueryParams{
		Pagination: odata.NewPagination(),
	}
}

type GetAccountsQueryParams struct {
	odata.Pagination
	CompanyID int `schema:"CompanyId"`
}

func (p GetAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountsRequest) QueryParams() *GetAccountsQueryParams {
	return r.queryParams
}

func (r GetAccountsRequest) NewGetAccountsPathParams() *GetAccountsPathParams {
	return &GetAccountsPathParams{}
}

type GetAccountsPathParams struct {
}

func (p *GetAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsRequest) PathParams() *GetAccountsPathParams {
	return r.pathParams
}

func (r *GetAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsRequest) Method() string {
	return r.method
}

func (r GetAccountsRequest) NewGetAccountsRequestBody() GetAccountsRequestBody {
	return GetAccountsRequestBody{}
}

type GetAccountsRequestBody struct{}

func (r *GetAccountsRequest) RequestBody() *GetAccountsRequestBody {
	return &r.requestBody
}

func (r *GetAccountsRequest) SetRequestBody(body GetAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsRequest) NewResponseBody() *GetAccountsResponseBody {
	return &GetAccountsResponseBody{}
}

type GetAccountsResponseBody struct {
	TotalResults    int `json:"TotalResults"`
	ReturnedResults int `json:"ReturnedResults"`
	Results         Accounts
}

func (r *GetAccountsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/Account/Get", r.PathParams())
}

func (r *GetAccountsRequest) Do() (GetAccountsResponseBody, error) {
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

type Accounts []Account

type Account struct {
	Name     string `json:"Name"`
	Category struct {
		Comment     string `json:"Comment"`
		Order       int    `json:"Order"`
		Description string `json:"Description"`
		ID          int    `json:"ID"`
		Modified    string `json:"Modified"`
		Created     string `json:"Created"`
	} `json:"Category"`
	Active             bool    `json:"Active"`
	Balance            float64 `json:"Balance"`
	Description        string  `json:"Description"`
	ReportingGroupID   int     `json:"ReportingGroupId"`
	UnallocatedAccount bool    `json:"UnallocatedAccount"`
	IsTaxLocked        bool    `json:"IsTaxLocked"`
	Modified           string  `json:"Modified"`
	Created            string  `json:"Created"`
	AccountType        int     `json:"AccountType"`
	HasActivity        bool    `json:"HasActivity"`
	DefaultTaxTypeID   int     `json:"DefaultTaxTypeId"`
	DefaultTaxType     struct {
		ID                int     `json:"ID"`
		Name              string  `json:"Name"`
		Percentage        float64 `json:"Percentage"`
		IsDefault         bool    `json:"IsDefault"`
		HasActivity       bool    `json:"HasActivity"`
		IsManualTax       bool    `json:"IsManualTax"`
		Active            bool    `json:"Active"`
		Created           string  `json:"Created"`
		Modified          string  `json:"Modified"`
		TaxTypeDefaultUID string  `json:"TaxTypeDefaultUID"`
	} `json:"DefaultTaxType"`
	ID int `json:"ID"`
}
