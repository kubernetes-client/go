# V1ResourceQuotaStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | **map[string]string** | Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] [default to null]
**Used** | **map[string]string** | Used is the current observed total usage of the resource in the namespace. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


