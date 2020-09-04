// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.Latest
{
    /// <summary>
    /// Type of the Storage Target.
    /// 
    /// ## Example Usage
    /// ### StorageTargets_CreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var storageTarget = new AzureRM.StorageCache.Latest.StorageTarget("storageTarget", new AzureRM.StorageCache.Latest.StorageTargetArgs
    ///         {
    ///             CacheName = "sc1",
    ///             Junctions = 
    ///             {
    ///                 new AzureRM.StorageCache.Latest.Inputs.NamespaceJunctionArgs
    ///                 {
    ///                     NamespacePath = "/path/on/cache",
    ///                     NfsExport = "exp1",
    ///                     TargetPath = "/path/on/exp1",
    ///                 },
    ///                 new AzureRM.StorageCache.Latest.Inputs.NamespaceJunctionArgs
    ///                 {
    ///                     NamespacePath = "/path2/on/cache",
    ///                     NfsExport = "exp2",
    ///                     TargetPath = "/path2/on/exp2",
    ///                 },
    ///             },
    ///             Nfs3 = new AzureRM.StorageCache.Latest.Inputs.Nfs3TargetArgs
    ///             {
    ///                 Target = "10.0.44.44",
    ///                 UsageModel = "READ_HEAVY_INFREQ",
    ///             },
    ///             ResourceGroupName = "scgroup",
    ///             StorageTargetName = "st1",
    ///             TargetType = "nfs3",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class StorageTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// Properties when targetType is clfs.
        /// </summary>
        [Output("clfs")]
        public Output<Outputs.ClfsTargetResponseResult?> Clfs { get; private set; } = null!;

        /// <summary>
        /// List of Cache namespace junctions to target for namespace associations.
        /// </summary>
        [Output("junctions")]
        public Output<ImmutableArray<Outputs.NamespaceJunctionResponseResult>> Junctions { get; private set; } = null!;

        /// <summary>
        /// Name of the Storage Target.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties when targetType is nfs3.
        /// </summary>
        [Output("nfs3")]
        public Output<Outputs.Nfs3TargetResponseResult?> Nfs3 { get; private set; } = null!;

        /// <summary>
        /// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Type of the Storage Target.
        /// </summary>
        [Output("targetType")]
        public Output<string> TargetType { get; private set; } = null!;

        /// <summary>
        /// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Properties when targetType is unknown.
        /// </summary>
        [Output("unknown")]
        public Output<Outputs.UnknownTargetResponseResult?> Unknown { get; private set; } = null!;


        /// <summary>
        /// Create a StorageTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageTarget(string name, StorageTargetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storagecache/latest:StorageTarget", name, args ?? new StorageTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageTarget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storagecache/latest:StorageTarget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storagecache/v20190801preview:StorageTarget"},
                    new Pulumi.Alias { Type = "azurerm:storagecache/v20191101:StorageTarget"},
                    new Pulumi.Alias { Type = "azurerm:storagecache/v20200301:StorageTarget"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageTarget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageTarget(name, id, options);
        }
    }

    public sealed class StorageTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
        /// </summary>
        [Input("cacheName", required: true)]
        public Input<string> CacheName { get; set; } = null!;

        /// <summary>
        /// Properties when targetType is clfs.
        /// </summary>
        [Input("clfs")]
        public Input<Inputs.ClfsTargetArgs>? Clfs { get; set; }

        [Input("junctions")]
        private InputList<Inputs.NamespaceJunctionArgs>? _junctions;

        /// <summary>
        /// List of Cache namespace junctions to target for namespace associations.
        /// </summary>
        public InputList<Inputs.NamespaceJunctionArgs> Junctions
        {
            get => _junctions ?? (_junctions = new InputList<Inputs.NamespaceJunctionArgs>());
            set => _junctions = value;
        }

        /// <summary>
        /// Properties when targetType is nfs3.
        /// </summary>
        [Input("nfs3")]
        public Input<Inputs.Nfs3TargetArgs>? Nfs3 { get; set; }

        /// <summary>
        /// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Target resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Storage Target. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
        /// </summary>
        [Input("storageTargetName", required: true)]
        public Input<string> StorageTargetName { get; set; } = null!;

        /// <summary>
        /// Type of the Storage Target.
        /// </summary>
        [Input("targetType", required: true)]
        public Input<string> TargetType { get; set; } = null!;

        /// <summary>
        /// Properties when targetType is unknown.
        /// </summary>
        [Input("unknown")]
        public Input<Inputs.UnknownTargetArgs>? Unknown { get; set; }

        public StorageTargetArgs()
        {
        }
    }
}
