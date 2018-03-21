# V1alpha1VolumeAttachmentSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attacher** | **string** | Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName(). | [default to null]
**NodeName** | **string** | The node that the volume should be attached to. | [default to null]
**Source** | [***V1alpha1VolumeAttachmentSource**](v1alpha1.VolumeAttachmentSource.md) | Source represents the volume that should be attached. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


