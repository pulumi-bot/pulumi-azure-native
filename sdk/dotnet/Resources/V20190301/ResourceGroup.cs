// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20190301
{
    /// <summary>
    /// Resource group information.
    /// 
    /// ## Example Usage
    /// ### Create or update a resource group
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var resourceGroup = new AzureRM.Resources.V20190301.ResourceGroup("resourceGroup", new AzureRM.Resources.V20190301.ResourceGroupArgs
    ///         {
    ///             Location = "eastus",
    ///             ResourceGroupName = "myResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ResourceGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource that manages this resource group.
        /// </summary>
        [Output("managedBy")]
        public Output<string?> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource group properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ResourceGroupPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The tags attached to the resource group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource group.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceGroup(string name, ResourceGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:resources/v20190301:ResourceGroup", name, args ?? new ResourceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:resources/v20190301:ResourceGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:resources/latest:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20151101:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20160201:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20160701:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20160901:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20170510:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20180201:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20180501:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20190501:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20190510:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20190701:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20190801:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20191001:ResourceGroup"},
                    new Pulumi.Alias { Type = "azurerm:resources/v20200601:ResourceGroup"},
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
        /// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The ID of the resource that manages this resource group.
        /// </summary>
        [Input("managedBy")]
        public Input<string>? ManagedBy { get; set; }

        /// <summary>
        /// The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags attached to the resource group.
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
