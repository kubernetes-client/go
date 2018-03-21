# V1beta1ReplicaSetStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | **int32** | The number of available replicas (ready for at least minReadySeconds) for this replica set. | [optional] [default to null]
**Conditions** | [**[]V1beta1ReplicaSetCondition**](v1beta1.ReplicaSetCondition.md) | Represents the latest available observations of a replica set&#39;s current state. | [optional] [default to null]
**FullyLabeledReplicas** | **int32** | The number of pods that have labels matching the labels of the pod template of the replicaset. | [optional] [default to null]
**ObservedGeneration** | **int64** | ObservedGeneration reflects the generation of the most recently observed ReplicaSet. | [optional] [default to null]
**ReadyReplicas** | **int32** | The number of ready replicas for this replica set. | [optional] [default to null]
**Replicas** | **int32** | Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


