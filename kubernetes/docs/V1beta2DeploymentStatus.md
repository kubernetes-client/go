# V1beta2DeploymentStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | **int32** | Total number of available pods (ready for at least minReadySeconds) targeted by this deployment. | [optional] 
**CollisionCount** | **int32** | Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet. | [optional] 
**Conditions** | [**[]V1beta2DeploymentCondition**](v1beta2.DeploymentCondition.md) | Represents the latest available observations of a deployment&#39;s current state. | [optional] 
**ObservedGeneration** | **int64** | The generation observed by the deployment controller. | [optional] 
**ReadyReplicas** | **int32** | Total number of ready pods targeted by this deployment. | [optional] 
**Replicas** | **int32** | Total number of non-terminated pods targeted by this deployment (their labels match the selector). | [optional] 
**UnavailableReplicas** | **int32** | Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created. | [optional] 
**UpdatedReplicas** | **int32** | Total number of non-terminated pods targeted by this deployment that have the desired template spec. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


