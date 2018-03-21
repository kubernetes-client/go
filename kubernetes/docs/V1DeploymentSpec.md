# V1DeploymentSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] [default to null]
**Paused** | **bool** | Indicates that the deployment is paused. | [optional] [default to null]
**ProgressDeadlineSeconds** | **int32** | The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s. | [optional] [default to null]
**Replicas** | **int32** | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. | [optional] [default to null]
**RevisionHistoryLimit** | **int32** | The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] [default to null]
**Selector** | [***V1LabelSelector**](v1.LabelSelector.md) | Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template&#39;s labels. | [default to null]
**Strategy** | [***V1DeploymentStrategy**](v1.DeploymentStrategy.md) | The deployment strategy to use to replace existing pods with new ones. | [optional] [default to null]
**Template** | [***V1PodTemplateSpec**](v1.PodTemplateSpec.md) | Template describes the pods that will be created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


