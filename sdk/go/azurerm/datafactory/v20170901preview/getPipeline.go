// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("azurerm:datafactory/v20170901preview:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The pipeline name.
	PipelineName string `pulumi:"pipelineName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Pipeline resource type.
type LookupPipelineResult struct {
	// List of activities in pipeline.
	Activities []ActivityResponse `pulumi:"activities"`
	// List of tags that can be used for describing the Pipeline.
	Annotations []map[string]interface{} `pulumi:"annotations"`
	// The max number of concurrent runs for the pipeline.
	Concurrency *int `pulumi:"concurrency"`
	// The description of the pipeline.
	Description *string `pulumi:"description"`
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource name.
	Name string `pulumi:"name"`
	// List of parameters for pipeline.
	Parameters map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	// The resource type.
	Type string `pulumi:"type"`
}
