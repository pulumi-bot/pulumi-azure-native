// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AzureData.V20190724Preview
{
    /// <summary>
    /// A SqlManagedInstance.
    /// </summary>
    public partial class SqlManagedInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// The instance admin user
        /// </summary>
        [Output("admin")]
        public Output<string?> Admin { get; private set; } = null!;

        /// <summary>
        /// null
        /// </summary>
        [Output("dataControllerId")]
        public Output<string?> DataControllerId { get; private set; } = null!;

        /// <summary>
        /// The instance end time
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// The on premise instance endpoint
        /// </summary>
        [Output("instanceEndpoint")]
        public Output<string?> InstanceEndpoint { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The instance start time
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// Read only system data
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The instance vCore
        /// </summary>
        [Output("vCore")]
        public Output<string?> VCore { get; private set; } = null!;


        /// <summary>
        /// Create a SqlManagedInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlManagedInstance(string name, SqlManagedInstanceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:azuredata/v20190724preview:SqlManagedInstance", name, args ?? new SqlManagedInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlManagedInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:azuredata/v20190724preview:SqlManagedInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:azuredata/v20200908preview:SqlManagedInstance"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlManagedInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlManagedInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlManagedInstance(name, id, options);
        }
    }

    public sealed class SqlManagedInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance admin user
        /// </summary>
        [Input("admin")]
        public Input<string>? Admin { get; set; }

        /// <summary>
        /// null
        /// </summary>
        [Input("dataControllerId")]
        public Input<string>? DataControllerId { get; set; }

        /// <summary>
        /// The instance end time
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The on premise instance endpoint
        /// </summary>
        [Input("instanceEndpoint")]
        public Input<string>? InstanceEndpoint { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of SQL Managed Instances
        /// </summary>
        [Input("sqlManagedInstanceName", required: true)]
        public Input<string> SqlManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// The instance start time
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The instance vCore
        /// </summary>
        [Input("vCore")]
        public Input<string>? VCore { get; set; }

        public SqlManagedInstanceArgs()
        {
        }
    }
}
