// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Gateway hostname configuration details.
type GatewayHostnameConfiguration struct {
	pulumi.CustomResourceState

	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate pulumi.BoolPtrOutput `pulumi:"negotiateClientCertificate"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGatewayHostnameConfiguration registers a new resource with the given unique name, arguments, and options.
func NewGatewayHostnameConfiguration(ctx *pulumi.Context,
	name string, args *GatewayHostnameConfigurationArgs, opts ...pulumi.ResourceOption) (*GatewayHostnameConfiguration, error) {
	if args == nil || args.GatewayId == nil {
		return nil, errors.New("missing required argument 'GatewayId'")
	}
	if args == nil || args.HcId == nil {
		return nil, errors.New("missing required argument 'HcId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &GatewayHostnameConfigurationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:GatewayHostnameConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayHostnameConfiguration
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201preview:GatewayHostnameConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayHostnameConfiguration gets an existing GatewayHostnameConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayHostnameConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayHostnameConfigurationState, opts ...pulumi.ResourceOption) (*GatewayHostnameConfiguration, error) {
	var resource GatewayHostnameConfiguration
	err := ctx.ReadResource("azurerm:apimanagement/v20191201preview:GatewayHostnameConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayHostnameConfiguration resources.
type gatewayHostnameConfigurationState struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId *string `pulumi:"certificateId"`
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname *string `pulumi:"hostname"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate *bool `pulumi:"negotiateClientCertificate"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type GatewayHostnameConfigurationState struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId pulumi.StringPtrInput
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate pulumi.BoolPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (GatewayHostnameConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayHostnameConfigurationState)(nil)).Elem()
}

type gatewayHostnameConfigurationArgs struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId *string `pulumi:"certificateId"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
	HcId string `pulumi:"hcId"`
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname *string `pulumi:"hostname"`
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate *bool `pulumi:"negotiateClientCertificate"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a GatewayHostnameConfiguration resource.
type GatewayHostnameConfigurationArgs struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId pulumi.StringPtrInput
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringInput
	// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
	HcId pulumi.StringInput
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname pulumi.StringPtrInput
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate pulumi.BoolPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (GatewayHostnameConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayHostnameConfigurationArgs)(nil)).Elem()
}
