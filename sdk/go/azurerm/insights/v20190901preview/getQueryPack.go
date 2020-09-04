// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupQueryPack(ctx *pulumi.Context, args *LookupQueryPackArgs, opts ...pulumi.InvokeOption) (*LookupQueryPackResult, error) {
	var rv LookupQueryPackResult
	err := ctx.Invoke("azurerm:insights/v20190901preview:getQueryPack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryPackArgs struct {
	// The name of the Log Analytics QueryPack resource.
	QueryPackName string `pulumi:"queryPackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Log Analytics QueryPack definition.
type LookupQueryPackResult struct {
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState string `pulumi:"provisioningState"`
	// The unique ID of your application. This field cannot be changed.
	QueryPackId string `pulumi:"queryPackId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
	TimeCreated string `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
	TimeModified string `pulumi:"timeModified"`
	// Azure resource type
	Type string `pulumi:"type"`
}
