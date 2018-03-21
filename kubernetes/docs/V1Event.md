# V1Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | What action was taken/failed regarding to the Regarding object. | [optional] [default to null]
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Count** | **int32** | The number of times this event has occurred. | [optional] [default to null]
**EventTime** | [**time.Time**](time.Time.md) | Time when this Event was first observed. | [optional] [default to null]
**FirstTimestamp** | [**time.Time**](time.Time.md) | The time at which the event was first recorded. (Time of server receipt is in TypeMeta.) | [optional] [default to null]
**InvolvedObject** | [***V1ObjectReference**](v1.ObjectReference.md) | The object that this event is about. | [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**LastTimestamp** | [**time.Time**](time.Time.md) | The time at which the most recent occurrence of this event was recorded. | [optional] [default to null]
**Message** | **string** | A human-readable description of the status of this operation. | [optional] [default to null]
**Metadata** | [***V1ObjectMeta**](v1.ObjectMeta.md) | Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata | [default to null]
**Reason** | **string** | This should be a short, machine understandable string that gives the reason for the transition into the object&#39;s current status. | [optional] [default to null]
**Related** | [***V1ObjectReference**](v1.ObjectReference.md) | Optional secondary object for more complex actions. | [optional] [default to null]
**ReportingComponent** | **string** | Name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. | [optional] [default to null]
**ReportingInstance** | **string** | ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. | [optional] [default to null]
**Series** | [***V1EventSeries**](v1.EventSeries.md) | Data about the Event series this event represents or nil if it&#39;s a singleton Event. | [optional] [default to null]
**Source** | [***V1EventSource**](v1.EventSource.md) | The component reporting this event. Should be a short machine understandable string. | [optional] [default to null]
**Type_** | **string** | Type of this event (Normal, Warning), new types could be added in the future | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


