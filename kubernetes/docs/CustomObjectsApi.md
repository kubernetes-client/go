# \CustomObjectsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterCustomObject**](CustomObjectsApi.md#CreateClusterCustomObject) | **Post** /apis/{group}/{version}/{plural} | 
[**CreateNamespacedCustomObject**](CustomObjectsApi.md#CreateNamespacedCustomObject) | **Post** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**DeleteClusterCustomObject**](CustomObjectsApi.md#DeleteClusterCustomObject) | **Delete** /apis/{group}/{version}/{plural}/{name} | 
[**DeleteNamespacedCustomObject**](CustomObjectsApi.md#DeleteNamespacedCustomObject) | **Delete** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**GetClusterCustomObject**](CustomObjectsApi.md#GetClusterCustomObject) | **Get** /apis/{group}/{version}/{plural}/{name} | 
[**GetNamespacedCustomObject**](CustomObjectsApi.md#GetNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**ListClusterCustomObject**](CustomObjectsApi.md#ListClusterCustomObject) | **Get** /apis/{group}/{version}/{plural} | 
[**ListNamespacedCustomObject**](CustomObjectsApi.md#ListNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**PatchClusterCustomObject**](CustomObjectsApi.md#PatchClusterCustomObject) | **Patch** /apis/{group}/{version}/{plural}/{name} | 
[**PatchNamespacedCustomObject**](CustomObjectsApi.md#PatchNamespacedCustomObject) | **Patch** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**ReplaceClusterCustomObject**](CustomObjectsApi.md#ReplaceClusterCustomObject) | **Put** /apis/{group}/{version}/{plural}/{name} | 
[**ReplaceNamespacedCustomObject**](CustomObjectsApi.md#ReplaceNamespacedCustomObject) | **Put** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 


# **CreateClusterCustomObject**
> interface{} CreateClusterCustomObject(ctx, group, version, plural, body, optional)


Creates a cluster scoped Custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespacedCustomObject**
> interface{} CreateNamespacedCustomObject(ctx, group, version, namespace, plural, body, optional)


Creates a namespace scoped Custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterCustomObject**
> interface{} DeleteClusterCustomObject(ctx, group, version, plural, name, body, optional)


Deletes the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespacedCustomObject**
> interface{} DeleteNamespacedCustomObject(ctx, group, version, namespace, plural, name, body, optional)


Deletes the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCustomObject**
> interface{} GetClusterCustomObject(ctx, group, version, plural, name)


Returns a cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacedCustomObject**
> interface{} GetNamespacedCustomObject(ctx, group, version, namespace, plural, name)


Returns a namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterCustomObject**
> interface{} ListClusterCustomObject(ctx, group, version, plural, optional)


list or watch cluster scoped custom objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedCustomObject**
> interface{} ListNamespacedCustomObject(ctx, group, version, namespace, plural, optional)


list or watch namespace scoped custom objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| The custom resource&#39;s group name | 
  **version** | **string**| The custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterCustomObject**
> interface{} PatchClusterCustomObject(ctx, group, version, plural, name, body)


patch the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to patch. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedCustomObject**
> interface{} PatchNamespacedCustomObject(ctx, group, version, namespace, plural, name, body)


patch the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to patch. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterCustomObject**
> interface{} ReplaceClusterCustomObject(ctx, group, version, plural, name, body)


replace the specified cluster scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to replace. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedCustomObject**
> interface{} ReplaceNamespacedCustomObject(ctx, group, version, namespace, plural, name, body)


replace the specified namespace scoped custom object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| the custom resource&#39;s group | 
  **version** | **string**| the custom resource&#39;s version | 
  **namespace** | **string**| The custom resource&#39;s namespace | 
  **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
  **name** | **string**| the custom object&#39;s name | 
  **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to replace. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

