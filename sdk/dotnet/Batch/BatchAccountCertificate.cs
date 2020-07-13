// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch
{
    /// <summary>
    /// Contains information about a certificate.
    /// </summary>
    public partial class BatchAccountCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties associated with the certificate.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.CertificatePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BatchAccountCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BatchAccountCertificate(string name, BatchAccountCertificateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:batch:BatchAccountCertificate", name, args ?? new BatchAccountCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BatchAccountCertificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:batch:BatchAccountCertificate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BatchAccountCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BatchAccountCertificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BatchAccountCertificate(name, id, options);
        }
    }

    public sealed class BatchAccountCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties associated with the certificate.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.CertificateCreateOrUpdatePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public BatchAccountCertificateArgs()
        {
        }
    }
}