/*
Export API

Export OpenAPI definitions

API version: 20.0.0
Contact: apiteam@pleo.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// ExportAPIV0ApiService ExportAPIV0Api service
type ExportAPIV0ApiService service

type ApiCreateExportJobRequest struct {
	ctx             context.Context
	ApiService      *ExportAPIV0ApiService
	createExportJob *CreateExportJob
}

func (r ApiCreateExportJobRequest) CreateExportJob(createExportJob CreateExportJob) ApiCreateExportJobRequest {
	r.createExportJob = &createExportJob
	return r
}

func (r ApiCreateExportJobRequest) Execute() (*DataResponseExportJob, *http.Response, error) {
	return r.ApiService.CreateExportJobExecute(r)
}

/*
CreateExportJob Create a new export job

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateExportJobRequest
*/
func (a *ExportAPIV0ApiService) CreateExportJob(ctx context.Context) ApiCreateExportJobRequest {
	return ApiCreateExportJobRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DataResponseExportJob
func (a *ExportAPIV0ApiService) CreateExportJobExecute(r ApiCreateExportJobRequest) (*DataResponseExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataResponseExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.CreateExportJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-jobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createExportJob
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v DataResponseExportJob
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiCreateExportJobEventRequest struct {
	ctx                  context.Context
	ApiService           *ExportAPIV0ApiService
	createExportJobEvent *CreateExportJobEvent
}

func (r ApiCreateExportJobEventRequest) CreateExportJobEvent(createExportJobEvent CreateExportJobEvent) ApiCreateExportJobEventRequest {
	r.createExportJobEvent = &createExportJobEvent
	return r
}

func (r ApiCreateExportJobEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateExportJobEventExecute(r)
}

/*
CreateExportJobEvent Create an Export Job Event

	           The Export Job Events are used to move the Export Job through the export lifecycle,
	           updating the status and results of the export.
	           This endpoint is used to update the status and failure reasons (if failed) of an export job,
	           i.e. start a job, mark a job as failed, complete a job, cancel a job, etc.

	           Events:

	           - started: Used to start processing an export job. This allows the integration
	                   service to start processing the export job items.
	           - failed: Use this to mark a job as failed when an irrecoverable error happens,
	                   and this could be while processing the export job items.
	           - completed: Used to indicate that all the export job items have been successfully processed.
	           - completed_with_errors: Used when all the export job items have been processed, but some of then failed.
	           - canceled: Used to indicate if the integration service is unable, for any reason,
	                   to continue processing the export job.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateExportJobEventRequest
*/
func (a *ExportAPIV0ApiService) CreateExportJobEvent(ctx context.Context) ApiCreateExportJobEventRequest {
	return ApiCreateExportJobEventRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExportAPIV0ApiService) CreateExportJobEventExecute(r ApiCreateExportJobEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.CreateExportJobEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-job-events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createExportJobEvent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetExportJobByIdRequest struct {
	ctx        context.Context
	ApiService *ExportAPIV0ApiService
	jobId      string
}

func (r ApiGetExportJobByIdRequest) Execute() (*DataResponseExportJob, *http.Response, error) {
	return r.ApiService.GetExportJobByIdExecute(r)
}

/*
GetExportJobById Get an Export Job

Retrieve an export job for a given export job ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jobId
	@return ApiGetExportJobByIdRequest
*/
func (a *ExportAPIV0ApiService) GetExportJobById(ctx context.Context, jobId string) ApiGetExportJobByIdRequest {
	return ApiGetExportJobByIdRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//
//	@return DataResponseExportJob
func (a *ExportAPIV0ApiService) GetExportJobByIdExecute(r ApiGetExportJobByIdRequest) (*DataResponseExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataResponseExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.GetExportJobById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v DataResponseExportJob
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGetExportJobItemsRequest struct {
	ctx          context.Context
	ApiService   *ExportAPIV0ApiService
	jobId        string
	status       *string
	before       *string
	after        *string
	offset       *int64
	limit        *int32
	sortingKeys  *[]string
	sortingOrder *[]PageOrder
}

// Filter by the export job item status
func (r ApiGetExportJobItemsRequest) Status(status string) ApiGetExportJobItemsRequest {
	r.status = &status
	return r
}

// Lower bound of the page of data to return (cannot be used together with [after] or [offset])
func (r ApiGetExportJobItemsRequest) Before(before string) ApiGetExportJobItemsRequest {
	r.before = &before
	return r
}

// Upper bound of the page of data to return (cannot be used together with [before] or [offset])
func (r ApiGetExportJobItemsRequest) After(after string) ApiGetExportJobItemsRequest {
	r.after = &after
	return r
}

// Offset of the page of data to return (cannot be used together with [before] or [after])
func (r ApiGetExportJobItemsRequest) Offset(offset int64) ApiGetExportJobItemsRequest {
	r.offset = &offset
	return r
}

// The maximum amount of items to return
func (r ApiGetExportJobItemsRequest) Limit(limit int32) ApiGetExportJobItemsRequest {
	r.limit = &limit
	return r
}

// The keys to sort the results by
func (r ApiGetExportJobItemsRequest) SortingKeys(sortingKeys []string) ApiGetExportJobItemsRequest {
	r.sortingKeys = &sortingKeys
	return r
}

// The order to sort the results by. Must be the same length as [sortingKeys]; one order per key
func (r ApiGetExportJobItemsRequest) SortingOrder(sortingOrder []PageOrder) ApiGetExportJobItemsRequest {
	r.sortingOrder = &sortingOrder
	return r
}

func (r ApiGetExportJobItemsRequest) Execute() (*CursorPaginatedResponseExportJobItem, *http.Response, error) {
	return r.ApiService.GetExportJobItemsExecute(r)
}

/*
GetExportJobItems Get Export Job Items

Fetch a list of export job items

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jobId
	@return ApiGetExportJobItemsRequest
*/
func (a *ExportAPIV0ApiService) GetExportJobItems(ctx context.Context, jobId string) ApiGetExportJobItemsRequest {
	return ApiGetExportJobItemsRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//
//	@return CursorPaginatedResponseExportJobItem
func (a *ExportAPIV0ApiService) GetExportJobItemsExecute(r ApiGetExportJobItemsRequest) (*CursorPaginatedResponseExportJobItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CursorPaginatedResponseExportJobItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.GetExportJobItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-jobs/{jobId}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortingKeys != nil {
		t := *r.sortingKeys
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sorting_keys", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sorting_keys", parameterToString(t, "multi"))
		}
	}
	if r.sortingOrder != nil {
		t := *r.sortingOrder
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sorting_order", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sorting_order", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v CursorPaginatedResponseExportJobItem
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGetExportJobsListRequest struct {
	ctx          context.Context
	ApiService   *ExportAPIV0ApiService
	companyId    *string
	status       *string
	before       *string
	after        *string
	offset       *int64
	limit        *int32
	sortingKeys  *[]string
	sortingOrder *[]PageOrder
}

// Fetch a list of jobs for a given company
func (r ApiGetExportJobsListRequest) CompanyId(companyId string) ApiGetExportJobsListRequest {
	r.companyId = &companyId
	return r
}

// Fetch a list of jobs for a specific status
func (r ApiGetExportJobsListRequest) Status(status string) ApiGetExportJobsListRequest {
	r.status = &status
	return r
}

// Lower bound of the page of data to return (cannot be used together with [after] or [offset])
func (r ApiGetExportJobsListRequest) Before(before string) ApiGetExportJobsListRequest {
	r.before = &before
	return r
}

// Upper bound of the page of data to return (cannot be used together with [before] or [offset])
func (r ApiGetExportJobsListRequest) After(after string) ApiGetExportJobsListRequest {
	r.after = &after
	return r
}

// Offset of the page of data to return (cannot be used together with [before] or [after])
func (r ApiGetExportJobsListRequest) Offset(offset int64) ApiGetExportJobsListRequest {
	r.offset = &offset
	return r
}

// The maximum amount of items to return
func (r ApiGetExportJobsListRequest) Limit(limit int32) ApiGetExportJobsListRequest {
	r.limit = &limit
	return r
}

// The keys to sort the results by
func (r ApiGetExportJobsListRequest) SortingKeys(sortingKeys []string) ApiGetExportJobsListRequest {
	r.sortingKeys = &sortingKeys
	return r
}

// The order to sort the results by. Must be the same length as [sortingKeys]; one order per key
func (r ApiGetExportJobsListRequest) SortingOrder(sortingOrder []PageOrder) ApiGetExportJobsListRequest {
	r.sortingOrder = &sortingOrder
	return r
}

func (r ApiGetExportJobsListRequest) Execute() (*CursorPaginatedResponseExportJob, *http.Response, error) {
	return r.ApiService.GetExportJobsListExecute(r)
}

/*
GetExportJobsList Get a list of Export Jobs

Fetch a list of export jobs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetExportJobsListRequest
*/
func (a *ExportAPIV0ApiService) GetExportJobsList(ctx context.Context) ApiGetExportJobsListRequest {
	return ApiGetExportJobsListRequest{
		ApiService: a,
		ctx:        ctx,
		companyId:  a.client.cfg.CompanyID,
	}
}

// Execute executes the request
//
//	@return CursorPaginatedResponseExportJob
func (a *ExportAPIV0ApiService) GetExportJobsListExecute(r ApiGetExportJobsListRequest) (*CursorPaginatedResponseExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CursorPaginatedResponseExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.GetExportJobsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-jobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.companyId != nil {
		localVarQueryParams.Add("company_id", parameterToString(*r.companyId, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortingKeys != nil {
		t := *r.sortingKeys
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sorting_keys", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sorting_keys", parameterToString(t, "multi"))
		}
	}
	if r.sortingOrder != nil {
		t := *r.sortingOrder
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sorting_order", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sorting_order", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v CursorPaginatedResponseExportJob
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiUpdateExportJobItemsRequest struct {
	ctx                 context.Context
	ApiService          *ExportAPIV0ApiService
	jobId               string
	updateExportJobItem *[]UpdateExportJobItem
}

func (r ApiUpdateExportJobItemsRequest) UpdateExportJobItem(updateExportJobItem []UpdateExportJobItem) ApiUpdateExportJobItemsRequest {
	r.updateExportJobItem = &updateExportJobItem
	return r
}

func (r ApiUpdateExportJobItemsRequest) Execute() (*ExportJobItemUpdate, *http.Response, error) {
	return r.ApiService.UpdateExportJobItemsExecute(r)
}

/*
UpdateExportJobItems Update Export Job Items

Update the status and other attributes of the export job items in batches of 100. This API only supports updating up to 100 items.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jobId
	@return ApiUpdateExportJobItemsRequest
*/
func (a *ExportAPIV0ApiService) UpdateExportJobItems(ctx context.Context, jobId string) ApiUpdateExportJobItemsRequest {
	return ApiUpdateExportJobItemsRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//
//	@return ExportJobItemUpdate
func (a *ExportAPIV0ApiService) UpdateExportJobItemsExecute(r ApiUpdateExportJobItemsRequest) (*ExportJobItemUpdate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportJobItemUpdate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAPIV0ApiService.UpdateExportJobItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/export-jobs/{jobId}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateExportJobItem
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ExportJobItemUpdate
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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
