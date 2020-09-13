// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
type Job struct {
	pulumi.CustomResourceState

	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData pulumi.StringMapOutput `pulumi:"correlationData"`
	// The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created pulumi.StringOutput `pulumi:"created"`
	// Optional customer supplied description of the Job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The inputs for the Job.
	Input pulumi.AnyOutput `pulumi:"input"`
	// The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outputs for the Job.
	Outputs JobOutputAssetResponseArrayOutput `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The current state of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Input == nil {
		return nil, errors.New("missing required argument 'Input'")
	}
	if args == nil || args.JobName == nil {
		return nil, errors.New("missing required argument 'JobName'")
	}
	if args == nil || args.Outputs == nil {
		return nil, errors.New("missing required argument 'Outputs'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TransformName == nil {
		return nil, errors.New("missing required argument 'TransformName'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:media/latest:Job"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180330preview:Job"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180701:Job"),
		},
		{
			Type: pulumi.String("azurerm:media/v20200501:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azurerm:media/v20180601preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azurerm:media/v20180601preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData map[string]string `pulumi:"correlationData"`
	// The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created *string `pulumi:"created"`
	// Optional customer supplied description of the Job.
	Description *string `pulumi:"description"`
	// The inputs for the Job.
	Input interface{} `pulumi:"input"`
	// The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified *string `pulumi:"lastModified"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The outputs for the Job.
	Outputs []JobOutputAssetResponse `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority *string `pulumi:"priority"`
	// The current state of the job.
	State *string `pulumi:"state"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type JobState struct {
	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData pulumi.StringMapInput
	// The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created pulumi.StringPtrInput
	// Optional customer supplied description of the Job.
	Description pulumi.StringPtrInput
	// The inputs for the Job.
	Input pulumi.Input
	// The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The outputs for the Job.
	Outputs JobOutputAssetResponseArrayInput
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority pulumi.StringPtrInput
	// The current state of the job.
	State pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData map[string]string `pulumi:"correlationData"`
	// Optional customer supplied description of the Job.
	Description *string `pulumi:"description"`
	// The inputs for the Job.
	Input interface{} `pulumi:"input"`
	// The Job name.
	JobName string `pulumi:"jobName"`
	// The outputs for the Job.
	Outputs []JobOutputAsset `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority *string `pulumi:"priority"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName string `pulumi:"transformName"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData pulumi.StringMapInput
	// Optional customer supplied description of the Job.
	Description pulumi.StringPtrInput
	// The inputs for the Job.
	Input pulumi.Input
	// The Job name.
	JobName pulumi.StringInput
	// The outputs for the Job.
	Outputs JobOutputAssetArrayInput
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Transform name.
	TransformName pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}
