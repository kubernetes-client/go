# V1beta1VolumeAttachment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [***V1ObjectMeta**](v1.ObjectMeta.md) | Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata | [optional] [default to null]
**Spec** | [***V1beta1VolumeAttachmentSpec**](v1beta1.VolumeAttachmentSpec.md) | Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system. | [default to null]
**Status** | [***V1beta1VolumeAttachmentStatus**](v1beta1.VolumeAttachmentStatus.md) | Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


