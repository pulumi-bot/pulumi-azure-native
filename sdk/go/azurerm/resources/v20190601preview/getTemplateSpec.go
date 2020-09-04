// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTemplateSpec(ctx *pulumi.Context, args *LookupTemplateSpecArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecResult, error) {
	var rv LookupTemplateSpecResult
	err := ctx.Invoke("azurerm:resources/v20190601preview:getTemplateSpec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Template Spec.
	TemplateSpecName string `pulumi:"templateSpecName"`
}

// Template Spec object.
type LookupTemplateSpecResult struct {
	// Template Spec description.
	Description *string `pulumi:"description"`
	// Template Spec display name.
	DisplayName *string `pulumi:"displayName"`
	// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
	Location string `pulumi:"location"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of this resource.
	Type string `pulumi:"type"`
}
