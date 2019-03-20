# V1ReplicationControllerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | **int32** | The number of available replicas (ready for at least minReadySeconds) for this replication controller. | [optional] 
**Conditions** | [**[]V1ReplicationControllerCondition**](v1.ReplicationControllerCondition.md) | Represents the latest available observations of a replication controller&#39;s current state. | [optional] 
**FullyLabeledReplicas** | **int32** | The number of pods that have labels matching the labels of the pod template of the replication controller. | [optional] 
**ObservedGeneration** | **int64** | ObservedGeneration reflects the generation of the most recently observed replication controller. | [optional] 
**ReadyReplicas** | **int32** | The number of ready replicas for this replication controller. | [optional] 
**Replicas** | **int32** | Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


