// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Application gateway resource
type ApplicationGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the application gateway.
	Properties ApplicationGatewayPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationGateway registers a new resource with the given unique name, arguments, and options.
func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationGatewayArgs{}
	}
	var resource ApplicationGateway
	err := ctx.RegisterResource("azurerm:network/v20171101:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGateway gets an existing ApplicationGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azurerm:network/v20171101:ApplicationGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGateway resources.
type applicationGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the application gateway.
	Properties *ApplicationGatewayPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ApplicationGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the application gateway.
	Properties ApplicationGatewayPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ApplicationGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayState)(nil)).Elem()
}

type applicationGatewayArgs struct {
	// Authentication certificates of the application gateway resource.
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificate `pulumi:"authenticationCertificates"`
	// Backend address pool of the application gateway resource.
	BackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettings `pulumi:"backendHttpSettingsCollection"`
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource.
	FrontendPorts []ApplicationGatewayFrontendPort `pulumi:"frontendPorts"`
	// Subnets of application the gateway resource.
	GatewayIPConfigurations []ApplicationGatewayIPConfiguration `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource.
	HttpListeners []ApplicationGatewayHttpListener `pulumi:"httpListeners"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the application gateway.
	Name string `pulumi:"name"`
	// Probes of the application gateway resource.
	Probes []ApplicationGatewayProbe `pulumi:"probes"`
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Redirect configurations of the application gateway resource.
	RedirectConfigurations []ApplicationGatewayRedirectConfiguration `pulumi:"redirectConfigurations"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule `pulumi:"requestRoutingRules"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource GUID property of the application gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// SKU of the application gateway resource.
	Sku *ApplicationGatewaySku `pulumi:"sku"`
	// SSL certificates of the application gateway resource.
	SslCertificates []ApplicationGatewaySslCertificate `pulumi:"sslCertificates"`
	// SSL policy of the application gateway resource.
	SslPolicy *ApplicationGatewaySslPolicy `pulumi:"sslPolicy"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// URL path map of the application gateway resource.
	UrlPathMaps []ApplicationGatewayUrlPathMap `pulumi:"urlPathMaps"`
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfiguration `pulumi:"webApplicationFirewallConfiguration"`
}

// The set of arguments for constructing a ApplicationGateway resource.
type ApplicationGatewayArgs struct {
	// Authentication certificates of the application gateway resource.
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateArrayInput
	// Backend address pool of the application gateway resource.
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsArrayInput
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationArrayInput
	// Frontend ports of the application gateway resource.
	FrontendPorts ApplicationGatewayFrontendPortArrayInput
	// Subnets of application the gateway resource.
	GatewayIPConfigurations ApplicationGatewayIPConfigurationArrayInput
	// Http listeners of the application gateway resource.
	HttpListeners ApplicationGatewayHttpListenerArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the application gateway.
	Name pulumi.StringInput
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeArrayInput
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// Redirect configurations of the application gateway resource.
	RedirectConfigurations ApplicationGatewayRedirectConfigurationArrayInput
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuPtrInput
	// SSL certificates of the application gateway resource.
	SslCertificates ApplicationGatewaySslCertificateArrayInput
	// SSL policy of the application gateway resource.
	SslPolicy ApplicationGatewaySslPolicyPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// URL path map of the application gateway resource.
	UrlPathMaps ApplicationGatewayUrlPathMapArrayInput
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration ApplicationGatewayWebApplicationFirewallConfigurationPtrInput
}

func (ApplicationGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayArgs)(nil)).Elem()
}
