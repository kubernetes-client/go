# PolicyV1beta1PodSecurityPolicySpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | **bool** | AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true. | [optional] [default to null]
**AllowedCapabilities** | **[]string** | AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author&#39;s discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities. | [optional] [default to null]
**AllowedFlexVolumes** | [**[]PolicyV1beta1AllowedFlexVolume**](policy.v1beta1.AllowedFlexVolume.md) | AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \&quot;Volumes\&quot; field. | [optional] [default to null]
**AllowedHostPaths** | [**[]PolicyV1beta1AllowedHostPath**](policy.v1beta1.AllowedHostPath.md) | is a white list of allowed host paths. Empty indicates that all host paths may be used. | [optional] [default to null]
**DefaultAddCapabilities** | **[]string** | DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both DefaultAddCapabilities and RequiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the AllowedCapabilities list. | [optional] [default to null]
**DefaultAllowPrivilegeEscalation** | **bool** | DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process. | [optional] [default to null]
**FsGroup** | [***PolicyV1beta1FsGroupStrategyOptions**](policy.v1beta1.FSGroupStrategyOptions.md) | FSGroup is the strategy that will dictate what fs group is used by the SecurityContext. | [default to null]
**HostIPC** | **bool** | hostIPC determines if the policy allows the use of HostIPC in the pod spec. | [optional] [default to null]
**HostNetwork** | **bool** | hostNetwork determines if the policy allows the use of HostNetwork in the pod spec. | [optional] [default to null]
**HostPID** | **bool** | hostPID determines if the policy allows the use of HostPID in the pod spec. | [optional] [default to null]
**HostPorts** | [**[]PolicyV1beta1HostPortRange**](policy.v1beta1.HostPortRange.md) | hostPorts determines which host port ranges are allowed to be exposed. | [optional] [default to null]
**Privileged** | **bool** | privileged determines if a pod can request to be run as privileged. | [optional] [default to null]
**ReadOnlyRootFilesystem** | **bool** | ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to. | [optional] [default to null]
**RequiredDropCapabilities** | **[]string** | RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added. | [optional] [default to null]
**RunAsUser** | [***PolicyV1beta1RunAsUserStrategyOptions**](policy.v1beta1.RunAsUserStrategyOptions.md) | runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set. | [default to null]
**SeLinux** | [***PolicyV1beta1SeLinuxStrategyOptions**](policy.v1beta1.SELinuxStrategyOptions.md) | seLinux is the strategy that will dictate the allowable labels that may be set. | [default to null]
**SupplementalGroups** | [***PolicyV1beta1SupplementalGroupsStrategyOptions**](policy.v1beta1.SupplementalGroupsStrategyOptions.md) | SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext. | [default to null]
**Volumes** | **[]string** | volumes is a white list of allowed volume plugins.  Empty indicates that all plugins may be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


