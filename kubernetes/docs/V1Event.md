# V1Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | What action was taken/failed regarding to the Regarding object. | [optional] 
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Count** | **int32** | The number of times this event has occurred. | [optional] 
**EventTime** | [**time.Time**](time.Time.md) | Time when this Event was first observed. | [optional] 
**FirstTimestamp** | [**time.Time**](time.Time.md) | The time at which the event was first recorded. (Time of server receipt is in TypeMeta.) | [optional] 
**InvolvedObject** | [**V1ObjectReference**](v1.ObjectReference.md) |  | 
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**LastTimestamp** | [**time.Time**](time.Time.md) | The time at which the most recent occurrence of this event was recorded. | [optional] 
**Message** | **string** | A human-readable description of the status of this operation. | [optional] 
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) |  | 
**Reason** | **string** | This should be a short, machine understandable string that gives the reason for the transition into the object&#39;s current status. | [optional] 
**Related** | [**V1ObjectReference**](v1.ObjectReference.md) |  | [optional] 
**ReportingComponent** | **string** | Name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. | [optional] 
**ReportingInstance** | **string** | ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. | [optional] 
**Series** | [**V1EventSeries**](v1.EventSeries.md) |  | [optional] 
**Source** | [**V1EventSource**](v1.EventSource.md) |  | [optional] 
**Type** | **string** | Type of this event (Normal, Warning), new types could be added in the future | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


