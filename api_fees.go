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
)


// FeesAPIService FeesAPI service
type FeesAPIService service

type ApiCreateCryptoWithdrawalFeeRequest struct {
	ctx context.Context
	ApiService *FeesAPIService
	createCryptoWithdrawalFeeRequest *CreateCryptoWithdrawalFeeRequest
}

func (r ApiCreateCryptoWithdrawalFeeRequest) CreateCryptoWithdrawalFeeRequest(createCryptoWithdrawalFeeRequest CreateCryptoWithdrawalFeeRequest) ApiCreateCryptoWithdrawalFeeRequest {
	r.createCryptoWithdrawalFeeRequest = &createCryptoWithdrawalFeeRequest
	return r
}

func (r ApiCreateCryptoWithdrawalFeeRequest) Execute() (*CryptoWithdrawalFee, *http.Response, error) {
	return r.ApiService.CreateCryptoWithdrawalFeeExecute(r)
}

/*
CreateCryptoWithdrawalFee Create Crypto Withdrawal Fee

Get a guaranteed fee for the given currency, valid for a period of time.
Specify exactly one of `amount` or `total`, otherwise an error is returned.

The [Create Crypto Withdrawal](#operation/CreateCryptoWithdrawal) request that uses the guaranteed fee endpoint
must specify an `amount` less than or equal to the guaranteed fee `amount`, otherwise the withdrawal is rejected.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCryptoWithdrawalFeeRequest
*/
func (a *FeesAPIService) CreateCryptoWithdrawalFee(ctx context.Context) ApiCreateCryptoWithdrawalFeeRequest {
	return ApiCreateCryptoWithdrawalFeeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CryptoWithdrawalFee
func (a *FeesAPIService) CreateCryptoWithdrawalFeeExecute(r ApiCreateCryptoWithdrawalFeeRequest) (*CryptoWithdrawalFee, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CryptoWithdrawalFee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.CreateCryptoWithdrawalFee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transfer/fees/crypto-withdrawal"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCryptoWithdrawalFeeRequest == nil {
		return localVarReturnValue, nil, reportError("createCryptoWithdrawalFeeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createCryptoWithdrawalFeeRequest
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
