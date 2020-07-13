// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VMwareCloudSimple
{
    /// <summary>
    /// Virtual machine model
    /// </summary>
    public partial class VirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// Azure region
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// {virtualMachineName}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Virtual machine properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualMachinePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The list of tags
        /// </summary>
        [Output("tags")]
        public Output<Outputs.TagsResponse?> Tags { get; private set; } = null!;

        /// <summary>
        /// {resourceProviderNamespace}/{resourceType}
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
            : base("azurerm:vmwarecloudsimple:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:vmwarecloudsimple:VirtualMachine", name, null, MakeResourceOptions(options, id))
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
        /// referer url
        /// </summary>
        [Input("Referer", required: true)]
        public Input<string> Referer { get; set; } = null!;

        /// <summary>
        /// Azure region
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// virtual machine name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Virtual machine properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.VirtualMachinePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The list of tags
        /// </summary>
        [Input("tags")]
        public Input<Inputs.TagsArgs>? Tags { get; set; }

        public VirtualMachineArgs()
        {
        }
    }
}
