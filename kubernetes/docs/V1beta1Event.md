# V1beta1Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | What action was taken/failed regarding to the regarding object. | [optional] 
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**DeprecatedCount** | **int32** | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] 
**DeprecatedFirstTimestamp** | [**time.Time**](time.Time.md) | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] 
**DeprecatedLastTimestamp** | [**time.Time**](time.Time.md) | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] 
**DeprecatedSource** | [**V1EventSource**](v1.EventSource.md) |  | [optional] 
**EventTime** | [**time.Time**](time.Time.md) | Required. Time when this Event was first observed. | 
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] 
**Note** | **string** | Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB. | [optional] 
**Reason** | **string** | Why the action was taken. | [optional] 
**Regarding** | [**V1ObjectReference**](v1.ObjectReference.md) |  | [optional] 
**Related** | [**V1ObjectReference**](v1.ObjectReference.md) |  | [optional] 
**ReportingController** | **string** | Name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. | [optional] 
**ReportingInstance** | **string** | ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. | [optional] 
**Series** | [**V1beta1EventSeries**](v1beta1.EventSeries.md) |  | [optional] 
**Type** | **string** | Type of this event (Normal, Warning), new types could be added in the future. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


