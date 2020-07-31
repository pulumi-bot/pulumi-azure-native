// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupOutput(ctx *pulumi.Context, args *LookupOutputArgs, opts ...pulumi.InvokeOption) (*LookupOutputResult, error) {
	var rv LookupOutputResult
	err := ctx.Invoke("azurerm:streamanalytics/v20160301:getOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutputArgs struct {
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the output.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
type LookupOutputResult struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with an output. Required on PUT (CreateOrReplace) requests.
	Properties OutputPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}