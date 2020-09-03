// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.Latest
{
    /// <summary>
    /// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
    /// 
    /// ## Example Usage
    /// ### CreateOrUpdateImmutabilityPolicy
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var blobContainerImmutabilityPolicy = new AzureRM.Storage.Latest.BlobContainerImmutabilityPolicy("blobContainerImmutabilityPolicy", new AzureRM.Storage.Latest.BlobContainerImmutabilityPolicyArgs
    ///         {
    ///             AccountName = "sto7069",
    ///             AllowProtectedAppendWrites = true,
    ///             ContainerName = "container6397",
    ///             ImmutabilityPeriodSinceCreationInDays = 3,
    ///             ImmutabilityPolicyName = "default",
    ///             ResourceGroupName = "res1782",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class BlobContainerImmutabilityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API
        /// </summary>
        [Output("allowProtectedAppendWrites")]
        public Output<bool?> AllowProtectedAppendWrites { get; private set; } = null!;

        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The immutability period for the blobs in the container since the policy creation, in days.
        /// </summary>
        [Output("immutabilityPeriodSinceCreationInDays")]
        public Output<int?> ImmutabilityPeriodSinceCreationInDays { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

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
            : base("azurerm:storage/latest:BlobContainerImmutabilityPolicy", name, args ?? new BlobContainerImmutabilityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlobContainerImmutabilityPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storage/latest:BlobContainerImmutabilityPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storage/v20180201:BlobContainerImmutabilityPolicy"},
                    new Pulumi.Alias { Type = "azurerm:storage/v20180301preview:BlobContainerImmutabilityPolicy"},
                    new Pulumi.Alias { Type = "azurerm:storage/v20180701:BlobContainerImmutabilityPolicy"},
                    new Pulumi.Alias { Type = "azurerm:storage/v20181101:BlobContainerImmutabilityPolicy"},
                    new Pulumi.Alias { Type = "azurerm:storage/v20190401:BlobContainerImmutabilityPolicy"},
                    new Pulumi.Alias { Type = "azurerm:storage/v20190601:BlobContainerImmutabilityPolicy"},
                },
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
        /// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API
        /// </summary>
        [Input("allowProtectedAppendWrites")]
        public Input<bool>? AllowProtectedAppendWrites { get; set; }

        /// <summary>
        /// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The immutability period for the blobs in the container since the policy creation, in days.
        /// </summary>
        [Input("immutabilityPeriodSinceCreationInDays")]
        public Input<int>? ImmutabilityPeriodSinceCreationInDays { get; set; }

        /// <summary>
        /// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
        /// </summary>
        [Input("immutabilityPolicyName", required: true)]
        public Input<string> ImmutabilityPolicyName { get; set; } = null!;

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
