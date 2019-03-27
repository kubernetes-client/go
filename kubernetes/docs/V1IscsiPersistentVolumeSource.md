# V1IscsiPersistentVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapAuthDiscovery** | **bool** | whether support iSCSI Discovery CHAP authentication | [optional] 
**ChapAuthSession** | **bool** | whether support iSCSI Session CHAP authentication | [optional] 
**FsType** | **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | [optional] 
**InitiatorName** | **string** | Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface &lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | [optional] 
**Iqn** | **string** | Target iSCSI Qualified Name. | 
**IscsiInterface** | **string** | iSCSI Interface Name that uses an iSCSI transport. Defaults to &#39;default&#39; (tcp). | [optional] 
**Lun** | **int32** | iSCSI Target Lun number. | 
**Portals** | **[]string** | iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | [optional] 
**ReadOnly** | **bool** | ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. | [optional] 
**SecretRef** | [**V1SecretReference**](v1.SecretReference.md) |  | [optional] 
**TargetPortal** | **string** | iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


