// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayRequestRoutingRulePropertiesFormatResponseResult
    {
        /// <summary>
        /// Backend address pool resource of the application gateway. 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendAddressPool;
        /// <summary>
        /// Frontend port resource of the application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendHttpSettings;
        /// <summary>
        /// Http listener resource of the application gateway. 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? HttpListener;
        /// <summary>
        /// Provisioning state of the request routing rule resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Rule type. Possible values are: 'Basic' and 'PathBasedRouting'.
        /// </summary>
        public readonly string? RuleType;
        /// <summary>
        /// URL path map resource of the application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? UrlPathMap;

        [OutputConstructor]
        private ApplicationGatewayRequestRoutingRulePropertiesFormatResponseResult(
            Outputs.SubResourceResponseResult? backendAddressPool,

            Outputs.SubResourceResponseResult? backendHttpSettings,

            Outputs.SubResourceResponseResult? httpListener,

            string? provisioningState,

            string? ruleType,

            Outputs.SubResourceResponseResult? urlPathMap)
        {
            BackendAddressPool = backendAddressPool;
            BackendHttpSettings = backendHttpSettings;
            HttpListener = httpListener;
            ProvisioningState = provisioningState;
            RuleType = ruleType;
            UrlPathMap = urlPathMap;
        }
    }
}