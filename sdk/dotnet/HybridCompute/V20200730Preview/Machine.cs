// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HybridCompute.V20200730Preview
{
    /// <summary>
    /// Describes a hybrid machine.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:hybridcompute/v20200730preview:Machine")]
    public partial class Machine : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the AD fully qualified display name.
        /// </summary>
        [Output("adFqdn")]
        public Output<string> AdFqdn { get; private set; } = null!;

        /// <summary>
        /// The hybrid machine agent full version.
        /// </summary>
        [Output("agentVersion")]
        public Output<string> AgentVersion { get; private set; } = null!;

        /// <summary>
        /// Public Key that the client provides to be used during initial resource onboarding
        /// </summary>
        [Output("clientPublicKey")]
        public Output<string?> ClientPublicKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the hybrid machine display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Specifies the DNS fully qualified display name.
        /// </summary>
        [Output("dnsFqdn")]
        public Output<string> DnsFqdn { get; private set; } = null!;

        /// <summary>
        /// Specifies the Windows domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Details about the error state.
        /// </summary>
        [Output("errorDetails")]
        public Output<ImmutableArray<Outputs.ErrorDetailResponse>> ErrorDetails { get; private set; } = null!;

        /// <summary>
        /// Machine Extensions information
        /// </summary>
        [Output("extensions")]
        public Output<ImmutableArray<Outputs.MachineExtensionInstanceViewResponse>> Extensions { get; private set; } = null!;

        [Output("identity")]
        public Output<Outputs.MachineResponseIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The time of the last status change.
        /// </summary>
        [Output("lastStatusChange")]
        public Output<string> LastStatusChange { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to the geographic location of the resource.
        /// </summary>
        [Output("locationData")]
        public Output<Outputs.LocationDataResponse?> LocationData { get; private set; } = null!;

        /// <summary>
        /// Specifies the hybrid machine FQDN.
        /// </summary>
        [Output("machineFqdn")]
        public Output<string> MachineFqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Operating System running on the hybrid machine.
        /// </summary>
        [Output("osName")]
        public Output<string> OsName { get; private set; } = null!;

        /// <summary>
        /// Specifies the operating system settings for the hybrid machine.
        /// </summary>
        [Output("osProfile")]
        public Output<Outputs.MachinePropertiesResponseOsProfile?> OsProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the Operating System product SKU.
        /// </summary>
        [Output("osSku")]
        public Output<string> OsSku { get; private set; } = null!;

        /// <summary>
        /// The version of Operating System running on the hybrid machine.
        /// </summary>
        [Output("osVersion")]
        public Output<string> OsVersion { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The status of the hybrid machine agent.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the hybrid machine unique ID.
        /// </summary>
        [Output("vmId")]
        public Output<string?> VmId { get; private set; } = null!;

        /// <summary>
        /// Specifies the Arc Machine's unique SMBIOS ID
        /// </summary>
        [Output("vmUuid")]
        public Output<string> VmUuid { get; private set; } = null!;


        /// <summary>
        /// Create a Machine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Machine(string name, MachineArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:hybridcompute/v20200730preview:Machine", name, args ?? new MachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Machine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:hybridcompute/v20200730preview:Machine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/latest:Machine"},
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/v20190318preview:Machine"},
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/v20190802preview:Machine"},
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/v20191212:Machine"},
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/v20200802:Machine"},
                    new Pulumi.Alias { Type = "azure-nextgen:hybridcompute/v20200815preview:Machine"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Machine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Machine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Machine(name, id, options);
        }
    }

    public sealed class MachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public Key that the client provides to be used during initial resource onboarding
        /// </summary>
        [Input("clientPublicKey")]
        public Input<string>? ClientPublicKey { get; set; }

        [Input("extensions")]
        private InputList<Inputs.MachineExtensionInstanceViewArgs>? _extensions;

        /// <summary>
        /// Machine Extensions information
        /// </summary>
        public InputList<Inputs.MachineExtensionInstanceViewArgs> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<Inputs.MachineExtensionInstanceViewArgs>());
            set => _extensions = value;
        }

        [Input("identity")]
        public Input<Inputs.MachineIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Metadata pertaining to the geographic location of the resource.
        /// </summary>
        [Input("locationData")]
        public Input<Inputs.LocationDataArgs>? LocationData { get; set; }

        /// <summary>
        /// The name of the hybrid machine.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
        /// Specifies the hybrid machine unique ID.
        /// </summary>
        [Input("vmId")]
        public Input<string>? VmId { get; set; }

        public MachineArgs()
        {
        }
    }
}
