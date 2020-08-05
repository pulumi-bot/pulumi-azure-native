// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VirtualWAN Resource.
type VirtualWAN struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for VirtualWAN
	Properties VirtualWanPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualWAN registers a new resource with the given unique name, arguments, and options.
func NewVirtualWAN(ctx *pulumi.Context,
	name string, args *VirtualWANArgs, opts ...pulumi.ResourceOption) (*VirtualWAN, error) {
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
		args = &VirtualWANArgs{}
	}
	var resource VirtualWAN
	err := ctx.RegisterResource("azurerm:network/v20180601:VirtualWAN", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualWAN gets an existing VirtualWAN resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualWAN(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWANState, opts ...pulumi.ResourceOption) (*VirtualWAN, error) {
	var resource VirtualWAN
	err := ctx.ReadResource("azurerm:network/v20180601:VirtualWAN", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualWAN resources.
type virtualWANState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Parameters for VirtualWAN
	Properties *VirtualWanPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VirtualWANState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Parameters for VirtualWAN
	Properties VirtualWanPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VirtualWANState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWANState)(nil)).Elem()
}

type virtualWANArgs struct {
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the VirtualWAN being created or updated.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group name of the VirtualWan.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualWAN resource.
type VirtualWANArgs struct {
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the VirtualWAN being created or updated.
	Name pulumi.StringInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The resource group name of the VirtualWan.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (VirtualWANArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWANArgs)(nil)).Elem()
}
