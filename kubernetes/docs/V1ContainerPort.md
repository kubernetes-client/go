# V1ContainerPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerPort** | **int32** | Number of port to expose on the pod&#39;s IP address. This must be a valid port number, 0 &lt; x &lt; 65536. | [default to null]
**HostIP** | **string** | What host IP to bind the external port to. | [optional] [default to null]
**HostPort** | **int32** | Number of port to expose on the host. If specified, this must be a valid port number, 0 &lt; x &lt; 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. | [optional] [default to null]
**Name** | **string** | If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services. | [optional] [default to null]
**Protocol** | **string** | Protocol for port. Must be UDP or TCP. Defaults to \&quot;TCP\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


