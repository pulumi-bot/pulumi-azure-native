// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridData.V20160601
{
    /// <summary>
    /// Data store.
    /// 
    /// ## Example Usage
    /// ### DataStores_CreateOrUpdate_DataSinkPUT162
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dataStore = new AzureRM.HybridData.V20160601.DataStore("dataStore", new AzureRM.HybridData.V20160601.DataStoreArgs
    ///         {
    ///             CustomerSecrets = 
    ///             {
    ///                 new AzureRM.HybridData.V20160601.Inputs.CustomerSecretArgs
    ///                 {
    ///                     Algorithm = "RSA1_5",
    ///                     KeyIdentifier = "StorageAccountAccessKey",
    ///                     KeyValue = "Pses92T1jTpPBH2roHgNKArWVv57WomubD/9ukE2d0M89gIVnzc/0bfeoJVf0Med6Uvt4mzmOghFEVxqoBNXzKLgsCLVqvPkqdst4WzZbeel3k/NVkfDdf04eNPAm1FwM/hWZQlG3lGr/olCihW8AFHoiLEZWK8DC6UmHl8ittuuhzY/Ct8R5VXdMWavLdFg8G66TCSyH2aF/eeqzHcOBP4XFgbF2NxuvmEd/cY+y5lEon3TfdDwI0JcOumf5s4zHTWM5+StWa3SsvKxGPpJ27xik2FBo2kEqrtAAByi6HxXSinJB8DwXZicHIwjaOiTeiJUMADwZXwv1PLpwo5d1Q==:QBzHKcxhJkiIeLtFx84oLqDUd5p+oM2AwoKtZIgYVZhfKFw5VW2BfsigL2K7AGxyTNatwn6JZm9ylo8YhRZrn0eIVqLR4gCiRSwDHI7i6R/tqTfx8ZO/aJy6rTh/WW8d6vZOXXGeuRDAz6fYfjQKb9J/OhTq3cjfVouLt6bKdZsZve08NVZq0sNBYZftCabcOhVg5hamuDhQhemwqFMn6l1xrCWcq4e5YgJ90fbK5N66Wj5LNr2dU+scHH7YfM8a3IIhq51TObhXZ59oNnLhLGGA8j0K43MMKtQAnqpBc+hmwgwc8/DZLod1wnaPbJW5/fQ2HkF7vH9xakIip4bZ9Q==",
    ///                 },
    ///                 new AzureRM.HybridData.V20160601.Inputs.CustomerSecretArgs
    ///                 {
    ///                     Algorithm = "RSA1_5",
    ///                     KeyIdentifier = "StorageAccountAccessKeyForQueue",
    ///                     KeyValue = "Pses92T1jTpPBH2roHgNKArWVv57WomubD/9ukE2d0M89gIVnzc/0bfeoJVf0Med6Uvt4mzmOghFEVxqoBNXzKLgsCLVqvPkqdst4WzZbeel3k/NVkfDdf04eNPAm1FwM/hWZQlG3lGr/olCihW8AFHoiLEZWK8DC6UmHl8ittuuhzY/Ct8R5VXdMWavLdFg8G66TCSyH2aF/eeqzHcOBP4XFgbF2NxuvmEd/cY+y5lEon3TfdDwI0JcOumf5s4zHTWM5+StWa3SsvKxGPpJ27xik2FBo2kEqrtAAByi6HxXSinJB8DwXZicHIwjaOiTeiJUMADwZXwv1PLpwo5d1Q==:QBzHKcxhJkiIeLtFx84oLqDUd5p+oM2AwoKtZIgYVZhfKFw5VW2BfsigL2K7AGxyTNatwn6JZm9ylo8YhRZrn0eIVqLR4gCiRSwDHI7i6R/tqTfx8ZO/aJy6rTh/WW8d6vZOXXGeuRDAz6fYfjQKb9J/OhTq3cjfVouLt6bKdZsZve08NVZq0sNBYZftCabcOhVg5hamuDhQhemwqFMn6l1xrCWcq4e5YgJ90fbK5N66Wj5LNr2dU+scHH7YfM8a3IIhq51TObhXZ59oNnLhLGGA8j0K43MMKtQAnqpBc+hmwgwc8/DZLod1wnaPbJW5/fQ2HkF7vH9xakIip4bZ9Q==",
    ///                 },
    ///             },
    ///             DataManagerName = "TestAzureSDKOperations",
    ///             DataStoreName = "TestAzureStorage1",
    ///             DataStoreTypeId = "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount",
    ///             ExtendedProperties = 
    ///             {
    ///                 { "extendedSaKey", null },
    ///                 { "extendedSaName", "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink" },
    ///                 { "storageAccountNameForQueue", "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink" },
    ///             },
    ///             RepositoryId = "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
    ///             ResourceGroupName = "ResourceGroupForSDKTest",
    ///             State = "Enabled",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### DataStores_CreateOrUpdate_DataSourcePUT162
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dataStore = new AzureRM.HybridData.V20160601.DataStore("dataStore", new AzureRM.HybridData.V20160601.DataStoreArgs
    ///         {
    ///             CustomerSecrets = 
    ///             {
    ///                 new AzureRM.HybridData.V20160601.Inputs.CustomerSecretArgs
    ///                 {
    ///                     Algorithm = "RSA1_5",
    ///                     KeyIdentifier = "ServiceEncryptionKey",
    ///                     KeyValue = "vZqtembBBg2RC/UyYwZiexGOqujLcMYCmaywqf0sURqIidjxlSp86FGz+T2eRnb1XlYCqFf1CzPzwLpwHEuTJ8LP5hTV1yUiM+YnyKHIGdlQajLcVcFy8ji9n+jSS4J9PjjHsr5AKzW1w+y76UgTEpX7K9kFDWFVyDGEujvuB2bYBlxlKolMCOu0WHZYkBBYLob6a3mQgCHbXYj1mqTmdhPW+J+8tyBCzG6cjlvRJ9hcp9Ss3HV9TV6hrbqlUU3lE1FX8O5Dr6/TXi7tIU7hGfmS5psE0Kz+2PsJTX1R1AbkBpKObPwPxPoC5jCXFxwfmZOrNQdjZ7nu5+JHaLZylw==:tS9oSCAvIwOrkYRyD/jLahSLZypl4eNexW5N/pGqf9vsVfzMhmxob+O/Io48uCPxvtdDksef09OVXpxgaC65K2Og49W9rtRt8cvGyyC41cx5D2DP9fxAu7d/lREP9cWHgrRJlZ4JJFcqy+m+yqYKl3WPrTA2yoZpISGbWAPkj0Hk3IwRr1lmqKfCWtp0jNHnrIJmQ5BQaDLGXpohKQSrrftqz7TdBCYuorSntQz8pqHgc8MTiYMgMtgZ+HRKQ1F5ctOlP+6LJMS6/OFl/tnYb5BD6rn/RufB4OHhVDe9ZD5GMtkwqkUvU9b1v2n31mb63JLApxIi/o8OsSpkA8ZTCg==",
    ///                 },
    ///             },
    ///             DataManagerName = "TestAzureSDKOperations",
    ///             DataStoreName = "TestStorSimpleSource1",
    ///             DataStoreTypeId = "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series",
    ///             ExtendedProperties = 
    ///             {
    ///                 { "extendedSaKey", null },
    ///                 { "resourceId", "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/dmsdatasource/providers/Microsoft.StorSimple/managers/dmsdatasource" },
    ///             },
    ///             RepositoryId = "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/dmsdatasource/providers/Microsoft.StorSimple/managers/dmsdatasource",
    ///             ResourceGroupName = "ResourceGroupForSDKTest",
    ///             State = "Enabled",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class DataStore : Pulumi.CustomResource
    {
        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        [Output("customerSecrets")]
        public Output<ImmutableArray<Outputs.CustomerSecretResponseResult>> CustomerSecrets { get; private set; } = null!;

        /// <summary>
        /// The arm id of the data store type.
        /// </summary>
        [Output("dataStoreTypeId")]
        public Output<string> DataStoreTypeId { get; private set; } = null!;

        /// <summary>
        /// A generic json used differently by each data source type.
        /// </summary>
        [Output("extendedProperties")]
        public Output<ImmutableDictionary<string, object>?> ExtendedProperties { get; private set; } = null!;

        /// <summary>
        /// Name of the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Arm Id for the manager resource to which the data source is associated. This is optional.
        /// </summary>
        [Output("repositoryId")]
        public Output<string?> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// State of the data source.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataStore(string name, DataStoreArgs args, CustomResourceOptions? options = null)
            : base("azurerm:hybriddata/v20160601:DataStore", name, args ?? new DataStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataStore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:hybriddata/v20160601:DataStore", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:hybriddata/latest:DataStore"},
                    new Pulumi.Alias { Type = "azurerm:hybriddata/v20190601:DataStore"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataStore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataStore(name, id, options);
        }
    }

    public sealed class DataStoreArgs : Pulumi.ResourceArgs
    {
        [Input("customerSecrets")]
        private InputList<Inputs.CustomerSecretArgs>? _customerSecrets;

        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        public InputList<Inputs.CustomerSecretArgs> CustomerSecrets
        {
            get => _customerSecrets ?? (_customerSecrets = new InputList<Inputs.CustomerSecretArgs>());
            set => _customerSecrets = value;
        }

        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public Input<string> DataManagerName { get; set; } = null!;

        /// <summary>
        /// The data store/repository name to be created or updated.
        /// </summary>
        [Input("dataStoreName", required: true)]
        public Input<string> DataStoreName { get; set; } = null!;

        /// <summary>
        /// The arm id of the data store type.
        /// </summary>
        [Input("dataStoreTypeId", required: true)]
        public Input<string> DataStoreTypeId { get; set; } = null!;

        [Input("extendedProperties")]
        private InputMap<object>? _extendedProperties;

        /// <summary>
        /// A generic json used differently by each data source type.
        /// </summary>
        public InputMap<object> ExtendedProperties
        {
            get => _extendedProperties ?? (_extendedProperties = new InputMap<object>());
            set => _extendedProperties = value;
        }

        /// <summary>
        /// Arm Id for the manager resource to which the data source is associated. This is optional.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// State of the data source.
        /// </summary>
        [Input("state", required: true)]
        public Input<string> State { get; set; } = null!;

        public DataStoreArgs()
        {
        }
    }
}
