# V1beta1PodDisruptionBudget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [***V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] [default to null]
**Spec** | [***V1beta1PodDisruptionBudgetSpec**](v1beta1.PodDisruptionBudgetSpec.md) | Specification of the desired behavior of the PodDisruptionBudget. | [optional] [default to null]
**Status** | [***V1beta1PodDisruptionBudgetStatus**](v1beta1.PodDisruptionBudgetStatus.md) | Most recently observed status of the PodDisruptionBudget. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


