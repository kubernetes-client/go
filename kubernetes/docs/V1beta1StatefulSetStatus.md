# V1beta1StatefulSetStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollisionCount** | **int32** | collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. | [optional] [default to null]
**Conditions** | [**[]V1beta1StatefulSetCondition**](v1beta1.StatefulSetCondition.md) | Represents the latest available observations of a statefulset&#39;s current state. | [optional] [default to null]
**CurrentReplicas** | **int32** | currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision. | [optional] [default to null]
**CurrentRevision** | **string** | currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas). | [optional] [default to null]
**ObservedGeneration** | **int64** | observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet&#39;s generation, which is updated on mutation by the API Server. | [optional] [default to null]
**ReadyReplicas** | **int32** | readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition. | [optional] [default to null]
**Replicas** | **int32** | replicas is the number of Pods created by the StatefulSet controller. | [default to null]
**UpdateRevision** | **string** | updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas) | [optional] [default to null]
**UpdatedReplicas** | **int32** | updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


