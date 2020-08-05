// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190801
{
    /// <summary>
    /// ExpressRoute gateway resource.
    /// </summary>
    public partial class ExpressRouteGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the express route gateway.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ExpressRouteGatewayPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteGateway(string name, ExpressRouteGatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190801:ExpressRouteGateway", name, args ?? new ExpressRouteGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190801:ExpressRouteGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ExpressRouteGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteGateway(name, id, options);
        }
    }

    public sealed class ExpressRouteGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for auto scaling.
        /// </summary>
        [Input("autoScaleConfiguration")]
        public Input<Inputs.ExpressRouteGatewayPropertiesPropertiesArgs>? AutoScaleConfiguration { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the ExpressRoute gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the express route gateway resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
        /// </summary>
        [Input("virtualHub", required: true)]
        public Input<Inputs.VirtualHubIdArgs> VirtualHub { get; set; } = null!;

        public ExpressRouteGatewayArgs()
        {
        }
    }
}
