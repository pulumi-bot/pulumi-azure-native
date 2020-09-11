// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The description of the provisioning service.
type IotDpsResource struct {
	pulumi.CustomResourceState

	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionResponseOutput `pulumi:"properties"`
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotDpsResource registers a new resource with the given unique name, arguments, and options.
func NewIotDpsResource(ctx *pulumi.Context,
	name string, args *IotDpsResourceArgs, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ProvisioningServiceName == nil {
		return nil, errors.New("missing required argument 'ProvisioningServiceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &IotDpsResourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:devices/latest:IotDpsResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20170821preview:IotDpsResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20171115:IotDpsResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20180122:IotDpsResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200101:IotDpsResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200901preview:IotDpsResource"),
		},
	})
	opts = append(opts, aliases)
	var resource IotDpsResource
	err := ctx.RegisterResource("azurerm:devices/v20200301:IotDpsResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotDpsResource gets an existing IotDpsResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotDpsResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDpsResourceState, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	var resource IotDpsResource
	err := ctx.ReadResource("azurerm:devices/v20200301:IotDpsResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotDpsResource resources.
type iotDpsResourceState struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Service specific properties for a provisioning service
	Properties *IotDpsPropertiesDescriptionResponse `pulumi:"properties"`
	// Sku info for a provisioning Service.
	Sku *IotDpsSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IotDpsResourceState struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionResponsePtrInput
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IotDpsResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceState)(nil)).Elem()
}

type iotDpsResourceArgs struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource location.
	Location string `pulumi:"location"`
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescription `pulumi:"properties"`
	// Name of provisioning service to create or update.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// Resource group identifier.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfo `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IotDpsResource resource.
type IotDpsResourceArgs struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringInput
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionInput
	// Name of provisioning service to create or update.
	ProvisioningServiceName pulumi.StringInput
	// Resource group identifier.
	ResourceGroupName pulumi.StringInput
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IotDpsResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceArgs)(nil)).Elem()
}
