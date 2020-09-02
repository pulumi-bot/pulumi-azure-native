// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupFactory(ctx *pulumi.Context, args *LookupFactoryArgs, opts ...pulumi.InvokeOption) (*LookupFactoryResult, error) {
	var rv LookupFactoryResult
	err := ctx.Invoke("azurerm:datafactory/v20170901preview:getFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFactoryArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Factory resource type.
type LookupFactoryResult struct {
	// Time the factory was created in ISO8601 format.
	CreateTime string `pulumi:"createTime"`
	// Managed service identity of the factory.
	Identity *FactoryIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Factory provisioning state, example Succeeded.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
	// Version of the factory.
	Version string `pulumi:"version"`
	// VSTS repo information of the factory.
	VstsConfiguration *FactoryVSTSConfigurationResponse `pulumi:"vstsConfiguration"`
}
