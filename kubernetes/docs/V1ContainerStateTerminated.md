# V1ContainerStateTerminated

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerID** | **string** | Container&#39;s ID in the format &#39;docker://&lt;container_id&gt;&#39; | [optional] [default to null]
**ExitCode** | **int32** | Exit status from the last termination of the container | [default to null]
**FinishedAt** | [**time.Time**](time.Time.md) | Time at which the container last terminated | [optional] [default to null]
**Message** | **string** | Message regarding the last termination of the container | [optional] [default to null]
**Reason** | **string** | (brief) reason from the last termination of the container | [optional] [default to null]
**Signal** | **int32** | Signal from the last termination of the container | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | Time at which previous execution of the container started | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


