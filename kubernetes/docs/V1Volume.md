# V1Volume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsElasticBlockStore** | [***V1AwsElasticBlockStoreVolumeSource**](v1.AWSElasticBlockStoreVolumeSource.md) | AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | [optional] [default to null]
**AzureDisk** | [***V1AzureDiskVolumeSource**](v1.AzureDiskVolumeSource.md) | AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod. | [optional] [default to null]
**AzureFile** | [***V1AzureFileVolumeSource**](v1.AzureFileVolumeSource.md) | AzureFile represents an Azure File Service mount on the host and bind mount to the pod. | [optional] [default to null]
**Cephfs** | [***V1CephFsVolumeSource**](v1.CephFSVolumeSource.md) | CephFS represents a Ceph FS mount on the host that shares a pod&#39;s lifetime | [optional] [default to null]
**Cinder** | [***V1CinderVolumeSource**](v1.CinderVolumeSource.md) | Cinder represents a cinder volume attached and mounted on kubelets host machine More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md | [optional] [default to null]
**ConfigMap** | [***V1ConfigMapVolumeSource**](v1.ConfigMapVolumeSource.md) | ConfigMap represents a configMap that should populate this volume | [optional] [default to null]
**DownwardAPI** | [***V1DownwardApiVolumeSource**](v1.DownwardAPIVolumeSource.md) | DownwardAPI represents downward API about the pod that should populate this volume | [optional] [default to null]
**EmptyDir** | [***V1EmptyDirVolumeSource**](v1.EmptyDirVolumeSource.md) | EmptyDir represents a temporary directory that shares a pod&#39;s lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | [optional] [default to null]
**Fc** | [***V1FcVolumeSource**](v1.FCVolumeSource.md) | FC represents a Fibre Channel resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. | [optional] [default to null]
**FlexVolume** | [***V1FlexVolumeSource**](v1.FlexVolumeSource.md) | FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin. | [optional] [default to null]
**Flocker** | [***V1FlockerVolumeSource**](v1.FlockerVolumeSource.md) | Flocker represents a Flocker volume attached to a kubelet&#39;s host machine. This depends on the Flocker control service being running | [optional] [default to null]
**GcePersistentDisk** | [***V1GcePersistentDiskVolumeSource**](v1.GCEPersistentDiskVolumeSource.md) | GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | [optional] [default to null]
**GitRepo** | [***V1GitRepoVolumeSource**](v1.GitRepoVolumeSource.md) | GitRepo represents a git repository at a particular revision. | [optional] [default to null]
**Glusterfs** | [***V1GlusterfsVolumeSource**](v1.GlusterfsVolumeSource.md) | Glusterfs represents a Glusterfs mount on the host that shares a pod&#39;s lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md | [optional] [default to null]
**HostPath** | [***V1HostPathVolumeSource**](v1.HostPathVolumeSource.md) | HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | [optional] [default to null]
**Iscsi** | [***V1IscsiVolumeSource**](v1.ISCSIVolumeSource.md) | ISCSI represents an ISCSI Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. More info: https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md | [optional] [default to null]
**Name** | **string** | Volume&#39;s name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [default to null]
**Nfs** | [***V1NfsVolumeSource**](v1.NFSVolumeSource.md) | NFS represents an NFS mount on the host that shares a pod&#39;s lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | [optional] [default to null]
**PersistentVolumeClaim** | [***V1PersistentVolumeClaimVolumeSource**](v1.PersistentVolumeClaimVolumeSource.md) | PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | [optional] [default to null]
**PhotonPersistentDisk** | [***V1PhotonPersistentDiskVolumeSource**](v1.PhotonPersistentDiskVolumeSource.md) | PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine | [optional] [default to null]
**PortworxVolume** | [***V1PortworxVolumeSource**](v1.PortworxVolumeSource.md) | PortworxVolume represents a portworx volume attached and mounted on kubelets host machine | [optional] [default to null]
**Projected** | [***V1ProjectedVolumeSource**](v1.ProjectedVolumeSource.md) | Items for all in one resources secrets, configmaps, and downward API | [optional] [default to null]
**Quobyte** | [***V1QuobyteVolumeSource**](v1.QuobyteVolumeSource.md) | Quobyte represents a Quobyte mount on the host that shares a pod&#39;s lifetime | [optional] [default to null]
**Rbd** | [***V1RbdVolumeSource**](v1.RBDVolumeSource.md) | RBD represents a Rados Block Device mount on the host that shares a pod&#39;s lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md | [optional] [default to null]
**ScaleIO** | [***V1ScaleIoVolumeSource**](v1.ScaleIOVolumeSource.md) | ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes. | [optional] [default to null]
**Secret** | [***V1SecretVolumeSource**](v1.SecretVolumeSource.md) | Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | [optional] [default to null]
**Storageos** | [***V1StorageOsVolumeSource**](v1.StorageOSVolumeSource.md) | StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes. | [optional] [default to null]
**VsphereVolume** | [***V1VsphereVirtualDiskVolumeSource**](v1.VsphereVirtualDiskVolumeSource.md) | VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


