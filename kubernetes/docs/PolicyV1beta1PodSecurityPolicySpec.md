# PolicyV1beta1PodSecurityPolicySpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | **bool** | allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true. | [optional] 
**AllowedCapabilities** | **[]string** | allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author&#39;s discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities. | [optional] 
**AllowedFlexVolumes** | [**[]PolicyV1beta1AllowedFlexVolume**](policy.v1beta1.AllowedFlexVolume.md) | allowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \&quot;volumes\&quot; field. | [optional] 
**AllowedHostPaths** | [**[]PolicyV1beta1AllowedHostPath**](policy.v1beta1.AllowedHostPath.md) | allowedHostPaths is a white list of allowed host paths. Empty indicates that all host paths may be used. | [optional] 
**AllowedProcMountTypes** | **[]string** | AllowedProcMountTypes is a whitelist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled. | [optional] 
**AllowedUnsafeSysctls** | **[]string** | allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \&quot;*\&quot; in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.  Examples: e.g. \&quot;foo/_*\&quot; allows \&quot;foo/bar\&quot;, \&quot;foo/baz\&quot;, etc. e.g. \&quot;foo.*\&quot; allows \&quot;foo.bar\&quot;, \&quot;foo.baz\&quot;, etc. | [optional] 
**DefaultAddCapabilities** | **[]string** | defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list. | [optional] 
**DefaultAllowPrivilegeEscalation** | **bool** | defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process. | [optional] 
**ForbiddenSysctls** | **[]string** | forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \&quot;*\&quot; in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.  Examples: e.g. \&quot;foo/_*\&quot; forbids \&quot;foo/bar\&quot;, \&quot;foo/baz\&quot;, etc. e.g. \&quot;foo.*\&quot; forbids \&quot;foo.bar\&quot;, \&quot;foo.baz\&quot;, etc. | [optional] 
**FsGroup** | [**PolicyV1beta1FsGroupStrategyOptions**](policy.v1beta1.FSGroupStrategyOptions.md) |  | 
**HostIPC** | **bool** | hostIPC determines if the policy allows the use of HostIPC in the pod spec. | [optional] 
**HostNetwork** | **bool** | hostNetwork determines if the policy allows the use of HostNetwork in the pod spec. | [optional] 
**HostPID** | **bool** | hostPID determines if the policy allows the use of HostPID in the pod spec. | [optional] 
**HostPorts** | [**[]PolicyV1beta1HostPortRange**](policy.v1beta1.HostPortRange.md) | hostPorts determines which host port ranges are allowed to be exposed. | [optional] 
**Privileged** | **bool** | privileged determines if a pod can request to be run as privileged. | [optional] 
**ReadOnlyRootFilesystem** | **bool** | readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to. | [optional] 
**RequiredDropCapabilities** | **[]string** | requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added. | [optional] 
**RunAsGroup** | [**PolicyV1beta1RunAsGroupStrategyOptions**](policy.v1beta1.RunAsGroupStrategyOptions.md) |  | [optional] 
**RunAsUser** | [**PolicyV1beta1RunAsUserStrategyOptions**](policy.v1beta1.RunAsUserStrategyOptions.md) |  | 
**SeLinux** | [**PolicyV1beta1SeLinuxStrategyOptions**](policy.v1beta1.SELinuxStrategyOptions.md) |  | 
**SupplementalGroups** | [**PolicyV1beta1SupplementalGroupsStrategyOptions**](policy.v1beta1.SupplementalGroupsStrategyOptions.md) |  | 
**Volumes** | **[]string** | volumes is a white list of allowed volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use &#39;*&#39;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


