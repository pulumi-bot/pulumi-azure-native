// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20151101Preview
{
    /// <summary>
    /// A user-defined logical grouping of machines.
    /// </summary>
    public partial class MachineGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
        /// </summary>
        [Output("count")]
        public Output<int?> Count { get; private set; } = null!;

        /// <summary>
        /// User defined name for the group
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource ETAG.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Type of the machine group
        /// </summary>
        [Output("groupType")]
        public Output<string?> GroupType { get; private set; } = null!;

        /// <summary>
        /// Additional resource type qualifier.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
        /// </summary>
        [Output("machines")]
        public Output<ImmutableArray<Outputs.MachineReferenceWithHintsResponseResult>> Machines { get; private set; } = null!;

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
        /// Create a MachineGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineGroup(string name, MachineGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20151101preview:MachineGroup", name, args ?? new MachineGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20151101preview:MachineGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MachineGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MachineGroup(name, id, options);
        }
    }

    public sealed class MachineGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// User defined name for the group
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Resource ETAG.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Type of the machine group
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        /// <summary>
        /// Additional resource type qualifier.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Machine Group resource name.
        /// </summary>
        [Input("machineGroupName", required: true)]
        public Input<string> MachineGroupName { get; set; } = null!;

        [Input("machines")]
        private InputList<Inputs.MachineReferenceWithHintsArgs>? _machines;

        /// <summary>
        /// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
        /// </summary>
        public InputList<Inputs.MachineReferenceWithHintsArgs> Machines
        {
            get => _machines ?? (_machines = new InputList<Inputs.MachineReferenceWithHintsArgs>());
            set => _machines = value;
        }

        /// <summary>
        /// Resource group name within the specified subscriptionId.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// OMS workspace containing the resources of interest.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public MachineGroupArgs()
        {
        }
    }
}
