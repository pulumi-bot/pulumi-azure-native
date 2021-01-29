// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contains information about a pool.
type Pool struct {
	pulumi.CustomResourceState

	AllocationState               pulumi.StringOutput `pulumi:"allocationState"`
	AllocationStateTransitionTime pulumi.StringOutput `pulumi:"allocationStateTransitionTime"`
	// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
	ApplicationLicenses pulumi.StringArrayOutput `pulumi:"applicationLicenses"`
	// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
	ApplicationPackages ApplicationPackageReferenceResponseArrayOutput `pulumi:"applicationPackages"`
	// This property is set only if the pool automatically scales, i.e. autoScaleSettings are used.
	AutoScaleRun AutoScaleRunResponseOutput `pulumi:"autoScaleRun"`
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
	Certificates            CertificateReferenceResponseArrayOutput `pulumi:"certificates"`
	CreationTime            pulumi.StringOutput                     `pulumi:"creationTime"`
	CurrentDedicatedNodes   pulumi.IntOutput                        `pulumi:"currentDedicatedNodes"`
	CurrentLowPriorityNodes pulumi.IntOutput                        `pulumi:"currentLowPriorityNodes"`
	// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
	DeploymentConfiguration DeploymentConfigurationResponsePtrOutput `pulumi:"deploymentConfiguration"`
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The type of identity used for the Batch Pool.
	Identity BatchPoolIdentityResponsePtrOutput `pulumi:"identity"`
	// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
	InterNodeCommunication pulumi.StringPtrOutput `pulumi:"interNodeCommunication"`
	// This is the last time at which the pool level data, such as the targetDedicatedNodes or autoScaleSettings, changed. It does not factor in node-level changes such as a compute node changing state.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata MetadataItemResponseArrayOutput `pulumi:"metadata"`
	// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
	MountConfiguration MountConfigurationResponseArrayOutput `pulumi:"mountConfiguration"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network configuration for a pool.
	NetworkConfiguration            NetworkConfigurationResponsePtrOutput `pulumi:"networkConfiguration"`
	ProvisioningState               pulumi.StringOutput                   `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime pulumi.StringOutput                   `pulumi:"provisioningStateTransitionTime"`
	// Describes either the current operation (if the pool AllocationState is Resizing) or the previously completed operation (if the AllocationState is Steady).
	ResizeOperationStatus ResizeOperationStatusResponseOutput `pulumi:"resizeOperationStatus"`
	// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
	ScaleSettings ScaleSettingsResponsePtrOutput `pulumi:"scaleSettings"`
	// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
	StartTask StartTaskResponsePtrOutput `pulumi:"startTask"`
	// If not specified, the default is spread.
	TaskSchedulingPolicy TaskSchedulingPolicyResponsePtrOutput `pulumi:"taskSchedulingPolicy"`
	// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
	TaskSlotsPerNode pulumi.IntPtrOutput `pulumi:"taskSlotsPerNode"`
	// The type of the resource.
	Type         pulumi.StringOutput            `pulumi:"type"`
	UserAccounts UserAccountResponseArrayOutput `pulumi:"userAccounts"`
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:batch/latest:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20170901:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20181201:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190401:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200901:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-nextgen:batch/v20210101:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-nextgen:batch/v20210101:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	AllocationState               *string `pulumi:"allocationState"`
	AllocationStateTransitionTime *string `pulumi:"allocationStateTransitionTime"`
	// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
	ApplicationLicenses []string `pulumi:"applicationLicenses"`
	// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
	ApplicationPackages []ApplicationPackageReferenceResponse `pulumi:"applicationPackages"`
	// This property is set only if the pool automatically scales, i.e. autoScaleSettings are used.
	AutoScaleRun *AutoScaleRunResponse `pulumi:"autoScaleRun"`
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
	Certificates            []CertificateReferenceResponse `pulumi:"certificates"`
	CreationTime            *string                        `pulumi:"creationTime"`
	CurrentDedicatedNodes   *int                           `pulumi:"currentDedicatedNodes"`
	CurrentLowPriorityNodes *int                           `pulumi:"currentLowPriorityNodes"`
	// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
	DeploymentConfiguration *DeploymentConfigurationResponse `pulumi:"deploymentConfiguration"`
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName *string `pulumi:"displayName"`
	// The ETag of the resource, used for concurrency statements.
	Etag *string `pulumi:"etag"`
	// The type of identity used for the Batch Pool.
	Identity *BatchPoolIdentityResponse `pulumi:"identity"`
	// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
	InterNodeCommunication *string `pulumi:"interNodeCommunication"`
	// This is the last time at which the pool level data, such as the targetDedicatedNodes or autoScaleSettings, changed. It does not factor in node-level changes such as a compute node changing state.
	LastModified *string `pulumi:"lastModified"`
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata []MetadataItemResponse `pulumi:"metadata"`
	// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
	MountConfiguration []MountConfigurationResponse `pulumi:"mountConfiguration"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The network configuration for a pool.
	NetworkConfiguration            *NetworkConfigurationResponse `pulumi:"networkConfiguration"`
	ProvisioningState               *string                       `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime *string                       `pulumi:"provisioningStateTransitionTime"`
	// Describes either the current operation (if the pool AllocationState is Resizing) or the previously completed operation (if the AllocationState is Steady).
	ResizeOperationStatus *ResizeOperationStatusResponse `pulumi:"resizeOperationStatus"`
	// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
	ScaleSettings *ScaleSettingsResponse `pulumi:"scaleSettings"`
	// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
	StartTask *StartTaskResponse `pulumi:"startTask"`
	// If not specified, the default is spread.
	TaskSchedulingPolicy *TaskSchedulingPolicyResponse `pulumi:"taskSchedulingPolicy"`
	// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
	TaskSlotsPerNode *int `pulumi:"taskSlotsPerNode"`
	// The type of the resource.
	Type         *string               `pulumi:"type"`
	UserAccounts []UserAccountResponse `pulumi:"userAccounts"`
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VmSize *string `pulumi:"vmSize"`
}

