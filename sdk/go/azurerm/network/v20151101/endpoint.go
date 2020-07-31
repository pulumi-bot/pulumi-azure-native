// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Traffic Manager endpoint.
type Endpoint struct {
	pulumi.CustomResourceState

	// Gets or sets the name of the Traffic Manager endpoint.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Class representing a Traffic Manager endpoint properties.
	Properties EndpointPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the endpoint type of the Traffic Manager endpoint.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.EndpointType == nil {
		return nil, errors.New("missing required argument 'EndpointType'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EndpointArgs{}
	}
	var resource Endpoint
	err := ctx.RegisterResource("azurerm:network/v20151101:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azurerm:network/v20151101:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// Gets or sets the name of the Traffic Manager endpoint.
	Name *string `pulumi:"name"`
	// Class representing a Traffic Manager endpoint properties.
	Properties *EndpointPropertiesResponse `pulumi:"properties"`
	// Gets or sets the endpoint type of the Traffic Manager endpoint.
	Type *string `pulumi:"type"`
}

type EndpointState struct {
	// Gets or sets the name of the Traffic Manager endpoint.
	Name pulumi.StringPtrInput
	// Class representing a Traffic Manager endpoint properties.
	Properties EndpointPropertiesResponsePtrInput
	// Gets or sets the endpoint type of the Traffic Manager endpoint.
	Type pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType string `pulumi:"endpointType"`
	// Gets or sets the ID of the Traffic Manager endpoint.
	Id *string `pulumi:"id"`
	// The name of the Traffic Manager endpoint to be created or updated.
	Name string `pulumi:"name"`
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// Class representing a Traffic Manager endpoint properties.
	Properties *EndpointProperties `pulumi:"properties"`
	// The name of the resource group containing the Traffic Manager endpoint to be created or updated.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the endpoint type of the Traffic Manager endpoint.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType pulumi.StringInput
	// Gets or sets the ID of the Traffic Manager endpoint.
	Id pulumi.StringPtrInput
	// The name of the Traffic Manager endpoint to be created or updated.
	Name pulumi.StringInput
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringInput
	// Class representing a Traffic Manager endpoint properties.
	Properties EndpointPropertiesPtrInput
	// The name of the resource group containing the Traffic Manager endpoint to be created or updated.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the endpoint type of the Traffic Manager endpoint.
	Type pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}