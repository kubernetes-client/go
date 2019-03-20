# V1ClusterRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationRule** | [**V1AggregationRule**](v1.AggregationRule.md) |  | [optional] 
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] 
**Rules** | [**[]V1PolicyRule**](v1.PolicyRule.md) | Rules holds all the PolicyRules for this ClusterRole | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


