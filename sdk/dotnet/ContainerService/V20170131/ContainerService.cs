// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20170131
{
    /// <summary>
    /// Container service.
    /// </summary>
    public partial class ContainerService : Pulumi.CustomResource
    {
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
        /// Properties of the container service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ContainerServicePropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a ContainerService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerService(string name, ContainerServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:containerservice/v20170131:ContainerService", name, args ?? new ContainerServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:containerservice/v20170131:ContainerService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ContainerService(name, id, options);
        }
    }

    public sealed class ContainerServiceArgs : Pulumi.ResourceArgs
    {
        [Input("agentPoolProfiles", required: true)]
        private InputList<Inputs.ContainerServiceAgentPoolProfileArgs>? _agentPoolProfiles;

        /// <summary>
        /// Properties of the agent pool.
        /// </summary>
        public InputList<Inputs.ContainerServiceAgentPoolProfileArgs> AgentPoolProfiles
        {
            get => _agentPoolProfiles ?? (_agentPoolProfiles = new InputList<Inputs.ContainerServiceAgentPoolProfileArgs>());
            set => _agentPoolProfiles = value;
        }

        /// <summary>
        /// Properties for custom clusters.
        /// </summary>
        [Input("customProfile")]
        public Input<Inputs.ContainerServiceCustomProfileArgs>? CustomProfile { get; set; }

        /// <summary>
        /// Properties of the diagnostic agent.
        /// </summary>
        [Input("diagnosticsProfile")]
        public Input<Inputs.ContainerServiceDiagnosticsProfileArgs>? DiagnosticsProfile { get; set; }

        /// <summary>
        /// Properties of Linux VMs.
        /// </summary>
        [Input("linuxProfile", required: true)]
        public Input<Inputs.ContainerServiceLinuxProfileArgs> LinuxProfile { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Properties of master agents.
        /// </summary>
        [Input("masterProfile", required: true)]
        public Input<Inputs.ContainerServiceMasterProfileArgs> MasterProfile { get; set; } = null!;

        /// <summary>
        /// The name of the container service in the specified subscription and resource group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties of the orchestrator.
        /// </summary>
        [Input("orchestratorProfile")]
        public Input<Inputs.ContainerServiceOrchestratorProfileArgs>? OrchestratorProfile { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Properties for cluster service principals.
        /// </summary>
        [Input("servicePrincipalProfile")]
        public Input<Inputs.ContainerServiceServicePrincipalProfileArgs>? ServicePrincipalProfile { get; set; }

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

        /// <summary>
        /// Properties of Windows VMs.
        /// </summary>
        [Input("windowsProfile")]
        public Input<Inputs.ContainerServiceWindowsProfileArgs>? WindowsProfile { get; set; }

        public ContainerServiceArgs()
        {
        }
    }
}
