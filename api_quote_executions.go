/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// QuoteExecutionsAPIService QuoteExecutionsAPI service
type QuoteExecutionsAPIService service

type ApiCreateQuoteExecutionRequest struct {
	ctx context.Context
	ApiService *QuoteExecutionsAPIService
	profileId string
	createQuoteExecutionRequest *CreateQuoteExecutionRequest
}

func (r ApiCreateQuoteExecutionRequest) CreateQuoteExecutionRequest(createQuoteExecutionRequest CreateQuoteExecutionRequest) ApiCreateQuoteExecutionRequest {
	r.createQuoteExecutionRequest = &createQuoteExecutionRequest
	return r
}

func (r ApiCreateQuoteExecutionRequest) Execute() (*QuoteExecution, *http.Response, error) {
	return r.ApiService.CreateQuoteExecutionExecute(r)
}

/*
CreateQuoteExecution Create Quote Execution

Execute on a quote for buying or selling an asset.

The side, market, and guaranteed price of the execution are specified by the quote
with ID `quote_id`.

The amount to buy or sell must be specified in either fiat or crypto by setting
exactly one of:
 - `base_amount` to specify the amount of crypto to buy or sell.
 - `quote_amount` to specify the amount of fiat to spend or acquire.

An otherwise-valid request to create a quote execution may fail with the following
types of errors:
 - [Expired](https://developer.paxos.com/docs/v2/problems/expired) if the quote
   with ID `quote_id` has expired.
 - [Insufficient Funds](https://developer.paxos.com/docs/v2/problems/insufficient-funds)
   if the profile with ID `profile_id` has insufficient available balance to
   fund the execution.
 - [Rejected](https://developer.paxos.com/docs/v2/problems/rejected) if extreme
   market conditions (e.g. a very large price swing) have invalidated the quote.
 - [Already Exists](https://developer.paxos.com/docs/v2/problems/already-exists)
   if a Quote Execution with the same `ref_id` has already been created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profileId The ID of the profile under which to execute this order.
 @return ApiCreateQuoteExecutionRequest
*/
func (a *QuoteExecutionsAPIService) CreateQuoteExecution(ctx context.Context, profileId string) ApiCreateQuoteExecutionRequest {
	return ApiCreateQuoteExecutionRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
	}
}

// Execute executes the request
//  @return QuoteExecution
func (a *QuoteExecutionsAPIService) CreateQuoteExecutionExecute(r ApiCreateQuoteExecutionRequest) (*QuoteExecution, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteExecution
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteExecutionsAPIService.CreateQuoteExecution")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile_id}/quote-executions"
	localVarPath = strings.Replace(localVarPath, "{"+"profile_id"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createQuoteExecutionRequest == nil {
		return localVarReturnValue, nil, reportError("createQuoteExecutionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createQuoteExecutionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetQuoteExecutionRequest struct {
	ctx context.Context
	ApiService *QuoteExecutionsAPIService
	profileId string
	id string
}

func (r ApiGetQuoteExecutionRequest) Execute() (*QuoteExecution, *http.Response, error) {
	return r.ApiService.GetQuoteExecutionExecute(r)
}

/*
GetQuoteExecution Get Quote Execution

Get an existing quote execution for buying or selling an asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profileId
 @param id
 @return ApiGetQuoteExecutionRequest
*/
func (a *QuoteExecutionsAPIService) GetQuoteExecution(ctx context.Context, profileId string, id string) ApiGetQuoteExecutionRequest {
	return ApiGetQuoteExecutionRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
		id: id,
	}
}

// Execute executes the request
//  @return QuoteExecution
func (a *QuoteExecutionsAPIService) GetQuoteExecutionExecute(r ApiGetQuoteExecutionRequest) (*QuoteExecution, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteExecution
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteExecutionsAPIService.GetQuoteExecution")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile_id}/quote-executions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"profile_id"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListQuoteExecutionsRequest struct {
	ctx context.Context
	ApiService *QuoteExecutionsAPIService
	profileId string
	side *string
	market *string
	status *string
	createdAtBegin *time.Time
	createdAtEnd *time.Time
	limit *int32
	order *string
	orderBy *string
	pageCursor *string
}

// Filter by buy or sell side.
func (r ApiListQuoteExecutionsRequest) Side(side string) ApiListQuoteExecutionsRequest {
	r.side = &side
	return r
}

// Filter by the trading pair.
func (r ApiListQuoteExecutionsRequest) Market(market string) ApiListQuoteExecutionsRequest {
	r.market = &market
	return r
}

// Filter by the status of the order.
func (r ApiListQuoteExecutionsRequest) Status(status string) ApiListQuoteExecutionsRequest {
	r.status = &status
	return r
}

// Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;.
func (r ApiListQuoteExecutionsRequest) CreatedAtBegin(createdAtBegin time.Time) ApiListQuoteExecutionsRequest {
	r.createdAtBegin = &createdAtBegin
	return r
}

// Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;.
func (r ApiListQuoteExecutionsRequest) CreatedAtEnd(createdAtEnd time.Time) ApiListQuoteExecutionsRequest {
	r.createdAtEnd = &createdAtEnd
	return r
}

// Number of results to return.
func (r ApiListQuoteExecutionsRequest) Limit(limit int32) ApiListQuoteExecutionsRequest {
	r.limit = &limit
	return r
}

// Return items in ascending (ASC) or descending (DESC) order. Default order is descending (DESC).
func (r ApiListQuoteExecutionsRequest) Order(order string) ApiListQuoteExecutionsRequest {
	r.order = &order
	return r
}

// The specific method by which the returned results will be ordered.
func (r ApiListQuoteExecutionsRequest) OrderBy(orderBy string) ApiListQuoteExecutionsRequest {
	r.orderBy = &orderBy
	return r
}

// Cursor token for fetching the next page. If using this then do not use &#x60;pagination.limit&#x60; and &#x60;pagination.offset&#x60; fields.
func (r ApiListQuoteExecutionsRequest) PageCursor(pageCursor string) ApiListQuoteExecutionsRequest {
	r.pageCursor = &pageCursor
	return r
}

func (r ApiListQuoteExecutionsRequest) Execute() (*ListQuoteExecutionsResponse, *http.Response, error) {
	return r.ApiService.ListQuoteExecutionsExecute(r)
}

/*
ListQuoteExecutions List Quote Executions

List quote executions within a particular profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profileId The profile ID associated with the orders.
 @return ApiListQuoteExecutionsRequest
*/
func (a *QuoteExecutionsAPIService) ListQuoteExecutions(ctx context.Context, profileId string) ApiListQuoteExecutionsRequest {
	return ApiListQuoteExecutionsRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
	}
}

// Execute executes the request
//  @return ListQuoteExecutionsResponse
func (a *QuoteExecutionsAPIService) ListQuoteExecutionsExecute(r ApiListQuoteExecutionsRequest) (*ListQuoteExecutionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListQuoteExecutionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteExecutionsAPIService.ListQuoteExecutions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile_id}/quote-executions"
	localVarPath = strings.Replace(localVarPath, "{"+"profile_id"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.side != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "")
	}
	if r.market != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "market", r.market, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.createdAtBegin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at.begin", r.createdAtBegin, "")
	}
	if r.createdAtEnd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at.end", r.createdAtEnd, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_cursor", r.pageCursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
