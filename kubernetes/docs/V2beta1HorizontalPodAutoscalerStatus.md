# V2beta1HorizontalPodAutoscalerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]V2beta1HorizontalPodAutoscalerCondition**](v2beta1.HorizontalPodAutoscalerCondition.md) | conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met. | 
**CurrentMetrics** | [**[]V2beta1MetricStatus**](v2beta1.MetricStatus.md) | currentMetrics is the last read state of the metrics used by this autoscaler. | [optional] 
**CurrentReplicas** | **int32** | currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler. | 
**DesiredReplicas** | **int32** | desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler. | 
**LastScaleTime** | [**time.Time**](time.Time.md) | lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed. | [optional] 
**ObservedGeneration** | **int64** | observedGeneration is the most recent generation observed by this autoscaler. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


