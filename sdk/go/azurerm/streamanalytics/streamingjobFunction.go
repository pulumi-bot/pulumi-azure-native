// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A function object, containing all information associated with the named function. All functions are contained under a streaming job.
type StreamingjobFunction struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties that are associated with a function.
	Properties FunctionPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingjobFunction registers a new resource with the given unique name, arguments, and options.
func NewStreamingjobFunction(ctx *pulumi.Context,
	name string, args *StreamingjobFunctionArgs, opts ...pulumi.ResourceOption) (*StreamingjobFunction, error) {
	if args == nil || args.JobName == nil {
		return nil, errors.New("missing required argument 'JobName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StreamingjobFunctionArgs{}
	}
	var resource StreamingjobFunction
	err := ctx.RegisterResource("azurerm:streamanalytics:StreamingjobFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingjobFunction gets an existing StreamingjobFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingjobFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingjobFunctionState, opts ...pulumi.ResourceOption) (*StreamingjobFunction, error) {
	var resource StreamingjobFunction
	err := ctx.ReadResource("azurerm:streamanalytics:StreamingjobFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingjobFunction resources.
type streamingjobFunctionState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties *FunctionPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StreamingjobFunctionState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with a function.
	Properties FunctionPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StreamingjobFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobFunctionState)(nil)).Elem()
}

type streamingjobFunctionArgs struct {
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the function.
	Name string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties *FunctionProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StreamingjobFunction resource.
type StreamingjobFunctionArgs struct {
	// The name of the streaming job.
	JobName pulumi.StringInput
	// The name of the function.
	Name pulumi.StringInput
	// The properties that are associated with a function.
	Properties FunctionPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (StreamingjobFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobFunctionArgs)(nil)).Elem()
}
