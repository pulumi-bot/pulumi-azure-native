// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azurerm:deploymentmanager/v20180901preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of the service topology .
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
}

// The resource representation of a service in a service topology.
type LookupServiceResult struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Azure location to which the resources in the service belong to or should be deployed to.
	TargetLocation string `pulumi:"targetLocation"`
	// The subscription to which the resources in the service belong to or should be deployed to.
	TargetSubscriptionId string `pulumi:"targetSubscriptionId"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
