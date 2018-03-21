# V1alpha1PodPresetSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) | Env defines the collection of EnvVar to inject into containers. | [optional] [default to null]
**EnvFrom** | [**[]V1EnvFromSource**](v1.EnvFromSource.md) | EnvFrom defines the collection of EnvFromSource to inject into containers. | [optional] [default to null]
**Selector** | [***V1LabelSelector**](v1.LabelSelector.md) | Selector is a label query over a set of resources, in this case pods. Required. | [optional] [default to null]
**VolumeMounts** | [**[]V1VolumeMount**](v1.VolumeMount.md) | VolumeMounts defines the collection of VolumeMount to inject into containers. | [optional] [default to null]
**Volumes** | [**[]V1Volume**](v1.Volume.md) | Volumes defines the collection of Volume to inject into the pod. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


