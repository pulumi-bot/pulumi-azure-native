// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPeeringService(ctx *pulumi.Context, args *LookupPeeringServiceArgs, opts ...pulumi.InvokeOption) (*LookupPeeringServiceResult, error) {
	var rv LookupPeeringServiceResult
	err := ctx.Invoke("azurerm:peering/v20190801preview:getPeeringService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringServiceArgs struct {
	// The name of the peering.
	PeeringServiceName string `pulumi:"peeringServiceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering Service
type LookupPeeringServiceResult struct {
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation *string `pulumi:"peeringServiceLocation"`
	// The MAPS Provider Name.
	PeeringServiceProvider *string `pulumi:"peeringServiceProvider"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
