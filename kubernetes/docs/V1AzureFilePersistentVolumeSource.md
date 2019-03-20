# V1AzureFilePersistentVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | **bool** | Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretName** | **string** | the name of secret that contains Azure Storage Account Name and Key | 
**SecretNamespace** | **string** | the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod | [optional] 
**ShareName** | **string** | Share Name | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


