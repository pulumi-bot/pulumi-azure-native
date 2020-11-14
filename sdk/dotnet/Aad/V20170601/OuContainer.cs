// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Aad.V20170601
{
    /// <summary>
    /// Resource for OuContainer.
    /// </summary>
    public partial class OuContainer : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of container accounts
        /// </summary>
        [Output("accounts")]
        public Output<ImmutableArray<Outputs.ContainerAccountResponse>> Accounts { get; private set; } = null!;

        /// <summary>
        /// The OuContainer name
        /// </summary>
        [Output("containerId")]
        public Output<string> ContainerId { get; private set; } = null!;

        /// <summary>
        /// The Deployment id
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// Distinguished Name of OuContainer instance
        /// </summary>
        [Output("distinguishedName")]
        public Output<string> DistinguishedName { get; private set; } = null!;

        /// <summary>
        /// The domain name of Domain Services.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Resource etag
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Status of OuContainer instance
        /// </summary>
        [Output("serviceStatus")]
        public Output<string> ServiceStatus { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure Active Directory tenant id
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a OuContainer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OuContainer(string name, OuContainerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:aad/v20170601:OuContainer", name, args ?? new OuContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OuContainer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:aad/v20170601:OuContainer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:aad/latest:OuContainer"},
                    new Pulumi.Alias { Type = "azure-nextgen:aad/v20200101:OuContainer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OuContainer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OuContainer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OuContainer(name, id, options);
        }
    }

    public sealed class OuContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account name
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The name of the domain service.
        /// </summary>
        [Input("domainServiceName", required: true)]
        public Input<string> DomainServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the OuContainer.
        /// </summary>
        [Input("ouContainerName", required: true)]
        public Input<string> OuContainerName { get; set; } = null!;

        /// <summary>
        /// The account password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The account spn
        /// </summary>
        [Input("spn")]
        public Input<string>? Spn { get; set; }

        public OuContainerArgs()
        {
        }
    }
}
