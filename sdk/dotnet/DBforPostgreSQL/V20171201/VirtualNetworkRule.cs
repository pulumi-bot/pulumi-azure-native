// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforPostgreSQL.V20171201
{
    /// <summary>
    /// A virtual network rule.
    /// </summary>
    public partial class VirtualNetworkRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkRulePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkRule(string name, VirtualNetworkRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:dbforpostgresql/v20171201:VirtualNetworkRule", name, args ?? new VirtualNetworkRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:dbforpostgresql/v20171201:VirtualNetworkRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetworkRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkRule(name, id, options);
        }
    }

    public sealed class VirtualNetworkRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create firewall rule before the virtual network has vnet service endpoint enabled.
        /// </summary>
        [Input("ignoreMissingVnetServiceEndpoint")]
        public Input<bool>? IgnoreMissingVnetServiceEndpoint { get; set; }

        /// <summary>
        /// The name of the virtual network rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The ARM resource id of the virtual network subnet.
        /// </summary>
        [Input("virtualNetworkSubnetId", required: true)]
        public Input<string> VirtualNetworkSubnetId { get; set; } = null!;

        public VirtualNetworkRuleArgs()
        {
        }
    }
}
