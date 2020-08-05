// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101
{
    /// <summary>
    /// Subscription details.
    /// </summary>
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Subscription contract properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SubscriptionContractPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:Subscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether tracing can be enabled
        /// </summary>
        [Input("allowTracing")]
        public Input<bool>? AllowTracing { get; set; }

        /// <summary>
        /// Determines the type of application which send the create user request. Default is legacy publisher portal.
        /// </summary>
        [Input("appType")]
        public Input<string>? AppType { get; set; }

        /// <summary>
        /// Subscription name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Notify change in Subscription State. 
        ///  - If false, do not send any email notification for change of state of subscription 
        ///  - If true, send email notification of change of state of subscription 
        /// </summary>
        [Input("notify")]
        public Input<bool>? Notify { get; set; }

        /// <summary>
        /// User (user id path) for whom subscription is being created in form /users/{userId}
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// Primary subscription key. If not specified during request key will be generated automatically.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Scope like /products/{productId} or /apis or /apis/{apiId}.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Secondary subscription key. If not specified during request key will be generated automatically.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public SubscriptionArgs()
        {
        }
    }
}
