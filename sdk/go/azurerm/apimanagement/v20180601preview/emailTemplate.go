// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Email Template details.
type EmailTemplate struct {
	pulumi.CustomResourceState

	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringOutput `pulumi:"body"`
	// Description of the Email Template.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the template is the default template provided by Api Management or has been edited.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesResponseArrayOutput `pulumi:"parameters"`
	// Subject of the Template.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Title of the Template.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEmailTemplate registers a new resource with the given unique name, arguments, and options.
func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.TemplateName == nil {
		return nil, errors.New("missing required argument 'TemplateName'")
	}
	if args == nil {
		args = &EmailTemplateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:EmailTemplate"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:EmailTemplate"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:EmailTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource EmailTemplate
	err := ctx.RegisterResource("azurerm:apimanagement/v20180601preview:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailTemplate gets an existing EmailTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("azurerm:apimanagement/v20180601preview:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailTemplate resources.
type emailTemplateState struct {
	// Email Template Body. This should be a valid XDocument
	Body *string `pulumi:"body"`
	// Description of the Email Template.
	Description *string `pulumi:"description"`
	// Whether the template is the default template provided by Api Management or has been edited.
	IsDefault *bool `pulumi:"isDefault"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractPropertiesResponse `pulumi:"parameters"`
	// Subject of the Template.
	Subject *string `pulumi:"subject"`
	// Title of the Template.
	Title *string `pulumi:"title"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type EmailTemplateState struct {
	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringPtrInput
	// Description of the Email Template.
	Description pulumi.StringPtrInput
	// Whether the template is the default template provided by Api Management or has been edited.
	IsDefault pulumi.BoolPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesResponseArrayInput
	// Subject of the Template.
	Subject pulumi.StringPtrInput
	// Title of the Template.
	Title pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (EmailTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateState)(nil)).Elem()
}

type emailTemplateArgs struct {
	// Email Template Body. This should be a valid XDocument
	Body *string `pulumi:"body"`
	// Description of the Email Template.
	Description *string `pulumi:"description"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractProperties `pulumi:"parameters"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subject of the Template.
	Subject *string `pulumi:"subject"`
	// Email Template Name Identifier.
	TemplateName string `pulumi:"templateName"`
	// Title of the Template.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a EmailTemplate resource.
type EmailTemplateArgs struct {
	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringPtrInput
	// Description of the Email Template.
	Description pulumi.StringPtrInput
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Subject of the Template.
	Subject pulumi.StringPtrInput
	// Email Template Name Identifier.
	TemplateName pulumi.StringInput
	// Title of the Template.
	Title pulumi.StringPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}
