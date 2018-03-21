# V2beta1ExternalMetricStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAverageValue** | **string** | currentAverageValue is the current value of metric averaged over autoscaled pods. | [optional] [default to null]
**CurrentValue** | **string** | currentValue is the current value of the metric (as a quantity) | [default to null]
**MetricName** | **string** | metricName is the name of a metric used for autoscaling in metric system. | [default to null]
**MetricSelector** | [***V1LabelSelector**](v1.LabelSelector.md) | metricSelector is used to identify a specific time series within a given metric. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


