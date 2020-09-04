// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Pipeline resource type.
type Pipeline struct {
	pulumi.CustomResourceState

	// List of activities in pipeline.
	Activities ActivityResponseArrayOutput `pulumi:"activities"`
	// List of tags that can be used for describing the Pipeline.
	Annotations pulumi.MapArrayOutput `pulumi:"annotations"`
	// The max number of concurrent runs for the pipeline.
	Concurrency pulumi.IntPtrOutput `pulumi:"concurrency"`
	// The description of the pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
	Folder PipelineResponseFolderPtrOutput `pulumi:"folder"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of parameters for pipeline.
	Parameters ParameterSpecificationResponseMapOutput `pulumi:"parameters"`
	// Dimensions emitted by Pipeline.
	RunDimensions pulumi.MapMapOutput `pulumi:"runDimensions"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of variables for pipeline.
	Variables VariableSpecificationResponseMapOutput `pulumi:"variables"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil || args.FactoryName == nil {
		return nil, errors.New("missing required argument 'FactoryName'")
	}
	if args == nil || args.PipelineName == nil {
		return nil, errors.New("missing required argument 'PipelineName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PipelineArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datafactory/latest:Pipeline"),
		},
		{
			Type: pulumi.String("azurerm:datafactory/v20170901preview:Pipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource Pipeline
	err := ctx.RegisterResource("azurerm:datafactory/v20180601:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azurerm:datafactory/v20180601:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// List of activities in pipeline.
	Activities []ActivityResponse `pulumi:"activities"`
	// List of tags that can be used for describing the Pipeline.
	Annotations []map[string]interface{} `pulumi:"annotations"`
	// The max number of concurrent runs for the pipeline.
	Concurrency *int `pulumi:"concurrency"`
	// The description of the pipeline.
	Description *string `pulumi:"description"`
	// Etag identifies change in the resource.
	Etag *string `pulumi:"etag"`
	// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
	Folder *PipelineResponseFolder `pulumi:"folder"`
	// The resource name.
	Name *string `pulumi:"name"`
	// List of parameters for pipeline.
	Parameters map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	// Dimensions emitted by Pipeline.
	RunDimensions map[string]map[string]interface{} `pulumi:"runDimensions"`
	// The resource type.
	Type *string `pulumi:"type"`
	// List of variables for pipeline.
	Variables map[string]VariableSpecificationResponse `pulumi:"variables"`
}

type PipelineState struct {
	// List of activities in pipeline.
	Activities ActivityResponseArrayInput
	// List of tags that can be used for describing the Pipeline.
	Annotations pulumi.MapArrayInput
	// The max number of concurrent runs for the pipeline.
	Concurrency pulumi.IntPtrInput
	// The description of the pipeline.
	Description pulumi.StringPtrInput
	// Etag identifies change in the resource.
	Etag pulumi.StringPtrInput
	// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
	Folder PipelineResponseFolderPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// List of parameters for pipeline.
	Parameters ParameterSpecificationResponseMapInput
	// Dimensions emitted by Pipeline.
	RunDimensions pulumi.MapMapInput
	// The resource type.
	Type pulumi.StringPtrInput
	// List of variables for pipeline.
	Variables VariableSpecificationResponseMapInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// List of activities in pipeline.
	Activities []Activity `pulumi:"activities"`
	// List of tags that can be used for describing the Pipeline.
	Annotations []map[string]interface{} `pulumi:"annotations"`
	// The max number of concurrent runs for the pipeline.
	Concurrency *int `pulumi:"concurrency"`
	// The description of the pipeline.
	Description *string `pulumi:"description"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
	Folder *PipelineFolder `pulumi:"folder"`
	// List of parameters for pipeline.
	Parameters map[string]ParameterSpecification `pulumi:"parameters"`
	// The pipeline name.
	PipelineName string `pulumi:"pipelineName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Dimensions emitted by Pipeline.
	RunDimensions map[string]map[string]interface{} `pulumi:"runDimensions"`
	// List of variables for pipeline.
	Variables map[string]VariableSpecification `pulumi:"variables"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// List of activities in pipeline.
	Activities ActivityArrayInput
	// List of tags that can be used for describing the Pipeline.
	Annotations pulumi.MapArrayInput
	// The max number of concurrent runs for the pipeline.
	Concurrency pulumi.IntPtrInput
	// The description of the pipeline.
	Description pulumi.StringPtrInput
	// The factory name.
	FactoryName pulumi.StringInput
	// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
	Folder PipelineFolderPtrInput
	// List of parameters for pipeline.
	Parameters ParameterSpecificationMapInput
	// The pipeline name.
	PipelineName pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Dimensions emitted by Pipeline.
	RunDimensions pulumi.MapMapInput
	// List of variables for pipeline.
	Variables VariableSpecificationMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}
