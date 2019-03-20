# V1JobCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | [**time.Time**](time.Time.md) | Last time the condition was checked. | [optional] 
**LastTransitionTime** | [**time.Time**](time.Time.md) | Last time the condition transit from one status to another. | [optional] 
**Message** | **string** | Human readable message indicating details about last transition. | [optional] 
**Reason** | **string** | (brief) reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of job condition, Complete or Failed. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


