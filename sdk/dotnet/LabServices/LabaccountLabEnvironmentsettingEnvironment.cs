// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.LabServices
{
    /// <summary>
    /// Represents an environment instance
    /// </summary>
    public partial class LabaccountLabEnvironmentsettingEnvironment : Pulumi.CustomResource
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
        /// The properties of the Environment resource
        /// </summary>
        [Output("properties")]
        public Output<Outputs.EnvironmentPropertiesResponse> Properties { get; private set; } = null!;

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
        /// Create a LabaccountLabEnvironmentsettingEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LabaccountLabEnvironmentsettingEnvironment(string name, LabaccountLabEnvironmentsettingEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:labservices:LabaccountLabEnvironmentsettingEnvironment", name, args ?? new LabaccountLabEnvironmentsettingEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LabaccountLabEnvironmentsettingEnvironment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:labservices:LabaccountLabEnvironmentsettingEnvironment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LabaccountLabEnvironmentsettingEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LabaccountLabEnvironmentsettingEnvironment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LabaccountLabEnvironmentsettingEnvironment(name, id, options);
        }
    }

    public sealed class LabaccountLabEnvironmentsettingEnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the environment Setting.
        /// </summary>
        [Input("environmentSettingName", required: true)]
        public Input<string> EnvironmentSettingName { get; set; } = null!;

        /// <summary>
        /// The name of the lab Account.
        /// </summary>
        [Input("labAccountName", required: true)]
        public Input<string> LabAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties of the Environment resource
        /// </summary>
        [Input("properties")]
        public Input<Inputs.EnvironmentPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public LabaccountLabEnvironmentsettingEnvironmentArgs()
        {
        }
    }
}
