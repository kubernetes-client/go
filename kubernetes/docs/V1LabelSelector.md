# V1LabelSelector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | [**[]V1LabelSelectorRequirement**](v1.LabelSelectorRequirement.md) | matchExpressions is a list of label selector requirements. The requirements are ANDed. | [optional] [default to null]
**MatchLabels** | **map[string]string** | matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \&quot;key\&quot;, the operator is \&quot;In\&quot;, and the values array contains only \&quot;value\&quot;. The requirements are ANDed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


