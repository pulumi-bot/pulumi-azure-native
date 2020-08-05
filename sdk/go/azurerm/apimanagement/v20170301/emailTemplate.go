// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Email Template details.
type EmailTemplate struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Email Template entity contract properties.
	Properties EmailTemplateContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEmailTemplate registers a new resource with the given unique name, arguments, and options.
func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &EmailTemplateArgs{}
	}
	var resource EmailTemplate
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:EmailTemplate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailTemplate resources.
type emailTemplateState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Email Template entity contract properties.
	Properties *EmailTemplateContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type EmailTemplateState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Email Template entity contract properties.
	Properties EmailTemplateContractPropertiesResponsePtrInput
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
	// Email Template Name Identifier.
	Name string `pulumi:"name"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractProperties `pulumi:"parameters"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subject of the Template.
	Subject *string `pulumi:"subject"`
	// Title of the Template.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a EmailTemplate resource.
type EmailTemplateArgs struct {
	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringPtrInput
	// Description of the Email Template.
	Description pulumi.StringPtrInput
	// Email Template Name Identifier.
	Name pulumi.StringInput
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Subject of the Template.
	Subject pulumi.StringPtrInput
	// Title of the Template.
	Title pulumi.StringPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}
