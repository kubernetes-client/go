# \AuthorizationV1beta1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedLocalSubjectAccessReview**](AuthorizationV1beta1Api.md#CreateNamespacedLocalSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/namespaces/{namespace}/localsubjectaccessreviews | 
[**CreateSelfSubjectAccessReview**](AuthorizationV1beta1Api.md#CreateSelfSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews | 
[**CreateSelfSubjectRulesReview**](AuthorizationV1beta1Api.md#CreateSelfSubjectRulesReview) | **Post** /apis/authorization.k8s.io/v1beta1/selfsubjectrulesreviews | 
[**CreateSubjectAccessReview**](AuthorizationV1beta1Api.md#CreateSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/subjectaccessreviews | 
[**GetAPIResources**](AuthorizationV1beta1Api.md#GetAPIResources) | **Get** /apis/authorization.k8s.io/v1beta1/ | 


# **CreateNamespacedLocalSubjectAccessReview**
> V1beta1LocalSubjectAccessReview CreateNamespacedLocalSubjectAccessReview(ctx, namespace, body, optional)


create a LocalSubjectAccessReview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| object name and auth scope, such as for teams and projects | 
  **body** | [**V1beta1LocalSubjectAccessReview**](V1beta1LocalSubjectAccessReview.md)|  | 
 **optional** | ***CreateNamespacedLocalSubjectAccessReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateNamespacedLocalSubjectAccessReviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **includeUninitialized** | **optional.Bool**| If IncludeUninitialized is specified, the object may be returned without completing initialization. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1LocalSubjectAccessReview**](v1beta1.LocalSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSelfSubjectAccessReview**
> V1beta1SelfSubjectAccessReview CreateSelfSubjectAccessReview(ctx, body, optional)


create a SelfSubjectAccessReview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1beta1SelfSubjectAccessReview**](V1beta1SelfSubjectAccessReview.md)|  | 
 **optional** | ***CreateSelfSubjectAccessReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSelfSubjectAccessReviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **includeUninitialized** | **optional.Bool**| If IncludeUninitialized is specified, the object may be returned without completing initialization. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1SelfSubjectAccessReview**](v1beta1.SelfSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSelfSubjectRulesReview**
> V1beta1SelfSubjectRulesReview CreateSelfSubjectRulesReview(ctx, body, optional)


create a SelfSubjectRulesReview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1beta1SelfSubjectRulesReview**](V1beta1SelfSubjectRulesReview.md)|  | 
 **optional** | ***CreateSelfSubjectRulesReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSelfSubjectRulesReviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **includeUninitialized** | **optional.Bool**| If IncludeUninitialized is specified, the object may be returned without completing initialization. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1SelfSubjectRulesReview**](v1beta1.SelfSubjectRulesReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubjectAccessReview**
> V1beta1SubjectAccessReview CreateSubjectAccessReview(ctx, body, optional)


create a SubjectAccessReview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1beta1SubjectAccessReview**](V1beta1SubjectAccessReview.md)|  | 
 **optional** | ***CreateSubjectAccessReviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSubjectAccessReviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **includeUninitialized** | **optional.Bool**| If IncludeUninitialized is specified, the object may be returned without completing initialization. | 
 **pretty** | **optional.String**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1SubjectAccessReview**](v1beta1.SubjectAccessReview.md)

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

