// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20201101Preview
{
    /// <summary>
    /// An object that represents a replication for a container registry.
    /// </summary>
    public partial class Replication : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the replication at the time the operation was called.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the replication's regional endpoint is enabled. Requests will not be routed to a replication whose regional endpoint is disabled, however its data will continue to be synced with other replications.
        /// </summary>
        [Output("regionEndpointEnabled")]
        public Output<bool?> RegionEndpointEnabled { get; private set; } = null!;

        /// <summary>
        /// The status of the replication at the time the operation was called.
        /// </summary>
        [Output("status")]
        public Output<Outputs.StatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Whether or not zone redundancy is enabled for this container registry replication
        /// </summary>
        [Output("zoneRedundancy")]
        public Output<string?> ZoneRedundancy { get; private set; } = null!;


        /// <summary>
        /// Create a Replication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Replication(string name, ReplicationArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:containerregistry/v20201101preview:Replication", name, args ?? new ReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Replication(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:containerregistry/v20201101preview:Replication", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:containerregistry/latest:Replication"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerregistry/v20170601preview:Replication"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerregistry/v20171001:Replication"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerregistry/v20190501:Replication"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerregistry/v20191201preview:Replication"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Replication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Replication Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Replication(name, id, options);
        }
    }

    public sealed class ReplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies whether the replication's regional endpoint is enabled. Requests will not be routed to a replication whose regional endpoint is disabled, however its data will continue to be synced with other replications.
        /// </summary>
        [Input("regionEndpointEnabled")]
        public Input<bool>? RegionEndpointEnabled { get; set; }

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the replication.
        /// </summary>
        [Input("replicationName", required: true)]
        public Input<string> ReplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether or not zone redundancy is enabled for this container registry replication
        /// </summary>
        [Input("zoneRedundancy")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerRegistry.V20201101Preview.ZoneRedundancy>? ZoneRedundancy { get; set; }

        public ReplicationArgs()
        {
        }
    }
}
