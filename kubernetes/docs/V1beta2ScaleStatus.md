# V1beta2ScaleStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | **int32** | actual number of observed instances of the scaled object. | [default to null]
**Selector** | **map[string]string** | label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors | [optional] [default to null]
**TargetSelector** | **string** | label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


