// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190615Preview.Outputs
{

    [OutputType]
    public sealed class EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkResult
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkResult(string? id)
        {
            Id = id;
        }
    }
}
