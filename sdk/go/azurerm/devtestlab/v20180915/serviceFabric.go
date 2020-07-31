// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Service Fabric.
type ServiceFabric struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the resource.
	Properties ServiceFabricPropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceFabric registers a new resource with the given unique name, arguments, and options.
func NewServiceFabric(ctx *pulumi.Context,
	name string, args *ServiceFabricArgs, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &ServiceFabricArgs{}
	}
	var resource ServiceFabric
	err := ctx.RegisterResource("azurerm:devtestlab/v20180915:ServiceFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceFabric gets an existing ServiceFabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceFabricState, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	var resource ServiceFabric
	err := ctx.ReadResource("azurerm:devtestlab/v20180915:ServiceFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceFabric resources.
type serviceFabricState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the resource.
	Properties *ServiceFabricPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ServiceFabricState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the resource.
	Properties ServiceFabricPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ServiceFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricState)(nil)).Elem()
}

type serviceFabricArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the service fabric.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties ServiceFabricProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the user profile.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a ServiceFabric resource.
type ServiceFabricArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the service fabric.
	Name pulumi.StringInput
	// The properties of the resource.
	Properties ServiceFabricPropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The name of the user profile.
	UserName pulumi.StringInput
}

func (ServiceFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricArgs)(nil)).Elem()
}