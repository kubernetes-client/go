# V1FlexPersistentVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | Driver is the name of the driver to use for this volume. | [default to null]
**FsType** | **string** | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. The default filesystem depends on FlexVolume script. | [optional] [default to null]
**Options** | **map[string]string** | Optional: Extra command options if any. | [optional] [default to null]
**ReadOnly** | **bool** | Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] [default to null]
**SecretRef** | [***V1SecretReference**](v1.SecretReference.md) | Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


