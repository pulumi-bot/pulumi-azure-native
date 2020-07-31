// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// IoT Security solution configuration and resource information.
type IotSecuritySolution struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Security Solution data
	Properties IoTSecuritySolutionPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotSecuritySolution registers a new resource with the given unique name, arguments, and options.
func NewIotSecuritySolution(ctx *pulumi.Context,
	name string, args *IotSecuritySolutionArgs, opts ...pulumi.ResourceOption) (*IotSecuritySolution, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IotSecuritySolutionArgs{}
	}
	var resource IotSecuritySolution
	err := ctx.RegisterResource("azurerm:security/v20190801:IotSecuritySolution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotSecuritySolution gets an existing IotSecuritySolution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotSecuritySolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotSecuritySolutionState, opts ...pulumi.ResourceOption) (*IotSecuritySolution, error) {
	var resource IotSecuritySolution
	err := ctx.ReadResource("azurerm:security/v20190801:IotSecuritySolution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotSecuritySolution resources.
type iotSecuritySolutionState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Security Solution data
	Properties *IoTSecuritySolutionPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type IotSecuritySolutionState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Security Solution data
	Properties IoTSecuritySolutionPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (IotSecuritySolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotSecuritySolutionState)(nil)).Elem()
}

type iotSecuritySolutionArgs struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the IoT Security solution.
	Name string `pulumi:"name"`
	// Security Solution data
	Properties *IoTSecuritySolutionProperties `pulumi:"properties"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IotSecuritySolution resource.
type IotSecuritySolutionArgs struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the IoT Security solution.
	Name pulumi.StringInput
	// Security Solution data
	Properties IoTSecuritySolutionPropertiesPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (IotSecuritySolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotSecuritySolutionArgs)(nil)).Elem()
}