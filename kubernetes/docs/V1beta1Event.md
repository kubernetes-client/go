# V1beta1Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | What action was taken/failed regarding to the regarding object. | [optional] [default to null]
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**DeprecatedCount** | **int32** | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] [default to null]
**DeprecatedFirstTimestamp** | [**time.Time**](time.Time.md) | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] [default to null]
**DeprecatedLastTimestamp** | [**time.Time**](time.Time.md) | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] [default to null]
**DeprecatedSource** | [***V1EventSource**](v1.EventSource.md) | Deprecated field assuring backward compatibility with core.v1 Event type | [optional] [default to null]
**EventTime** | [**time.Time**](time.Time.md) | Required. Time when this Event was first observed. | [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [***V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] [default to null]
**Note** | **string** | Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB. | [optional] [default to null]
**Reason** | **string** | Why the action was taken. | [optional] [default to null]
**Regarding** | [***V1ObjectReference**](v1.ObjectReference.md) | The object this Event is about. In most cases it&#39;s an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object. | [optional] [default to null]
**Related** | [***V1ObjectReference**](v1.ObjectReference.md) | Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object. | [optional] [default to null]
**ReportingController** | **string** | Name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. | [optional] [default to null]
**ReportingInstance** | **string** | ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. | [optional] [default to null]
**Series** | [***V1beta1EventSeries**](v1beta1.EventSeries.md) | Data about the Event series this event represents or nil if it&#39;s a singleton Event. | [optional] [default to null]
**Type_** | **string** | Type of this event (Normal, Warning), new types could be added in the future. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


