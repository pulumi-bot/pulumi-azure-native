// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentManagedApiArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentManagedApiResult, error) {
	var rv LookupIntegrationServiceEnvironmentManagedApiResult
	err := ctx.Invoke("azurerm:logic/latest:getIntegrationServiceEnvironmentManagedApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationServiceEnvironmentManagedApiArgs struct {
	// The api name.
	ApiName string `pulumi:"apiName"`
	// The integration service environment name.
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	// The resource group name.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// The managed api definition.
type LookupIntegrationServiceEnvironmentManagedApiResult struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The api resource properties.
	Properties ApiResourcePropertiesResponse `pulumi:"properties"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}