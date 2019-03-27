# \CustomObjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterCustomObject**](CustomObjectsApi.md#CreateClusterCustomObject) | **Post** /apis/{group}/{version}/{plural} | 
[**CreateNamespacedCustomObject**](CustomObjectsApi.md#CreateNamespacedCustomObject) | **Post** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**DeleteClusterCustomObject**](CustomObjectsApi.md#DeleteClusterCustomObject) | **Delete** /apis/{group}/{version}/{plural}/{name} | 
[**DeleteNamespacedCustomObject**](CustomObjectsApi.md#DeleteNamespacedCustomObject) | **Delete** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**GetClusterCustomObject**](CustomObjectsApi.md#GetClusterCustomObject) | **Get** /apis/{group}/{version}/{plural}/{name} | 
[**GetClusterCustomObjectScale**](CustomObjectsApi.md#GetClusterCustomObjectScale) | **Get** /apis/{group}/{version}/{plural}/{name}/scale | 
[**GetClusterCustomObjectStatus**](CustomObjectsApi.md#GetClusterCustomObjectStatus) | **Get** /apis/{group}/{version}/{plural}/{name}/status | 
[**GetNamespacedCustomObject**](CustomObjectsApi.md#GetNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**GetNamespacedCustomObjectScale**](CustomObjectsApi.md#GetNamespacedCustomObjectScale) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/scale | 
[**GetNamespacedCustomObjectStatus**](CustomObjectsApi.md#GetNamespacedCustomObjectStatus) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/status | 
[**ListClusterCustomObject**](CustomObjectsApi.md#ListClusterCustomObject) | **Get** /apis/{group}/{version}/{plural} | 
[**ListNamespacedCustomObject**](CustomObjectsApi.md#ListNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**PatchClusterCustomObject**](CustomObjectsApi.md#PatchClusterCustomObject) | **Patch** /apis/{group}/{version}/{plural}/{name} | 
[**PatchClusterCustomObjectScale**](CustomObjectsApi.md#PatchClusterCustomObjectScale) | **Patch** /apis/{group}/{version}/{plural}/{name}/scale | 
[**PatchClusterCustomObjectStatus**](CustomObjectsApi.md#PatchClusterCustomObjectStatus) | **Patch** /apis/{group}/{version}/{plural}/{name}/status | 
[**PatchNamespacedCustomObject**](CustomObjectsApi.md#PatchNamespacedCustomObject) | **Patch** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**PatchNamespacedCustomObjectScale**](CustomObjectsApi.md#PatchNamespacedCustomObjectScale) | **Patch** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/scale | 
[**PatchNamespacedCustomObjectStatus**](CustomObjectsApi.md#PatchNamespacedCustomObjectStatus) | **Patch** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/status | 
[**ReplaceClusterCustomObject**](CustomObjectsApi.md#ReplaceClusterCustomObject) | **Put** /apis/{group}/{version}/{plural}/{name} | 
[**ReplaceClusterCustomObjectScale**](CustomObjectsApi.md#ReplaceClusterCustomObjectScale) | **Put** /apis/{group}/{version}/{plural}/{name}/scale | 
[**ReplaceClusterCustomObjectStatus**](CustomObjectsApi.md#ReplaceClusterCustomObjectStatus) | **Put** /apis/{group}/{version}/{plural}/{name}/status | 
[**ReplaceNamespacedCustomObject**](CustomObjectsApi.md#ReplaceNamespacedCustomObject) | **Put** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**ReplaceNamespacedCustomObjectScale**](CustomObjectsApi.md#ReplaceNamespacedCustomObjectScale) | **Put** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/scale | 
[**ReplaceNamespacedCustomObjectStatus**](CustomObjectsApi.md#ReplaceNamespacedCustomObjectStatus) | **Put** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name}/status | 


# **CreateClusterCustomObject**
> map[string]interface{} CreateClusterCustomObject(ctx, group, version, plural, body, optional)


Creates a cluster scoped Custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to create. | 
 **optional** | ***CreateClusterCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateClusterCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespacedCustomObject**
> map[string]interface{} CreateNamespacedCustomObject(ctx, group, version, namespace, plural, body, optional)


Creates a namespace scoped Custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to create. | 
 **optional** | ***CreateNamespacedCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateNamespacedCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterCustomObject**
> map[string]interface{} DeleteClusterCustomObject(ctx, group, version, plural, name, body, optional)


Deletes the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **optional** | ***DeleteClusterCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteClusterCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **gracePeriodSeconds** | **optional.Int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **optional.Bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **optional.String**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespacedCustomObject**
> map[string]interface{} DeleteNamespacedCustomObject(ctx, group, version, namespace, plural, name, body, optional)


Deletes the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **optional** | ***DeleteNamespacedCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteNamespacedCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **gracePeriodSeconds** | **optional.Int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **optional.Bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **optional.String**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCustomObject**
> map[string]interface{} GetClusterCustomObject(ctx, group, version, plural, name)


Returns a cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCustomObjectScale**
> map[string]interface{} GetClusterCustomObjectScale(ctx, group, version, plural, name)


read scale of the specified custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCustomObjectStatus**
> map[string]interface{} GetClusterCustomObjectStatus(ctx, group, version, plural, name)


read status of the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacedCustomObject**
> map[string]interface{} GetNamespacedCustomObject(ctx, group, version, namespace, plural, name)


Returns a namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacedCustomObjectScale**
> map[string]interface{} GetNamespacedCustomObjectScale(ctx, group, version, namespace, plural, name)


read scale of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacedCustomObjectStatus**
> map[string]interface{} GetNamespacedCustomObjectStatus(ctx, group, version, namespace, plural, name)


read status of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterCustomObject**
> map[string]interface{} ListClusterCustomObject(ctx, group, version, plural, optional)


list or watch cluster scoped custom objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **optional** | ***ListClusterCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListClusterCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **fieldSelector** | **optional.String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **labelSelector** | **optional.String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **resourceVersion** | **optional.String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **optional.Int32**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **watch** | **optional.Bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedCustomObject**
> map[string]interface{} ListNamespacedCustomObject(ctx, group, version, namespace, plural, optional)


list or watch namespace scoped custom objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **optional** | ***ListNamespacedCustomObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNamespacedCustomObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 
 **fieldSelector** | **optional.String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **labelSelector** | **optional.String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **resourceVersion** | **optional.String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **optional.Int32**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **watch** | **optional.Bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterCustomObject**
> map[string]interface{} PatchClusterCustomObject(ctx, group, version, plural, name, body)


patch the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to patch. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterCustomObjectScale**
> map[string]interface{} PatchClusterCustomObjectScale(ctx, group, version, plural, name, body)


partially update scale of the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterCustomObjectStatus**
> map[string]interface{} PatchClusterCustomObjectStatus(ctx, group, version, plural, name, body)


partially update status of the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedCustomObject**
> map[string]interface{} PatchNamespacedCustomObject(ctx, group, version, namespace, plural, name, body)


patch the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to patch. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedCustomObjectScale**
> map[string]interface{} PatchNamespacedCustomObjectScale(ctx, group, version, namespace, plural, name, body)


partially update scale of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedCustomObjectStatus**
> map[string]interface{} PatchNamespacedCustomObjectStatus(ctx, group, version, namespace, plural, name, body)


partially update status of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterCustomObject**
> map[string]interface{} ReplaceClusterCustomObject(ctx, group, version, plural, name, body)


replace the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to replace. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterCustomObjectScale**
> map[string]interface{} ReplaceClusterCustomObjectScale(ctx, group, version, plural, name, body)


replace scale of the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterCustomObjectStatus**
> map[string]interface{} ReplaceClusterCustomObjectStatus(ctx, group, version, plural, name, body)


replace status of the cluster scoped specified custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedCustomObject**
> map[string]interface{} ReplaceNamespacedCustomObject(ctx, group, version, namespace, plural, name, body)


replace the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The JSON schema of the Resource to replace. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedCustomObjectScale**
> map[string]interface{} ReplaceNamespacedCustomObjectScale(ctx, group, version, namespace, plural, name, body)


replace scale of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedCustomObjectStatus**
> map[string]interface{} ReplaceNamespacedCustomObjectStatus(ctx, group, version, namespace, plural, name, body)


replace status of the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

