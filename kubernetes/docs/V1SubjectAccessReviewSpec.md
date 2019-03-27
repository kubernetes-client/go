# V1SubjectAccessReviewSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | [**map[string][]string**](array.md) | Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here. | [optional] 
**Groups** | **[]string** | Groups is the groups you&#39;re testing for. | [optional] 
**NonResourceAttributes** | [**V1NonResourceAttributes**](v1.NonResourceAttributes.md) |  | [optional] 
**ResourceAttributes** | [**V1ResourceAttributes**](v1.ResourceAttributes.md) |  | [optional] 
**Uid** | **string** | UID information about the requesting user. | [optional] 
**User** | **string** | User is the user you&#39;re testing for. If you specify \&quot;User\&quot; but not \&quot;Groups\&quot;, then is it interpreted as \&quot;What if User were not a member of any groups | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


