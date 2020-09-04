// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEmailTemplate(ctx *pulumi.Context, args *LookupEmailTemplateArgs, opts ...pulumi.InvokeOption) (*LookupEmailTemplateResult, error) {
	var rv LookupEmailTemplateResult
	err := ctx.Invoke("azurerm:apimanagement/v20191201preview:getEmailTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailTemplateArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Email Template Name Identifier.
	TemplateName string `pulumi:"templateName"`
}

// Email Template details.
type LookupEmailTemplateResult struct {
	// Email Template Body. This should be a valid XDocument
	Body string `pulumi:"body"`
	// Description of the Email Template.
	Description *string `pulumi:"description"`
	// Whether the template is the default template provided by Api Management or has been edited.
	IsDefault bool `pulumi:"isDefault"`
	// Resource name.
	Name string `pulumi:"name"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractPropertiesResponse `pulumi:"parameters"`
	// Subject of the Template.
	Subject string `pulumi:"subject"`
	// Title of the Template.
	Title *string `pulumi:"title"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
