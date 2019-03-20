# \AutoscalingV2beta1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#CreateNamespacedHorizontalPodAutoscaler) | **Post** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers | 
[**DeleteCollectionNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#DeleteCollectionNamespacedHorizontalPodAutoscaler) | **Delete** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers | 
[**DeleteNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#DeleteNamespacedHorizontalPodAutoscaler) | **Delete** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} | 
[**GetAPIResources**](AutoscalingV2beta1Api.md#GetAPIResources) | **Get** /apis/autoscaling/v2beta1/ | 
[**ListHorizontalPodAutoscalerForAllNamespaces**](AutoscalingV2beta1Api.md#ListHorizontalPodAutoscalerForAllNamespaces) | **Get** /apis/autoscaling/v2beta1/horizontalpodautoscalers | 
[**ListNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#ListNamespacedHorizontalPodAutoscaler) | **Get** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers | 
[**PatchNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#PatchNamespacedHorizontalPodAutoscaler) | **Patch** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} | 
[**PatchNamespacedHorizontalPodAutoscalerStatus**](AutoscalingV2beta1Api.md#PatchNamespacedHorizontalPodAutoscalerStatus) | **Patch** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status | 
[**ReadNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#ReadNamespacedHorizontalPodAutoscaler) | **Get** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} | 
[**ReadNamespacedHorizontalPodAutoscalerStatus**](AutoscalingV2beta1Api.md#ReadNamespacedHorizontalPodAutoscalerStatus) | **Get** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status | 
[**ReplaceNamespacedHorizontalPodAutoscaler**](AutoscalingV2beta1Api.md#ReplaceNamespacedHorizontalPodAutoscaler) | **Put** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} | 
[**ReplaceNamespacedHorizontalPodAutoscalerStatus**](AutoscalingV2beta1Api.md#ReplaceNamespacedHorizontalPodAutoscalerStatus) | **Put** /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status | 


# **CreateNamespacedHorizontalPodAutoscaler**
> V2beta1HorizontalPodAutoscaler CreateNamespacedHorizontalPodAutoscaler(ctx, namespace, body, optional)


create a HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V2beta1HorizontalPodAutoscaler**](V2beta1HorizontalPodAutoscaler.md)|  | 
 **optional** | ***CreateNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateNamespacedHorizontalPodAutoscalerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionNamespacedHorizontalPodAutoscaler**
> V1Status DeleteCollectionNamespacedHorizontalPodAutoscaler(ctx, namespace, optional)


delete collection of HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***DeleteCollectionNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionNamespacedHorizontalPodAutoscalerOpts struct

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

# **DeleteNamespacedHorizontalPodAutoscaler**
> V1Status DeleteNamespacedHorizontalPodAutoscaler(ctx, name, namespace, optional)


delete a HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***DeleteNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteNamespacedHorizontalPodAutoscalerOpts struct

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

# **ListHorizontalPodAutoscalerForAllNamespaces**
> V2beta1HorizontalPodAutoscalerList ListHorizontalPodAutoscalerForAllNamespaces(ctx, optional)


list or watch objects of kind HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListHorizontalPodAutoscalerForAllNamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListHorizontalPodAutoscalerForAllNamespacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continue_** | **optional.String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **fieldSelector** | **optional.String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **labelSelector** | **optional.String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **limit** | **optional.Int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **resourceVersion** | **optional.String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **optional.Int32**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **watch** | **optional.Bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 

### Return type

[**V2beta1HorizontalPodAutoscalerList**](v2beta1.HorizontalPodAutoscalerList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedHorizontalPodAutoscaler**
> V2beta1HorizontalPodAutoscalerList ListNamespacedHorizontalPodAutoscaler(ctx, namespace, optional)


list or watch objects of kind HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ListNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNamespacedHorizontalPodAutoscalerOpts struct

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

[**V2beta1HorizontalPodAutoscalerList**](v2beta1.HorizontalPodAutoscalerList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedHorizontalPodAutoscaler**
> V2beta1HorizontalPodAutoscaler PatchNamespacedHorizontalPodAutoscaler(ctx, name, namespace, body, optional)


partially update the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchNamespacedHorizontalPodAutoscalerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedHorizontalPodAutoscalerStatus**
> V2beta1HorizontalPodAutoscaler PatchNamespacedHorizontalPodAutoscalerStatus(ctx, name, namespace, body, optional)


partially update status of the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchNamespacedHorizontalPodAutoscalerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchNamespacedHorizontalPodAutoscalerStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedHorizontalPodAutoscaler**
> V2beta1HorizontalPodAutoscaler ReadNamespacedHorizontalPodAutoscaler(ctx, name, namespace, optional)


read the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ReadNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadNamespacedHorizontalPodAutoscalerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedHorizontalPodAutoscalerStatus**
> V2beta1HorizontalPodAutoscaler ReadNamespacedHorizontalPodAutoscalerStatus(ctx, name, namespace, optional)


read status of the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ReadNamespacedHorizontalPodAutoscalerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadNamespacedHorizontalPodAutoscalerStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedHorizontalPodAutoscaler**
> V2beta1HorizontalPodAutoscaler ReplaceNamespacedHorizontalPodAutoscaler(ctx, name, namespace, body, optional)


replace the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V2beta1HorizontalPodAutoscaler**](V2beta1HorizontalPodAutoscaler.md)|  | 
 **optional** | ***ReplaceNamespacedHorizontalPodAutoscalerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceNamespacedHorizontalPodAutoscalerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedHorizontalPodAutoscalerStatus**
> V2beta1HorizontalPodAutoscaler ReplaceNamespacedHorizontalPodAutoscalerStatus(ctx, name, namespace, body, optional)


replace status of the specified HorizontalPodAutoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the HorizontalPodAutoscaler | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V2beta1HorizontalPodAutoscaler**](V2beta1HorizontalPodAutoscaler.md)|  | 
 **optional** | ***ReplaceNamespacedHorizontalPodAutoscalerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceNamespacedHorizontalPodAutoscalerStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V2beta1HorizontalPodAutoscaler**](v2beta1.HorizontalPodAutoscaler.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

