// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account map.
type Map struct {
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

// NewMap registers a new resource with the given unique name, arguments, and options.
func NewMap(ctx *pulumi.Context,
	name string, args *MapArgs, opts ...pulumi.ResourceOption) (*Map, error) {
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
		args = &MapArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/latest:Map"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:Map"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:Map"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:Map"),
		},
	})
	opts = append(opts, aliases)
	var resource Map
	err := ctx.RegisterResource("azure-nextgen:logic/v20160601:Map", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMap gets an existing Map resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MapState, opts ...pulumi.ResourceOption) (*Map, error) {
	var resource Map
	err := ctx.ReadResource("azure-nextgen:logic/v20160601:Map", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Map resources.
type mapState struct {
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

type MapState struct {
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

func (MapState) ElementType() reflect.Type {
	return reflect.TypeOf((*mapState)(nil)).Elem()
}

type mapArgs struct {
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

// The set of arguments for constructing a Map resource.
type MapArgs struct {
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

func (MapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mapArgs)(nil)).Elem()
}

type MapInput interface {
	pulumi.Input

	ToMapOutput() MapOutput
	ToMapOutputWithContext(ctx context.Context) MapOutput
}

func (Map) ElementType() reflect.Type {
	return reflect.TypeOf((*Map)(nil)).Elem()
}

func (i Map) ToMapOutput() MapOutput {
	return i.ToMapOutputWithContext(context.Background())
}

func (i Map) ToMapOutputWithContext(ctx context.Context) MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapOutput)
}

type MapOutput struct {
	*pulumi.OutputState
}

func (MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapOutput)(nil)).Elem()
}

func (o MapOutput) ToMapOutput() MapOutput {
	return o
}

func (o MapOutput) ToMapOutputWithContext(ctx context.Context) MapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MapOutput{})
}
