# V1HorizontalPodAutoscalerSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReplicas** | **int32** | upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas. | [default to null]
**MinReplicas** | **int32** | lower limit for the number of pods that can be set by the autoscaler, default 1. | [optional] [default to null]
**ScaleTargetRef** | [***V1CrossVersionObjectReference**](v1.CrossVersionObjectReference.md) | reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource. | [default to null]
**TargetCPUUtilizationPercentage** | **int32** | target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


