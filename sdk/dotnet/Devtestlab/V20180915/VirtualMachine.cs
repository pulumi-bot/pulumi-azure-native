// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20180915
{
    /// <summary>
    /// A virtual machine.
    /// </summary>
    public partial class VirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.LabVirtualMachinePropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20180915:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20180915:VirtualMachine", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, options);
        }
    }

    public sealed class VirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether another user can take ownership of the virtual machine
        /// </summary>
        [Input("allowClaim")]
        public Input<bool>? AllowClaim { get; set; }

        /// <summary>
        /// The artifact deployment status for the virtual machine.
        /// </summary>
        [Input("artifactDeploymentStatus")]
        public Input<Inputs.ArtifactDeploymentStatusPropertiesArgs>? ArtifactDeploymentStatus { get; set; }

        [Input("artifacts")]
        private InputList<Inputs.ArtifactInstallPropertiesArgs>? _artifacts;

        /// <summary>
        /// The artifacts to be installed on the virtual machine.
        /// </summary>
        public InputList<Inputs.ArtifactInstallPropertiesArgs> Artifacts
        {
            get => _artifacts ?? (_artifacts = new InputList<Inputs.ArtifactInstallPropertiesArgs>());
            set => _artifacts = value;
        }

        /// <summary>
        /// The resource identifier (Microsoft.Compute) of the virtual machine.
        /// </summary>
        [Input("computeId")]
        public Input<string>? ComputeId { get; set; }

        /// <summary>
        /// The email address of creator of the virtual machine.
        /// </summary>
        [Input("createdByUser")]
        public Input<string>? CreatedByUser { get; set; }

        /// <summary>
        /// The object identifier of the creator of the virtual machine.
        /// </summary>
        [Input("createdByUserId")]
        public Input<string>? CreatedByUserId { get; set; }

        /// <summary>
        /// The creation date of the virtual machine.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The custom image identifier of the virtual machine.
        /// </summary>
        [Input("customImageId")]
        public Input<string>? CustomImageId { get; set; }

        [Input("dataDiskParameters")]
        private InputList<Inputs.DataDiskPropertiesArgs>? _dataDiskParameters;

        /// <summary>
        /// New or existing data disks to attach to the virtual machine after creation
        /// </summary>
        public InputList<Inputs.DataDiskPropertiesArgs> DataDiskParameters
        {
            get => _dataDiskParameters ?? (_dataDiskParameters = new InputList<Inputs.DataDiskPropertiesArgs>());
            set => _dataDiskParameters = value;
        }

        /// <summary>
        /// Indicates whether the virtual machine is to be created without a public IP address.
        /// </summary>
        [Input("disallowPublicIpAddress")]
        public Input<bool>? DisallowPublicIpAddress { get; set; }

        /// <summary>
        /// The resource ID of the environment that contains this virtual machine, if any.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// The expiration date for VM.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// The fully-qualified domain name of the virtual machine.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The Microsoft Azure Marketplace image reference of the virtual machine.
        /// </summary>
        [Input("galleryImageReference")]
        public Input<Inputs.GalleryImageReferenceArgs>? GalleryImageReference { get; set; }

        /// <summary>
        /// Indicates whether this virtual machine uses an SSH key for authentication.
        /// </summary>
        [Input("isAuthenticationWithSshKey")]
        public Input<bool>? IsAuthenticationWithSshKey { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The lab subnet name of the virtual machine.
        /// </summary>
        [Input("labSubnetName")]
        public Input<string>? LabSubnetName { get; set; }

        /// <summary>
        /// The lab virtual network identifier of the virtual machine.
        /// </summary>
        [Input("labVirtualNetworkId")]
        public Input<string>? LabVirtualNetworkId { get; set; }

        /// <summary>
        /// Last known compute power state captured in DTL
        /// </summary>
        [Input("lastKnownPowerState")]
        public Input<string>? LastKnownPowerState { get; set; }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The network interface properties.
        /// </summary>
        [Input("networkInterface")]
        public Input<Inputs.NetworkInterfacePropertiesArgs>? NetworkInterface { get; set; }

        /// <summary>
        /// The notes of the virtual machine.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The OS type of the virtual machine.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The object identifier of the owner of the virtual machine.
        /// </summary>
        [Input("ownerObjectId")]
        public Input<string>? OwnerObjectId { get; set; }

        /// <summary>
        /// The user principal name of the virtual machine owner.
        /// </summary>
        [Input("ownerUserPrincipalName")]
        public Input<string>? OwnerUserPrincipalName { get; set; }

        /// <summary>
        /// The password of the virtual machine administrator.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The id of the plan associated with the virtual machine image
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("scheduleParameters")]
        private InputList<Inputs.ScheduleCreationParameterArgs>? _scheduleParameters;

        /// <summary>
        /// Virtual Machine schedules to be created
        /// </summary>
        public InputList<Inputs.ScheduleCreationParameterArgs> ScheduleParameters
        {
            get => _scheduleParameters ?? (_scheduleParameters = new InputList<Inputs.ScheduleCreationParameterArgs>());
            set => _scheduleParameters = value;
        }

        /// <summary>
        /// The size of the virtual machine.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The SSH key of the virtual machine administrator.
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        /// <summary>
        /// Storage type to use for virtual machine (i.e. Standard, Premium).
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

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
        /// The user name of the virtual machine.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Tells source of creation of lab virtual machine. Output property only.
        /// </summary>
        [Input("virtualMachineCreationSource")]
        public Input<string>? VirtualMachineCreationSource { get; set; }

        public VirtualMachineArgs()
        {
        }
    }
}
