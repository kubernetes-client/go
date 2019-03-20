# V1NodeCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatTime** | [**time.Time**](time.Time.md) | Last time we got an update on a given condition. | [optional] 
**LastTransitionTime** | [**time.Time**](time.Time.md) | Last time the condition transit from one status to another. | [optional] 
**Message** | **string** | Human readable message indicating details about last transition. | [optional] 
**Reason** | **string** | (brief) reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of node condition. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


