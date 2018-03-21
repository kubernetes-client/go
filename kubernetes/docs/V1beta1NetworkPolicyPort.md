# V1beta1NetworkPolicyPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | [***interface{}**](interface{}.md) | If specified, the port on the given protocol.  This can either be a numerical or named port on a pod.  If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched. | [optional] [default to null]
**Protocol** | **string** | Optional.  The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


