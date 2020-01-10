package aktiva

import (
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"
	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewGetGLBatchRequest() GetGLBatchRequest {
	r := GetGLBatchRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetGLBatchQueryParams()
	r.pathParams = r.NewGetGLBatchPathParams()
	r.requestBody = r.NewGetGLBatchRequestBody()
	return r
}

type GetGLBatchRequest struct {
	client      *Client
	queryParams *GetGLBatchQueryParams
	pathParams  *GetGLBatchPathParams
	method      string
	headers     http.Header
	requestBody GetGLBatchRequestBody
}

func (r GetGLBatchRequest) NewGetGLBatchQueryParams() *GetGLBatchQueryParams {
	return &GetGLBatchQueryParams{}
}

type GetGLBatchQueryParams struct {
}

func (p GetGLBatchQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetGLBatchRequest) QueryParams() *GetGLBatchQueryParams {
	return r.queryParams
}

func (r GetGLBatchRequest) NewGetGLBatchPathParams() *GetGLBatchPathParams {
	return &GetGLBatchPathParams{}
}

type GetGLBatchPathParams struct {
}

func (p *GetGLBatchPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetGLBatchRequest) PathParams() *GetGLBatchPathParams {
	return r.pathParams
}

func (r *GetGLBatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetGLBatchRequest) Method() string {
	return r.method
}

func (r GetGLBatchRequest) NewGetGLBatchRequestBody() GetGLBatchRequestBody {
	return GetGLBatchRequestBody{}
}

type GetGLBatchRequestBody struct {
	ID uuid.UUID
}

func (r *GetGLBatchRequest) RequestBody() *GetGLBatchRequestBody {
	return &r.requestBody
}

func (r *GetGLBatchRequest) SetRequestBody(body GetGLBatchRequestBody) {
	r.requestBody = body
}

func (r *GetGLBatchRequest) NewResponseBody() *GetGLBatchResponseBody {
	return &GetGLBatchResponseBody{}
}

type GetGLBatchResponseBody struct {
	Header struct {
		GLBID        string      `json:"GLBId"`
		BatchCode    string      `json:"BatchCode"`
		No           int         `json:"No"`
		Document     interface{} `json:"Document"`
		BatchDate    string      `json:"BatchDate"`
		CurrencyCode string      `json:"CurrencyCode"`
		CurrencyRate float64     `json:"CurrencyRate"`
		TotalAmount  float64     `json:"TotalAmount"`
		PriceInclVat int         `json:"PriceInclVat"`
	} `json:"Header"`
	Lines []struct {
		AccountCode    string      `json:"AccountCode"`
		Memo           string      `json:"Memo"`
		DepartmentCode interface{} `json:"DepartmentCode"`
		TaxName        string      `json:"TaxName"`
		DebitAmount    float64     `json:"DebitAmount"`
		DebitCurrency  float64     `json:"DebitCurrency"`
		CreditAmount   float64     `json:"CreditAmount"`
		CreditCurrency float64     `json:"CreditCurrency"`
		TypeID         int         `json:"TypeId"`
	} `json:"Lines"`
}

func (r *GetGLBatchRequest) URL() url.URL {
	return r.client.GetEndpointURL("getglbatch", r.PathParams())
}

func (r *GetGLBatchRequest) Do() (GetGLBatchResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
