# V2beta1ResourceMetricStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAverageUtilization** | **int32** | currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if &#x60;targetAverageValue&#x60; was set in the corresponding metric specification. | [optional] [default to null]
**CurrentAverageValue** | **string** | currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the \&quot;pods\&quot; metric source type. It will always be set, regardless of the corresponding metric specification. | [default to null]
**Name** | **string** | name is the name of the resource in question. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


