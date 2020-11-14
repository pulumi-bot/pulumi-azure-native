// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Managed Network resource
type ManagedNetwork struct {
	pulumi.CustomResourceState

	// The collection of groups and policies concerned with connectivity
	Connectivity ConnectivityCollectionResponseOutput `pulumi:"connectivity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope ScopeResponsePtrOutput `pulumi:"scope"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedNetwork registers a new resource with the given unique name, arguments, and options.
func NewManagedNetwork(ctx *pulumi.Context,
	name string, args *ManagedNetworkArgs, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ManagedNetworkName == nil {
		return nil, errors.New("missing required argument 'ManagedNetworkName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagedNetworkArgs{}
	}
	var resource ManagedNetwork
	err := ctx.RegisterResource("azure-nextgen:managednetwork/v20190601preview:ManagedNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedNetwork gets an existing ManagedNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkState, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	var resource ManagedNetwork
	err := ctx.ReadResource("azure-nextgen:managednetwork/v20190601preview:ManagedNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedNetwork resources.
type managedNetworkState struct {
	// The collection of groups and policies concerned with connectivity
	Connectivity *ConnectivityCollectionResponse `pulumi:"connectivity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope *ScopeResponse `pulumi:"scope"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ManagedNetworkState struct {
	// The collection of groups and policies concerned with connectivity
	Connectivity ConnectivityCollectionResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState pulumi.StringPtrInput
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope ScopeResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ManagedNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkState)(nil)).Elem()
}

type managedNetworkArgs struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope *Scope `pulumi:"scope"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedNetwork resource.
type ManagedNetworkArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope ScopePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ManagedNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkArgs)(nil)).Elem()
}

type ManagedNetworkInput interface {
	pulumi.Input

	ToManagedNetworkOutput() ManagedNetworkOutput
	ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput
}

func (ManagedNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetwork)(nil)).Elem()
}

func (i ManagedNetwork) ToManagedNetworkOutput() ManagedNetworkOutput {
	return i.ToManagedNetworkOutputWithContext(context.Background())
}

func (i ManagedNetwork) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkOutput)
}

type ManagedNetworkOutput struct {
	*pulumi.OutputState
}

func (ManagedNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkOutput)(nil)).Elem()
}

func (o ManagedNetworkOutput) ToManagedNetworkOutput() ManagedNetworkOutput {
	return o
}

func (o ManagedNetworkOutput) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkOutput{})
}
