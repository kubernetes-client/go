# V2beta1ObjectMetricStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageValue** | **string** | averageValue is the current value of the average of the metric across all relevant pods (as a quantity) | [optional] 
**CurrentValue** | **string** | currentValue is the current value of the metric (as a quantity). | 
**MetricName** | **string** | metricName is the name of the metric in question. | 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | [optional] 
**Target** | [**V2beta1CrossVersionObjectReference**](v2beta1.CrossVersionObjectReference.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


