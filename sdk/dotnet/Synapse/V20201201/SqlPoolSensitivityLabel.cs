// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20201201
{
    /// <summary>
    /// A sensitivity label.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:synapse/v20201201:SqlPoolSensitivityLabel")]
    public partial class SqlPoolSensitivityLabel : Pulumi.CustomResource
    {
        /// <summary>
        /// The column name.
        /// </summary>
        [Output("columnName")]
        public Output<string> ColumnName { get; private set; } = null!;

        /// <summary>
        /// The information type.
        /// </summary>
        [Output("informationType")]
        public Output<string?> InformationType { get; private set; } = null!;

        /// <summary>
        /// The information type ID.
        /// </summary>
        [Output("informationTypeId")]
        public Output<string?> InformationTypeId { get; private set; } = null!;

        /// <summary>
        /// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
        /// </summary>
        [Output("isDisabled")]
        public Output<bool> IsDisabled { get; private set; } = null!;

        /// <summary>
        /// The label ID.
        /// </summary>
        [Output("labelId")]
        public Output<string?> LabelId { get; private set; } = null!;

        /// <summary>
        /// The label name.
        /// </summary>
        [Output("labelName")]
        public Output<string?> LabelName { get; private set; } = null!;

        /// <summary>
        /// managed by
        /// </summary>
        [Output("managedBy")]
        public Output<string> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("rank")]
        public Output<string?> Rank { get; private set; } = null!;

        /// <summary>
        /// The schema name.
        /// </summary>
        [Output("schemaName")]
        public Output<string> SchemaName { get; private set; } = null!;

        /// <summary>
        /// The table name.
        /// </summary>
        [Output("tableName")]
        public Output<string> TableName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SqlPoolSensitivityLabel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlPoolSensitivityLabel(string name, SqlPoolSensitivityLabelArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20201201:SqlPoolSensitivityLabel", name, args ?? new SqlPoolSensitivityLabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlPoolSensitivityLabel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20201201:SqlPoolSensitivityLabel", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/latest:SqlPoolSensitivityLabel"},
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/v20190601preview:SqlPoolSensitivityLabel"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlPoolSensitivityLabel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlPoolSensitivityLabel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlPoolSensitivityLabel(name, id, options);
        }
    }

    public sealed class SqlPoolSensitivityLabelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the column.
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        /// <summary>
        /// The information type.
        /// </summary>
        [Input("informationType")]
        public Input<string>? InformationType { get; set; }

        /// <summary>
        /// The information type ID.
        /// </summary>
        [Input("informationTypeId")]
        public Input<string>? InformationTypeId { get; set; }

        /// <summary>
        /// The label ID.
        /// </summary>
        [Input("labelId")]
        public Input<string>? LabelId { get; set; }

        /// <summary>
        /// The label name.
        /// </summary>
        [Input("labelName")]
        public Input<string>? LabelName { get; set; }

        [Input("rank")]
        public Input<Pulumi.AzureNextGen.Synapse.V20201201.SensitivityLabelRank>? Rank { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the schema.
        /// </summary>
        [Input("schemaName", required: true)]
        public Input<string> SchemaName { get; set; } = null!;

        /// <summary>
        /// The source of the sensitivity label.
        /// </summary>
        [Input("sensitivityLabelSource", required: true)]
        public Input<string> SensitivityLabelSource { get; set; } = null!;

        /// <summary>
        /// SQL pool name
        /// </summary>
        [Input("sqlPoolName", required: true)]
        public Input<string> SqlPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the table.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public SqlPoolSensitivityLabelArgs()
        {
        }
    }
}
