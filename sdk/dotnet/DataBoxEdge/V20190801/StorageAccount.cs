// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190801
{
    /// <summary>
    /// Represents a Storage Account on the  Data Box Edge/Gateway device.
    /// 
    /// ## Example Usage
    /// ### StorageAccountPut
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var storageAccount = new AzureRM.DataBoxEdge.V20190801.StorageAccount("storageAccount", new AzureRM.DataBoxEdge.V20190801.StorageAccountArgs
    ///         {
    ///             DataPolicy = "Cloud",
    ///             Description = "It's an awesome storage account",
    ///             DeviceName = "testedgedevice",
    ///             ResourceGroupName = "GroupForEdgeAutomation",
    ///             StorageAccountCredentialId = "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt",
    ///             StorageAccountName = "blobstorageaccount1",
    ///             StorageAccountStatus = "OK",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class StorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// BlobEndpoint of Storage Account
        /// </summary>
        [Output("blobEndpoint")]
        public Output<string> BlobEndpoint { get; private set; } = null!;

        /// <summary>
        /// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
        /// </summary>
        [Output("containerCount")]
        public Output<int> ContainerCount { get; private set; } = null!;

        /// <summary>
        /// Data policy of the storage Account.
        /// </summary>
        [Output("dataPolicy")]
        public Output<string?> DataPolicy { get; private set; } = null!;

        /// <summary>
        /// Description for the storage Account.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Storage Account Credential Id
        /// </summary>
        [Output("storageAccountCredentialId")]
        public Output<string?> StorageAccountCredentialId { get; private set; } = null!;

        /// <summary>
        /// Current status of the storage account
        /// </summary>
        [Output("storageAccountStatus")]
        public Output<string?> StorageAccountStatus { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StorageAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageAccount(string name, StorageAccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190801:StorageAccount", name, args ?? new StorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190801:StorageAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:databoxedge/latest:StorageAccount"},
                    new Pulumi.Alias { Type = "azurerm:databoxedge/v20200501preview:StorageAccount"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageAccount(name, id, options);
        }
    }

    public sealed class StorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data policy of the storage Account.
        /// </summary>
        [Input("dataPolicy")]
        public Input<string>? DataPolicy { get; set; }

        /// <summary>
        /// Description for the storage Account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Storage Account Credential Id
        /// </summary>
        [Input("storageAccountCredentialId")]
        public Input<string>? StorageAccountCredentialId { get; set; }

        /// <summary>
        /// The StorageAccount name.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// Current status of the storage account
        /// </summary>
        [Input("storageAccountStatus")]
        public Input<string>? StorageAccountStatus { get; set; }

        public StorageAccountArgs()
        {
        }
    }
}
