# \AuthenticationV1beta1Api

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenReview**](AuthenticationV1beta1Api.md#CreateTokenReview) | **Post** /apis/authentication.k8s.io/v1beta1/tokenreviews | 
[**GetAPIResources**](AuthenticationV1beta1Api.md#GetAPIResources) | **Get** /apis/authentication.k8s.io/v1beta1/ | 


# **CreateTokenReview**
> V1beta1TokenReview CreateTokenReview(ctx, body, optional)


create a TokenReview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**V1beta1TokenReview**](V1beta1TokenReview.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1TokenReview**](V1beta1TokenReview.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1beta1TokenReview**](v1beta1.TokenReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
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

 - **Content-Type**: application/json, application/yaml, application/vnd.kubernetes.protobuf
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

