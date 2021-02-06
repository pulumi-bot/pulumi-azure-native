// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.Latest
{
    /// <summary>
    /// Specifies information about the Dedicated host.
    /// Latest API Version: 2020-12-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:compute/latest:DedicatedHost")]
    public partial class DedicatedHost : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        /// </summary>
        [Output("autoReplaceOnFailure")]
        public Output<bool?> AutoReplaceOnFailure { get; private set; } = null!;

        /// <summary>
        /// A unique id generated and assigned to the dedicated host by the platform. &lt;br&gt;&lt;br&gt; Does not change throughout the lifetime of the host.
        /// </summary>
        [Output("hostId")]
        public Output<string> HostId { get; private set; } = null!;

        /// <summary>
        /// The dedicated host instance view.
        /// </summary>
        [Output("instanceView")]
        public Output<Outputs.DedicatedHostInstanceViewResponse> InstanceView { get; private set; } = null!;

        /// <summary>
        /// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **None** &lt;br&gt;&lt;br&gt; **Windows_Server_Hybrid** &lt;br&gt;&lt;br&gt; **Windows_Server_Perpetual** &lt;br&gt;&lt;br&gt; Default: **None**
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

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
        /// Fault domain of the dedicated host within a dedicated host group.
        /// </summary>
        [Output("platformFaultDomain")]
        public Output<int?> PlatformFaultDomain { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The date when the host was first provisioned.
        /// </summary>
        [Output("provisioningTime")]
        public Output<string> ProvisioningTime { get; private set; } = null!;

        /// <summary>
        /// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

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
        /// A list of references to all virtual machines in the Dedicated Host.
        /// </summary>
        [Output("virtualMachines")]
        public Output<ImmutableArray<Outputs.SubResourceReadOnlyResponse>> VirtualMachines { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHost(string name, DedicatedHostArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/latest:DedicatedHost", name, args ?? new DedicatedHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHost(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/latest:DedicatedHost", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190301:DedicatedHost"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190701:DedicatedHost"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20191201:DedicatedHost"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200601:DedicatedHost"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20201201:DedicatedHost"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHost Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DedicatedHost(name, id, options);
        }
    }

    public sealed class DedicatedHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        /// </summary>
        [Input("autoReplaceOnFailure")]
        public Input<bool>? AutoReplaceOnFailure { get; set; }

        /// <summary>
        /// The name of the dedicated host group.
        /// </summary>
        [Input("hostGroupName", required: true)]
        public Input<string> HostGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the dedicated host .
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **None** &lt;br&gt;&lt;br&gt; **Windows_Server_Hybrid** &lt;br&gt;&lt;br&gt; **Windows_Server_Perpetual** &lt;br&gt;&lt;br&gt; Default: **None**
        /// </summary>
        [Input("licenseType")]
        public Input<Pulumi.AzureNextGen.Compute.Latest.DedicatedHostLicenseTypes>? LicenseType { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Fault domain of the dedicated host within a dedicated host group.
        /// </summary>
        [Input("platformFaultDomain")]
        public Input<int>? PlatformFaultDomain { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

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

        public DedicatedHostArgs()
        {
        }
    }
}
