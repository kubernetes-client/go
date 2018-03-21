# ExtensionsV1beta1SeLinuxStrategyOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** | type is the strategy that will dictate the allowable labels that may be set. | [default to null]
**SeLinuxOptions** | [***V1SeLinuxOptions**](v1.SELinuxOptions.md) | seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


