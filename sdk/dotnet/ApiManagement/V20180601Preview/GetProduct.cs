// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview
{
    public static class GetProduct
    {
        public static Task<GetProductResult> InvokeAsync(GetProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductResult>("azurerm:apimanagement/v20180601preview:getProduct", args ?? new GetProductArgs(), options.WithVersion());
    }


    public sealed class GetProductArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetProductArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProductResult
    {
        /// <summary>
        /// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of true.
        /// </summary>
        public readonly bool? ApprovalRequired;
        /// <summary>
        /// Product description. May include HTML formatting tags.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Product name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
        /// </summary>
        public readonly bool? SubscriptionRequired;
        /// <summary>
        /// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of true.
        /// </summary>
        public readonly int? SubscriptionsLimit;
        /// <summary>
        /// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
        /// </summary>
        public readonly string? Terms;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProductResult(
            bool? approvalRequired,

            string? description,

            string displayName,

            string name,

            string? state,

            bool? subscriptionRequired,

            int? subscriptionsLimit,

            string? terms,

            string type)
        {
            ApprovalRequired = approvalRequired;
            Description = description;
            DisplayName = displayName;
            Name = name;
            State = state;
            SubscriptionRequired = subscriptionRequired;
            SubscriptionsLimit = subscriptionsLimit;
            Terms = terms;
            Type = type;
        }
    }
}
