// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Fault Domain count.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount pulumi.IntPtrOutput `pulumi:"platformUpdateDomainCount"`
	// Sku of the availability set
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The resource status information.
	Statuses InstanceViewStatusResponseArrayOutput `pulumi:"statuses"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines SubResourceResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilitySetName == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilitySetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/latest:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20150615:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20181001:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:AvailabilitySet"),
		},
	})
	opts = append(opts, aliases)
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure-nextgen:compute/v20171201:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure-nextgen:compute/v20171201:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type availabilitySetState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Fault Domain count.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// Sku of the availability set
	Sku *SkuResponse `pulumi:"sku"`
	// The resource status information.
	Statuses []InstanceViewStatusResponse `pulumi:"statuses"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines []SubResourceResponse `pulumi:"virtualMachines"`
}

type AvailabilitySetState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Fault Domain count.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Update Domain count.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// Sku of the availability set
	Sku SkuResponsePtrInput
	// The resource status information.
	Statuses InstanceViewStatusResponseArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// A list of references to all virtual machines in the availability set.
	VirtualMachines SubResourceResponseArrayInput
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Fault Domain count.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku of the availability set
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines []SubResource `pulumi:"virtualMachines"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// Fault Domain count.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Update Domain count.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Sku of the availability set
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// A list of references to all virtual machines in the availability set.
	VirtualMachines SubResourceArrayInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}

type AvailabilitySetInput interface {
	pulumi.Input

	ToAvailabilitySetOutput() AvailabilitySetOutput
	ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput
}

func (*AvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySet)(nil))
}

func (i *AvailabilitySet) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return i.ToAvailabilitySetOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetOutput)
}

type AvailabilitySetOutput struct {
	*pulumi.OutputState
}

func (AvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySet)(nil))
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
}
