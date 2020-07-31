// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Diagnostic details.
type ApiDiagnostic struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Diagnostic entity contract properties.
	Properties DiagnosticContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
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
		args = &ApiDiagnosticArgs{}
	}
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azurerm:apimanagement/v20190101:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDiagnostic gets an existing ApiDiagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azurerm:apimanagement/v20190101:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDiagnostic resources.
type apiDiagnosticState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Diagnostic entity contract properties.
	Properties *DiagnosticContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ApiDiagnosticState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Diagnostic entity contract properties.
	Properties DiagnosticContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// Diagnostic entity contract properties.
	Properties *DiagnosticContractProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ApiDiagnostic resource.
type ApiDiagnosticArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// Diagnostic entity contract properties.
	Properties DiagnosticContractPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}