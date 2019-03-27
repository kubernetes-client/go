# V1ContainerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerID** | **string** | Container&#39;s ID in the format &#39;docker://&lt;container_id&gt;&#39;. | [optional] 
**Image** | **string** | The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images | 
**ImageID** | **string** | ImageID of the container&#39;s image. | 
**LastState** | [**V1ContainerState**](v1.ContainerState.md) |  | [optional] 
**Name** | **string** | This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated. | 
**Ready** | **bool** | Specifies whether the container has passed its readiness probe. | 
**RestartCount** | **int32** | The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC. | 
**State** | [**V1ContainerState**](v1.ContainerState.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


