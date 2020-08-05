// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601
{
    /// <summary>
    /// Factory resource type.
    /// </summary>
    public partial class Factory : Pulumi.CustomResource
    {
        /// <summary>
        /// Etag identifies change in the resource.
        /// </summary>
        [Output("eTag")]
        public Output<string> ETag { get; private set; } = null!;

        /// <summary>
        /// Managed service identity of the factory.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.FactoryIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the factory.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.FactoryPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Factory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Factory(string name, FactoryArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/v20180601:Factory", name, args ?? new FactoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Factory(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/v20180601:Factory", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Factory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Factory Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Factory(name, id, options);
        }
    }

    public sealed class FactoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of parameters for factory.
        /// </summary>
        [Input("globalParameters")]
        public Input<Inputs.GlobalParameterDefinitionSpecificationArgs>? GlobalParameters { get; set; }

        /// <summary>
        /// Managed service identity of the factory.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.FactoryIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The factory name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Git repo information of the factory.
        /// </summary>
        [Input("repoConfiguration")]
        public Input<Inputs.FactoryRepoConfigurationArgs>? RepoConfiguration { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FactoryArgs()
        {
        }
    }
}
