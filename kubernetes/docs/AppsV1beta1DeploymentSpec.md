# AppsV1beta1DeploymentSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Paused** | **bool** | Indicates that the deployment is paused. | [optional] 
**ProgressDeadlineSeconds** | **int32** | The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s. | [optional] 
**Replicas** | **int32** | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. | [optional] 
**RevisionHistoryLimit** | **int32** | The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 2. | [optional] 
**RollbackTo** | [**AppsV1beta1RollbackConfig**](apps.v1beta1.RollbackConfig.md) |  | [optional] 
**Selector** | [**V1LabelSelector**](v1.LabelSelector.md) |  | [optional] 
**Strategy** | [**AppsV1beta1DeploymentStrategy**](apps.v1beta1.DeploymentStrategy.md) |  | [optional] 
**Template** | [**V1PodTemplateSpec**](v1.PodTemplateSpec.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


