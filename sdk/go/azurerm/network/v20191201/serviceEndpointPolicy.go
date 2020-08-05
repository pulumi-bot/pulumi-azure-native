// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Service End point policy resource.
type ServiceEndpointPolicy struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the service end point policy.
	Properties ServiceEndpointPolicyPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceEndpointPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPolicy(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ServiceEndpointPolicyArgs{}
	}
	var resource ServiceEndpointPolicy
	err := ctx.RegisterResource("azurerm:network/v20191201:ServiceEndpointPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPolicy gets an existing ServiceEndpointPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	var resource ServiceEndpointPolicy
	err := ctx.ReadResource("azurerm:network/v20191201:ServiceEndpointPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPolicy resources.
type serviceEndpointPolicyState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the service end point policy.
	Properties *ServiceEndpointPolicyPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ServiceEndpointPolicyState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the service end point policy.
	Properties ServiceEndpointPolicyPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ServiceEndpointPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyState)(nil)).Elem()
}

type serviceEndpointPolicyArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the service endpoint policy.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionType `pulumi:"serviceEndpointPolicyDefinitions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceEndpointPolicy resource.
type ServiceEndpointPolicyArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the service endpoint policy.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceEndpointPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyArgs)(nil)).Elem()
}
