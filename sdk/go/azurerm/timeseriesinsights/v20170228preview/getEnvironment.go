// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170228preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azurerm:timeseriesinsights/v20170228preview:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource.
type LookupEnvironmentResult struct {
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn string `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId string `pulumi:"dataAccessId"`
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime string `pulumi:"dataRetentionTime"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku *SkuResponse `pulumi:"sku"`
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}
