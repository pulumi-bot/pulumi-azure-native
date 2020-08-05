// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// External OAuth authorization server settings.
type AuthorizationServer struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the External OAuth authorization server Contract.
	Properties AuthorizationServerContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthorizationServer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationServer(ctx *pulumi.Context,
	name string, args *AuthorizationServerArgs, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	if args == nil || args.AuthorizationEndpoint == nil {
		return nil, errors.New("missing required argument 'AuthorizationEndpoint'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientRegistrationEndpoint == nil {
		return nil, errors.New("missing required argument 'ClientRegistrationEndpoint'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.GrantTypes == nil {
		return nil, errors.New("missing required argument 'GrantTypes'")
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
		args = &AuthorizationServerArgs{}
	}
	var resource AuthorizationServer
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:AuthorizationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationServer gets an existing AuthorizationServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationServerState, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	var resource AuthorizationServer
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:AuthorizationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationServer resources.
type authorizationServerState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the External OAuth authorization server Contract.
	Properties *AuthorizationServerContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type AuthorizationServerState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the External OAuth authorization server Contract.
	Properties AuthorizationServerContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (AuthorizationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerState)(nil)).Elem()
}

type authorizationServerArgs struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods []string `pulumi:"authorizationMethods"`
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod []string `pulumi:"clientAuthenticationMethod"`
	// Client or app id registered with this authorization server.
	ClientId string `pulumi:"clientId"`
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint string `pulumi:"clientRegistrationEndpoint"`
	// Client or app secret registered with this authorization server.
	ClientSecret *string `pulumi:"clientSecret"`
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// Description of the authorization server. Can contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// User-friendly authorization server name.
	DisplayName string `pulumi:"displayName"`
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes []string `pulumi:"grantTypes"`
	// Identifier of the authorization server.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters []TokenBodyParameterContract `pulumi:"tokenBodyParameters"`
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
}

// The set of arguments for constructing a AuthorizationServer resource.
type AuthorizationServerArgs struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint pulumi.StringInput
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods pulumi.StringArrayInput
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods pulumi.StringArrayInput
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod pulumi.StringArrayInput
	// Client or app id registered with this authorization server.
	ClientId pulumi.StringInput
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint pulumi.StringInput
	// Client or app secret registered with this authorization server.
	ClientSecret pulumi.StringPtrInput
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope pulumi.StringPtrInput
	// Description of the authorization server. Can contain HTML formatting tags.
	Description pulumi.StringPtrInput
	// User-friendly authorization server name.
	DisplayName pulumi.StringInput
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes pulumi.StringArrayInput
	// Identifier of the authorization server.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword pulumi.StringPtrInput
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState pulumi.BoolPtrInput
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters TokenBodyParameterContractArrayInput
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint pulumi.StringPtrInput
}

func (AuthorizationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerArgs)(nil)).Elem()
}
