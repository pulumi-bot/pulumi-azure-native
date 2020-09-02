// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azurerm:apimanagement/v20180601preview:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Policy Contract details.
type LookupPolicyResult struct {
	// Format of the policyContent.
	ContentFormat *string `pulumi:"contentFormat"`
	// Resource name.
	Name string `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
