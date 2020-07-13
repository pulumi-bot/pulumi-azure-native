// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.KeyVault.Outputs
{

    [OutputType]
    public sealed class IPRuleResponse
    {
        /// <summary>
        /// An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private IPRuleResponse(string value)
        {
            Value = value;
        }
    }
}