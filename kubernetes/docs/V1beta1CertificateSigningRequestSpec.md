# V1beta1CertificateSigningRequestSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | [**map[string][]string**](array.md) | Extra information about the requesting user. See user.Info interface for details. | [optional] [default to null]
**Groups** | **[]string** | Group information about the requesting user. See user.Info interface for details. | [optional] [default to null]
**Request** | **string** | Base64-encoded PKCS#10 CSR data | [default to null]
**Uid** | **string** | UID information about the requesting user. See user.Info interface for details. | [optional] [default to null]
**Usages** | **[]string** | allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3      https://tools.ietf.org/html/rfc5280#section-4.2.1.12 | [optional] [default to null]
**Username** | **string** | Information about the requesting user. See user.Info interface for details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


