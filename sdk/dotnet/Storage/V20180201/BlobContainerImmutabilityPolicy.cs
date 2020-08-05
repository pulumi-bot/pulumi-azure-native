// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20180201
{
    /// <summary>
    /// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
    /// </summary>
    public partial class BlobContainerImmutabilityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of an ImmutabilityPolicy of a blob container.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ImmutabilityPolicyPropertyResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BlobContainerImmutabilityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlobContainerImmutabilityPolicy(string name, BlobContainerImmutabilityPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storage/v20180201:BlobContainerImmutabilityPolicy", name, args ?? new BlobContainerImmutabilityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlobContainerImmutabilityPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storage/v20180201:BlobContainerImmutabilityPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BlobContainerImmutabilityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlobContainerImmutabilityPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BlobContainerImmutabilityPolicy(name, id, options);
        }
    }

    public sealed class BlobContainerImmutabilityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The immutability period for the blobs in the container since the policy creation, in days.
        /// </summary>
        [Input("immutabilityPeriodSinceCreationInDays", required: true)]
        public Input<int> ImmutabilityPeriodSinceCreationInDays { get; set; } = null!;

        /// <summary>
        /// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public BlobContainerImmutabilityPolicyArgs()
        {
        }
    }
}
