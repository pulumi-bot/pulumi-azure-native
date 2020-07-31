// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupExpressRouteCircuit(ctx *pulumi.Context, args *LookupExpressRouteCircuitArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitResult, error) {
	var rv LookupExpressRouteCircuitResult
	err := ctx.Invoke("azurerm:network/v20191101:getExpressRouteCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitArgs struct {
	// The name of express route circuit.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRouteCircuit resource.
type LookupExpressRouteCircuitResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Properties of the express route circuit.
	Properties ExpressRouteCircuitPropertiesFormatResponse `pulumi:"properties"`
	// The SKU.
	Sku *ExpressRouteCircuitSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}