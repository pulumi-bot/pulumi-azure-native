// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.Latest
{
    /// <summary>
    /// ## Example Usage
    /// ### HyperVCollectors_Create
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var hyperVCollector = new AzureRM.Migrate.Latest.HyperVCollector("hyperVCollector", new AzureRM.Migrate.Latest.HyperVCollectorArgs
    ///         {
    ///             ETag = "\"00000981-0000-0300-0000-5d74cd5f0000\"",
    ///             HyperVCollectorName = "migrateprojectce73collector",
    ///             ProjectName = "migrateprojectce73project",
    ///             ResourceGroupName = "contosoithyperv",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class HyperVCollector : Pulumi.CustomResource
    {
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.CollectorPropertiesResponseResult> Properties { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a HyperVCollector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HyperVCollector(string name, HyperVCollectorArgs args, CustomResourceOptions? options = null)
            : base("azurerm:migrate/latest:HyperVCollector", name, args ?? new HyperVCollectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HyperVCollector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:migrate/latest:HyperVCollector", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:migrate/v20191001:HyperVCollector"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HyperVCollector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HyperVCollector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HyperVCollector(name, id, options);
        }
    }

    public sealed class HyperVCollectorArgs : Pulumi.ResourceArgs
    {
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Unique name of a Hyper-V collector within a project.
        /// </summary>
        [Input("hyperVCollectorName", required: true)]
        public Input<string> HyperVCollectorName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("properties")]
        public Input<Inputs.CollectorPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public HyperVCollectorArgs()
        {
        }
    }
}
