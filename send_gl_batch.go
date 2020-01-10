package aktiva

import (
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"
	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewSendGLBatchRequest() SendGLBatchRequest {
	r := SendGLBatchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSendGLBatchQueryParams()
	r.pathParams = r.NewSendGLBatchPathParams()
	r.requestBody = r.NewSendGLBatchRequestBody()
	return r
}

type SendGLBatchRequest struct {
	client      *Client
	queryParams *SendGLBatchQueryParams
	pathParams  *SendGLBatchPathParams
	method      string
	headers     http.Header
	requestBody SendGLBatchRequestBody
}

func (r SendGLBatchRequest) NewSendGLBatchQueryParams() *SendGLBatchQueryParams {
	return &SendGLBatchQueryParams{}
}

type SendGLBatchQueryParams struct{}

func (p SendGLBatchQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SendGLBatchRequest) QueryParams() *SendGLBatchQueryParams {
	return r.queryParams
}

func (r SendGLBatchRequest) NewSendGLBatchPathParams() *SendGLBatchPathParams {
	return &SendGLBatchPathParams{}
}

type SendGLBatchPathParams struct {
}

func (p *SendGLBatchPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SendGLBatchRequest) PathParams() *SendGLBatchPathParams {
	return r.pathParams
}

func (r *SendGLBatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *SendGLBatchRequest) Method() string {
	return r.method
}

func (r SendGLBatchRequest) NewSendGLBatchRequestBody() SendGLBatchRequestBody {
	return SendGLBatchRequestBody{}
}

type SendGLBatchRequestBody NewGLBatch

func (r *SendGLBatchRequest) RequestBody() *SendGLBatchRequestBody {
	return &r.requestBody
}

func (r *SendGLBatchRequest) SetRequestBody(body SendGLBatchRequestBody) {
	r.requestBody = body
}

func (r *SendGLBatchRequest) NewResponseBody() *SendGLBatchResponseBody {
	return &SendGLBatchResponseBody{}
}

type SendGLBatchResponseBody struct {
	BatchID   uuid.UUID `json:"BatchId"`
	BatchInfo string    `json:"BatchInfo"`
}

func (r *SendGLBatchRequest) URL() url.URL {
	return r.client.GetEndpointURL("sendglbatch", r.PathParams())
}

func (r *SendGLBatchRequest) Do() (SendGLBatchResponseBody, error) {
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

type NewGLBatch struct {
	DocNo     string
	BatchDate Date
	EntryRow  []EntryRow
}

type EntryRow struct {
	AccountCode    string
	DepartmentCode string `json:"DepartmentCode,omitempty"`
	Debit          float64
	Credit         float64
	ProjectCode    string `json:"ProjectCode,omitempty"`
	CostCenterCode string `json:"CostCenterCode,omitempty"`
}
