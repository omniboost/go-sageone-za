package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/odata"
	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetChartOfAccountsRequest() GetChartOfAccountsRequest {
	r := GetChartOfAccountsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetChartOfAccountsQueryParams()
	r.pathParams = r.NewGetChartOfAccountsPathParams()
	r.requestBody = r.NewGetChartOfAccountsRequestBody()
	return r
}

type GetChartOfAccountsRequest struct {
	client      *Client
	queryParams *GetChartOfAccountsQueryParams
	pathParams  *GetChartOfAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetChartOfAccountsRequestBody
}

func (r GetChartOfAccountsRequest) NewGetChartOfAccountsQueryParams() *GetChartOfAccountsQueryParams {
	return &GetChartOfAccountsQueryParams{
		Pagination: odata.NewPagination(),
	}
}

type GetChartOfAccountsQueryParams struct {
	odata.Pagination
	CompanyID int `schema:"CompanyId"`
}

func (p GetChartOfAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetChartOfAccountsRequest) QueryParams() *GetChartOfAccountsQueryParams {
	return r.queryParams
}

func (r GetChartOfAccountsRequest) NewGetChartOfAccountsPathParams() *GetChartOfAccountsPathParams {
	return &GetChartOfAccountsPathParams{}
}

type GetChartOfAccountsPathParams struct {
}

func (p *GetChartOfAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetChartOfAccountsRequest) PathParams() *GetChartOfAccountsPathParams {
	return r.pathParams
}

func (r *GetChartOfAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetChartOfAccountsRequest) Method() string {
	return r.method
}

func (r GetChartOfAccountsRequest) NewGetChartOfAccountsRequestBody() GetChartOfAccountsRequestBody {
	return GetChartOfAccountsRequestBody{}
}

type GetChartOfAccountsRequestBody struct{}

func (r *GetChartOfAccountsRequest) RequestBody() *GetChartOfAccountsRequestBody {
	return &r.requestBody
}

func (r *GetChartOfAccountsRequest) SetRequestBody(body GetChartOfAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetChartOfAccountsRequest) NewResponseBody() *GetChartOfAccountsResponseBody {
	return &GetChartOfAccountsResponseBody{}
}

type GetChartOfAccountsResponseBody struct {
	TotalResults    int `json:"TotalResults"`
	ReturnedResults int `json:"ReturnedResults"`
	Results         Accounts
}

func (r *GetChartOfAccountsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/Account/Get", r.PathParams())
}

func (r *GetChartOfAccountsRequest) Do() (GetChartOfAccountsResponseBody, error) {
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
