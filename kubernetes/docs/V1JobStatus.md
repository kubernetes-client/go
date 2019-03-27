# V1JobStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int32** | The number of actively running pods. | [optional] 
**CompletionTime** | [**time.Time**](time.Time.md) | Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. | [optional] 
**Conditions** | [**[]V1JobCondition**](v1.JobCondition.md) | The latest available observations of an object&#39;s current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/ | [optional] 
**Failed** | **int32** | The number of pods which reached phase Failed. | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. | [optional] 
**Succeeded** | **int32** | The number of pods which reached phase Succeeded. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


