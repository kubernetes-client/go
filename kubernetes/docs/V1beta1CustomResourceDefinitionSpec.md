# V1beta1CustomResourceDefinitionSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | Group is the group this resource belongs in | [default to null]
**Names** | [***V1beta1CustomResourceDefinitionNames**](v1beta1.CustomResourceDefinitionNames.md) | Names are the names used to describe this custom resource | [default to null]
**Scope** | **string** | Scope indicates whether this resource is cluster or namespace scoped.  Default is namespaced | [default to null]
**Subresources** | [***V1beta1CustomResourceSubresources**](v1beta1.CustomResourceSubresources.md) | Subresources describes the subresources for CustomResources This field is alpha-level and should only be sent to servers that enable subresources via the CustomResourceSubresources feature gate. | [optional] [default to null]
**Validation** | [***V1beta1CustomResourceValidation**](v1beta1.CustomResourceValidation.md) | Validation describes the validation methods for CustomResources | [optional] [default to null]
**Version** | **string** | Version is the version this resource belongs in | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


