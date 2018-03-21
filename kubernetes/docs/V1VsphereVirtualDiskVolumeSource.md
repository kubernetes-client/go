# V1VsphereVirtualDiskVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. | [optional] [default to null]
**StoragePolicyID** | **string** | Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. | [optional] [default to null]
**StoragePolicyName** | **string** | Storage Policy Based Management (SPBM) profile name. | [optional] [default to null]
**VolumePath** | **string** | Path that identifies vSphere volume vmdk | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


