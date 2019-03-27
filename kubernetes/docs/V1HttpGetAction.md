# V1HttpGetAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Host name to connect to, defaults to the pod IP. You probably want to set \&quot;Host\&quot; in httpHeaders instead. | [optional] 
**HttpHeaders** | [**[]V1HttpHeader**](v1.HTTPHeader.md) | Custom headers to set in the request. HTTP allows repeated headers. | [optional] 
**Path** | **string** | Path to access on the HTTP server. | [optional] 
**Port** | [**map[string]interface{}**](.md) | Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. | 
**Scheme** | **string** | Scheme to use for connecting to the host. Defaults to HTTP. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


