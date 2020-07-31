// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The IoT Central application.
type App struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ARM resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of an IoT Central application.
	Properties AppPropertiesResponseOutput `pulumi:"properties"`
	// A valid instance SKU.
	Sku AppSkuInfoResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &AppArgs{}
	}
	var resource App
	err := ctx.RegisterResource("azurerm:iotcentral/v20180901:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("azurerm:iotcentral/v20180901:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// The ARM resource name.
	Name *string `pulumi:"name"`
	// The common properties of an IoT Central application.
	Properties *AppPropertiesResponse `pulumi:"properties"`
	// A valid instance SKU.
	Sku *AppSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type AppState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// The ARM resource name.
	Name pulumi.StringPtrInput
	// The common properties of an IoT Central application.
	Properties AppPropertiesResponsePtrInput
	// A valid instance SKU.
	Sku AppSkuInfoResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// The resource location.
	Location string `pulumi:"location"`
	// The ARM resource name of the IoT Central application.
	Name string `pulumi:"name"`
	// The common properties of an IoT Central application.
	Properties *AppProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A valid instance SKU.
	Sku AppSkuInfo `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// The resource location.
	Location pulumi.StringInput
	// The ARM resource name of the IoT Central application.
	Name pulumi.StringInput
	// The common properties of an IoT Central application.
	Properties AppPropertiesPtrInput
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName pulumi.StringInput
	// A valid instance SKU.
	Sku AppSkuInfoInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}