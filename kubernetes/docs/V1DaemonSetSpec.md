# V1DaemonSetSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | **int32** | The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready). | [optional] 
**RevisionHistoryLimit** | **int32** | The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | 
**Template** | [**V1PodTemplateSpec**](v1.PodTemplateSpec.md) |  | 
**UpdateStrategy** | [**V1DaemonSetUpdateStrategy**](v1.DaemonSetUpdateStrategy.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


