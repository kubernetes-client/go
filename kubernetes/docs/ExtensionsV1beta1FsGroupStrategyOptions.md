# ExtensionsV1beta1FsGroupStrategyOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ranges** | [**[]ExtensionsV1beta1IdRange**](extensions.v1beta1.IDRange.md) | Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. | [optional] [default to null]
**Rule** | **string** | Rule is the strategy that will dictate what FSGroup is used in the SecurityContext. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


