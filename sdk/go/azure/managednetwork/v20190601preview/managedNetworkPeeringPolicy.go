// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Managed Network Peering Policy resource
type ManagedNetworkPeeringPolicy struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of a Managed Network Policy
	Properties ManagedNetworkPeeringPolicyPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedNetworkPeeringPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, args *ManagedNetworkPeeringPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	if args == nil || args.ManagedNetworkName == nil {
		return nil, errors.New("missing required argument 'ManagedNetworkName'")
	}
	if args == nil || args.ManagedNetworkPeeringPolicyName == nil {
		return nil, errors.New("missing required argument 'ManagedNetworkPeeringPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagedNetworkPeeringPolicyArgs{}
	}
	var resource ManagedNetworkPeeringPolicy
	err := ctx.RegisterResource("azure-nextgen:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedNetworkPeeringPolicy gets an existing ManagedNetworkPeeringPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkPeeringPolicyState, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	var resource ManagedNetworkPeeringPolicy
	err := ctx.ReadResource("azure-nextgen:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedNetworkPeeringPolicy resources.
type managedNetworkPeeringPolicyState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the properties of a Managed Network Policy
	Properties *ManagedNetworkPeeringPolicyPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ManagedNetworkPeeringPolicyState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the properties of a Managed Network Policy
	Properties ManagedNetworkPeeringPolicyPropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ManagedNetworkPeeringPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyState)(nil)).Elem()
}

type managedNetworkPeeringPolicyArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the Managed Network Peering Policy.
	ManagedNetworkPeeringPolicyName string `pulumi:"managedNetworkPeeringPolicyName"`
	// Gets or sets the properties of a Managed Network Policy
	Properties *ManagedNetworkPeeringPolicyProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedNetworkPeeringPolicy resource.
type ManagedNetworkPeeringPolicyArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringInput
	// The name of the Managed Network Peering Policy.
	ManagedNetworkPeeringPolicyName pulumi.StringInput
	// Gets or sets the properties of a Managed Network Policy
	Properties ManagedNetworkPeeringPolicyPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ManagedNetworkPeeringPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyArgs)(nil)).Elem()
}

type ManagedNetworkPeeringPolicyInput interface {
	pulumi.Input

	ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput
	ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput
}

func (ManagedNetworkPeeringPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicy)(nil)).Elem()
}

func (i ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return i.ToManagedNetworkPeeringPolicyOutputWithContext(context.Background())
}

func (i ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyOutput)
}

type ManagedNetworkPeeringPolicyOutput struct {
	*pulumi.OutputState
}

func (ManagedNetworkPeeringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicyOutput)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyOutput{})
}
