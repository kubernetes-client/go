# V1LimitRangeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default_** | **map[string]string** | Default resource requirement limit value by resource name if resource limit is omitted. | [optional] [default to null]
**DefaultRequest** | **map[string]string** | DefaultRequest is the default resource requirement request value by resource name if resource request is omitted. | [optional] [default to null]
**Max** | **map[string]string** | Max usage constraints on this kind by resource name. | [optional] [default to null]
**MaxLimitRequestRatio** | **map[string]string** | MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource. | [optional] [default to null]
**Min** | **map[string]string** | Min usage constraints on this kind by resource name. | [optional] [default to null]
**Type_** | **string** | Type of resource that this limit applies to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


