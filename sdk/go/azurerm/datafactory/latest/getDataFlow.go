// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDataFlow(ctx *pulumi.Context, args *LookupDataFlowArgs, opts ...pulumi.InvokeOption) (*LookupDataFlowResult, error) {
	var rv LookupDataFlowResult
	err := ctx.Invoke("azurerm:datafactory/latest:getDataFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataFlowArgs struct {
	// The data flow name.
	DataFlowName string `pulumi:"dataFlowName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data flow resource type.
type LookupDataFlowResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource name.
	Name string `pulumi:"name"`
	// Data flow properties.
	Properties DataFlowResponse `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}