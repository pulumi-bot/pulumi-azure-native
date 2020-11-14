// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Inputs
{

    /// <summary>
    /// Set of actions in the Rewrite Rule in Application Gateway.
    /// </summary>
    public sealed class ApplicationGatewayRewriteRuleActionSetArgs : Pulumi.ResourceArgs
    {
        [Input("requestHeaderConfigurations")]
        private InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs>? _requestHeaderConfigurations;

        /// <summary>
        /// Request Header Actions in the Action Set.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs> RequestHeaderConfigurations
        {
            get => _requestHeaderConfigurations ?? (_requestHeaderConfigurations = new InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs>());
            set => _requestHeaderConfigurations = value;
        }

        [Input("responseHeaderConfigurations")]
        private InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs>? _responseHeaderConfigurations;

        /// <summary>
        /// Response Header Actions in the Action Set.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs> ResponseHeaderConfigurations
        {
            get => _responseHeaderConfigurations ?? (_responseHeaderConfigurations = new InputList<Inputs.ApplicationGatewayHeaderConfigurationArgs>());
            set => _responseHeaderConfigurations = value;
        }

        /// <summary>
        /// Url Configuration Action in the Action Set.
        /// </summary>
        [Input("urlConfiguration")]
        public Input<Inputs.ApplicationGatewayUrlConfigurationArgs>? UrlConfiguration { get; set; }

        public ApplicationGatewayRewriteRuleActionSetArgs()
        {
        }
    }
}
