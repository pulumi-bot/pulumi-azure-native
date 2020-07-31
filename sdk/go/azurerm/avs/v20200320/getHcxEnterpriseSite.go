// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200320

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHcxEnterpriseSite(ctx *pulumi.Context, args *LookupHcxEnterpriseSiteArgs, opts ...pulumi.InvokeOption) (*LookupHcxEnterpriseSiteResult, error) {
	var rv LookupHcxEnterpriseSiteResult
	err := ctx.Invoke("azurerm:avs/v20200320:getHcxEnterpriseSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHcxEnterpriseSiteArgs struct {
	// Name of the HCX Enterprise Site in the private cloud
	Name string `pulumi:"name"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An HCX Enterprise Site resource
type LookupHcxEnterpriseSiteResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of an HCX Enterprise Site resource
	Properties HcxEnterpriseSitePropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}