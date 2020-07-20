// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type JobCollection struct {
	pulumi.CustomResourceState

	// Gets or sets the storage account location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the job collection resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobCollection registers a new resource with the given unique name, arguments, and options.
func NewJobCollection(ctx *pulumi.Context,
	name string, args *JobCollectionArgs, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &JobCollectionArgs{}
	}
	var resource JobCollection
	err := ctx.RegisterResource("azurerm:scheduler:JobCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobCollection gets an existing JobCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCollectionState, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	var resource JobCollection
	err := ctx.ReadResource("azurerm:scheduler:JobCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobCollection resources.
type jobCollectionState struct {
	// Gets or sets the storage account location.
	Location *string `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties *JobCollectionPropertiesResponse `pulumi:"properties"`
	// Gets or sets the tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the job collection resource type.
	Type *string `pulumi:"type"`
}

type JobCollectionState struct {
	// Gets or sets the storage account location.
	Location pulumi.StringPtrInput
	// Gets or sets the job collection resource name.
	Name pulumi.StringPtrInput
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesResponsePtrInput
	// Gets or sets the tags.
	Tags pulumi.StringMapInput
	// Gets the job collection resource type.
	Type pulumi.StringPtrInput
}

func (JobCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionState)(nil)).Elem()
}

type jobCollectionArgs struct {
	// Gets or sets the storage account location.
	Location *string `pulumi:"location"`
	// The job collection name.
	Name string `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties *JobCollectionProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JobCollection resource.
type JobCollectionArgs struct {
	// Gets or sets the storage account location.
	Location pulumi.StringPtrInput
	// The job collection name.
	Name pulumi.StringInput
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags.
	Tags pulumi.StringMapInput
}

func (JobCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionArgs)(nil)).Elem()
}
