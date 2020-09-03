// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.V20181001
{
    /// <summary>
    /// Cloud shell console
    /// 
    /// ## Example Usage
    /// ### PutConsole
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var consoleWithLocation = new AzureRM.Portal.V20181001.ConsoleWithLocation("consoleWithLocation", new AzureRM.Portal.V20181001.ConsoleWithLocationArgs
    ///         {
    ///             ConsoleName = "default",
    ///             Location = "eastus",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ConsoleWithLocation : Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud shell console properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ConsolePropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a ConsoleWithLocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsoleWithLocation(string name, ConsoleWithLocationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001:ConsoleWithLocation", name, args ?? new ConsoleWithLocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsoleWithLocation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001:ConsoleWithLocation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:portal/latest:ConsoleWithLocation"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConsoleWithLocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsoleWithLocation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConsoleWithLocation(name, id, options);
        }
    }

    public sealed class ConsoleWithLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the console
        /// </summary>
        [Input("consoleName", required: true)]
        public Input<string> ConsoleName { get; set; } = null!;

        /// <summary>
        /// The provider location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public ConsoleWithLocationArgs()
        {
        }
    }
}
