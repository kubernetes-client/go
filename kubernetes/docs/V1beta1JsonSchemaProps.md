# V1beta1JsonSchemaProps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** |  | [optional] 
**Schema** | **string** |  | [optional] 
**AdditionalItems** | [**map[string]interface{}**](.md) | JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property. | [optional] 
**AdditionalProperties** | [**map[string]interface{}**](.md) | JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property. | [optional] 
**AllOf** | [**[]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**AnyOf** | [**[]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**Default** | [**map[string]interface{}**](.md) | JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil. | [optional] 
**Definitions** | [**map[string]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**Dependencies** | [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Description** | **string** |  | [optional] 
**Enum** | [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Example** | [**map[string]interface{}**](.md) | JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil. | [optional] 
**ExclusiveMaximum** | **bool** |  | [optional] 
**ExclusiveMinimum** | **bool** |  | [optional] 
**ExternalDocs** | [**V1beta1ExternalDocumentation**](v1beta1.ExternalDocumentation.md) |  | [optional] 
**Format** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Items** | [**map[string]interface{}**](.md) | JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes. | [optional] 
**MaxItems** | **int64** |  | [optional] 
**MaxLength** | **int64** |  | [optional] 
**MaxProperties** | **int64** |  | [optional] 
**Maximum** | **float64** |  | [optional] 
**MinItems** | **int64** |  | [optional] 
**MinLength** | **int64** |  | [optional] 
**MinProperties** | **int64** |  | [optional] 
**Minimum** | **float64** |  | [optional] 
**MultipleOf** | **float64** |  | [optional] 
**Not** | [**V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**OneOf** | [**[]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**Pattern** | **string** |  | [optional] 
**PatternProperties** | [**map[string]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**Properties** | [**map[string]V1beta1JsonSchemaProps**](v1beta1.JSONSchemaProps.md) |  | [optional] 
**Required** | **[]string** |  | [optional] 
**Title** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**UniqueItems** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


