// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The description of the IoTSpaces service.
type IoTSpace struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of a IoTSpaces service.
	Properties IoTSpacesPropertiesResponseOutput `pulumi:"properties"`
	// A valid instance SKU.
	Sku IoTSpacesSkuInfoResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIoTSpace registers a new resource with the given unique name, arguments, and options.
func NewIoTSpace(ctx *pulumi.Context,
	name string, args *IoTSpaceArgs, opts ...pulumi.ResourceOption) (*IoTSpace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource IoTSpace
	err := ctx.RegisterResource("azure-nextgen:iotspaces/v20171001preview:IoTSpace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIoTSpace gets an existing IoTSpace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTSpaceState, opts ...pulumi.ResourceOption) (*IoTSpace, error) {
	var resource IoTSpace
	err := ctx.ReadResource("azure-nextgen:iotspaces/v20171001preview:IoTSpace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IoTSpace resources.
type ioTSpaceState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The common properties of a IoTSpaces service.
	Properties *IoTSpacesPropertiesResponse `pulumi:"properties"`
	// A valid instance SKU.
	Sku *IoTSpacesSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IoTSpaceState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The common properties of a IoTSpaces service.
	Properties IoTSpacesPropertiesResponsePtrInput
	// A valid instance SKU.
	Sku IoTSpacesSkuInfoResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IoTSpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTSpaceState)(nil)).Elem()
}

type ioTSpaceArgs struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// The common properties of a IoTSpaces service.
	Properties *IoTSpacesProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoTSpaces instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoTSpaces instance.
	ResourceName string `pulumi:"resourceName"`
	// A valid instance SKU.
	Sku IoTSpacesSkuInfo `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IoTSpace resource.
type IoTSpaceArgs struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// The common properties of a IoTSpaces service.
	Properties IoTSpacesPropertiesPtrInput
	// The name of the resource group that contains the IoTSpaces instance.
	ResourceGroupName pulumi.StringInput
	// The name of the IoTSpaces instance.
	ResourceName pulumi.StringInput
	// A valid instance SKU.
	Sku IoTSpacesSkuInfoInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IoTSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTSpaceArgs)(nil)).Elem()
}

type IoTSpaceInput interface {
	pulumi.Input

	ToIoTSpaceOutput() IoTSpaceOutput
	ToIoTSpaceOutputWithContext(ctx context.Context) IoTSpaceOutput
}

func (*IoTSpace) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTSpace)(nil))
}

func (i *IoTSpace) ToIoTSpaceOutput() IoTSpaceOutput {
	return i.ToIoTSpaceOutputWithContext(context.Background())
}

func (i *IoTSpace) ToIoTSpaceOutputWithContext(ctx context.Context) IoTSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTSpaceOutput)
}

type IoTSpaceOutput struct {
	*pulumi.OutputState
}

func (IoTSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTSpace)(nil))
}

func (o IoTSpaceOutput) ToIoTSpaceOutput() IoTSpaceOutput {
	return o
}

func (o IoTSpaceOutput) ToIoTSpaceOutputWithContext(ctx context.Context) IoTSpaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IoTSpaceOutput{})
}
