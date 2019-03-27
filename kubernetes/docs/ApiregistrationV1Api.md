# \ApiregistrationV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIService**](ApiregistrationV1Api.md#CreateAPIService) | **Post** /apis/apiregistration.k8s.io/v1/apiservices | 
[**DeleteAPIService**](ApiregistrationV1Api.md#DeleteAPIService) | **Delete** /apis/apiregistration.k8s.io/v1/apiservices/{name} | 
[**DeleteCollectionAPIService**](ApiregistrationV1Api.md#DeleteCollectionAPIService) | **Delete** /apis/apiregistration.k8s.io/v1/apiservices | 
[**GetAPIResources**](ApiregistrationV1Api.md#GetAPIResources) | **Get** /apis/apiregistration.k8s.io/v1/ | 
[**ListAPIService**](ApiregistrationV1Api.md#ListAPIService) | **Get** /apis/apiregistration.k8s.io/v1/apiservices | 
[**PatchAPIService**](ApiregistrationV1Api.md#PatchAPIService) | **Patch** /apis/apiregistration.k8s.io/v1/apiservices/{name} | 
[**PatchAPIServiceStatus**](ApiregistrationV1Api.md#PatchAPIServiceStatus) | **Patch** /apis/apiregistration.k8s.io/v1/apiservices/{name}/status | 
[**ReadAPIService**](ApiregistrationV1Api.md#ReadAPIService) | **Get** /apis/apiregistration.k8s.io/v1/apiservices/{name} | 
[**ReadAPIServiceStatus**](ApiregistrationV1Api.md#ReadAPIServiceStatus) | **Get** /apis/apiregistration.k8s.io/v1/apiservices/{name}/status | 
[**ReplaceAPIService**](ApiregistrationV1Api.md#ReplaceAPIService) | **Put** /apis/apiregistration.k8s.io/v1/apiservices/{name} | 
[**ReplaceAPIServiceStatus**](ApiregistrationV1Api.md#ReplaceAPIServiceStatus) | **Put** /apis/apiregistration.k8s.io/v1/apiservices/{name}/status | 


# **CreateAPIService**
> V1ApiService CreateAPIService(ctx, body, optional)


create an APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ApiService**](V1ApiService.md)|  | 
 **optional** | ***CreateAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAPIService**
> V1Status DeleteAPIService(ctx, name, optional)


delete an APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
 **optional** | ***DeleteAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **gracePeriodSeconds** | **optional.Int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **optional.Bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **optional.String**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. | 
 **body** | [**optional.Interface of V1DeleteOptions**](V1DeleteOptions.md)|  | 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionAPIService**
> V1Status DeleteCollectionAPIService(ctx, optional)


delete collection of APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **continue_** | **optional.String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **fieldSelector** | **optional.String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **labelSelector** | **optional.String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **limit** | **optional.Int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **resourceVersion** | **optional.String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **optional.Int32**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **watch** | **optional.Bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAPIResources**
> V1ApiResourceList GetAPIResources(ctx, )


get available resources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1ApiResourceList**](v1.APIResourceList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAPIService**
> V1ApiServiceList ListAPIService(ctx, optional)


list or watch objects of kind APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **continue_** | **optional.String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **fieldSelector** | **optional.String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **labelSelector** | **optional.String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **limit** | **optional.Int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **resourceVersion** | **optional.String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **optional.Int32**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **watch** | **optional.Bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 

### Return type

[**V1ApiServiceList**](v1.APIServiceList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAPIService**
> V1ApiService PatchAPIService(ctx, name, body, optional)


partially update the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAPIServiceStatus**
> V1ApiService PatchAPIServiceStatus(ctx, name, body, optional)


partially update status of the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchAPIServiceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchAPIServiceStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAPIService**
> V1ApiService ReadAPIService(ctx, name, optional)


read the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
 **optional** | ***ReadAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAPIServiceStatus**
> V1ApiService ReadAPIServiceStatus(ctx, name, optional)


read status of the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
 **optional** | ***ReadAPIServiceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAPIServiceStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAPIService**
> V1ApiService ReplaceAPIService(ctx, name, body, optional)


replace the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
  **body** | [**V1ApiService**](V1ApiService.md)|  | 
 **optional** | ***ReplaceAPIServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAPIServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAPIServiceStatus**
> V1ApiService ReplaceAPIServiceStatus(ctx, name, body, optional)


replace status of the specified APIService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the APIService | 
  **body** | [**V1ApiService**](V1ApiService.md)|  | 
 **optional** | ***ReplaceAPIServiceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAPIServiceStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1ApiService**](v1.APIService.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

