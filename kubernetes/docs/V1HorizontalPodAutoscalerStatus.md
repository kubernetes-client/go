# V1HorizontalPodAutoscalerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCPUUtilizationPercentage** | **int32** | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU. | [optional] [default to null]
**CurrentReplicas** | **int32** | current number of replicas of pods managed by this autoscaler. | [default to null]
**DesiredReplicas** | **int32** | desired number of replicas of pods managed by this autoscaler. | [default to null]
**LastScaleTime** | [**time.Time**](time.Time.md) | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed. | [optional] [default to null]
**ObservedGeneration** | **int64** | most recent generation observed by this autoscaler. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


