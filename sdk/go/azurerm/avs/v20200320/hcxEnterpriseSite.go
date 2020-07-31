// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200320

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An HCX Enterprise Site resource
type HcxEnterpriseSite struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an HCX Enterprise Site resource
	Properties HcxEnterpriseSitePropertiesResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHcxEnterpriseSite registers a new resource with the given unique name, arguments, and options.
func NewHcxEnterpriseSite(ctx *pulumi.Context,
	name string, args *HcxEnterpriseSiteArgs, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PrivateCloudName == nil {
		return nil, errors.New("missing required argument 'PrivateCloudName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HcxEnterpriseSiteArgs{}
	}
	var resource HcxEnterpriseSite
	err := ctx.RegisterResource("azurerm:avs/v20200320:HcxEnterpriseSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHcxEnterpriseSite gets an existing HcxEnterpriseSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHcxEnterpriseSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HcxEnterpriseSiteState, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
	var resource HcxEnterpriseSite
	err := ctx.ReadResource("azurerm:avs/v20200320:HcxEnterpriseSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HcxEnterpriseSite resources.
type hcxEnterpriseSiteState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The properties of an HCX Enterprise Site resource
	Properties *HcxEnterpriseSitePropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type HcxEnterpriseSiteState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The properties of an HCX Enterprise Site resource
	Properties HcxEnterpriseSitePropertiesResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (HcxEnterpriseSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteState)(nil)).Elem()
}

type hcxEnterpriseSiteArgs struct {
	// Name of the HCX Enterprise Site in the private cloud
	Name string `pulumi:"name"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HcxEnterpriseSite resource.
type HcxEnterpriseSiteArgs struct {
	// Name of the HCX Enterprise Site in the private cloud
	Name pulumi.StringInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (HcxEnterpriseSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteArgs)(nil)).Elem()
}