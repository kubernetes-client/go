# V1OwnerReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version of the referent. | [default to null]
**BlockOwnerDeletion** | **bool** | If true, AND if the owner has the \&quot;foregroundDeletion\&quot; finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \&quot;delete\&quot; permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. | [optional] [default to null]
**Controller** | **bool** | If true, this reference points to the managing controller. | [optional] [default to null]
**Kind** | **string** | Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [default to null]
**Name** | **string** | Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names | [default to null]
**Uid** | **string** | UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


