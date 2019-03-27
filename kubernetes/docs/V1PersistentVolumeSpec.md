# V1PersistentVolumeSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes | [optional] 
**AwsElasticBlockStore** | [**V1AwsElasticBlockStoreVolumeSource**](v1.AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | [**V1AzureDiskVolumeSource**](v1.AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | [**V1AzureFilePersistentVolumeSource**](v1.AzureFilePersistentVolumeSource.md) |  | [optional] 
**Capacity** | **map[string]string** | A description of the persistent volume&#39;s resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Cephfs** | [**V1CephFsPersistentVolumeSource**](v1.CephFSPersistentVolumeSource.md) |  | [optional] 
**Cinder** | [**V1CinderPersistentVolumeSource**](v1.CinderPersistentVolumeSource.md) |  | [optional] 
**ClaimRef** | [**V1ObjectReference**](v1.ObjectReference.md) |  | [optional] 
**Csi** | [**V1CsiPersistentVolumeSource**](v1.CSIPersistentVolumeSource.md) |  | [optional] 
**Fc** | [**V1FcVolumeSource**](v1.FCVolumeSource.md) |  | [optional] 
**FlexVolume** | [**V1FlexPersistentVolumeSource**](v1.FlexPersistentVolumeSource.md) |  | [optional] 
**Flocker** | [**V1FlockerVolumeSource**](v1.FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | [**V1GcePersistentDiskVolumeSource**](v1.GCEPersistentDiskVolumeSource.md) |  | [optional] 
**Glusterfs** | [**V1GlusterfsPersistentVolumeSource**](v1.GlusterfsPersistentVolumeSource.md) |  | [optional] 
**HostPath** | [**V1HostPathVolumeSource**](v1.HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | [**V1IscsiPersistentVolumeSource**](v1.ISCSIPersistentVolumeSource.md) |  | [optional] 
**Local** | [**V1LocalVolumeSource**](v1.LocalVolumeSource.md) |  | [optional] 
**MountOptions** | **[]string** | A list of mount options, e.g. [\&quot;ro\&quot;, \&quot;soft\&quot;]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options | [optional] 
**Nfs** | [**V1NfsVolumeSource**](v1.NFSVolumeSource.md) |  | [optional] 
**NodeAffinity** | [**V1VolumeNodeAffinity**](v1.VolumeNodeAffinity.md) |  | [optional] 
**PersistentVolumeReclaimPolicy** | **string** | What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming | [optional] 
**PhotonPersistentDisk** | [**V1PhotonPersistentDiskVolumeSource**](v1.PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | [**V1PortworxVolumeSource**](v1.PortworxVolumeSource.md) |  | [optional] 
**Quobyte** | [**V1QuobyteVolumeSource**](v1.QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | [**V1RbdPersistentVolumeSource**](v1.RBDPersistentVolumeSource.md) |  | [optional] 
**ScaleIO** | [**V1ScaleIoPersistentVolumeSource**](v1.ScaleIOPersistentVolumeSource.md) |  | [optional] 
**StorageClassName** | **string** | Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass. | [optional] 
**Storageos** | [**V1StorageOsPersistentVolumeSource**](v1.StorageOSPersistentVolumeSource.md) |  | [optional] 
**VolumeMode** | **string** | volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec. This is a beta feature. | [optional] 
**VsphereVolume** | [**V1VsphereVirtualDiskVolumeSource**](v1.VsphereVirtualDiskVolumeSource.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


