// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180122

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The description of the IoT hub.
type IotHubResource struct {
	pulumi.CustomResourceState

	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an IoT hub.
	Properties IotHubPropertiesResponseOutput `pulumi:"properties"`
	// Information about the SKU of the IoT hub.
	Sku IotHubSkuInfoResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotHubResource registers a new resource with the given unique name, arguments, and options.
func NewIotHubResource(ctx *pulumi.Context,
	name string, args *IotHubResourceArgs, opts ...pulumi.ResourceOption) (*IotHubResource, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &IotHubResourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:devices/latest:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20160203:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20170119:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20170701:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20180401:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20181201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190322:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190322preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20191104:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200301:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200401:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200615:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200710preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200801:IotHubResource"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubResource
	err := ctx.RegisterResource("azurerm:devices/v20180122:IotHubResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubResource gets an existing IotHubResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceState, opts ...pulumi.ResourceOption) (*IotHubResource, error) {
	var resource IotHubResource
	err := ctx.ReadResource("azurerm:devices/v20180122:IotHubResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubResource resources.
type iotHubResourceState struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The properties of an IoT hub.
	Properties *IotHubPropertiesResponse `pulumi:"properties"`
	// Information about the SKU of the IoT hub.
	Sku *IotHubSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IotHubResourceState struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The properties of an IoT hub.
	Properties IotHubPropertiesResponsePtrInput
	// Information about the SKU of the IoT hub.
	Sku IotHubSkuInfoResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IotHubResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceState)(nil)).Elem()
}

type iotHubResourceArgs struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource location.
	Location string `pulumi:"location"`
	// The properties of an IoT hub.
	Properties *IotHubProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
	// Information about the SKU of the IoT hub.
	Sku IotHubSkuInfo `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IotHubResource resource.
type IotHubResourceArgs struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringInput
	// The properties of an IoT hub.
	Properties IotHubPropertiesPtrInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
	// Information about the SKU of the IoT hub.
	Sku IotHubSkuInfoInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IotHubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceArgs)(nil)).Elem()
}
