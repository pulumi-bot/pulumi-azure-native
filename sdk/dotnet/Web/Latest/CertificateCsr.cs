// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.Latest
{
    /// <summary>
    /// Certificate signing request object
    /// Latest API Version: 2015-08-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:web/latest:CertificateCsr")]
    public partial class CertificateCsr : Pulumi.CustomResource
    {
        /// <summary>
        /// Actual CSR string created
        /// </summary>
        [Output("csrString")]
        public Output<string?> CsrString { get; private set; } = null!;

        /// <summary>
        /// Distinguished name of certificate to be created
        /// </summary>
        [Output("distinguishedName")]
        public Output<string?> DistinguishedName { get; private set; } = null!;

        /// <summary>
        /// Hosting environment
        /// </summary>
        [Output("hostingEnvironment")]
        public Output<string?> HostingEnvironment { get; private set; } = null!;

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// PFX password
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// PFX certificate of created certificate
        /// </summary>
        [Output("pfxBlob")]
        public Output<string?> PfxBlob { get; private set; } = null!;

        /// <summary>
        /// Hash of the certificates public key
        /// </summary>
        [Output("publicKeyHash")]
        public Output<string?> PublicKeyHash { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateCsr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateCsr(string name, CertificateCsrArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/latest:CertificateCsr", name, args ?? new CertificateCsrArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateCsr(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/latest:CertificateCsr", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20150801:CertificateCsr"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateCsr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateCsr Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateCsr(name, id, options);
        }
    }

    public sealed class CertificateCsrArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actual CSR string created
        /// </summary>
        [Input("csrString")]
        public Input<string>? CsrString { get; set; }

        /// <summary>
        /// Distinguished name of certificate to be created
        /// </summary>
        [Input("distinguishedName")]
        public Input<string>? DistinguishedName { get; set; }

        /// <summary>
        /// Hosting environment
        /// </summary>
        [Input("hostingEnvironment")]
        public Input<string>? HostingEnvironment { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// PFX password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// PFX certificate of created certificate
        /// </summary>
        [Input("pfxBlob")]
        public Input<string>? PfxBlob { get; set; }

        /// <summary>
        /// Hash of the certificates public key
        /// </summary>
        [Input("publicKeyHash")]
        public Input<string>? PublicKeyHash { get; set; }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateCsrArgs()
        {
        }
    }
}
