// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A streaming job object, containing all information associated with the named streaming job.
type Streamingjob struct {
	pulumi.CustomResourceState

	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties StreamingJobPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingjob registers a new resource with the given unique name, arguments, and options.
func NewStreamingjob(ctx *pulumi.Context,
	name string, args *StreamingjobArgs, opts ...pulumi.ResourceOption) (*Streamingjob, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StreamingjobArgs{}
	}
	var resource Streamingjob
	err := ctx.RegisterResource("azurerm:streamanalytics:Streamingjob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingjob gets an existing Streamingjob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingjob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingjobState, opts ...pulumi.ResourceOption) (*Streamingjob, error) {
	var resource Streamingjob
	err := ctx.ReadResource("azurerm:streamanalytics:Streamingjob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Streamingjob resources.
type streamingjobState struct {
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties *StreamingJobPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StreamingjobState struct {
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties StreamingJobPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StreamingjobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobState)(nil)).Elem()
}

type streamingjobArgs struct {
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// The name of the streaming job.
	Name string `pulumi:"name"`
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties *StreamingJobProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Streamingjob resource.
type StreamingjobArgs struct {
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrInput
	// The name of the streaming job.
	Name pulumi.StringInput
	// The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
	Properties StreamingJobPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (StreamingjobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingjobArgs)(nil)).Elem()
}