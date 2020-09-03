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
    ///         var console = new AzureRM.Portal.V20181001.Console("console", new AzureRM.Portal.V20181001.ConsoleArgs
    ///         {
    ///             ConsoleName = "default",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Console : Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud shell console properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ConsolePropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a Console resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Console(string name, ConsoleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001:Console", name, args ?? new ConsoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Console(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001:Console", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:portal/latest:Console"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Console resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Console Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Console(name, id, options);
        }
    }

    public sealed class ConsoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the console
        /// </summary>
        [Input("consoleName", required: true)]
        public Input<string> ConsoleName { get; set; } = null!;

        /// <summary>
        /// Cloud shell properties for creating a console.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.ConsoleCreatePropertiesArgs> Properties { get; set; } = null!;

        public ConsoleArgs()
        {
        }
    }
}
