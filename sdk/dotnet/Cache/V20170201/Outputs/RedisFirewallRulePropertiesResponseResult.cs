// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20170201.Outputs
{

    [OutputType]
    public sealed class RedisFirewallRulePropertiesResponseResult
    {
        /// <summary>
        /// highest IP address included in the range
        /// </summary>
        public readonly string EndIP;
        /// <summary>
        /// lowest IP address included in the range
        /// </summary>
        public readonly string StartIP;

        [OutputConstructor]
        private RedisFirewallRulePropertiesResponseResult(
            string endIP,

            string startIP)
        {
            EndIP = endIP;
            StartIP = startIP;
        }
    }
}