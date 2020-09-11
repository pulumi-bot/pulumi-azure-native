// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	var rv LookupDashboardResult
	err := ctx.Invoke("azurerm:portal/v20200901preview:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDashboardArgs struct {
	// The name of the dashboard.
	DashboardName string `pulumi:"dashboardName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The shared dashboard resource definition.
type LookupDashboardResult struct {
	// The dashboard lenses.
	Lenses []map[string]interface{} `pulumi:"lenses"`
	// Resource location
	Location string `pulumi:"location"`
	// The dashboard metadata.
	Metadata map[string]map[string]interface{} `pulumi:"metadata"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}
