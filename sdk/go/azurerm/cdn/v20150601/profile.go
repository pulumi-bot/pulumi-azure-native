// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CDN profile represents the top level resource and the entry point into the CDN API. This allows users to set up a logical grouping of endpoints in addition to creating shared configuration settings and selecting pricing tiers and providers.
type Profile struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning status of the profile.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The SKU (pricing tier) of the CDN profile.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ProfileArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:cdn/latest:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20160402:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20161002:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20170402:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20171012:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190415:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190615:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190615preview:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20191231:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20200331:Profile"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20200415:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azurerm:cdn/v20150601:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azurerm:cdn/v20150601:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Provisioning status of the profile.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState *string `pulumi:"resourceState"`
	// The SKU (pricing tier) of the CDN profile.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ProfileState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Provisioning status of the profile.
	ProvisioningState pulumi.StringPtrInput
	// Resource status of the profile.
	ResourceState pulumi.StringPtrInput
	// The SKU (pricing tier) of the CDN profile.
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Profile location
	Location string `pulumi:"location"`
	// Name of the CDN profile within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Profile SKU
	Sku Sku `pulumi:"sku"`
	// Profile tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Profile location
	Location pulumi.StringInput
	// Name of the CDN profile within the resource group.
	ProfileName pulumi.StringInput
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Profile SKU
	Sku SkuInput
	// Profile tags
	Tags pulumi.StringMapInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}
