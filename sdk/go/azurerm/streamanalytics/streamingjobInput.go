// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An input object, containing all information associated with the named input. All inputs are contained under a streaming job.
type StreamingjobInput struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties InputPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingjobInput registers a new resource with the given unique name, arguments, and options.
func NewStreamingjobInput(ctx *pulumi.Context,
	name string, args *StreamingjobInputArgs, opts ...pulumi.ResourceOption) (*StreamingjobInput, error) {
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
		args = &StreamingjobInputArgs{}
	}
	var resource StreamingjobInput
	err := ctx.RegisterResource("azurerm:streamanalytics:StreamingjobInput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingjobInput gets an existing StreamingjobInput resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingjobInput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingjobInputState, opts ...pulumi.ResourceOption) (*StreamingjobInput, error) {
	var resource StreamingjobInput
	err := ctx.ReadResource("azurerm:streamanalytics:StreamingjobInput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingjobInput resources.
type streamingjobInputState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties *InputPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StreamingjobInputState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties InputPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StreamingjobInputState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobInputState)(nil)).Elem()
}

type streamingjobInputArgs struct {
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the input.
	Name string `pulumi:"name"`
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties *InputProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StreamingjobInput resource.
type StreamingjobInputArgs struct {
	// The name of the streaming job.
	JobName pulumi.StringInput
	// The name of the input.
	Name pulumi.StringInput
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties InputPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (StreamingjobInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobInputArgs)(nil)).Elem()
}
