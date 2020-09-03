// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Latest
{
    /// <summary>
    /// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. &lt;br&gt;&lt;br&gt; Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
    /// 
    /// ## Example Usage
    /// ### Create or update a dedicated host group.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dedicatedHostGroup = new AzureRM.Compute.Latest.DedicatedHostGroup("dedicatedHostGroup", new AzureRM.Compute.Latest.DedicatedHostGroupArgs
    ///         {
    ///             HostGroupName = "myDedicatedHostGroup",
    ///             Location = "westus",
    ///             PlatformFaultDomainCount = 3,
    ///             ResourceGroupName = "myResourceGroup",
    ///             SupportAutomaticPlacement = true,
    ///             Tags = 
    ///             {
    ///                 { "department", "finance" },
    ///             },
    ///             Zones = 
    ///             {
    ///                 "1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class DedicatedHostGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of references to all dedicated hosts in the dedicated host group.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<Outputs.SubResourceReadOnlyResponseResult>> Hosts { get; private set; } = null!;

        /// <summary>
        /// The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
        /// </summary>
        [Output("instanceView")]
        public Output<Outputs.DedicatedHostGroupInstanceViewResponseResult> InstanceView { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of fault domains that the host group can span.
        /// </summary>
        [Output("platformFaultDomainCount")]
        public Output<int> PlatformFaultDomainCount { get; private set; } = null!;

        /// <summary>
        /// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'true' when not provided. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01.
        /// </summary>
        [Output("supportAutomaticPlacement")]
        public Output<bool?> SupportAutomaticPlacement { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHostGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHostGroup(string name, DedicatedHostGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/latest:DedicatedHostGroup", name, args ?? new DedicatedHostGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHostGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/latest:DedicatedHostGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/v20190301:DedicatedHostGroup"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:DedicatedHostGroup"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20191201:DedicatedHostGroup"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20200601:DedicatedHostGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DedicatedHostGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHostGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DedicatedHostGroup(name, id, options);
        }
    }

    public sealed class DedicatedHostGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the dedicated host group.
        /// </summary>
        [Input("hostGroupName", required: true)]
        public Input<string> HostGroupName { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Number of fault domains that the host group can span.
        /// </summary>
        [Input("platformFaultDomainCount", required: true)]
        public Input<int> PlatformFaultDomainCount { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'true' when not provided. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01.
        /// </summary>
        [Input("supportAutomaticPlacement")]
        public Input<bool>? SupportAutomaticPlacement { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public DedicatedHostGroupArgs()
        {
        }
    }
}
