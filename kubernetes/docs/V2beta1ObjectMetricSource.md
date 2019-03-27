# V2beta1ObjectMetricSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageValue** | **string** | averageValue is the target value of the average of the metric across all relevant pods (as a quantity) | [optional] 
**MetricName** | **string** | metricName is the name of the metric in question. | 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | [optional] 
**Target** | [**V2beta1CrossVersionObjectReference**](v2beta1.CrossVersionObjectReference.md) |  | 
**TargetValue** | **string** | targetValue is the target value of the metric (as a quantity). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


