// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIdentityProvider(ctx *pulumi.Context, args *LookupIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupIdentityProviderResult, error) {
	var rv LookupIdentityProviderResult
	err := ctx.Invoke("azurerm:apimanagement/v20170301:getIdentityProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentityProviderArgs struct {
	// Identity Provider Type identifier.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Identity Provider details.
type LookupIdentityProviderResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// Identity Provider contract properties.
	Properties IdentityProviderContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}