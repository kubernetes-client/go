# V1HorizontalPodAutoscalerSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReplicas** | **int32** | upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas. | 
**MinReplicas** | **int32** | lower limit for the number of pods that can be set by the autoscaler, default 1. | [optional] 
**ScaleTargetRef** | [**V1CrossVersionObjectReference**](v1.CrossVersionObjectReference.md) |  | 
**TargetCPUUtilizationPercentage** | **int32** | target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


