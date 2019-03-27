# V1beta1IngressSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | [**V1beta1IngressBackend**](v1beta1.IngressBackend.md) |  | [optional] 
**Rules** | [**[]V1beta1IngressRule**](v1beta1.IngressRule.md) | A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. | [optional] 
**Tls** | [**[]V1beta1IngressTls**](v1beta1.IngressTLS.md) | TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


