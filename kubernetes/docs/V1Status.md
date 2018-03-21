# V1Status

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Code** | **int32** | Suggested HTTP return code for this status, 0 if not set. | [optional] [default to null]
**Details** | [***V1StatusDetails**](v1.StatusDetails.md) | Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type. | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Message** | **string** | A human-readable description of the status of this operation. | [optional] [default to null]
**Metadata** | [***V1ListMeta**](v1.ListMeta.md) | Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Reason** | **string** | A machine-readable description of why this operation is in the \&quot;Failure\&quot; status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it. | [optional] [default to null]
**Status** | **string** | Status of the operation. One of: \&quot;Success\&quot; or \&quot;Failure\&quot;. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


