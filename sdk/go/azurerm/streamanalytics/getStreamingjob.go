// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStreamingjob(ctx *pulumi.Context, args *LookupStreamingjobArgs, opts ...pulumi.InvokeOption) (*LookupStreamingjobResult, error) {
	var rv LookupStreamingjobResult
	err := ctx.Invoke("azurerm:streamanalytics:getStreamingjob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingjobArgs struct {
	// The name of the streaming job.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A streaming job object, containing all information associated with the named streaming job.
type LookupStreamingjobResult struct {
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties StreamingJobPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}