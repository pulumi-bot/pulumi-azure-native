// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The description of the DigitalTwins service.
type DigitalTwin struct {
	pulumi.CustomResourceState

	// Time when DigitalTwinsInstance was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Api endpoint to work with DigitalTwinsInstance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Time when DigitalTwinsInstance was updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDigitalTwin registers a new resource with the given unique name, arguments, and options.
func NewDigitalTwin(ctx *pulumi.Context,
	name string, args *DigitalTwinArgs, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &DigitalTwinArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/latest:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20200301preview:DigitalTwin"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwin
	err := ctx.RegisterResource("azure-nextgen:digitaltwins/v20201031:DigitalTwin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDigitalTwin gets an existing DigitalTwin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDigitalTwin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinState, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	var resource DigitalTwin
	err := ctx.ReadResource("azure-nextgen:digitaltwins/v20201031:DigitalTwin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DigitalTwin resources.
type digitalTwinState struct {
	// Time when DigitalTwinsInstance was created.
	CreatedTime *string `pulumi:"createdTime"`
	// Api endpoint to work with DigitalTwinsInstance.
	HostName *string `pulumi:"hostName"`
	// Time when DigitalTwinsInstance was updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type DigitalTwinState struct {
	// Time when DigitalTwinsInstance was created.
	CreatedTime pulumi.StringPtrInput
	// Api endpoint to work with DigitalTwinsInstance.
	HostName pulumi.StringPtrInput
	// Time when DigitalTwinsInstance was updated.
	LastUpdatedTime pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (DigitalTwinState) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinState)(nil)).Elem()
}

type digitalTwinArgs struct {
	// The resource location.
	Location string `pulumi:"location"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DigitalTwin resource.
type DigitalTwinArgs struct {
	// The resource location.
	Location pulumi.StringInput
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (DigitalTwinArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinArgs)(nil)).Elem()
}

type DigitalTwinInput interface {
	pulumi.Input

	ToDigitalTwinOutput() DigitalTwinOutput
	ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput
}

func (DigitalTwin) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwin)(nil)).Elem()
}

func (i DigitalTwin) ToDigitalTwinOutput() DigitalTwinOutput {
	return i.ToDigitalTwinOutputWithContext(context.Background())
}

func (i DigitalTwin) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinOutput)
}

type DigitalTwinOutput struct {
	*pulumi.OutputState
}

func (DigitalTwinOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinOutput)(nil)).Elem()
}

func (o DigitalTwinOutput) ToDigitalTwinOutput() DigitalTwinOutput {
	return o
}

func (o DigitalTwinOutput) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinOutput{})
}
