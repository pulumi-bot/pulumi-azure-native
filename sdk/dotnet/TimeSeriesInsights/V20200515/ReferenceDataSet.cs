// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.TimeSeriesInsights.V20200515
{
    /// <summary>
    /// A reference data set provides metadata about the events in an environment. Metadata in the reference data set will be joined with events as they are read from event sources. The metadata that makes up the reference data set is uploaded or modified through the Time Series Insights data plane APIs.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:timeseriesinsights/v20200515:ReferenceDataSet")]
    public partial class ReferenceDataSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the resource was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
        /// </summary>
        [Output("dataStringComparisonBehavior")]
        public Output<string?> DataStringComparisonBehavior { get; private set; } = null!;

        /// <summary>
        /// The list of key properties for the reference data set.
        /// </summary>
        [Output("keyProperties")]
        public Output<ImmutableArray<Outputs.ReferenceDataSetKeyPropertyResponse>> KeyProperties { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a ReferenceDataSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReferenceDataSet(string name, ReferenceDataSetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:timeseriesinsights/v20200515:ReferenceDataSet", name, args ?? new ReferenceDataSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReferenceDataSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:timeseriesinsights/v20200515:ReferenceDataSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/latest:ReferenceDataSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20170228preview:ReferenceDataSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20171115:ReferenceDataSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20180815preview:ReferenceDataSet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReferenceDataSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReferenceDataSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReferenceDataSet(name, id, options);
        }
    }

    public sealed class ReferenceDataSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
        /// </summary>
        [Input("dataStringComparisonBehavior")]
        public InputUnion<string, Pulumi.AzureNextGen.TimeSeriesInsights.V20200515.DataStringComparisonBehavior>? DataStringComparisonBehavior { get; set; }

        /// <summary>
        /// The name of the Time Series Insights environment associated with the specified resource group.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        [Input("keyProperties", required: true)]
        private InputList<Inputs.ReferenceDataSetKeyPropertyArgs>? _keyProperties;

        /// <summary>
        /// The list of key properties for the reference data set.
        /// </summary>
        public InputList<Inputs.ReferenceDataSetKeyPropertyArgs> KeyProperties
        {
            get => _keyProperties ?? (_keyProperties = new InputList<Inputs.ReferenceDataSetKeyPropertyArgs>());
            set => _keyProperties = value;
        }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the reference data set.
        /// </summary>
        [Input("referenceDataSetName", required: true)]
        public Input<string> ReferenceDataSetName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of additional properties for the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ReferenceDataSetArgs()
        {
        }
    }
}
