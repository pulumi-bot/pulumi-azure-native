// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account map.
type IntegrationAccountMap struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The content.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponseOutput `pulumi:"contentLink"`
	// The content type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The map type.
	MapType pulumi.StringOutput `pulumi:"mapType"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput `pulumi:"parametersSchema"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountMap registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountMap(ctx *pulumi.Context,
	name string, args *IntegrationAccountMapArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.MapName == nil {
		return nil, errors.New("missing required argument 'MapName'")
	}
	if args == nil || args.MapType == nil {
		return nil, errors.New("missing required argument 'MapType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IntegrationAccountMapArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountMap"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountMap
	err := ctx.RegisterResource("azure-nextgen:logic/latest:IntegrationAccountMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountMap gets an existing IntegrationAccountMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountMapState, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	var resource IntegrationAccountMap
	err := ctx.ReadResource("azure-nextgen:logic/latest:IntegrationAccountMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountMap resources.
type integrationAccountMapState struct {
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The content.
	Content *string `pulumi:"content"`
	// The content link.
	ContentLink *ContentLinkResponse `pulumi:"contentLink"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The map type.
	MapType *string `pulumi:"mapType"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesResponseParametersSchema `pulumi:"parametersSchema"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountMapState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The content.
	Content pulumi.StringPtrInput
	// The content link.
	ContentLink ContentLinkResponsePtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The map type.
	MapType pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapState)(nil)).Elem()
}

type integrationAccountMapArgs struct {
	// The content.
	Content *string `pulumi:"content"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration account map name.
	MapName string `pulumi:"mapName"`
	// The map type.
	MapType string `pulumi:"mapType"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesParametersSchema `pulumi:"parametersSchema"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountMap resource.
type IntegrationAccountMapArgs struct {
	// The content.
	Content pulumi.StringPtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration account map name.
	MapName pulumi.StringInput
	// The map type.
	MapType pulumi.StringInput
	// The metadata.
	Metadata pulumi.Input
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesParametersSchemaPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapArgs)(nil)).Elem()
}

type IntegrationAccountMapInput interface {
	pulumi.Input

	ToIntegrationAccountMapOutput() IntegrationAccountMapOutput
	ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput
}

func (IntegrationAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMap)(nil)).Elem()
}

func (i IntegrationAccountMap) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return i.ToIntegrationAccountMapOutputWithContext(context.Background())
}

func (i IntegrationAccountMap) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapOutput)
}

type IntegrationAccountMapOutput struct {
	*pulumi.OutputState
}

func (IntegrationAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMapOutput)(nil)).Elem()
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return o
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountMapOutput{})
}
