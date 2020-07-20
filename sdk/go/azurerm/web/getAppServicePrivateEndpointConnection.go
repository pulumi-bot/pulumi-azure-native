// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServicePrivateEndpointConnection(ctx *pulumi.Context, args *LookupAppServicePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePrivateEndpointConnectionResult, error) {
	var rv LookupAppServicePrivateEndpointConnectionResult
	err := ctx.Invoke("azurerm:web:getAppServicePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServicePrivateEndpointConnectionArgs struct {
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private Endpoint Connection ARM resource.
type LookupAppServicePrivateEndpointConnectionResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Core resource properties
	Properties RemotePrivateEndpointConnectionResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
