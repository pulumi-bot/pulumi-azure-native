// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The default value is Pool. If the pool is running Windows a value of Task should be specified if stricter isolation between tasks is required. For example, if the task mutates the registry in a way which could impact other tasks, or if certificates have been specified on the pool which should not be accessible by normal tasks but should be accessible by start tasks.
type AutoUserScope pulumi.String

const (
	// Specifies that the service should create a new user for the task.
	AutoUserScopeTask = AutoUserScope("Task")
	// Specifies that the task runs as the common auto user account which is created on every node in a pool.
	AutoUserScopePool = AutoUserScope("Pool")
)

func (AutoUserScope) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AutoUserScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUserScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUserScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoUserScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Values are:
//
//  none - The caching mode for the disk is not enabled.
//  readOnly - The caching mode for the disk is read only.
//  readWrite - The caching mode for the disk is read and write.
//
//  The default value for caching is none. For information about the caching options see: https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
type CachingType pulumi.String

const (
	// The caching mode for the disk is not enabled.
	CachingTypeNone = CachingType("None")
	// The caching mode for the disk is read only.
	CachingTypeReadOnly = CachingType("ReadOnly")
	// The caching mode for the disk is read and write.
	CachingTypeReadWrite = CachingType("ReadWrite")
)

func (CachingType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CachingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CachingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
type CertificateFormat pulumi.String

const (
	// The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
	CertificateFormatPfx = CertificateFormat("Pfx")
	// The certificate is a base64-encoded X.509 certificate.
	CertificateFormatCer = CertificateFormat("Cer")
)

func (CertificateFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CertificateFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The default value is currentUser. This property is applicable only for pools configured with Windows nodes (that is, created with cloudServiceConfiguration, or with virtualMachineConfiguration using a Windows image reference). For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
type CertificateStoreLocation pulumi.String

const (
	// Certificates should be installed to the CurrentUser certificate store.
	CertificateStoreLocationCurrentUser = CertificateStoreLocation("CurrentUser")
	// Certificates should be installed to the LocalMachine certificate store.
	CertificateStoreLocationLocalMachine = CertificateStoreLocation("LocalMachine")
)

func (CertificateStoreLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CertificateStoreLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStoreLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStoreLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateStoreLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateVisibility pulumi.String

const (
	// The certificate should be visible to the user account under which the start task is run. Note that if AutoUser Scope is Pool for both the StartTask and a Task, this certificate will be visible to the Task as well.
	CertificateVisibilityStartTask = CertificateVisibility("StartTask")
	// The certificate should be visible to the user accounts under which job tasks are run.
	CertificateVisibilityTask = CertificateVisibility("Task")
	// The certificate should be visible to the user accounts under which users remotely access the node.
	CertificateVisibilityRemoteUser = CertificateVisibility("RemoteUser")
)

func (CertificateVisibility) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CertificateVisibility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateVisibility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateVisibility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateVisibility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// If omitted, the default value is Requeue.
type ComputeNodeDeallocationOption pulumi.String

const (
	// Terminate running task processes and requeue the tasks. The tasks will run again when a node is available. Remove nodes as soon as tasks have been terminated.
	ComputeNodeDeallocationOptionRequeue = ComputeNodeDeallocationOption("Requeue")
	// Terminate running tasks. The tasks will be completed with failureInfo indicating that they were terminated, and will not run again. Remove nodes as soon as tasks have been terminated.
	ComputeNodeDeallocationOptionTerminate = ComputeNodeDeallocationOption("Terminate")
	// Allow currently running tasks to complete. Schedule no new tasks while waiting. Remove nodes when all tasks have completed.
	ComputeNodeDeallocationOptionTaskCompletion = ComputeNodeDeallocationOption("TaskCompletion")
	// Allow currently running tasks to complete, then wait for all task data retention periods to expire. Schedule no new tasks while waiting. Remove nodes when all task retention periods have expired.
	ComputeNodeDeallocationOptionRetainedData = ComputeNodeDeallocationOption("RetainedData")
)

func (ComputeNodeDeallocationOption) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ComputeNodeDeallocationOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeDeallocationOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeDeallocationOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeNodeDeallocationOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeNodeFillType pulumi.String

const (
	// Tasks should be assigned evenly across all nodes in the pool.
	ComputeNodeFillTypeSpread = ComputeNodeFillType("Spread")
	// As many tasks as possible (taskSlotsPerNode) should be assigned to each node in the pool before any tasks are assigned to the next node in the pool.
	ComputeNodeFillTypePack = ComputeNodeFillType("Pack")
)

func (ComputeNodeFillType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ComputeNodeFillType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeFillType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeFillType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeNodeFillType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerType pulumi.String

const (
	// A Docker compatible container technology will be used to launch the containers.
	ContainerTypeDockerCompatible = ContainerType("DockerCompatible")
)

func (ContainerType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ContainerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerWorkingDirectory pulumi.String

const (
	// Use the standard Batch service task working directory, which will contain the Task resource files populated by Batch.
	ContainerWorkingDirectoryTaskWorkingDirectory = ContainerWorkingDirectory("TaskWorkingDirectory")
	// Using container image defined working directory. Beware that this directory will not contain the resource files downloaded by Batch.
	ContainerWorkingDirectoryContainerImageDefault = ContainerWorkingDirectory("ContainerImageDefault")
)

func (ContainerWorkingDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ContainerWorkingDirectory) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerWorkingDirectory) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerWorkingDirectory) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerWorkingDirectory) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// If omitted, no disks on the compute nodes in the pool will be encrypted.
type DiskEncryptionTarget pulumi.String

const (
	// The OS Disk on the compute node is encrypted.
	DiskEncryptionTargetOsDisk = DiskEncryptionTarget("OsDisk")
	// The temporary disk on the compute node is encrypted. On Linux this encryption applies to other partitions (such as those on mounted data disks) when encryption occurs at boot time.
	DiskEncryptionTargetTemporaryDisk = DiskEncryptionTarget("TemporaryDisk")
)

func (DiskEncryptionTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DiskEncryptionTarget) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionTarget) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskEncryptionTarget) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskEncryptionTarget) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// nonAdmin - The auto user is a standard user without elevated access. admin - The auto user is a user with elevated access and operates with full Administrator permissions. The default value is nonAdmin.
type ElevationLevel pulumi.String

