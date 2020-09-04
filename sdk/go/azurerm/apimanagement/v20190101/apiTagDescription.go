// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contract details.
type ApiTagDescription struct {
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

// NewApiTagDescription registers a new resource with the given unique name, arguments, and options.
func NewApiTagDescription(ctx *pulumi.Context,
	name string, args *ApiTagDescriptionArgs, opts ...pulumi.ResourceOption) (*ApiTagDescription, error) {
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
		args = &ApiTagDescriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiTagDescription"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiTagDescription
	err := ctx.RegisterResource("azurerm:apimanagement/v20190101:ApiTagDescription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiTagDescription gets an existing ApiTagDescription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiTagDescription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiTagDescriptionState, opts ...pulumi.ResourceOption) (*ApiTagDescription, error) {
	var resource ApiTagDescription
	err := ctx.ReadResource("azurerm:apimanagement/v20190101:ApiTagDescription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiTagDescription resources.
type apiTagDescriptionState struct {
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

type ApiTagDescriptionState struct {
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

func (ApiTagDescriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTagDescriptionState)(nil)).Elem()
}

type apiTagDescriptionArgs struct {
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

// The set of arguments for constructing a ApiTagDescription resource.
type ApiTagDescriptionArgs struct {
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

func (ApiTagDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTagDescriptionArgs)(nil)).Elem()
}
