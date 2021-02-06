// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest
{
    /// <summary>
    /// Describes a DNS zone.
    /// Latest API Version: 2018-05-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/latest:Zone")]
    public partial class Zone : Pulumi.CustomResource
    {
        /// <summary>
        /// The etag of the zone.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("maxNumberOfRecordSets")]
        public Output<double> MaxNumberOfRecordSets { get; private set; } = null!;

        /// <summary>
        /// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("maxNumberOfRecordsPerRecordSet")]
        public Output<double> MaxNumberOfRecordsPerRecordSet { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("numberOfRecordSets")]
        public Output<double> NumberOfRecordSets { get; private set; } = null!;

        /// <summary>
        /// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
        /// </summary>
        [Output("registrationVirtualNetworks")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> RegistrationVirtualNetworks { get; private set; } = null!;

        /// <summary>
        /// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
        /// </summary>
        [Output("resolutionVirtualNetworks")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> ResolutionVirtualNetworks { get; private set; } = null!;

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
        /// The type of this DNS zone (Public or Private).
        /// </summary>
        [Output("zoneType")]
        public Output<string?> ZoneType { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:Zone", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150504preview:Zone"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160401:Zone"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:Zone"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:Zone"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180301preview:Zone"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180501:Zone"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, options);
        }
    }

    public sealed class ZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The etag of the zone.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("registrationVirtualNetworks")]
        private InputList<Inputs.SubResourceArgs>? _registrationVirtualNetworks;

        /// <summary>
        /// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> RegistrationVirtualNetworks
        {
            get => _registrationVirtualNetworks ?? (_registrationVirtualNetworks = new InputList<Inputs.SubResourceArgs>());
            set => _registrationVirtualNetworks = value;
        }

        [Input("resolutionVirtualNetworks")]
        private InputList<Inputs.SubResourceArgs>? _resolutionVirtualNetworks;

        /// <summary>
        /// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> ResolutionVirtualNetworks
        {
            get => _resolutionVirtualNetworks ?? (_resolutionVirtualNetworks = new InputList<Inputs.SubResourceArgs>());
            set => _resolutionVirtualNetworks = value;
        }

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
        /// The name of the DNS zone (without a terminating dot).
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        /// <summary>
        /// The type of this DNS zone (Public or Private).
        /// </summary>
        [Input("zoneType")]
        public Input<Pulumi.AzureNextGen.Network.Latest.ZoneType>? ZoneType { get; set; }

        public ZoneArgs()
        {
            ZoneType = Pulumi.AzureNextGen.Network.Latest.ZoneType.Public;
        }
    }
}
