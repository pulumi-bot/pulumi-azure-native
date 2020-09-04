// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCognitiveServicesAccount(ctx *pulumi.Context, args *LookupCognitiveServicesAccountArgs, opts ...pulumi.InvokeOption) (*LookupCognitiveServicesAccountResult, error) {
	var rv LookupCognitiveServicesAccountResult
	err := ctx.Invoke("azurerm:cognitiveservices/v20160201preview:getCognitiveServicesAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCognitiveServicesAccountArgs struct {
	// The name of the cognitive services account within the specified resource group. Cognitive Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Cognitive Services Account is an Azure resource representing the provisioned account, its type, location and SKU.
type LookupCognitiveServicesAccountResult struct {
	// Endpoint of the created account
	Endpoint *string `pulumi:"endpoint"`
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Type of cognitive service account.
	Kind *string `pulumi:"kind"`
	// The location of the resource
	Location *string `pulumi:"location"`
	// The name of the created account
	Name *string `pulumi:"name"`
	// Gets the status of the cognitive services account at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the cognitive services account.
	Sku *SkuResponse `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}
