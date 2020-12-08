// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Inputs
{

    /// <summary>
    /// API OAuth2 Authentication settings details.
    /// </summary>
    public sealed class OpenIdAuthenticationSettingsContractArgs : Pulumi.ResourceArgs
    {
        [Input("bearerTokenSendingMethods")]
        private InputList<Union<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.BearerTokenSendingMethods>>? _bearerTokenSendingMethods;

        /// <summary>
        /// How to send token to the server.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.BearerTokenSendingMethods>> BearerTokenSendingMethods
        {
            get => _bearerTokenSendingMethods ?? (_bearerTokenSendingMethods = new InputList<Union<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.BearerTokenSendingMethods>>());
            set => _bearerTokenSendingMethods = value;
        }

        /// <summary>
        /// OAuth authorization server identifier.
        /// </summary>
        [Input("openidProviderId")]
        public Input<string>? OpenidProviderId { get; set; }

        public OpenIdAuthenticationSettingsContractArgs()
        {
        }
    }
}
