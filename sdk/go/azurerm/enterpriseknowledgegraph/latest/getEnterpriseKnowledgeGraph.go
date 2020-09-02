// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEnterpriseKnowledgeGraph(ctx *pulumi.Context, args *LookupEnterpriseKnowledgeGraphArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseKnowledgeGraphResult, error) {
	var rv LookupEnterpriseKnowledgeGraphResult
	err := ctx.Invoke("azurerm:enterpriseknowledgegraph/latest:getEnterpriseKnowledgeGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseKnowledgeGraphArgs struct {
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName string `pulumi:"resourceName"`
}

// EnterpriseKnowledgeGraph resource definition
type LookupEnterpriseKnowledgeGraphResult struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}
