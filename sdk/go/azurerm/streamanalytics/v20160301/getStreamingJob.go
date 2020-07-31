// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStreamingJob(ctx *pulumi.Context, args *LookupStreamingJobArgs, opts ...pulumi.InvokeOption) (*LookupStreamingJobResult, error) {
	var rv LookupStreamingJobResult
	err := ctx.Invoke("azurerm:streamanalytics/v20160301:getStreamingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingJobArgs struct {
	// The name of the streaming job.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A streaming job object, containing all information associated with the named streaming job.
type LookupStreamingJobResult struct {
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