// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.VMwareCloudSimple.Latest
{
    /// <summary>
    /// Dedicated cloud node model
    /// </summary>
    public partial class DedicatedCloudNode : Pulumi.CustomResource
    {
        /// <summary>
        /// Availability Zone id, e.g. "az1"
        /// </summary>
        [Output("availabilityZoneId")]
        public Output<string> AvailabilityZoneId { get; private set; } = null!;

        /// <summary>
        /// Availability Zone name, e.g. "Availability Zone 1"
        /// </summary>
        [Output("availabilityZoneName")]
        public Output<string> AvailabilityZoneName { get; private set; } = null!;

        /// <summary>
        /// VMWare Cloud Rack Name
        /// </summary>
        [Output("cloudRackName")]
        public Output<string> CloudRackName { get; private set; } = null!;

        /// <summary>
        /// Azure region
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// SKU's name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// count of nodes to create
        /// </summary>
        [Output("nodesCount")]
        public Output<int> NodesCount { get; private set; } = null!;

        /// <summary>
        /// Placement Group id, e.g. "n1"
        /// </summary>
        [Output("placementGroupId")]
        public Output<string> PlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// Placement Name, e.g. "Placement Group 1"
        /// </summary>
        [Output("placementGroupName")]
        public Output<string> PlacementGroupName { get; private set; } = null!;

        /// <summary>
        /// Private Cloud Id
        /// </summary>
        [Output("privateCloudId")]
        public Output<string> PrivateCloudId { get; private set; } = null!;

        /// <summary>
        /// Resource Pool Name
        /// </summary>
        [Output("privateCloudName")]
        public Output<string> PrivateCloudName { get; private set; } = null!;

        /// <summary>
        /// The provisioning status of the resource
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// purchase id
        /// </summary>
        [Output("purchaseId")]
        public Output<string> PurchaseId { get; private set; } = null!;

        /// <summary>
        /// Dedicated Cloud Nodes SKU
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Node status, indicates is private cloud set up on this node or not
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Dedicated Cloud Nodes tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// {resourceProviderNamespace}/{resourceType}
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VMWare Cluster Name
        /// </summary>
        [Output("vmwareClusterName")]
        public Output<string> VmwareClusterName { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedCloudNode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedCloudNode(string name, DedicatedCloudNodeArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:vmwarecloudsimple/latest:DedicatedCloudNode", name, args ?? new DedicatedCloudNodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedCloudNode(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:vmwarecloudsimple/latest:DedicatedCloudNode", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:vmwarecloudsimple/v20190401:DedicatedCloudNode"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DedicatedCloudNode resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedCloudNode Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DedicatedCloudNode(name, id, options);
        }
    }

    public sealed class DedicatedCloudNodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability Zone id, e.g. "az1"
        /// </summary>
        [Input("availabilityZoneId", required: true)]
        public Input<string> AvailabilityZoneId { get; set; } = null!;

        /// <summary>
        /// dedicated cloud node name
        /// </summary>
        [Input("dedicatedCloudNodeName", required: true)]
        public Input<string> DedicatedCloudNodeName { get; set; } = null!;

        /// <summary>
        /// SKU's id
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Azure region
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// SKU's name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// count of nodes to create
        /// </summary>
        [Input("nodesCount", required: true)]
        public Input<int> NodesCount { get; set; } = null!;

        /// <summary>
        /// Placement Group id, e.g. "n1"
        /// </summary>
        [Input("placementGroupId", required: true)]
        public Input<string> PlacementGroupId { get; set; } = null!;

        /// <summary>
        /// purchase id
        /// </summary>
        [Input("purchaseId", required: true)]
        public Input<string> PurchaseId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Dedicated Cloud Nodes SKU
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Dedicated Cloud Nodes tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DedicatedCloudNodeArgs()
        {
        }
    }
}
