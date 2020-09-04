// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140801preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupJobCollection(ctx *pulumi.Context, args *LookupJobCollectionArgs, opts ...pulumi.InvokeOption) (*LookupJobCollectionResult, error) {
	var rv LookupJobCollectionResult
	err := ctx.Invoke("azurerm:scheduler/v20140801preview:getJobCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCollectionArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobCollectionResult struct {
	// Gets or sets the storage account location.
	Location *string `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesResponse `pulumi:"properties"`
	// Gets or sets the tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the job collection resource type.
	Type string `pulumi:"type"`
}
