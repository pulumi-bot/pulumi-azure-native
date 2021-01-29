// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IntegrationAccountMap struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The content.
	Content pulumi.AnyOutput `pulumi:"content"`
	// The content link.
	ContentLink IntegrationAccountContentLinkResponseOutput `pulumi:"contentLink"`
	// The content type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The map type.
	MapType pulumi.StringPtrOutput `pulumi:"mapType"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewIntegrationAccountMap registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountMap(ctx *pulumi.Context,
	name string, args *IntegrationAccountMapArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.MapName == nil {
		return nil, errors.New("invalid value for required argument 'MapName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/latest:IntegrationAccountMap"),
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
	err := ctx.RegisterResource("azure-nextgen:logic/v20150801preview:IntegrationAccountMap", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:logic/v20150801preview:IntegrationAccountMap", name, id, state, &resource, opts...)
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
	Content interface{} `pulumi:"content"`
	// The content link.
	ContentLink *IntegrationAccountContentLinkResponse `pulumi:"contentLink"`
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
	// The resource name.
	Name *string `pulumi:"name"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountMapState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The content.
	Content pulumi.Input
	// The content link.
	ContentLink IntegrationAccountContentLinkResponsePtrInput
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
	// The resource name.
	Name pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapState)(nil)).Elem()
}

type integrationAccountMapArgs struct {
	// The content.
	Content interface{} `pulumi:"content"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The resource id.
	Id *string `pulumi:"id"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration account map name.
	MapName string `pulumi:"mapName"`
	// The map type.
	MapType *string `pulumi:"mapType"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IntegrationAccountMap resource.
type IntegrationAccountMapArgs struct {
	// The content.
	Content pulumi.Input
	// The content type.
	ContentType pulumi.StringPtrInput
	// The resource id.
	Id pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration account map name.
	MapName pulumi.StringInput
	// The map type.
	MapType *MapType
	// The metadata.
	Metadata pulumi.Input
	// The resource name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapArgs)(nil)).Elem()
}

type IntegrationAccountMapInput interface {
	pulumi.Input

	ToIntegrationAccountMapOutput() IntegrationAccountMapOutput
	ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput
}

func (*IntegrationAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMap)(nil))
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return i.ToIntegrationAccountMapOutputWithContext(context.Background())
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapOutput)
}

type IntegrationAccountMapOutput struct {
	*pulumi.OutputState
}

func (IntegrationAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMap)(nil))
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
