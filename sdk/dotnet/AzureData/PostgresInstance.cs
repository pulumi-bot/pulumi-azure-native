// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureData
{
    /// <summary>
    /// A Postgres Instance.
    /// API Version: 2020-09-08-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:azuredata:PostgresInstance")]
    public partial class PostgresInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// The instance admin
        /// </summary>
        [Output("admin")]
        public Output<string?> Admin { get; private set; } = null!;

        /// <summary>
        /// The data controller id
        /// </summary>
        [Output("dataControllerId")]
        public Output<string?> DataControllerId { get; private set; } = null!;

        /// <summary>
        /// The raw kubernetes information
        /// </summary>
        [Output("k8sRaw")]
        public Output<object?> K8sRaw { get; private set; } = null!;

        /// <summary>
        /// Last uploaded date from on premise cluster. Defaults to current date time
        /// </summary>
        [Output("lastUploadedDate")]
        public Output<string?> LastUploadedDate { get; private set; } = null!;

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
        /// Create a PostgresInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PostgresInstance(string name, PostgresInstanceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:azuredata:PostgresInstance", name, args ?? new PostgresInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PostgresInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:azuredata:PostgresInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:azuredata:PostgresInstance"},
                    new Pulumi.Alias { Type = "azure-native:azuredata/v20190724preview:PostgresInstance"},
                    new Pulumi.Alias { Type = "azure-nextgen:azuredata/v20190724preview:PostgresInstance"},
                    new Pulumi.Alias { Type = "azure-native:azuredata/v20200908preview:PostgresInstance"},
                    new Pulumi.Alias { Type = "azure-nextgen:azuredata/v20200908preview:PostgresInstance"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PostgresInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PostgresInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PostgresInstance(name, id, options);
        }
    }

    public sealed class PostgresInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance admin
        /// </summary>
        [Input("admin")]
        public Input<string>? Admin { get; set; }

        /// <summary>
        /// The data controller id
        /// </summary>
        [Input("dataControllerId")]
        public Input<string>? DataControllerId { get; set; }

        /// <summary>
        /// The raw kubernetes information
        /// </summary>
        [Input("k8sRaw")]
        public Input<object>? K8sRaw { get; set; }

        /// <summary>
        /// Last uploaded date from on premise cluster. Defaults to current date time
        /// </summary>
        [Input("lastUploadedDate")]
        public Input<string>? LastUploadedDate { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of PostgresInstance
        /// </summary>
        [Input("postgresInstanceName")]
        public Input<string>? PostgresInstanceName { get; set; }

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public PostgresInstanceArgs()
        {
        }
    }
}
