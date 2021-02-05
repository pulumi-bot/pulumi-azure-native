// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataShare.Latest
{
    /// <summary>
    /// A share data transfer object.
    /// Latest API Version: 2020-09-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:datashare/latest:Share")]
    public partial class Share : Pulumi.CustomResource
    {
        /// <summary>
        /// Time at which the share was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Share description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the provisioning state
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Share kind.
        /// </summary>
        [Output("shareKind")]
        public Output<string?> ShareKind { get; private set; } = null!;

        /// <summary>
        /// System Data of the Azure resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.ProxyDtoResponseSystemData> SystemData { get; private set; } = null!;

        /// <summary>
        /// Share terms.
        /// </summary>
        [Output("terms")]
        public Output<string?> Terms { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Email of the user who created the resource
        /// </summary>
        [Output("userEmail")]
        public Output<string> UserEmail { get; private set; } = null!;

        /// <summary>
        /// Name of the user who created the resource
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a Share resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Share(string name, ShareArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:datashare/latest:Share", name, args ?? new ShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Share(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:datashare/latest:Share", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:datashare/v20181101preview:Share"},
                    new Pulumi.Alias { Type = "azure-nextgen:datashare/v20191101:Share"},
                    new Pulumi.Alias { Type = "azure-nextgen:datashare/v20200901:Share"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Share resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Share Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Share(name, id, options);
        }
    }

    public sealed class ShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Share description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Share kind.
        /// </summary>
        [Input("shareKind")]
        public InputUnion<string, Pulumi.AzureNextGen.DataShare.Latest.ShareKind>? ShareKind { get; set; }

        /// <summary>
        /// The name of the share.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        /// <summary>
        /// Share terms.
        /// </summary>
        [Input("terms")]
        public Input<string>? Terms { get; set; }

        public ShareArgs()
        {
        }
    }
}
