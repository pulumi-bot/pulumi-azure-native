// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Latest
{
    /// <summary>
    /// The integration account partner.
    /// 
    /// ## Example Usage
    /// ### Create or update a partner
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var integrationAccountPartner = new AzureRM.Logic.Latest.IntegrationAccountPartner("integrationAccountPartner", new AzureRM.Logic.Latest.IntegrationAccountPartnerArgs
    ///         {
    ///             Content = new AzureRM.Logic.Latest.Inputs.PartnerContentArgs
    ///             {
    ///                 B2b = new AzureRM.Logic.Latest.Inputs.B2BPartnerContentArgs
    ///                 {
    ///                     BusinessIdentities = 
    ///                     {
    ///                         new AzureRM.Logic.Latest.Inputs.BusinessIdentityArgs
    ///                         {
    ///                             Qualifier = "AA",
    ///                             Value = "ZZ",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             IntegrationAccountName = "testIntegrationAccount",
    ///             Location = "westus",
    ///             Metadata = ,
    ///             PartnerName = "testPartner",
    ///             PartnerType = "B2B",
    ///             ResourceGroupName = "testResourceGroup",
    ///             Tags = ,
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class IntegrationAccountPartner : Pulumi.CustomResource
    {
        /// <summary>
        /// The changed time.
        /// </summary>
        [Output("changedTime")]
        public Output<string> ChangedTime { get; private set; } = null!;

        /// <summary>
        /// The partner content.
        /// </summary>
        [Output("content")]
        public Output<Outputs.PartnerContentResponseResult> Content { get; private set; } = null!;

        /// <summary>
        /// The created time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Gets the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The partner type.
        /// </summary>
        [Output("partnerType")]
        public Output<string> PartnerType { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Gets the resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationAccountPartner resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAccountPartner(string name, IntegrationAccountPartnerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/latest:IntegrationAccountPartner", name, args ?? new IntegrationAccountPartnerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAccountPartner(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/latest:IntegrationAccountPartner", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:logic/v20150801preview:IntegrationAccountPartner"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20160601:IntegrationAccountPartner"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20180701preview:IntegrationAccountPartner"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20190501:IntegrationAccountPartner"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationAccountPartner resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAccountPartner Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IntegrationAccountPartner(name, id, options);
        }
    }

    public sealed class IntegrationAccountPartnerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The partner content.
        /// </summary>
        [Input("content", required: true)]
        public Input<Inputs.PartnerContentArgs> Content { get; set; } = null!;

        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public Input<string> IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The integration account partner name.
        /// </summary>
        [Input("partnerName", required: true)]
        public Input<string> PartnerName { get; set; } = null!;

        /// <summary>
        /// The partner type.
        /// </summary>
        [Input("partnerType", required: true)]
        public Input<string> PartnerType { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IntegrationAccountPartnerArgs()
        {
        }
    }
}
