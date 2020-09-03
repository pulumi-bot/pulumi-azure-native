// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AzureStack.Latest
{
    /// <summary>
    /// Customer subscription.
    /// 
    /// ## Example Usage
    /// ### Creates a new customer subscription under a registration.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var customerSubscription = new AzureRM.AzureStack.Latest.CustomerSubscription("customerSubscription", new AzureRM.AzureStack.Latest.CustomerSubscriptionArgs
    ///         {
    ///             CustomerSubscriptionName = "E09A4E93-29A7-4EBA-A6D4-76202383F07F",
    ///             RegistrationName = "testregistration",
    ///             ResourceGroup = "azurestack",
    ///             TenantId = "dbab3982-796f-4d03-9908-044c08aef8a2",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class CustomerSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The entity tag used for optimistic concurrency when modifying the resource.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tenant Id.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Type of Resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CustomerSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomerSubscription(string name, CustomerSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:azurestack/latest:CustomerSubscription", name, args ?? new CustomerSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomerSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:azurestack/latest:CustomerSubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:azurestack/v20170601:CustomerSubscription"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomerSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomerSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomerSubscription(name, id, options);
        }
    }

    public sealed class CustomerSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("customerSubscriptionName", required: true)]
        public Input<string> CustomerSubscriptionName { get; set; } = null!;

        /// <summary>
        /// The entity tag used for optimistic concurrency when modifying the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public Input<string> RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        /// <summary>
        /// Tenant Id.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public CustomerSubscriptionArgs()
        {
        }
    }
}
