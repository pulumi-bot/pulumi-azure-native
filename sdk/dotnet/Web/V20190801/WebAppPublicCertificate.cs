// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20190801
{
    /// <summary>
    /// Public certificate object
    /// </summary>
    public partial class WebAppPublicCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// PublicCertificate resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PublicCertificateResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppPublicCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppPublicCertificate(string name, WebAppPublicCertificateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20190801:WebAppPublicCertificate", name, args ?? new WebAppPublicCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppPublicCertificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20190801:WebAppPublicCertificate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebAppPublicCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppPublicCertificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppPublicCertificate(name, id, options);
        }
    }

    public sealed class WebAppPublicCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public Certificate byte array
        /// </summary>
        [Input("blob")]
        public Input<string>? Blob { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Public certificate name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Public Certificate Location
        /// </summary>
        [Input("publicCertificateLocation")]
        public Input<string>? PublicCertificateLocation { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public WebAppPublicCertificateArgs()
        {
        }
    }
}
