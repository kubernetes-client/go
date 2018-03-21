# V1beta1StatefulSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [***V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] [default to null]
**Spec** | [***V1beta1StatefulSetSpec**](v1beta1.StatefulSetSpec.md) | Spec defines the desired identities of pods in this set. | [optional] [default to null]
**Status** | [***V1beta1StatefulSetStatus**](v1beta1.StatefulSetStatus.md) | Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


