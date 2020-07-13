// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHubKpi(ctx *pulumi.Context, args *LookupHubKpiArgs, opts ...pulumi.InvokeOption) (*LookupHubKpiResult, error) {
	var rv LookupHubKpiResult
	err := ctx.Invoke("azurerm:customerinsights:getHubKpi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubKpiArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the KPI.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The KPI resource format.
type LookupHubKpiResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// Defines the KPI Threshold limits.
	Properties KpiDefinitionResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}