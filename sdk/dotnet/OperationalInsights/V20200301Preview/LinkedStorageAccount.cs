// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.OperationalInsights.V20200301Preview
{
    /// <summary>
    /// Linked storage accounts top level resource container.
    /// </summary>
    public partial class LinkedStorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Linked storage accounts type.
        /// </summary>
        [Output("dataSourceType")]
        public Output<string> DataSourceType { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Linked storage accounts resources ids.
        /// </summary>
        [Output("storageAccountIds")]
        public Output<ImmutableArray<string>> StorageAccountIds { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedStorageAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedStorageAccount(string name, LinkedStorageAccountArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:operationalinsights/v20200301preview:LinkedStorageAccount", name, args ?? new LinkedStorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedStorageAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:operationalinsights/v20200301preview:LinkedStorageAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:operationalinsights/latest:LinkedStorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:operationalinsights/v20190801preview:LinkedStorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:operationalinsights/v20200801:LinkedStorageAccount"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LinkedStorageAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkedStorageAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LinkedStorageAccount(name, id, options);
        }
    }

    public sealed class LinkedStorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Linked storage accounts type.
        /// </summary>
        [Input("dataSourceType", required: true)]
        public Input<string> DataSourceType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("storageAccountIds")]
        private InputList<string>? _storageAccountIds;

        /// <summary>
        /// Linked storage accounts resources ids.
        /// </summary>
        public InputList<string> StorageAccountIds
        {
            get => _storageAccountIds ?? (_storageAccountIds = new InputList<string>());
            set => _storageAccountIds = value;
        }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public LinkedStorageAccountArgs()
        {
        }
    }
}
