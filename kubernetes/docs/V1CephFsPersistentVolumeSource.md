# V1CephFsPersistentVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitors** | **[]string** | Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [default to null]
**Path** | **string** | Optional: Used as the mounted root, rather than the full Ceph tree, default is / | [optional] [default to null]
**ReadOnly** | **bool** | Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] [default to null]
**SecretFile** | **string** | Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] [default to null]
**SecretRef** | [***V1SecretReference**](v1.SecretReference.md) | Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] [default to null]
**User** | **string** | Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


