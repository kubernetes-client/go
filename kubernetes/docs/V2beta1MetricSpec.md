# V2beta1MetricSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | [**V2beta1ExternalMetricSource**](v2beta1.ExternalMetricSource.md) |  | [optional] 
**Object** | [**V2beta1ObjectMetricSource**](v2beta1.ObjectMetricSource.md) |  | [optional] 
**Pods** | [**V2beta1PodsMetricSource**](v2beta1.PodsMetricSource.md) |  | [optional] 
**Resource** | [**V2beta1ResourceMetricSource**](v2beta1.ResourceMetricSource.md) |  | [optional] 
**Type** | **string** | type is the type of metric source.  It should be one of \&quot;Object\&quot;, \&quot;Pods\&quot; or \&quot;Resource\&quot;, each mapping to a matching field in the object. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


