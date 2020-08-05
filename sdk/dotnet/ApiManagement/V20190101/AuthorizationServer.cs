// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101
{
    /// <summary>
    /// External OAuth authorization server settings.
    /// </summary>
    public partial class AuthorizationServer : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the External OAuth authorization server Contract.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AuthorizationServerContractPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizationServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizationServer(string name, AuthorizationServerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:AuthorizationServer", name, args ?? new AuthorizationServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizationServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:AuthorizationServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthorizationServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizationServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AuthorizationServer(name, id, options);
        }
    }

    public sealed class AuthorizationServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        [Input("authorizationMethods")]
        private InputList<string>? _authorizationMethods;

        /// <summary>
        /// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
        /// </summary>
        public InputList<string> AuthorizationMethods
        {
            get => _authorizationMethods ?? (_authorizationMethods = new InputList<string>());
            set => _authorizationMethods = value;
        }

        [Input("bearerTokenSendingMethods")]
        private InputList<string>? _bearerTokenSendingMethods;

        /// <summary>
        /// Specifies the mechanism by which access token is passed to the API. 
        /// </summary>
        public InputList<string> BearerTokenSendingMethods
        {
            get => _bearerTokenSendingMethods ?? (_bearerTokenSendingMethods = new InputList<string>());
            set => _bearerTokenSendingMethods = value;
        }

        [Input("clientAuthenticationMethod")]
        private InputList<string>? _clientAuthenticationMethod;

        /// <summary>
        /// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
        /// </summary>
        public InputList<string> ClientAuthenticationMethod
        {
            get => _clientAuthenticationMethod ?? (_clientAuthenticationMethod = new InputList<string>());
            set => _clientAuthenticationMethod = value;
        }

        /// <summary>
        /// Client or app id registered with this authorization server.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
        /// </summary>
        [Input("clientRegistrationEndpoint", required: true)]
        public Input<string> ClientRegistrationEndpoint { get; set; } = null!;

        /// <summary>
        /// Client or app secret registered with this authorization server.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
        /// </summary>
        [Input("defaultScope")]
        public Input<string>? DefaultScope { get; set; }

        /// <summary>
        /// Description of the authorization server. Can contain HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-friendly authorization server name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("grantTypes", required: true)]
        private InputList<string>? _grantTypes;

        /// <summary>
        /// Form of an authorization grant, which the client uses to request the access token.
        /// </summary>
        public InputList<string> GrantTypes
        {
            get => _grantTypes ?? (_grantTypes = new InputList<string>());
            set => _grantTypes = value;
        }

        /// <summary>
        /// Identifier of the authorization server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
        /// </summary>
        [Input("resourceOwnerPassword")]
        public Input<string>? ResourceOwnerPassword { get; set; }

        /// <summary>
        /// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
        /// </summary>
        [Input("resourceOwnerUsername")]
        public Input<string>? ResourceOwnerUsername { get; set; }

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
        /// </summary>
        [Input("supportState")]
        public Input<bool>? SupportState { get; set; }

        [Input("tokenBodyParameters")]
        private InputList<Inputs.TokenBodyParameterContractArgs>? _tokenBodyParameters;

        /// <summary>
        /// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
        /// </summary>
        public InputList<Inputs.TokenBodyParameterContractArgs> TokenBodyParameters
        {
            get => _tokenBodyParameters ?? (_tokenBodyParameters = new InputList<Inputs.TokenBodyParameterContractArgs>());
            set => _tokenBodyParameters = value;
        }

        /// <summary>
        /// OAuth token endpoint. Contains absolute URI to entity being referenced.
        /// </summary>
        [Input("tokenEndpoint")]
        public Input<string>? TokenEndpoint { get; set; }

        public AuthorizationServerArgs()
        {
        }
    }
}
