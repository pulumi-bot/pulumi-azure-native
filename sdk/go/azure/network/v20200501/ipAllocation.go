// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// IpAllocation resource.
type IpAllocation struct {
	pulumi.CustomResourceState

	// IpAllocation tags.
	AllocationTags pulumi.StringMapOutput `pulumi:"allocationTags"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The IPAM allocation ID.
	IpamAllocationId pulumi.StringPtrOutput `pulumi:"ipamAllocationId"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The address prefix for the IpAllocation.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType pulumi.StringPtrOutput `pulumi:"prefixType"`
	// The Subnet that using the prefix of this IpAllocation resource.
	Subnet SubResourceResponseOutput `pulumi:"subnet"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualNetwork that using the prefix of this IpAllocation resource.
	VirtualNetwork SubResourceResponseOutput `pulumi:"virtualNetwork"`
}

// NewIpAllocation registers a new resource with the given unique name, arguments, and options.
func NewIpAllocation(ctx *pulumi.Context,
	name string, args *IpAllocationArgs, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAllocationName == nil {
		return nil, errors.New("invalid value for required argument 'IpAllocationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:IpAllocation"),
		},
	})
	opts = append(opts, aliases)
	var resource IpAllocation
	err := ctx.RegisterResource("azure-nextgen:network/v20200501:IpAllocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAllocation gets an existing IpAllocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAllocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAllocationState, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	var resource IpAllocation
	err := ctx.ReadResource("azure-nextgen:network/v20200501:IpAllocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAllocation resources.
type ipAllocationState struct {
	// IpAllocation tags.
	AllocationTags map[string]string `pulumi:"allocationTags"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The IPAM allocation ID.
	IpamAllocationId *string `pulumi:"ipamAllocationId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The address prefix for the IpAllocation.
	Prefix *string `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength *int `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType *string `pulumi:"prefixType"`
	// The Subnet that using the prefix of this IpAllocation resource.
	Subnet *SubResourceResponse `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The VirtualNetwork that using the prefix of this IpAllocation resource.
	VirtualNetwork *SubResourceResponse `pulumi:"virtualNetwork"`
}

type IpAllocationState struct {
	// IpAllocation tags.
	AllocationTags pulumi.StringMapInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The IPAM allocation ID.
	IpamAllocationId pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The address prefix for the IpAllocation.
	Prefix pulumi.StringPtrInput
	// The address prefix length for the IpAllocation.
	PrefixLength pulumi.IntPtrInput
	// The address prefix Type for the IpAllocation.
	PrefixType pulumi.StringPtrInput
	// The Subnet that using the prefix of this IpAllocation resource.
	Subnet SubResourceResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The VirtualNetwork that using the prefix of this IpAllocation resource.
	VirtualNetwork SubResourceResponsePtrInput
}

func (IpAllocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationState)(nil)).Elem()
}

type ipAllocationArgs struct {
	// IpAllocation tags.
	AllocationTags map[string]string `pulumi:"allocationTags"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the IpAllocation.
	IpAllocationName string `pulumi:"ipAllocationName"`
	// The IPAM allocation ID.
	IpamAllocationId *string `pulumi:"ipamAllocationId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The address prefix for the IpAllocation.
	Prefix *string `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength *int `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType *string `pulumi:"prefixType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type for the IpAllocation.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IpAllocation resource.
type IpAllocationArgs struct {
	// IpAllocation tags.
	AllocationTags pulumi.StringMapInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the IpAllocation.
	IpAllocationName pulumi.StringInput
	// The IPAM allocation ID.
	IpamAllocationId pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The address prefix for the IpAllocation.
	Prefix pulumi.StringPtrInput
	// The address prefix length for the IpAllocation.
	PrefixLength pulumi.IntPtrInput
	// The address prefix Type for the IpAllocation.
	PrefixType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type for the IpAllocation.
	Type pulumi.StringPtrInput
}

func (IpAllocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationArgs)(nil)).Elem()
}

type IpAllocationInput interface {
	pulumi.Input

	ToIpAllocationOutput() IpAllocationOutput
	ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput
}

func (*IpAllocation) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocation)(nil))
}

func (i *IpAllocation) ToIpAllocationOutput() IpAllocationOutput {
	return i.ToIpAllocationOutputWithContext(context.Background())
}

func (i *IpAllocation) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAllocationOutput)
}

type IpAllocationOutput struct {
	*pulumi.OutputState
}

func (IpAllocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocation)(nil))
}

func (o IpAllocationOutput) ToIpAllocationOutput() IpAllocationOutput {
	return o
}

func (o IpAllocationOutput) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IpAllocationOutput{})
}
