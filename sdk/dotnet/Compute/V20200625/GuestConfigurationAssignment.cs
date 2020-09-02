// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20200625
{
    /// <summary>
    /// Guest configuration assignment is an association between a machine and guest configuration.
    /// </summary>
    public partial class GuestConfigurationAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// Region where the VM is located.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the guest configuration assignment.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the Guest configuration assignment.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.GuestConfigurationAssignmentPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GuestConfigurationAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GuestConfigurationAssignment(string name, GuestConfigurationAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20200625:GuestConfigurationAssignment", name, args ?? new GuestConfigurationAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GuestConfigurationAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20200625:GuestConfigurationAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/latest:GuestConfigurationAssignment"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20180630preview:GuestConfigurationAssignment"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20181120:GuestConfigurationAssignment"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GuestConfigurationAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GuestConfigurationAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GuestConfigurationAssignment(name, id, options);
        }
    }

    public sealed class GuestConfigurationAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the guest configuration assignment.
        /// </summary>
        [Input("guestConfigurationAssignmentName", required: true)]
        public Input<string> GuestConfigurationAssignmentName { get; set; } = null!;

        /// <summary>
        /// Region where the VM is located.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the guest configuration assignment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Properties of the Guest configuration assignment.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.GuestConfigurationAssignmentPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("vmName", required: true)]
        public Input<string> VmName { get; set; } = null!;

        public GuestConfigurationAssignmentArgs()
        {
        }
    }
}
