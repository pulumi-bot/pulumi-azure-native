// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contract details.
type TagDescription struct {
	pulumi.CustomResourceState

	// Description of the Tag.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Tag name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Description of the external resources describing the tag.
	ExternalDocsDescription pulumi.StringPtrOutput `pulumi:"externalDocsDescription"`
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl pulumi.StringPtrOutput `pulumi:"externalDocsUrl"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagDescription registers a new resource with the given unique name, arguments, and options.
func NewTagDescription(ctx *pulumi.Context,
	name string, args *TagDescriptionArgs, opts ...pulumi.ResourceOption) (*TagDescription, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
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
		args = &TagDescriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/latest:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:TagDescription"),
		},
	})
	opts = append(opts, aliases)
	var resource TagDescription
	err := ctx.RegisterResource("azure-nextgen:apimanagement/v20180601preview:TagDescription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagDescription gets an existing TagDescription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagDescription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagDescriptionState, opts ...pulumi.ResourceOption) (*TagDescription, error) {
	var resource TagDescription
	err := ctx.ReadResource("azure-nextgen:apimanagement/v20180601preview:TagDescription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagDescription resources.
type tagDescriptionState struct {
	// Description of the Tag.
	Description *string `pulumi:"description"`
	// Tag name.
	DisplayName *string `pulumi:"displayName"`
	// Description of the external resources describing the tag.
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl *string `pulumi:"externalDocsUrl"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type TagDescriptionState struct {
	// Description of the Tag.
	Description pulumi.StringPtrInput
	// Tag name.
	DisplayName pulumi.StringPtrInput
	// Description of the external resources describing the tag.
	ExternalDocsDescription pulumi.StringPtrInput
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (TagDescriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagDescriptionState)(nil)).Elem()
}

type tagDescriptionArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Description of the Tag.
	Description *string `pulumi:"description"`
	// Description of the external resources describing the tag.
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl *string `pulumi:"externalDocsUrl"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagDescription resource.
type TagDescriptionArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Description of the Tag.
	Description pulumi.StringPtrInput
	// Description of the external resources describing the tag.
	ExternalDocsDescription pulumi.StringPtrInput
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
}

func (TagDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagDescriptionArgs)(nil)).Elem()
}

type TagDescriptionInput interface {
	pulumi.Input

	ToTagDescriptionOutput() TagDescriptionOutput
	ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput
}

func (TagDescription) ElementType() reflect.Type {
	return reflect.TypeOf((*TagDescription)(nil)).Elem()
}

func (i TagDescription) ToTagDescriptionOutput() TagDescriptionOutput {
	return i.ToTagDescriptionOutputWithContext(context.Background())
}

func (i TagDescription) ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagDescriptionOutput)
}

type TagDescriptionOutput struct {
	*pulumi.OutputState
}

func (TagDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagDescriptionOutput)(nil)).Elem()
}

func (o TagDescriptionOutput) ToTagDescriptionOutput() TagDescriptionOutput {
	return o
}

func (o TagDescriptionOutput) ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagDescriptionOutput{})
}
