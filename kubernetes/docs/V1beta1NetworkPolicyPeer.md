# V1beta1NetworkPolicyPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | [***V1beta1IpBlock**](v1beta1.IPBlock.md) | IPBlock defines policy on a particular IPBlock | [optional] [default to null]
**NamespaceSelector** | [***V1LabelSelector**](v1.LabelSelector.md) | Selects Namespaces using cluster scoped-labels.  This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If present but empty, this selector selects all namespaces. | [optional] [default to null]
**PodSelector** | [***V1LabelSelector**](v1.LabelSelector.md) | This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If present but empty, this selector selects all pods in this namespace. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


