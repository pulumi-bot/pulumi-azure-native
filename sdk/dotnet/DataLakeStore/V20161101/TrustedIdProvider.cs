// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeStore.V20161101
{
    /// <summary>
    /// Data Lake Store trusted identity provider information.
    /// 
    /// ## Example Usage
    /// ### Creates or updates the specified trusted identity provider. During update, the trusted identity provider with the specified name will be replaced with this new provider
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var trustedIdProvider = new AzureRM.DataLakeStore.V20161101.TrustedIdProvider("trustedIdProvider", new AzureRM.DataLakeStore.V20161101.TrustedIdProviderArgs
    ///         {
    ///             AccountName = "contosoadla",
    ///             IdProvider = "https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1",
    ///             ResourceGroupName = "contosorg",
    ///             TrustedIdProviderName = "test_trusted_id_provider_name",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class TrustedIdProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// The URL of this trusted identity provider.
        /// </summary>
        [Output("idProvider")]
        public Output<string> IdProvider { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TrustedIdProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrustedIdProvider(string name, TrustedIdProviderArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datalakestore/v20161101:TrustedIdProvider", name, args ?? new TrustedIdProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrustedIdProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datalakestore/v20161101:TrustedIdProvider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datalakestore/latest:TrustedIdProvider"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrustedIdProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrustedIdProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TrustedIdProvider(name, id, options);
        }
    }

    public sealed class TrustedIdProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Store account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The URL of this trusted identity provider.
        /// </summary>
        [Input("idProvider", required: true)]
        public Input<string> IdProvider { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the trusted identity provider. This is used for differentiation of providers in the account.
        /// </summary>
        [Input("trustedIdProviderName", required: true)]
        public Input<string> TrustedIdProviderName { get; set; } = null!;

        public TrustedIdProviderArgs()
        {
        }
    }
}
