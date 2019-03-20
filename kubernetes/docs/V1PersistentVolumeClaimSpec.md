# V1PersistentVolumeClaimSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | [**V1TypedLocalObjectReference**](v1.TypedLocalObjectReference.md) |  | [optional] 
**Resources** | [**V1ResourceRequirements**](v1.ResourceRequirements.md) |  | [optional] 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | [optional] 
**StorageClassName** | **string** | Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeMode** | **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. This is a beta feature. | [optional] 
**VolumeName** | **string** | VolumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


