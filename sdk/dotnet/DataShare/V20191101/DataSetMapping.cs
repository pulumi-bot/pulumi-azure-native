// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.V20191101
{
    /// <summary>
    /// A data set mapping data transfer object.
    /// </summary>
    public partial class DataSetMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of data set mapping.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataSetMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSetMapping(string name, DataSetMappingArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datashare/v20191101:DataSetMapping", name, args ?? new DataSetMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSetMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datashare/v20191101:DataSetMapping", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datashare/latest:DataSetMapping"},
                    new Pulumi.Alias { Type = "azurerm:datashare/v20181101preview:DataSetMapping"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSetMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSetMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataSetMapping(name, id, options);
        }
    }

    public sealed class DataSetMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the data set mapping to be created.
        /// </summary>
        [Input("dataSetMappingName", required: true)]
        public Input<string> DataSetMappingName { get; set; } = null!;

        /// <summary>
        /// Kind of data set mapping.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share subscription which will hold the data set sink.
        /// </summary>
        [Input("shareSubscriptionName", required: true)]
        public Input<string> ShareSubscriptionName { get; set; } = null!;

        public DataSetMappingArgs()
        {
        }
    }
}
