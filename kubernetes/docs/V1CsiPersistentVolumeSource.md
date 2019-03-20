# V1CsiPersistentVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerPublishSecretRef** | [**V1SecretReference**](v1.SecretReference.md) |  | [optional] 
**Driver** | **string** | Driver is the name of the driver to use for this volume. Required. | 
**FsType** | **string** | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. | [optional] 
**NodePublishSecretRef** | [**V1SecretReference**](v1.SecretReference.md) |  | [optional] 
**NodeStageSecretRef** | [**V1SecretReference**](v1.SecretReference.md) |  | [optional] 
**ReadOnly** | **bool** | Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write). | [optional] 
**VolumeAttributes** | **map[string]string** | Attributes of the volume to publish. | [optional] 
**VolumeHandle** | **string** | VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


