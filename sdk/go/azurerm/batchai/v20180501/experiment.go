// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Experiment information.
type Experiment struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties associated with the experiment.
	Properties ExperimentPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExperiment registers a new resource with the given unique name, arguments, and options.
func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &ExperimentArgs{}
	}
	var resource Experiment
	err := ctx.RegisterResource("azurerm:batchai/v20180501:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperiment gets an existing Experiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("azurerm:batchai/v20180501:Experiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Experiment resources.
type experimentState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties associated with the experiment.
	Properties *ExperimentPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ExperimentState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties associated with the experiment.
	Properties ExperimentPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentState)(nil)).Elem()
}

type experimentArgs struct {
	// The name of the experiment. Experiment names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Experiment resource.
type ExperimentArgs struct {
	// The name of the experiment. Experiment names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName pulumi.StringInput
}

func (ExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentArgs)(nil)).Elem()
}