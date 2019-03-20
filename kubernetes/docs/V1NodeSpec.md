# V1NodeSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSource** | [**V1NodeConfigSource**](v1.NodeConfigSource.md) |  | [optional] 
**ExternalID** | **string** | Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966 | [optional] 
**PodCIDR** | **string** | PodCIDR represents the pod IP range assigned to the node. | [optional] 
**ProviderID** | **string** | ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt; | [optional] 
**Taints** | [**[]V1Taint**](v1.Taint.md) | If specified, the node&#39;s taints. | [optional] 
**Unschedulable** | **bool** | Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


