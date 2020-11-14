// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tag Contract details.
type TagByProduct struct {
	pulumi.CustomResourceState

	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagByProduct registers a new resource with the given unique name, arguments, and options.
func NewTagByProduct(ctx *pulumi.Context,
	name string, args *TagByProductArgs, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.TagId == nil {
		return nil, errors.New("missing required argument 'TagId'")
	}
	if args == nil {
		args = &TagByProductArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/latest:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:TagByProduct"),
		},
	})
	opts = append(opts, aliases)
	var resource TagByProduct
	err := ctx.RegisterResource("azure-nextgen:apimanagement/v20180601preview:TagByProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagByProduct gets an existing TagByProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagByProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByProductState, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
	var resource TagByProduct
	err := ctx.ReadResource("azure-nextgen:apimanagement/v20180601preview:TagByProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagByProduct resources.
type tagByProductState struct {
	// Tag name.
	DisplayName *string `pulumi:"displayName"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type TagByProductState struct {
	// Tag name.
	DisplayName pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (TagByProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductState)(nil)).Elem()
}

type tagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagByProduct resource.
type TagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
}

func (TagByProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductArgs)(nil)).Elem()
}

type TagByProductInput interface {
	pulumi.Input

	ToTagByProductOutput() TagByProductOutput
	ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput
}

func (TagByProduct) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByProduct)(nil)).Elem()
}

func (i TagByProduct) ToTagByProductOutput() TagByProductOutput {
	return i.ToTagByProductOutputWithContext(context.Background())
}

func (i TagByProduct) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByProductOutput)
}

type TagByProductOutput struct {
	*pulumi.OutputState
}

func (TagByProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByProductOutput)(nil)).Elem()
}

func (o TagByProductOutput) ToTagByProductOutput() TagByProductOutput {
	return o
}

func (o TagByProductOutput) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagByProductOutput{})
}
