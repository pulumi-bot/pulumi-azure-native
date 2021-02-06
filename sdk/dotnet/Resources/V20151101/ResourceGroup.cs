// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20151101
{
    /// <summary>
    /// Resource group information.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:resources/v20151101:ResourceGroup")]
    public partial class ResourceGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the location of the resource group. It cannot be changed after the resource group has been created. Has to be one of the supported Azure Locations, such as West US, East US, West Europe, East Asia, etc.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the Name of the resource group.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The resource group properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ResourceGroupPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the tags attached to the resource group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceGroup(string name, ResourceGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20151101:ResourceGroup", name, args ?? new ResourceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20151101:ResourceGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:resources/latest:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20160201:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20160701:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20160901:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20170510:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20180201:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20180501:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190301:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190501:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190510:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190701:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190801:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20191001:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20200601:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20200801:ResourceGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20201001:ResourceGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceGroup(name, id, options);
        }
    }

    public sealed class ResourceGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the location of the resource group. It cannot be changed after the resource group has been created. Has to be one of the supported Azure Locations, such as West US, East US, West Europe, East Asia, etc.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Gets or sets the Name of the resource group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group to be created or updated.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets the tags attached to the resource group.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResourceGroupArgs()
        {
        }
    }
}
