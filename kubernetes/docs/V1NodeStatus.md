# V1NodeStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]V1NodeAddress**](v1.NodeAddress.md) | List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses | [optional] 
**Allocatable** | **map[string]string** | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity. | [optional] 
**Capacity** | **map[string]string** | Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Conditions** | [**[]V1NodeCondition**](v1.NodeCondition.md) | Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition | [optional] 
**Config** | [**V1NodeConfigStatus**](v1.NodeConfigStatus.md) |  | [optional] 
**DaemonEndpoints** | [**V1NodeDaemonEndpoints**](v1.NodeDaemonEndpoints.md) |  | [optional] 
**Images** | [**[]V1ContainerImage**](v1.ContainerImage.md) | List of container images on this node | [optional] 
**NodeInfo** | [**V1NodeSystemInfo**](v1.NodeSystemInfo.md) |  | [optional] 
**Phase** | **string** | NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated. | [optional] 
**VolumesAttached** | [**[]V1AttachedVolume**](v1.AttachedVolume.md) | List of volumes that are attached to the node. | [optional] 
**VolumesInUse** | **[]string** | List of attachable volumes in use (mounted) by the node. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


