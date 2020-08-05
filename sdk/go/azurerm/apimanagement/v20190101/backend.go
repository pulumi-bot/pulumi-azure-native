// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Backend details.
type Backend struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Backend entity contract properties.
	Properties BackendContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &BackendArgs{}
	}
	var resource Backend
	err := ctx.RegisterResource("azurerm:apimanagement/v20190101:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("azurerm:apimanagement/v20190101:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Backend entity contract properties.
	Properties *BackendContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type BackendState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Backend entity contract properties.
	Properties BackendContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract `pulumi:"credentials"`
	// Backend Description.
	Description *string `pulumi:"description"`
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// Backend Properties contract
	Properties *BackendProperties `pulumi:"properties"`
	// Backend communication protocol.
	Protocol string `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy *BackendProxyContract `pulumi:"proxy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Backend Title.
	Title *string `pulumi:"title"`
	// Backend TLS Properties
	Tls *BackendTlsProperties `pulumi:"tls"`
	// Runtime Url of the Backend.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractPtrInput
	// Backend Description.
	Description pulumi.StringPtrInput
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// Backend Properties contract
	Properties BackendPropertiesPtrInput
	// Backend communication protocol.
	Protocol pulumi.StringInput
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Backend Title.
	Title pulumi.StringPtrInput
	// Backend TLS Properties
	Tls BackendTlsPropertiesPtrInput
	// Runtime Url of the Backend.
	Url pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}
