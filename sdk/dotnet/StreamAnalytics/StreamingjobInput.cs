// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StreamAnalytics
{
    /// <summary>
    /// An input object, containing all information associated with the named input. All inputs are contained under a streaming job.
    /// </summary>
    public partial class StreamingjobInput : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.InputPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StreamingjobInput resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamingjobInput(string name, StreamingjobInputArgs args, CustomResourceOptions? options = null)
            : base("azurerm:streamanalytics:StreamingjobInput", name, args ?? new StreamingjobInputArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamingjobInput(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:streamanalytics:StreamingjobInput", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamingjobInput resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamingjobInput Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamingjobInput(name, id, options);
        }
    }

    public sealed class StreamingjobInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the streaming job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        /// <summary>
        /// The name of the input.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.InputPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public StreamingjobInputArgs()
        {
        }
    }
}
