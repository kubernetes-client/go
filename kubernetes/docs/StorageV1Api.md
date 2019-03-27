# \StorageV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorageClass**](StorageV1Api.md#CreateStorageClass) | **Post** /apis/storage.k8s.io/v1/storageclasses | 
[**CreateVolumeAttachment**](StorageV1Api.md#CreateVolumeAttachment) | **Post** /apis/storage.k8s.io/v1/volumeattachments | 
[**DeleteCollectionStorageClass**](StorageV1Api.md#DeleteCollectionStorageClass) | **Delete** /apis/storage.k8s.io/v1/storageclasses | 
[**DeleteCollectionVolumeAttachment**](StorageV1Api.md#DeleteCollectionVolumeAttachment) | **Delete** /apis/storage.k8s.io/v1/volumeattachments | 
[**DeleteStorageClass**](StorageV1Api.md#DeleteStorageClass) | **Delete** /apis/storage.k8s.io/v1/storageclasses/{name} | 
[**DeleteVolumeAttachment**](StorageV1Api.md#DeleteVolumeAttachment) | **Delete** /apis/storage.k8s.io/v1/volumeattachments/{name} | 
[**GetAPIResources**](StorageV1Api.md#GetAPIResources) | **Get** /apis/storage.k8s.io/v1/ | 
[**ListStorageClass**](StorageV1Api.md#ListStorageClass) | **Get** /apis/storage.k8s.io/v1/storageclasses | 
[**ListVolumeAttachment**](StorageV1Api.md#ListVolumeAttachment) | **Get** /apis/storage.k8s.io/v1/volumeattachments | 
[**PatchStorageClass**](StorageV1Api.md#PatchStorageClass) | **Patch** /apis/storage.k8s.io/v1/storageclasses/{name} | 
[**PatchVolumeAttachment**](StorageV1Api.md#PatchVolumeAttachment) | **Patch** /apis/storage.k8s.io/v1/volumeattachments/{name} | 
[**PatchVolumeAttachmentStatus**](StorageV1Api.md#PatchVolumeAttachmentStatus) | **Patch** /apis/storage.k8s.io/v1/volumeattachments/{name}/status | 
[**ReadStorageClass**](StorageV1Api.md#ReadStorageClass) | **Get** /apis/storage.k8s.io/v1/storageclasses/{name} | 
[**ReadVolumeAttachment**](StorageV1Api.md#ReadVolumeAttachment) | **Get** /apis/storage.k8s.io/v1/volumeattachments/{name} | 
[**ReadVolumeAttachmentStatus**](StorageV1Api.md#ReadVolumeAttachmentStatus) | **Get** /apis/storage.k8s.io/v1/volumeattachments/{name}/status | 
[**ReplaceStorageClass**](StorageV1Api.md#ReplaceStorageClass) | **Put** /apis/storage.k8s.io/v1/storageclasses/{name} | 
[**ReplaceVolumeAttachment**](StorageV1Api.md#ReplaceVolumeAttachment) | **Put** /apis/storage.k8s.io/v1/volumeattachments/{name} | 
[**ReplaceVolumeAttachmentStatus**](StorageV1Api.md#ReplaceVolumeAttachmentStatus) | **Put** /apis/storage.k8s.io/v1/volumeattachments/{name}/status | 


# **CreateStorageClass**
> V1StorageClass CreateStorageClass(ctx, body, optional)


create a StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1StorageClass**](V1StorageClass.md)|  | 
 **optional** | ***CreateStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateStorageClassOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1StorageClass**](v1.StorageClass.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVolumeAttachment**
> V1VolumeAttachment CreateVolumeAttachment(ctx, body, optional)


create a VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1VolumeAttachment**](V1VolumeAttachment.md)|  | 
 **optional** | ***CreateVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateVolumeAttachmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUninitialized** | **optional.Bool**| If true, partially initialized resources are included in the response. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionStorageClass**
> V1Status DeleteCollectionStorageClass(ctx, optional)


delete collection of StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionStorageClassOpts struct

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

# **DeleteCollectionVolumeAttachment**
> V1Status DeleteCollectionVolumeAttachment(ctx, optional)


delete collection of VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCollectionVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCollectionVolumeAttachmentOpts struct

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

# **DeleteStorageClass**
> V1Status DeleteStorageClass(ctx, name, optional)


delete a StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the StorageClass | 
 **optional** | ***DeleteStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteStorageClassOpts struct

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

# **DeleteVolumeAttachment**
> V1Status DeleteVolumeAttachment(ctx, name, optional)


delete a VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
 **optional** | ***DeleteVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteVolumeAttachmentOpts struct

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

# **ListStorageClass**
> V1StorageClassList ListStorageClass(ctx, optional)


list or watch objects of kind StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListStorageClassOpts struct

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

[**V1StorageClassList**](v1.StorageClassList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeAttachment**
> V1VolumeAttachmentList ListVolumeAttachment(ctx, optional)


list or watch objects of kind VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVolumeAttachmentOpts struct

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

[**V1VolumeAttachmentList**](v1.VolumeAttachmentList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchStorageClass**
> V1StorageClass PatchStorageClass(ctx, name, body, optional)


partially update the specified StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the StorageClass | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchStorageClassOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1StorageClass**](v1.StorageClass.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchVolumeAttachment**
> V1VolumeAttachment PatchVolumeAttachment(ctx, name, body, optional)


partially update the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchVolumeAttachmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchVolumeAttachmentStatus**
> V1VolumeAttachment PatchVolumeAttachmentStatus(ctx, name, body, optional)


partially update status of the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 
 **optional** | ***PatchVolumeAttachmentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchVolumeAttachmentStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadStorageClass**
> V1StorageClass ReadStorageClass(ctx, name, optional)


read the specified StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the StorageClass | 
 **optional** | ***ReadStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadStorageClassOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1StorageClass**](v1.StorageClass.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVolumeAttachment**
> V1VolumeAttachment ReadVolumeAttachment(ctx, name, optional)


read the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
 **optional** | ***ReadVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadVolumeAttachmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **exact** | **optional.Bool**| Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **optional.Bool**| Should this value be exported.  Export strips fields that a user can not specify. | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVolumeAttachmentStatus**
> V1VolumeAttachment ReadVolumeAttachmentStatus(ctx, name, optional)


read status of the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
 **optional** | ***ReadVolumeAttachmentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadVolumeAttachmentStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStorageClass**
> V1StorageClass ReplaceStorageClass(ctx, name, body, optional)


replace the specified StorageClass

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the StorageClass | 
  **body** | [**V1StorageClass**](V1StorageClass.md)|  | 
 **optional** | ***ReplaceStorageClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceStorageClassOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1StorageClass**](v1.StorageClass.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceVolumeAttachment**
> V1VolumeAttachment ReplaceVolumeAttachment(ctx, name, body, optional)


replace the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
  **body** | [**V1VolumeAttachment**](V1VolumeAttachment.md)|  | 
 **optional** | ***ReplaceVolumeAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceVolumeAttachmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceVolumeAttachmentStatus**
> V1VolumeAttachment ReplaceVolumeAttachmentStatus(ctx, name, body, optional)


replace status of the specified VolumeAttachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the VolumeAttachment | 
  **body** | [**V1VolumeAttachment**](V1VolumeAttachment.md)|  | 
 **optional** | ***ReplaceVolumeAttachmentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceVolumeAttachmentStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 

### Return type

[**V1VolumeAttachment**](v1.VolumeAttachment.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

