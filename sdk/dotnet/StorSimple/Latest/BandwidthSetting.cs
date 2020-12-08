// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorSimple.Latest
{
    /// <summary>
    /// The bandwidth setting.
    /// </summary>
    public partial class BandwidthSetting : Pulumi.CustomResource
    {
        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The schedules.
        /// </summary>
        [Output("schedules")]
        public Output<ImmutableArray<Outputs.BandwidthScheduleResponse>> Schedules { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The number of volumes that uses the bandwidth setting.
        /// </summary>
        [Output("volumeCount")]
        public Output<int> VolumeCount { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthSetting(string name, BandwidthSettingArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:storsimple/latest:BandwidthSetting", name, args ?? new BandwidthSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthSetting(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:storsimple/latest:BandwidthSetting", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:storsimple/v20170601:BandwidthSetting"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BandwidthSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthSetting Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BandwidthSetting(name, id, options);
        }
    }

    public sealed class BandwidthSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth setting name.
        /// </summary>
        [Input("bandwidthSettingName", required: true)]
        public Input<string> BandwidthSettingName { get; set; } = null!;

        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        [Input("kind")]
        public Input<Pulumi.AzureNextGen.StorSimple.Latest.Kind>? Kind { get; set; }

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("schedules", required: true)]
        private InputList<Inputs.BandwidthScheduleArgs>? _schedules;

        /// <summary>
        /// The schedules.
        /// </summary>
        public InputList<Inputs.BandwidthScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.BandwidthScheduleArgs>());
            set => _schedules = value;
        }

        public BandwidthSettingArgs()
        {
        }
    }
}
