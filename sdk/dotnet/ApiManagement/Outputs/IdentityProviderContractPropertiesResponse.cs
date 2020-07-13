// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Outputs
{

    [OutputType]
    public sealed class IdentityProviderContractPropertiesResponse
    {
        /// <summary>
        /// List of Allowed Tenants when configuring Azure Active Directory login.
        /// </summary>
        public readonly ImmutableArray<string> AllowedTenants;
        /// <summary>
        /// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
        /// </summary>
        public readonly string? Authority;
        /// <summary>
        /// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        public readonly string? PasswordResetPolicyName;
        /// <summary>
        /// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        public readonly string? ProfileEditingPolicyName;
        /// <summary>
        /// Signin Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        public readonly string? SigninPolicyName;
        /// <summary>
        /// The TenantId to use instead of Common when logging into Active Directory
        /// </summary>
        public readonly string? SigninTenant;
        /// <summary>
        /// Signup Policy Name. Only applies to AAD B2C Identity Provider.
        /// </summary>
        public readonly string? SignupPolicyName;
        /// <summary>
        /// Identity Provider Type identifier.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private IdentityProviderContractPropertiesResponse(
            ImmutableArray<string> allowedTenants,

            string? authority,

            string clientId,

            string? clientSecret,

            string? passwordResetPolicyName,

            string? profileEditingPolicyName,

            string? signinPolicyName,

            string? signinTenant,

            string? signupPolicyName,

            string? type)
        {
            AllowedTenants = allowedTenants;
            Authority = authority;
            ClientId = clientId;
            ClientSecret = clientSecret;
            PasswordResetPolicyName = passwordResetPolicyName;
            ProfileEditingPolicyName = profileEditingPolicyName;
            SigninPolicyName = signinPolicyName;
            SigninTenant = signinTenant;
            SignupPolicyName = signupPolicyName;
            Type = type;
        }
    }
}