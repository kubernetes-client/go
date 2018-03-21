# V1JobCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | [**time.Time**](time.Time.md) | Last time the condition was checked. | [optional] [default to null]
**LastTransitionTime** | [**time.Time**](time.Time.md) | Last time the condition transit from one status to another. | [optional] [default to null]
**Message** | **string** | Human readable message indicating details about last transition. | [optional] [default to null]
**Reason** | **string** | (brief) reason for the condition&#39;s last transition. | [optional] [default to null]
**Status** | **string** | Status of the condition, one of True, False, Unknown. | [default to null]
**Type_** | **string** | Type of job condition, Complete or Failed. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


