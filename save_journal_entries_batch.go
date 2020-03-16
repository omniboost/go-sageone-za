package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewSaveJournalEntriesBatchRequest() SaveJournalEntriesBatchRequest {
	r := SaveJournalEntriesBatchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSaveJournalEntriesBatchQueryParams()
	r.pathParams = r.NewSaveJournalEntriesBatchPathParams()
	r.requestBody = r.NewSaveJournalEntriesBatchRequestBody()
	return r
}

type SaveJournalEntriesBatchRequest struct {
	client      *Client
	queryParams *SaveJournalEntriesBatchQueryParams
	pathParams  *SaveJournalEntriesBatchPathParams
	method      string
	headers     http.Header
	requestBody SaveJournalEntriesBatchRequestBody
}

func (r SaveJournalEntriesBatchRequest) NewSaveJournalEntriesBatchQueryParams() *SaveJournalEntriesBatchQueryParams {
	return &SaveJournalEntriesBatchQueryParams{}
}

type SaveJournalEntriesBatchQueryParams struct {
	CompanyID int `schema:"CompanyId"`
}

func (p SaveJournalEntriesBatchQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveJournalEntriesBatchRequest) QueryParams() *SaveJournalEntriesBatchQueryParams {
	return r.queryParams
}

func (r SaveJournalEntriesBatchRequest) NewSaveJournalEntriesBatchPathParams() *SaveJournalEntriesBatchPathParams {
	return &SaveJournalEntriesBatchPathParams{}
}

type SaveJournalEntriesBatchPathParams struct {
}

func (p *SaveJournalEntriesBatchPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveJournalEntriesBatchRequest) PathParams() *SaveJournalEntriesBatchPathParams {
	return r.pathParams
}

func (r *SaveJournalEntriesBatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveJournalEntriesBatchRequest) Method() string {
	return r.method
}

func (r SaveJournalEntriesBatchRequest) NewSaveJournalEntriesBatchRequestBody() SaveJournalEntriesBatchRequestBody {
	return SaveJournalEntriesBatchRequestBody{}
}

type SaveJournalEntriesBatchRequestBody []NewJournalEntry

func (r *SaveJournalEntriesBatchRequest) RequestBody() *SaveJournalEntriesBatchRequestBody {
	return &r.requestBody
}

func (r *SaveJournalEntriesBatchRequest) SetRequestBody(body SaveJournalEntriesBatchRequestBody) {
	r.requestBody = body
}

func (r *SaveJournalEntriesBatchRequest) NewResponseBody() *SaveJournalEntriesBatchResponseBody {
	return &SaveJournalEntriesBatchResponseBody{}
}

type SaveJournalEntriesBatchResponseBody JournalEntries

func (r *SaveJournalEntriesBatchRequest) URL() url.URL {
	return r.client.GetEndpointURL("/JournalEntry/SaveBatch", r.PathParams())
}

func (r *SaveJournalEntriesBatchRequest) Do() (SaveJournalEntriesBatchResponseBody, error) {
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
