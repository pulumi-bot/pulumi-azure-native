// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180201
{
    /// <summary>
    /// Peering in an ExpressRoute Cross Connection resource.
    /// </summary>
    public partial class ExpressRouteCrossConnectionPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.ExpressRouteCrossConnectionPeeringPropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteCrossConnectionPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCrossConnectionPeering(string name, ExpressRouteCrossConnectionPeeringArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180201:ExpressRouteCrossConnectionPeering", name, args ?? new ExpressRouteCrossConnectionPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCrossConnectionPeering(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180201:ExpressRouteCrossConnectionPeering", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ExpressRouteCrossConnectionPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCrossConnectionPeering Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCrossConnectionPeering(name, id, options);
        }
    }

    public sealed class ExpressRouteCrossConnectionPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ExpressRouteCrossConnection.
        /// </summary>
        [Input("crossConnectionName", required: true)]
        public Input<string> CrossConnectionName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("properties")]
        public Input<Inputs.ExpressRouteCrossConnectionPeeringPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ExpressRouteCrossConnectionPeeringArgs()
        {
        }
    }
}