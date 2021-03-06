// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190724preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Data controller resource
func LookupDataController(ctx *pulumi.Context, args *LookupDataControllerArgs, opts ...pulumi.InvokeOption) (*LookupDataControllerResult, error) {
	var rv LookupDataControllerResult
	err := ctx.Invoke("azure-native:azuredata/v20190724preview:getDataController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataControllerArgs struct {
	DataControllerName string `pulumi:"dataControllerName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data controller resource
type LookupDataControllerResult struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties from the on premise data controller
	OnPremiseProperty OnPremisePropertyResponse `pulumi:"onPremiseProperty"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
