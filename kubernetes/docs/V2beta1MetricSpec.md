# V2beta1MetricSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | [***V2beta1ExternalMetricSource**](v2beta1.ExternalMetricSource.md) | external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). | [optional] [default to null]
**Object** | [***V2beta1ObjectMetricSource**](v2beta1.ObjectMetricSource.md) | object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object). | [optional] [default to null]
**Pods** | [***V2beta1PodsMetricSource**](v2beta1.PodsMetricSource.md) | pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value. | [optional] [default to null]
**Resource** | [***V2beta1ResourceMetricSource**](v2beta1.ResourceMetricSource.md) | resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \&quot;pods\&quot; source. | [optional] [default to null]
**Type_** | **string** | type is the type of metric source.  It should be one of \&quot;Object\&quot;, \&quot;Pods\&quot; or \&quot;Resource\&quot;, each mapping to a matching field in the object. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


