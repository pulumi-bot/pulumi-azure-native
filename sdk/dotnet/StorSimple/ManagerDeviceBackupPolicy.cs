// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple
{
    /// <summary>
    /// The backup policy.
    /// </summary>
    public partial class ManagerDeviceBackupPolicy : Pulumi.CustomResource
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
        /// The properties of the backup policy.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.BackupPolicyPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagerDeviceBackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagerDeviceBackupPolicy(string name, ManagerDeviceBackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storsimple:ManagerDeviceBackupPolicy", name, args ?? new ManagerDeviceBackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagerDeviceBackupPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storsimple:ManagerDeviceBackupPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ManagerDeviceBackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagerDeviceBackupPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagerDeviceBackupPolicy(name, id, options);
        }
    }

    public sealed class ManagerDeviceBackupPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the backup policy to be created/updated.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties of the backup policy.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.BackupPolicyPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ManagerDeviceBackupPolicyArgs()
        {
        }
    }
}
