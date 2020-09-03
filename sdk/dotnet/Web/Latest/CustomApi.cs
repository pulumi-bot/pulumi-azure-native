// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.Latest
{
    /// <summary>
    /// A custom API
    /// 
    /// ## Example Usage
    /// ### Replace a custom API
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var customApi = new AzureRM.Web.Latest.CustomApi("customApi", new AzureRM.Web.Latest.CustomApiArgs
    ///         {
    ///             ApiName = "testCustomApi",
    ///             ResourceGroupName = "testResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class CustomApi : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource ETag
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom API properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.CustomApiPropertiesDefinitionResponseResult> Properties { get; private set; } = null!;

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
        /// Create a CustomApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomApi(string name, CustomApiArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/latest:CustomApi", name, args ?? new CustomApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomApi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/latest:CustomApi", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:web/v20160601:CustomApi"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomApi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomApi(name, id, options);
        }
    }

    public sealed class CustomApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API name
        /// </summary>
        [Input("apiName", required: true)]
        public Input<string> ApiName { get; set; } = null!;

        /// <summary>
        /// Resource ETag
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Custom API properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.CustomApiPropertiesDefinitionArgs>? Properties { get; set; }

        /// <summary>
        /// The resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public CustomApiArgs()
        {
        }
    }
}
