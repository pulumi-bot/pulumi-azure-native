// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes the cloud service.
type CloudService struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Cloud service properties
	Properties CloudServicePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCloudService registers a new resource with the given unique name, arguments, and options.
func NewCloudService(ctx *pulumi.Context,
	name string, args *CloudServiceArgs, opts ...pulumi.ResourceOption) (*CloudService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudServiceName == nil {
		return nil, errors.New("invalid value for required argument 'CloudServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CloudService
	err := ctx.RegisterResource("azure-nextgen:compute/v20201001preview:CloudService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudService gets an existing CloudService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudServiceState, opts ...pulumi.ResourceOption) (*CloudService, error) {
	var resource CloudService
	err := ctx.ReadResource("azure-nextgen:compute/v20201001preview:CloudService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudService resources.
type cloudServiceState struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Cloud service properties
	Properties *CloudServicePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type CloudServiceState struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Cloud service properties
	Properties CloudServicePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (CloudServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServiceState)(nil)).Elem()
}

type cloudServiceArgs struct {
	// Name of the cloud service.
	CloudServiceName string `pulumi:"cloudServiceName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Cloud service properties
	Properties *CloudServiceProperties `pulumi:"properties"`
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CloudService resource.
type CloudServiceArgs struct {
	// Name of the cloud service.
	CloudServiceName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Cloud service properties
	Properties CloudServicePropertiesPtrInput
	// Name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CloudServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServiceArgs)(nil)).Elem()
}

type CloudServiceInput interface {
	pulumi.Input

	ToCloudServiceOutput() CloudServiceOutput
	ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput
}

func (*CloudService) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudService)(nil))
}

func (i *CloudService) ToCloudServiceOutput() CloudServiceOutput {
	return i.ToCloudServiceOutputWithContext(context.Background())
}

func (i *CloudService) ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceOutput)
}

type CloudServiceOutput struct {
	*pulumi.OutputState
}

func (CloudServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudService)(nil))
}

func (o CloudServiceOutput) ToCloudServiceOutput() CloudServiceOutput {
	return o
}

func (o CloudServiceOutput) ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CloudServiceOutput{})
}
