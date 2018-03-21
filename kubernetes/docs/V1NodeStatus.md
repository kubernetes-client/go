# V1NodeStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]V1NodeAddress**](v1.NodeAddress.md) | List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses | [optional] [default to null]
**Allocatable** | **map[string]string** | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity. | [optional] [default to null]
**Capacity** | **map[string]string** | Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] [default to null]
**Conditions** | [**[]V1NodeCondition**](v1.NodeCondition.md) | Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition | [optional] [default to null]
**DaemonEndpoints** | [***V1NodeDaemonEndpoints**](v1.NodeDaemonEndpoints.md) | Endpoints of daemons running on the Node. | [optional] [default to null]
**Images** | [**[]V1ContainerImage**](v1.ContainerImage.md) | List of container images on this node | [optional] [default to null]
**NodeInfo** | [***V1NodeSystemInfo**](v1.NodeSystemInfo.md) | Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info | [optional] [default to null]
**Phase** | **string** | NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated. | [optional] [default to null]
**VolumesAttached** | [**[]V1AttachedVolume**](v1.AttachedVolume.md) | List of volumes that are attached to the node. | [optional] [default to null]
**VolumesInUse** | **[]string** | List of attachable volumes in use (mounted) by the node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