type PoolState struct {
	AllocationState               pulumi.StringPtrInput
	AllocationStateTransitionTime pulumi.StringPtrInput
	// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
	ApplicationLicenses pulumi.StringArrayInput
	// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
	ApplicationPackages ApplicationPackageReferenceResponseArrayInput
	// This property is set only if the pool automatically scales, i.e. autoScaleSettings are used.
	AutoScaleRun AutoScaleRunResponsePtrInput
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
	Certificates            CertificateReferenceResponseArrayInput
	CreationTime            pulumi.StringPtrInput
	CurrentDedicatedNodes   pulumi.IntPtrInput
	CurrentLowPriorityNodes pulumi.IntPtrInput
	// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
	DeploymentConfiguration DeploymentConfigurationResponsePtrInput
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName pulumi.StringPtrInput
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringPtrInput
	// The type of identity used for the Batch Pool.
	Identity BatchPoolIdentityResponsePtrInput
	// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
	InterNodeCommunication pulumi.StringPtrInput
	// This is the last time at which the pool level data, such as the targetDedicatedNodes or autoScaleSettings, changed. It does not factor in node-level changes such as a compute node changing state.
	LastModified pulumi.StringPtrInput
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata MetadataItemResponseArrayInput
	// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
	MountConfiguration MountConfigurationResponseArrayInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The network configuration for a pool.
	NetworkConfiguration            NetworkConfigurationResponsePtrInput
	ProvisioningState               pulumi.StringPtrInput
	ProvisioningStateTransitionTime pulumi.StringPtrInput
	// Describes either the current operation (if the pool AllocationState is Resizing) or the previously completed operation (if the AllocationState is Steady).
	ResizeOperationStatus ResizeOperationStatusResponsePtrInput
	// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
	ScaleSettings ScaleSettingsResponsePtrInput
	// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
	StartTask StartTaskResponsePtrInput
	// If not specified, the default is spread.
	TaskSchedulingPolicy TaskSchedulingPolicyResponsePtrInput
	// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
	TaskSlotsPerNode pulumi.IntPtrInput
	// The type of the resource.
	Type         pulumi.StringPtrInput
	UserAccounts UserAccountResponseArrayInput
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VmSize pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
	ApplicationLicenses []string `pulumi:"applicationLicenses"`
	// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
	ApplicationPackages []ApplicationPackageReference `pulumi:"applicationPackages"`
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
	Certificates []CertificateReference `pulumi:"certificates"`
	// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
	DeploymentConfiguration *DeploymentConfiguration `pulumi:"deploymentConfiguration"`
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName *string `pulumi:"displayName"`
	// The type of identity used for the Batch Pool.
	Identity *BatchPoolIdentity `pulumi:"identity"`
	// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
	InterNodeCommunication *string `pulumi:"interNodeCommunication"`
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata []MetadataItem `pulumi:"metadata"`
	// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
	MountConfiguration []MountConfiguration `pulumi:"mountConfiguration"`
	// The network configuration for a pool.
	NetworkConfiguration *NetworkConfiguration `pulumi:"networkConfiguration"`
	// The pool name. This must be unique within the account.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
	ScaleSettings *ScaleSettings `pulumi:"scaleSettings"`
	// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
	StartTask *StartTask `pulumi:"startTask"`
	// If not specified, the default is spread.
	TaskSchedulingPolicy *TaskSchedulingPolicy `pulumi:"taskSchedulingPolicy"`
	// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
	TaskSlotsPerNode *int          `pulumi:"taskSlotsPerNode"`
	UserAccounts     []UserAccount `pulumi:"userAccounts"`
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VmSize *string `pulumi:"vmSize"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput
	// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
	ApplicationLicenses pulumi.StringArrayInput
	// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
	ApplicationPackages ApplicationPackageReferenceArrayInput
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
	Certificates CertificateReferenceArrayInput
	// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
	DeploymentConfiguration DeploymentConfigurationPtrInput
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName pulumi.StringPtrInput
	// The type of identity used for the Batch Pool.
	Identity BatchPoolIdentityPtrInput
	// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
	InterNodeCommunication *InterNodeCommunicationState
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata MetadataItemArrayInput
	// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
	MountConfiguration MountConfigurationArrayInput
	// The network configuration for a pool.
	NetworkConfiguration NetworkConfigurationPtrInput
	// The pool name. This must be unique within the account.
	PoolName pulumi.StringInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
	// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
	ScaleSettings ScaleSettingsPtrInput
	// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
	StartTask StartTaskPtrInput
	// If not specified, the default is spread.
	TaskSchedulingPolicy TaskSchedulingPolicyPtrInput
	// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
	TaskSlotsPerNode pulumi.IntPtrInput
	UserAccounts     UserAccountArrayInput
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VmSize pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct {
	*pulumi.OutputState
}

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
