// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Service Endpoint policy definitions.
type ServiceEndpointPolicyDefinition struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the service endpoint policy definition.
	Properties ServiceEndpointPolicyDefinitionPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewServiceEndpointPolicyDefinition registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceEndpointPolicyName == nil {
		return nil, errors.New("missing required argument 'ServiceEndpointPolicyName'")
	}
	if args == nil {
		args = &ServiceEndpointPolicyDefinitionArgs{}
	}
	var resource ServiceEndpointPolicyDefinition
	err := ctx.RegisterResource("azurerm:network/v20190801:ServiceEndpointPolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPolicyDefinition gets an existing ServiceEndpointPolicyDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyDefinitionState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	var resource ServiceEndpointPolicyDefinition
	err := ctx.ReadResource("azurerm:network/v20190801:ServiceEndpointPolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPolicyDefinition resources.
type serviceEndpointPolicyDefinitionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the service endpoint policy definition.
	Properties *ServiceEndpointPolicyDefinitionPropertiesFormatResponse `pulumi:"properties"`
}

type ServiceEndpointPolicyDefinitionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the service endpoint policy definition.
	Properties ServiceEndpointPolicyDefinitionPropertiesFormatResponsePtrInput
}

func (ServiceEndpointPolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionState)(nil)).Elem()
}

type serviceEndpointPolicyDefinitionArgs struct {
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the service endpoint policy definition name.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Service endpoint name.
	Service *string `pulumi:"service"`
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName string `pulumi:"serviceEndpointPolicyName"`
	// A list of service resources.
	ServiceResources []string `pulumi:"serviceResources"`
}

// The set of arguments for constructing a ServiceEndpointPolicyDefinition resource.
type ServiceEndpointPolicyDefinitionArgs struct {
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the service endpoint policy definition name.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Service endpoint name.
	Service pulumi.StringPtrInput
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName pulumi.StringInput
	// A list of service resources.
	ServiceResources pulumi.StringArrayInput
}

func (ServiceEndpointPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionArgs)(nil)).Elem()
}
