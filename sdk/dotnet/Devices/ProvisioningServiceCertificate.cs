// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices
{
    /// <summary>
    /// The X509 Certificate.
    /// </summary>
    public partial class ProvisioningServiceCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The entity tag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// properties of a certificate
        /// </summary>
        [Output("properties")]
        public Output<Outputs.CertificatePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProvisioningServiceCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProvisioningServiceCertificate(string name, ProvisioningServiceCertificateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devices:ProvisioningServiceCertificate", name, args ?? new ProvisioningServiceCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProvisioningServiceCertificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devices:ProvisioningServiceCertificate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ProvisioningServiceCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProvisioningServiceCertificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProvisioningServiceCertificate(name, id, options);
        }
    }

    public sealed class ProvisioningServiceCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// The name of the certificate create or update.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the provisioning service.
        /// </summary>
        [Input("provisioningServiceName", required: true)]
        public Input<string> ProvisioningServiceName { get; set; } = null!;

        /// <summary>
        /// Resource group identifier.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ProvisioningServiceCertificateArgs()
        {
        }
    }
}
