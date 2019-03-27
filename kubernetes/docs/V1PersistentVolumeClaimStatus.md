# V1PersistentVolumeClaimStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**Capacity** | **map[string]string** | Represents the actual resources of the underlying volume. | [optional] 
**Conditions** | [**[]V1PersistentVolumeClaimCondition**](v1.PersistentVolumeClaimCondition.md) | Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;. | [optional] 
**Phase** | **string** | Phase represents the current phase of PersistentVolumeClaim. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


