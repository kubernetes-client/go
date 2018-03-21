# V1DownwardApiVolumeFile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldRef** | [***V1ObjectFieldSelector**](v1.ObjectFieldSelector.md) | Required: Selects a field of the pod: only annotations, labels, name and namespace are supported. | [optional] [default to null]
**Mode** | **int32** | Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] [default to null]
**Path** | **string** | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the &#39;..&#39; path. Must be utf-8 encoded. The first item of the relative path must not start with &#39;..&#39; | [default to null]
**ResourceFieldRef** | [***V1ResourceFieldSelector**](v1.ResourceFieldSelector.md) | Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


