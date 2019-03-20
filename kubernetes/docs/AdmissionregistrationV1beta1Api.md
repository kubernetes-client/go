# \AdmissionregistrationV1beta1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#CreateMutatingWebhookConfiguration) | **Post** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations | 
[**CreateValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#CreateValidatingWebhookConfiguration) | **Post** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations | 
[**DeleteCollectionMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#DeleteCollectionMutatingWebhookConfiguration) | **Delete** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations | 
[**DeleteCollectionValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#DeleteCollectionValidatingWebhookConfiguration) | **Delete** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations | 
[**DeleteMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#DeleteMutatingWebhookConfiguration) | **Delete** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name} | 
[**DeleteValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#DeleteValidatingWebhookConfiguration) | **Delete** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name} | 
[**GetAPIResources**](AdmissionregistrationV1beta1Api.md#GetAPIResources) | **Get** /apis/admissionregistration.k8s.io/v1beta1/ | 
[**ListMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ListMutatingWebhookConfiguration) | **Get** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations | 
[**ListValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ListValidatingWebhookConfiguration) | **Get** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations | 
[**PatchMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#PatchMutatingWebhookConfiguration) | **Patch** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name} | 
[**PatchValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#PatchValidatingWebhookConfiguration) | **Patch** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name} | 
[**ReadMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ReadMutatingWebhookConfiguration) | **Get** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name} | 
[**ReadValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ReadValidatingWebhookConfiguration) | **Get** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name} | 
[**ReplaceMutatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ReplaceMutatingWebhookConfiguration) | **Put** /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name} | 
[**ReplaceValidatingWebhookConfiguration**](AdmissionregistrationV1beta1Api.md#ReplaceValidatingWebhookConfiguration) | **Put** /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name} | 


# **CreateMutatingWebhookConfiguration**
> V1beta1MutatingWebhookConfiguration CreateMutatingWebhookConfiguration(ctx, body, optional)


create a MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1beta1MutatingWebhookConfiguration**](V1beta1MutatingWebhookConfiguration.md)|  | 
 **optional** | ***CreateMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateMutatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1MutatingWebhookConfiguration**](v1beta1.MutatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateValidatingWebhookConfiguration**
> V1beta1ValidatingWebhookConfiguration CreateValidatingWebhookConfiguration(ctx, body, optional)


create a ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1beta1ValidatingWebhookConfiguration**](V1beta1ValidatingWebhookConfiguration.md)|  | 
 **optional** | ***CreateValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateValidatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1ValidatingWebhookConfiguration**](v1beta1.ValidatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionMutatingWebhookConfiguration**
> V1Status DeleteCollectionMutatingWebhookConfiguration(ctx, optional)


delete collection of MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionMutatingWebhookConfigurationOpts struct

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

# **DeleteCollectionValidatingWebhookConfiguration**
> V1Status DeleteCollectionValidatingWebhookConfiguration(ctx, optional)


delete collection of ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionValidatingWebhookConfigurationOpts struct

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

# **DeleteMutatingWebhookConfiguration**
> V1Status DeleteMutatingWebhookConfiguration(ctx, name, optional)


delete a MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the MutatingWebhookConfiguration | 
 **optional** | ***DeleteMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteMutatingWebhookConfigurationOpts struct

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

# **DeleteValidatingWebhookConfiguration**
> V1Status DeleteValidatingWebhookConfiguration(ctx, name, optional)


delete a ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the ValidatingWebhookConfiguration | 
 **optional** | ***DeleteValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteValidatingWebhookConfigurationOpts struct

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

# **ListMutatingWebhookConfiguration**
> V1beta1MutatingWebhookConfigurationList ListMutatingWebhookConfiguration(ctx, optional)


list or watch objects of kind MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMutatingWebhookConfigurationOpts struct

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

[**V1beta1MutatingWebhookConfigurationList**](v1beta1.MutatingWebhookConfigurationList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListValidatingWebhookConfiguration**
> V1beta1ValidatingWebhookConfigurationList ListValidatingWebhookConfiguration(ctx, optional)


list or watch objects of kind ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListValidatingWebhookConfigurationOpts struct

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

[**V1beta1ValidatingWebhookConfigurationList**](v1beta1.ValidatingWebhookConfigurationList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMutatingWebhookConfiguration**
> V1beta1MutatingWebhookConfiguration PatchMutatingWebhookConfiguration(ctx, name, body, optional)


partially update the specified MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the MutatingWebhookConfiguration | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchMutatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1MutatingWebhookConfiguration**](v1beta1.MutatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchValidatingWebhookConfiguration**
> V1beta1ValidatingWebhookConfiguration PatchValidatingWebhookConfiguration(ctx, name, body, optional)


partially update the specified ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the ValidatingWebhookConfiguration | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchValidatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1ValidatingWebhookConfiguration**](v1beta1.ValidatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMutatingWebhookConfiguration**
> V1beta1MutatingWebhookConfiguration ReadMutatingWebhookConfiguration(ctx, name, optional)


read the specified MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the MutatingWebhookConfiguration | 
 **optional** | ***ReadMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadMutatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1beta1MutatingWebhookConfiguration**](v1beta1.MutatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadValidatingWebhookConfiguration**
> V1beta1ValidatingWebhookConfiguration ReadValidatingWebhookConfiguration(ctx, name, optional)


read the specified ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the ValidatingWebhookConfiguration | 
 **optional** | ***ReadValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadValidatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1beta1ValidatingWebhookConfiguration**](v1beta1.ValidatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMutatingWebhookConfiguration**
> V1beta1MutatingWebhookConfiguration ReplaceMutatingWebhookConfiguration(ctx, name, body, optional)


replace the specified MutatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the MutatingWebhookConfiguration | 
  **body** | [**V1beta1MutatingWebhookConfiguration**](V1beta1MutatingWebhookConfiguration.md)|  | 
 **optional** | ***ReplaceMutatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceMutatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1MutatingWebhookConfiguration**](v1beta1.MutatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceValidatingWebhookConfiguration**
> V1beta1ValidatingWebhookConfiguration ReplaceValidatingWebhookConfiguration(ctx, name, body, optional)


replace the specified ValidatingWebhookConfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the ValidatingWebhookConfiguration | 
  **body** | [**V1beta1ValidatingWebhookConfiguration**](V1beta1ValidatingWebhookConfiguration.md)|  | 
 **optional** | ***ReplaceValidatingWebhookConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceValidatingWebhookConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1beta1ValidatingWebhookConfiguration**](v1beta1.ValidatingWebhookConfiguration.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

