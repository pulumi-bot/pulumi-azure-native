// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListServiceNamedValueValue(ctx *pulumi.Context, args *ListServiceNamedValueValueArgs, opts ...pulumi.InvokeOption) (*ListServiceNamedValueValueResult, error) {
	var rv ListServiceNamedValueValueResult
	err := ctx.Invoke("azurerm:apimanagement:listServiceNamedValueValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceNamedValueValueArgs struct {
	// Identifier of the NamedValue.
	NamedValueId string `pulumi:"namedValueId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
type ListServiceNamedValueValueResult struct {
	// This is secret value of the NamedValue entity.
	Value *string `pulumi:"value"`
}