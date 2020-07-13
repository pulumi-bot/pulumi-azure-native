// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayRewriteRuleSetResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the rewrite rule set that is unique within an Application Gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of the application gateway rewrite rule set.
        /// </summary>
        public readonly Outputs.ApplicationGatewayRewriteRuleSetPropertiesFormatResponse? Properties;

        [OutputConstructor]
        private ApplicationGatewayRewriteRuleSetResponse(
            string etag,

            string? id,

            string? name,

            Outputs.ApplicationGatewayRewriteRuleSetPropertiesFormatResponse? properties)
        {
            Etag = etag;
            Id = id;
            Name = name;
            Properties = properties;
        }
    }
}