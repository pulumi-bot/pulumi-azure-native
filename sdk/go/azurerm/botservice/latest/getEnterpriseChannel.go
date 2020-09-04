// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEnterpriseChannel(ctx *pulumi.Context, args *LookupEnterpriseChannelArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseChannelResult, error) {
	var rv LookupEnterpriseChannelResult
	err := ctx.Invoke("azurerm:botservice/latest:getEnterpriseChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseChannelArgs struct {
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
}

// Enterprise Channel resource definition
type LookupEnterpriseChannelResult struct {
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to an Enterprise Channel resource.
	Properties EnterpriseChannelPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}