const (
	// The user is a standard user without elevated access.
	ElevationLevelNonAdmin = ElevationLevel("NonAdmin")
	// The user is a user with elevated access and operates with full Administrator permissions.
	ElevationLevelAdmin = ElevationLevel("Admin")
)

func (ElevationLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ElevationLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElevationLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElevationLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ElevationLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The default value is BatchManaged
type IPAddressProvisioningType pulumi.String

const (
	// A public IP will be created and managed by Batch. There may be multiple public IPs depending on the size of the Pool.
	IPAddressProvisioningTypeBatchManaged = IPAddressProvisioningType("BatchManaged")
	// Public IPs are provided by the user and will be used to provision the Compute Nodes.
	IPAddressProvisioningTypeUserManaged = IPAddressProvisioningType("UserManaged")
	// No public IP Address will be created for the Compute Nodes in the Pool.
	IPAddressProvisioningTypeNoPublicIPAddresses = IPAddressProvisioningType("NoPublicIPAddresses")
)

func (IPAddressProvisioningType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IPAddressProvisioningType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressProvisioningType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressProvisioningType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAddressProvisioningType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InboundEndpointProtocol pulumi.String

const (
	// Use TCP for the endpoint.
	InboundEndpointProtocolTCP = InboundEndpointProtocol("TCP")
	// Use UDP for the endpoint.
	InboundEndpointProtocolUDP = InboundEndpointProtocol("UDP")
)

func (InboundEndpointProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e InboundEndpointProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundEndpointProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundEndpointProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InboundEndpointProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
type InterNodeCommunicationState pulumi.String

const (
	// Enable network communication between virtual machines.
	InterNodeCommunicationStateEnabled = InterNodeCommunicationState("Enabled")
	// Disable network communication between virtual machines.
	InterNodeCommunicationStateDisabled = InterNodeCommunicationState("Disabled")
)

func (InterNodeCommunicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e InterNodeCommunicationState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InterNodeCommunicationState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InterNodeCommunicationState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InterNodeCommunicationState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Type of the key source.
type KeySource pulumi.String

const (
	// Batch creates and manages the encryption keys used to protect the account data.
	KeySource_Microsoft_Batch = KeySource("Microsoft.Batch")
	// The encryption keys used to protect the account data are stored in an external key vault. If this is set then the Batch Account identity must be set to `SystemAssigned` and a valid Key Identifier must also be supplied under the keyVaultProperties.
	KeySource_Microsoft_KeyVault = KeySource("Microsoft.KeyVault")
)

func (KeySource) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e KeySource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeySource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Specifies login mode for the user. The default value for VirtualMachineConfiguration pools is interactive mode and for CloudServiceConfiguration pools is batch mode.
type LoginMode pulumi.String

const (
	// The LOGON32_LOGON_BATCH Win32 login mode. The batch login mode is recommended for long running parallel processes.
	LoginModeBatch = LoginMode("Batch")
	// The LOGON32_LOGON_INTERACTIVE Win32 login mode. Some applications require having permissions associated with the interactive login mode. If this is the case for an application used in your task, then this option is recommended.
	LoginModeInteractive = LoginMode("Interactive")
)

func (LoginMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e LoginMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoginMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoginMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoginMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkSecurityGroupRuleAccess pulumi.String

const (
	// Allow access.
	NetworkSecurityGroupRuleAccessAllow = NetworkSecurityGroupRuleAccess("Allow")
	// Deny access.
	NetworkSecurityGroupRuleAccessDeny = NetworkSecurityGroupRuleAccess("Deny")
)

func (NetworkSecurityGroupRuleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e NetworkSecurityGroupRuleAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkSecurityGroupRuleAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Allocation policy used by Batch Service to provision the nodes. If not specified, Batch will use the regional policy.
type NodePlacementPolicyType pulumi.String

const (
	// All nodes in the pool will be allocated in the same region.
	NodePlacementPolicyTypeRegional = NodePlacementPolicyType("Regional")
	// Nodes in the pool will be spread across different zones with best effort balancing.
	NodePlacementPolicyTypeZonal = NodePlacementPolicyType("Zonal")
)

func (NodePlacementPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e NodePlacementPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodePlacementPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodePlacementPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NodePlacementPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
type PoolAllocationMode pulumi.String

const (
	// Pools will be allocated in subscriptions owned by the Batch service.
	PoolAllocationModeBatchService = PoolAllocationMode("BatchService")
	// Pools will be allocated in a subscription owned by the user.
	PoolAllocationModeUserSubscription = PoolAllocationMode("UserSubscription")
)

func (PoolAllocationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PoolAllocationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolAllocationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolAllocationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PoolAllocationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of identity used for the Batch Pool.
type PoolIdentityType pulumi.String

const (
	// Batch pool has user assigned identities with it.
	PoolIdentityTypeUserAssigned = PoolIdentityType("UserAssigned")
	// Batch pool has no identity associated with it. Setting `None` in update pool will remove existing identities.
	PoolIdentityTypeNone = PoolIdentityType("None")
)

func (PoolIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PoolIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PoolIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// If not specified, the default value is 'enabled'.
type PublicNetworkAccessType pulumi.String

const (
	// Enables connectivity to Azure Batch through public DNS.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public connectivity and enables private connectivity to Azure Batch Service through private endpoint resource.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func (PublicNetworkAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PublicNetworkAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of identity used for the Batch account.
type ResourceIdentityType pulumi.String

const (
	// Batch account has a system assigned identity with it.
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	// Batch account has user assigned identities with it.
	ResourceIdentityTypeUserAssigned = ResourceIdentityType("UserAssigned")
	// Batch account has no identity associated with it. Setting `None` in update account will remove existing identities.
	ResourceIdentityTypeNone = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// If omitted, the default is "Standard_LRS". Values are:
//
//  Standard_LRS - The data disk should use standard locally redundant storage.
//  Premium_LRS - The data disk should use premium locally redundant storage.
type StorageAccountType pulumi.String

const (
	// The data disk should use standard locally redundant storage.
	StorageAccountType_Standard_LRS = StorageAccountType("Standard_LRS")
	// The data disk should use premium locally redundant storage.
	StorageAccountType_Premium_LRS = StorageAccountType("Premium_LRS")
)

func (StorageAccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e StorageAccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
