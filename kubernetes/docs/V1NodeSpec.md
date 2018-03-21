# V1NodeSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSource** | [***V1NodeConfigSource**](v1.NodeConfigSource.md) | If specified, the source to get node configuration from The DynamicKubeletConfig feature gate must be enabled for the Kubelet to use this field | [optional] [default to null]
**ExternalID** | **string** | External ID of the node assigned by some machine database (e.g. a cloud provider). Deprecated. | [optional] [default to null]
**PodCIDR** | **string** | PodCIDR represents the pod IP range assigned to the node. | [optional] [default to null]
**ProviderID** | **string** | ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt; | [optional] [default to null]
**Taints** | [**[]V1Taint**](v1.Taint.md) | If specified, the node&#39;s taints. | [optional] [default to null]
**Unschedulable** | **bool** | Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


