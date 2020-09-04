// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20170301Preview
{
    /// <summary>
    /// A server DNS alias.
    /// </summary>
    public partial class ServerDnsAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// The fully qualified DNS record for alias
        /// </summary>
        [Output("azureDnsRecord")]
        public Output<string> AzureDnsRecord { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerDnsAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerDnsAlias(string name, ServerDnsAliasArgs args, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20170301preview:ServerDnsAlias", name, args ?? new ServerDnsAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerDnsAlias(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20170301preview:ServerDnsAlias", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServerDnsAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerDnsAlias Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerDnsAlias(name, id, options);
        }
    }

    public sealed class ServerDnsAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the server DNS alias.
        /// </summary>
        [Input("dnsAliasName", required: true)]
        public Input<string> DnsAliasName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server that the alias is pointing to.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public ServerDnsAliasArgs()
        {
        }
    }
}
