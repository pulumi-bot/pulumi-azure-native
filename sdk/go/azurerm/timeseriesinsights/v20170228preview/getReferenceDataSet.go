// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170228preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupReferenceDataSet(ctx *pulumi.Context, args *LookupReferenceDataSetArgs, opts ...pulumi.InvokeOption) (*LookupReferenceDataSetResult, error) {
	var rv LookupReferenceDataSetResult
	err := ctx.Invoke("azurerm:timeseriesinsights/v20170228preview:getReferenceDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReferenceDataSetArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the Time Series Insights reference data set associated with the specified environment.
	ReferenceDataSetName string `pulumi:"referenceDataSetName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A reference data set provides metadata about the events in an environment. Metadata in the reference data set will be joined with events as they are read from event sources. The metadata that makes up the reference data set is uploaded or modified through the Time Series Insights data plane APIs.
type LookupReferenceDataSetResult struct {
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The list of key properties for the reference data set.
	KeyProperties []ReferenceDataSetKeyPropertyResponse `pulumi:"keyProperties"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}
