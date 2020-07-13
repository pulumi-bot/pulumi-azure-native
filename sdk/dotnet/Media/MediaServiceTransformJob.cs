// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media
{
    /// <summary>
    /// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
    /// </summary>
    public partial class MediaServiceTransformJob : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.JobPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MediaServiceTransformJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaServiceTransformJob(string name, MediaServiceTransformJobArgs args, CustomResourceOptions? options = null)
            : base("azurerm:media:MediaServiceTransformJob", name, args ?? new MediaServiceTransformJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaServiceTransformJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:media:MediaServiceTransformJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MediaServiceTransformJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaServiceTransformJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MediaServiceTransformJob(name, id, options);
        }
    }

    public sealed class MediaServiceTransformJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The Job name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The resource properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.JobPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Transform name.
        /// </summary>
        [Input("transformName", required: true)]
        public Input<string> TransformName { get; set; } = null!;

        public MediaServiceTransformJobArgs()
        {
        }
    }
}