// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CostManagement.V20180531
{
    /// <summary>
    /// A report config resource.
    /// </summary>
    public partial class ReportConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Has definition for the report config.
        /// </summary>
        [Output("definition")]
        public Output<Outputs.ReportConfigDefinitionResponse> Definition { get; private set; } = null!;

        /// <summary>
        /// Has delivery information for the report config.
        /// </summary>
        [Output("deliveryInfo")]
        public Output<Outputs.ReportConfigDeliveryInfoResponse> DeliveryInfo { get; private set; } = null!;

        /// <summary>
        /// The format of the report being delivered.
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Has schedule information for the report config.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ReportConfigScheduleResponse?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ReportConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReportConfig(string name, ReportConfigArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:costmanagement/v20180531:ReportConfig", name, args ?? new ReportConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReportConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:costmanagement/v20180531:ReportConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:costmanagement/latest:ReportConfig"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReportConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReportConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReportConfig(name, id, options);
        }
    }

    public sealed class ReportConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Has definition for the report config.
        /// </summary>
        [Input("definition", required: true)]
        public Input<Inputs.ReportConfigDefinitionArgs> Definition { get; set; } = null!;

        /// <summary>
        /// Has delivery information for the report config.
        /// </summary>
        [Input("deliveryInfo", required: true)]
        public Input<Inputs.ReportConfigDeliveryInfoArgs> DeliveryInfo { get; set; } = null!;

        /// <summary>
        /// The format of the report being delivered.
        /// </summary>
        [Input("format")]
        public InputUnion<string, Pulumi.AzureNextGen.CostManagement.V20180531.FormatType>? Format { get; set; }

        /// <summary>
        /// Report Config Name.
        /// </summary>
        [Input("reportConfigName", required: true)]
        public Input<string> ReportConfigName { get; set; } = null!;

        /// <summary>
        /// Has schedule information for the report config.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ReportConfigScheduleArgs>? Schedule { get; set; }

        public ReportConfigArgs()
        {
        }
    }
}
