# V1EnvVarSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | [***V1ConfigMapKeySelector**](v1.ConfigMapKeySelector.md) | Selects a key of a ConfigMap. | [optional] [default to null]
**FieldRef** | [***V1ObjectFieldSelector**](v1.ObjectFieldSelector.md) | Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP. | [optional] [default to null]
**ResourceFieldRef** | [***V1ResourceFieldSelector**](v1.ResourceFieldSelector.md) | Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. | [optional] [default to null]
**SecretKeyRef** | [***V1SecretKeySelector**](v1.SecretKeySelector.md) | Selects a key of a secret in the pod&#39;s namespace | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


