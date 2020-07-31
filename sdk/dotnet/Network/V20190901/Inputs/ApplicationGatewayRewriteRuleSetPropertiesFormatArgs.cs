// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Inputs
{

    /// <summary>
    /// Properties of rewrite rule set of the application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRewriteRuleSetPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        [Input("rewriteRules")]
        private InputList<Inputs.ApplicationGatewayRewriteRuleArgs>? _rewriteRules;

        /// <summary>
        /// Rewrite rules in the rewrite rule set.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRewriteRuleArgs> RewriteRules
        {
            get => _rewriteRules ?? (_rewriteRules = new InputList<Inputs.ApplicationGatewayRewriteRuleArgs>());
            set => _rewriteRules = value;
        }

        public ApplicationGatewayRewriteRuleSetPropertiesFormatArgs()
        {
        }
    }
}