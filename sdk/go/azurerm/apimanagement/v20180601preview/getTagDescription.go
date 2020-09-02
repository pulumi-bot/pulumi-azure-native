// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTagDescription(ctx *pulumi.Context, args *LookupTagDescriptionArgs, opts ...pulumi.InvokeOption) (*LookupTagDescriptionResult, error) {
	var rv LookupTagDescriptionResult
	err := ctx.Invoke("azurerm:apimanagement/v20180601preview:getTagDescription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagDescriptionArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Contract details.
type LookupTagDescriptionResult struct {
	// Description of the Tag.
	Description *string `pulumi:"description"`
	// Tag name.
	DisplayName *string `pulumi:"displayName"`
	// Description of the external resources describing the tag.
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl *string `pulumi:"externalDocsUrl"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
