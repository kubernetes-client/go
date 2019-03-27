# \PolicyV1beta1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#CreateNamespacedPodDisruptionBudget) | **Post** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets | 
[**CreatePodSecurityPolicy**](PolicyV1beta1Api.md#CreatePodSecurityPolicy) | **Post** /apis/policy/v1beta1/podsecuritypolicies | 
[**DeleteCollectionNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#DeleteCollectionNamespacedPodDisruptionBudget) | **Delete** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets | 
[**DeleteCollectionPodSecurityPolicy**](PolicyV1beta1Api.md#DeleteCollectionPodSecurityPolicy) | **Delete** /apis/policy/v1beta1/podsecuritypolicies | 
[**DeleteNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#DeleteNamespacedPodDisruptionBudget) | **Delete** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name} | 
[**DeletePodSecurityPolicy**](PolicyV1beta1Api.md#DeletePodSecurityPolicy) | **Delete** /apis/policy/v1beta1/podsecuritypolicies/{name} | 
[**GetAPIResources**](PolicyV1beta1Api.md#GetAPIResources) | **Get** /apis/policy/v1beta1/ | 
[**ListNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#ListNamespacedPodDisruptionBudget) | **Get** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets | 
[**ListPodDisruptionBudgetForAllNamespaces**](PolicyV1beta1Api.md#ListPodDisruptionBudgetForAllNamespaces) | **Get** /apis/policy/v1beta1/poddisruptionbudgets | 
[**ListPodSecurityPolicy**](PolicyV1beta1Api.md#ListPodSecurityPolicy) | **Get** /apis/policy/v1beta1/podsecuritypolicies | 
[**PatchNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#PatchNamespacedPodDisruptionBudget) | **Patch** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name} | 
[**PatchNamespacedPodDisruptionBudgetStatus**](PolicyV1beta1Api.md#PatchNamespacedPodDisruptionBudgetStatus) | **Patch** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status | 
[**PatchPodSecurityPolicy**](PolicyV1beta1Api.md#PatchPodSecurityPolicy) | **Patch** /apis/policy/v1beta1/podsecuritypolicies/{name} | 
[**ReadNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#ReadNamespacedPodDisruptionBudget) | **Get** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name} | 
[**ReadNamespacedPodDisruptionBudgetStatus**](PolicyV1beta1Api.md#ReadNamespacedPodDisruptionBudgetStatus) | **Get** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status | 
[**ReadPodSecurityPolicy**](PolicyV1beta1Api.md#ReadPodSecurityPolicy) | **Get** /apis/policy/v1beta1/podsecuritypolicies/{name} | 
[**ReplaceNamespacedPodDisruptionBudget**](PolicyV1beta1Api.md#ReplaceNamespacedPodDisruptionBudget) | **Put** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name} | 
[**ReplaceNamespacedPodDisruptionBudgetStatus**](PolicyV1beta1Api.md#ReplaceNamespacedPodDisruptionBudgetStatus) | **Put** /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status | 
[**ReplacePodSecurityPolicy**](PolicyV1beta1Api.md#ReplacePodSecurityPolicy) | **Put** /apis/policy/v1beta1/podsecuritypolicies/{name} | 


# **CreateNamespacedPodDisruptionBudget**
> V1beta1PodDisruptionBudget CreateNamespacedPodDisruptionBudget(ctx, namespace, body, optional)


create a PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V1beta1PodDisruptionBudget**](V1beta1PodDisruptionBudget.md)|  | 
 **optional** | ***CreateNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateNamespacedPodDisruptionBudgetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePodSecurityPolicy**
> PolicyV1beta1PodSecurityPolicy CreatePodSecurityPolicy(ctx, body, optional)


create a PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyV1beta1PodSecurityPolicy**](PolicyV1beta1PodSecurityPolicy.md)|  | 
 **optional** | ***CreatePodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePodSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**PolicyV1beta1PodSecurityPolicy**](policy.v1beta1.PodSecurityPolicy.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionNamespacedPodDisruptionBudget**
> V1Status DeleteCollectionNamespacedPodDisruptionBudget(ctx, namespace, optional)


delete collection of PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***DeleteCollectionNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionNamespacedPodDisruptionBudgetOpts struct

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

# **DeleteCollectionPodSecurityPolicy**
> V1Status DeleteCollectionPodSecurityPolicy(ctx, optional)


delete collection of PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionPodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionPodSecurityPolicyOpts struct

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

# **DeleteNamespacedPodDisruptionBudget**
> V1Status DeleteNamespacedPodDisruptionBudget(ctx, name, namespace, optional)


delete a PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***DeleteNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteNamespacedPodDisruptionBudgetOpts struct

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

# **DeletePodSecurityPolicy**
> V1Status DeletePodSecurityPolicy(ctx, name, optional)


delete a PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodSecurityPolicy | 
 **optional** | ***DeletePodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeletePodSecurityPolicyOpts struct

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

# **ListNamespacedPodDisruptionBudget**
> V1beta1PodDisruptionBudgetList ListNamespacedPodDisruptionBudget(ctx, namespace, optional)


list or watch objects of kind PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ListNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNamespacedPodDisruptionBudgetOpts struct

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

[**V1beta1PodDisruptionBudgetList**](v1beta1.PodDisruptionBudgetList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPodDisruptionBudgetForAllNamespaces**
> V1beta1PodDisruptionBudgetList ListPodDisruptionBudgetForAllNamespaces(ctx, optional)


list or watch objects of kind PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPodDisruptionBudgetForAllNamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPodDisruptionBudgetForAllNamespacesOpts struct

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

[**V1beta1PodDisruptionBudgetList**](v1beta1.PodDisruptionBudgetList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPodSecurityPolicy**
> PolicyV1beta1PodSecurityPolicyList ListPodSecurityPolicy(ctx, optional)


list or watch objects of kind PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPodSecurityPolicyOpts struct

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

[**PolicyV1beta1PodSecurityPolicyList**](policy.v1beta1.PodSecurityPolicyList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedPodDisruptionBudget**
> V1beta1PodDisruptionBudget PatchNamespacedPodDisruptionBudget(ctx, name, namespace, body, optional)


partially update the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchNamespacedPodDisruptionBudgetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedPodDisruptionBudgetStatus**
> V1beta1PodDisruptionBudget PatchNamespacedPodDisruptionBudgetStatus(ctx, name, namespace, body, optional)


partially update status of the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchNamespacedPodDisruptionBudgetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchNamespacedPodDisruptionBudgetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPodSecurityPolicy**
> PolicyV1beta1PodSecurityPolicy PatchPodSecurityPolicy(ctx, name, body, optional)


partially update the specified PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodSecurityPolicy | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchPodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPodSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**PolicyV1beta1PodSecurityPolicy**](policy.v1beta1.PodSecurityPolicy.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedPodDisruptionBudget**
> V1beta1PodDisruptionBudget ReadNamespacedPodDisruptionBudget(ctx, name, namespace, optional)


read the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ReadNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadNamespacedPodDisruptionBudgetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedPodDisruptionBudgetStatus**
> V1beta1PodDisruptionBudget ReadNamespacedPodDisruptionBudgetStatus(ctx, name, namespace, optional)


read status of the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **optional** | ***ReadNamespacedPodDisruptionBudgetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadNamespacedPodDisruptionBudgetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPodSecurityPolicy**
> PolicyV1beta1PodSecurityPolicy ReadPodSecurityPolicy(ctx, name, optional)


read the specified PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodSecurityPolicy | 
 **optional** | ***ReadPodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadPodSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**PolicyV1beta1PodSecurityPolicy**](policy.v1beta1.PodSecurityPolicy.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedPodDisruptionBudget**
> V1beta1PodDisruptionBudget ReplaceNamespacedPodDisruptionBudget(ctx, name, namespace, body, optional)


replace the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V1beta1PodDisruptionBudget**](V1beta1PodDisruptionBudget.md)|  | 
 **optional** | ***ReplaceNamespacedPodDisruptionBudgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceNamespacedPodDisruptionBudgetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedPodDisruptionBudgetStatus**
> V1beta1PodDisruptionBudget ReplaceNamespacedPodDisruptionBudgetStatus(ctx, name, namespace, body, optional)


replace status of the specified PodDisruptionBudget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodDisruptionBudget | 
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V1beta1PodDisruptionBudget**](V1beta1PodDisruptionBudget.md)|  | 
 **optional** | ***ReplaceNamespacedPodDisruptionBudgetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceNamespacedPodDisruptionBudgetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1PodDisruptionBudget**](v1beta1.PodDisruptionBudget.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplacePodSecurityPolicy**
> PolicyV1beta1PodSecurityPolicy ReplacePodSecurityPolicy(ctx, name, body, optional)


replace the specified PodSecurityPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the PodSecurityPolicy | 
  **body** | [**PolicyV1beta1PodSecurityPolicy**](PolicyV1beta1PodSecurityPolicy.md)|  | 
 **optional** | ***ReplacePodSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplacePodSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**PolicyV1beta1PodSecurityPolicy**](policy.v1beta1.PodSecurityPolicy.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

