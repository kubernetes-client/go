# V1PersistentVolumeClaimCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | [**time.Time**](time.Time.md) | Last time we probed the condition. | [optional] 
**LastTransitionTime** | [**time.Time**](time.Time.md) | Last time the condition transitioned from one status to another. | [optional] 
**Message** | **string** | Human-readable message indicating details about last transition. | [optional] 
**Reason** | **string** | Unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports \&quot;ResizeStarted\&quot; that means the underlying persistent volume is being resized. | [optional] 
**Status** | **string** |  | 
**Type** | **string** |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


