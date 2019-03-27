# V1alpha1PodPresetSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) | Env defines the collection of EnvVar to inject into containers. | [optional] 
**EnvFrom** | [**[]V1EnvFromSource**](v1.EnvFromSource.md) | EnvFrom defines the collection of EnvFromSource to inject into containers. | [optional] 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | [optional] 
**VolumeMounts** | [**[]V1VolumeMount**](v1.VolumeMount.md) | VolumeMounts defines the collection of VolumeMount to inject into containers. | [optional] 
**Volumes** | [**[]V1Volume**](v1.Volume.md) | Volumes defines the collection of Volume to inject into the pod. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


