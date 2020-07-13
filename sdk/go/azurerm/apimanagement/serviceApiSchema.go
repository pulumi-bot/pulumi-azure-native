// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Schema Contract details.
type ServiceApiSchema struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Schema.
	Properties SchemaContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceApiSchema registers a new resource with the given unique name, arguments, and options.
func NewServiceApiSchema(ctx *pulumi.Context,
	name string, args *ServiceApiSchemaArgs, opts ...pulumi.ResourceOption) (*ServiceApiSchema, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceApiSchemaArgs{}
	}
	var resource ServiceApiSchema
	err := ctx.RegisterResource("azurerm:apimanagement:ServiceApiSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceApiSchema gets an existing ServiceApiSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceApiSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceApiSchemaState, opts ...pulumi.ResourceOption) (*ServiceApiSchema, error) {
	var resource ServiceApiSchema
	err := ctx.ReadResource("azurerm:apimanagement:ServiceApiSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceApiSchema resources.
type serviceApiSchemaState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the Schema.
	Properties *SchemaContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ServiceApiSchemaState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the Schema.
	Properties SchemaContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ServiceApiSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceApiSchemaState)(nil)).Elem()
}

type serviceApiSchemaArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Schema identifier within an API. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// Properties of the Schema.
	Properties *SchemaContractProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServiceApiSchema resource.
type ServiceApiSchemaArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Schema identifier within an API. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// Properties of the Schema.
	Properties SchemaContractPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ServiceApiSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceApiSchemaArgs)(nil)).Elem()
}