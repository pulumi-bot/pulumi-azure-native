// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20160401
{
    /// <summary>
    /// A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect
    /// </summary>
    public partial class RedisFirewallRule : Pulumi.CustomResource
    {
        /// <summary>
        /// name of the firewall rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// redis cache firewall rule properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.RedisFirewallRulePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// type (of the firewall rule resource = 'Microsoft.Cache/redis/firewallRule')
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RedisFirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RedisFirewallRule(string name, RedisFirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:cache/v20160401:RedisFirewallRule", name, args ?? new RedisFirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RedisFirewallRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:cache/v20160401:RedisFirewallRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RedisFirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RedisFirewallRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RedisFirewallRule(name, id, options);
        }
    }

    public sealed class RedisFirewallRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("cacheName", required: true)]
        public Input<string> CacheName { get; set; } = null!;

        /// <summary>
        /// highest IP address included in the range
        /// </summary>
        [Input("endIP", required: true)]
        public Input<string> EndIP { get; set; } = null!;

        /// <summary>
        /// The name of the firewall rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// lowest IP address included in the range
        /// </summary>
        [Input("startIP", required: true)]
        public Input<string> StartIP { get; set; } = null!;

        public RedisFirewallRuleArgs()
        {
        }
    }
}
