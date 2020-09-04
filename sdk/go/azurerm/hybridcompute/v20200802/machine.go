// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200802

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a hybrid machine.
type Machine struct {
	pulumi.CustomResourceState

	// Specifies the AD fully qualified display name.
	AdFqdn pulumi.StringOutput `pulumi:"adFqdn"`
	// The hybrid machine agent full version.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey pulumi.StringPtrOutput `pulumi:"clientPublicKey"`
	// Specifies the hybrid machine display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies the DNS fully qualified display name.
	DnsFqdn pulumi.StringOutput `pulumi:"dnsFqdn"`
	// Specifies the Windows domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Details about the error state.
	ErrorDetails ErrorDetailResponseArrayOutput `pulumi:"errorDetails"`
	// Machine Extensions information
	Extensions MachineExtensionInstanceViewResponseArrayOutput `pulumi:"extensions"`
	Identity   MachineResponseIdentityPtrOutput                `pulumi:"identity"`
	// The time of the last status change.
	LastStatusChange pulumi.StringOutput `pulumi:"lastStatusChange"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata pertaining to the geographic location of the resource.
	LocationData LocationDataResponsePtrOutput `pulumi:"locationData"`
	// Specifies the hybrid machine FQDN.
	MachineFqdn pulumi.StringOutput `pulumi:"machineFqdn"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operating System running on the hybrid machine.
	OsName pulumi.StringOutput `pulumi:"osName"`
	// Specifies the operating system settings for the hybrid machine.
	OsProfile MachinePropertiesResponseOsProfilePtrOutput `pulumi:"osProfile"`
	// Specifies the Operating System product SKU.
	OsSku pulumi.StringOutput `pulumi:"osSku"`
	// The version of Operating System running on the hybrid machine.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of the hybrid machine agent.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the hybrid machine unique ID.
	VmId pulumi.StringPtrOutput `pulumi:"vmId"`
	// Specifies the Arc Machine's unique SMBIOS ID
	VmUuid pulumi.StringOutput `pulumi:"vmUuid"`
}

// NewMachine registers a new resource with the given unique name, arguments, and options.
func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MachineArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:hybridcompute/latest:Machine"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20190802preview:Machine"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20191212:Machine"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20200815preview:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azurerm:hybridcompute/v20200802:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachine gets an existing Machine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azurerm:hybridcompute/v20200802:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Machine resources.
type machineState struct {
	// Specifies the AD fully qualified display name.
	AdFqdn *string `pulumi:"adFqdn"`
	// The hybrid machine agent full version.
	AgentVersion *string `pulumi:"agentVersion"`
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey *string `pulumi:"clientPublicKey"`
	// Specifies the hybrid machine display name.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the DNS fully qualified display name.
	DnsFqdn *string `pulumi:"dnsFqdn"`
	// Specifies the Windows domain name.
	DomainName *string `pulumi:"domainName"`
	// Details about the error state.
	ErrorDetails []ErrorDetailResponse `pulumi:"errorDetails"`
	// Machine Extensions information
	Extensions []MachineExtensionInstanceViewResponse `pulumi:"extensions"`
	Identity   *MachineResponseIdentity               `pulumi:"identity"`
	// The time of the last status change.
	LastStatusChange *string `pulumi:"lastStatusChange"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Metadata pertaining to the geographic location of the resource.
	LocationData *LocationDataResponse `pulumi:"locationData"`
	// Specifies the hybrid machine FQDN.
	MachineFqdn *string `pulumi:"machineFqdn"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Operating System running on the hybrid machine.
	OsName *string `pulumi:"osName"`
	// Specifies the operating system settings for the hybrid machine.
	OsProfile *MachinePropertiesResponseOsProfile `pulumi:"osProfile"`
	// Specifies the Operating System product SKU.
	OsSku *string `pulumi:"osSku"`
	// The version of Operating System running on the hybrid machine.
	OsVersion *string `pulumi:"osVersion"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The status of the hybrid machine agent.
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Specifies the hybrid machine unique ID.
	VmId *string `pulumi:"vmId"`
	// Specifies the Arc Machine's unique SMBIOS ID
	VmUuid *string `pulumi:"vmUuid"`
}

type MachineState struct {
	// Specifies the AD fully qualified display name.
	AdFqdn pulumi.StringPtrInput
	// The hybrid machine agent full version.
	AgentVersion pulumi.StringPtrInput
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey pulumi.StringPtrInput
	// Specifies the hybrid machine display name.
	DisplayName pulumi.StringPtrInput
	// Specifies the DNS fully qualified display name.
	DnsFqdn pulumi.StringPtrInput
	// Specifies the Windows domain name.
	DomainName pulumi.StringPtrInput
	// Details about the error state.
	ErrorDetails ErrorDetailResponseArrayInput
	// Machine Extensions information
	Extensions MachineExtensionInstanceViewResponseArrayInput
	Identity   MachineResponseIdentityPtrInput
	// The time of the last status change.
	LastStatusChange pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Metadata pertaining to the geographic location of the resource.
	LocationData LocationDataResponsePtrInput
	// Specifies the hybrid machine FQDN.
	MachineFqdn pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Operating System running on the hybrid machine.
	OsName pulumi.StringPtrInput
	// Specifies the operating system settings for the hybrid machine.
	OsProfile MachinePropertiesResponseOsProfilePtrInput
	// Specifies the Operating System product SKU.
	OsSku pulumi.StringPtrInput
	// The version of Operating System running on the hybrid machine.
	OsVersion pulumi.StringPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The status of the hybrid machine agent.
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Specifies the hybrid machine unique ID.
	VmId pulumi.StringPtrInput
	// Specifies the Arc Machine's unique SMBIOS ID
	VmUuid pulumi.StringPtrInput
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey *string          `pulumi:"clientPublicKey"`
	Identity        *MachineIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Metadata pertaining to the geographic location of the resource.
	LocationData *LocationData `pulumi:"locationData"`
	// The name of the hybrid machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the hybrid machine unique ID.
	VmId *string `pulumi:"vmId"`
}

// The set of arguments for constructing a Machine resource.
type MachineArgs struct {
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey pulumi.StringPtrInput
	Identity        MachineIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Metadata pertaining to the geographic location of the resource.
	LocationData LocationDataPtrInput
	// The name of the hybrid machine.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies the hybrid machine unique ID.
	VmId pulumi.StringPtrInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}
