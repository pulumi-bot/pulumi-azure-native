// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301
{
    /// <summary>
    /// Identity Provider details.
    /// </summary>
    public partial class IdentityProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// List of Allowed Tenants when configuring Azure Active Directory login.
        /// </summary>
        [Output("allowedTenants")]
        public Output<ImmutableArray<string>> AllowedTenants { get; private set; } = null!;

        /// <summary>
        /// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Output("passwordResetPolicyName")]
        public Output<string?> PasswordResetPolicyName { get; private set; } = null!;

        /// <summary>
        /// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Output("profileEditingPolicyName")]
        public Output<string?> ProfileEditingPolicyName { get; private set; } = null!;

        /// <summary>
        /// Signin Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Output("signinPolicyName")]
        public Output<string?> SigninPolicyName { get; private set; } = null!;

        /// <summary>
        /// Signup Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Output("signupPolicyName")]
        public Output<string?> SignupPolicyName { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProvider(string name, IdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20170301:IdentityProvider", name, args ?? new IdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20170301:IdentityProvider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/latest:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20160707:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20161010:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180101:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180601preview:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20190101:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:IdentityProvider"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201preview:IdentityProvider"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IdentityProvider(name, id, options);
        }
    }

    public sealed class IdentityProviderArgs : Pulumi.ResourceArgs
    {
        [Input("allowedTenants")]
        private InputList<string>? _allowedTenants;

        /// <summary>
        /// List of Allowed Tenants when configuring Azure Active Directory login.
        /// </summary>
        public InputList<string> AllowedTenants
        {
            get => _allowedTenants ?? (_allowedTenants = new InputList<string>());
            set => _allowedTenants = value;
        }

        /// <summary>
        /// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Identity Provider Type identifier.
        /// </summary>
        [Input("identityProviderName", required: true)]
        public Input<string> IdentityProviderName { get; set; } = null!;

        /// <summary>
        /// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Input("passwordResetPolicyName")]
        public Input<string>? PasswordResetPolicyName { get; set; }

        /// <summary>
        /// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Input("profileEditingPolicyName")]
        public Input<string>? ProfileEditingPolicyName { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Signin Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Input("signinPolicyName")]
        public Input<string>? SigninPolicyName { get; set; }

        /// <summary>
        /// Signup Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        [Input("signupPolicyName")]
        public Input<string>? SignupPolicyName { get; set; }

        /// <summary>
        /// Identity Provider Type identifier.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IdentityProviderArgs()
        {
        }
    }
}
