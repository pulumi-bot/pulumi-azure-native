// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NotificationHubs.V20170401
{
    /// <summary>
    /// Description of a NotificationHub Resource.
    /// </summary>
    public partial class NotificationHub : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the NotificationHub.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.NotificationHubPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The sku of the created namespace
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationHub(string name, NotificationHubArgs args, CustomResourceOptions? options = null)
            : base("azurerm:notificationhubs/v20170401:NotificationHub", name, args ?? new NotificationHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationHub(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:notificationhubs/v20170401:NotificationHub", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NotificationHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationHub Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NotificationHub(name, id, options);
        }
    }

    public sealed class NotificationHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AdmCredential of the created NotificationHub
        /// </summary>
        [Input("admCredential")]
        public Input<Inputs.AdmCredentialArgs>? AdmCredential { get; set; }

        /// <summary>
        /// The ApnsCredential of the created NotificationHub
        /// </summary>
        [Input("apnsCredential")]
        public Input<Inputs.ApnsCredentialArgs>? ApnsCredential { get; set; }

        [Input("authorizationRules")]
        private InputList<Inputs.SharedAccessAuthorizationRulePropertiesArgs>? _authorizationRules;

        /// <summary>
        /// The AuthorizationRules of the created NotificationHub
        /// </summary>
        public InputList<Inputs.SharedAccessAuthorizationRulePropertiesArgs> AuthorizationRules
        {
            get => _authorizationRules ?? (_authorizationRules = new InputList<Inputs.SharedAccessAuthorizationRulePropertiesArgs>());
            set => _authorizationRules = value;
        }

        /// <summary>
        /// The BaiduCredential of the created NotificationHub
        /// </summary>
        [Input("baiduCredential")]
        public Input<Inputs.BaiduCredentialArgs>? BaiduCredential { get; set; }

        /// <summary>
        /// The GcmCredential of the created NotificationHub
        /// </summary>
        [Input("gcmCredential")]
        public Input<Inputs.GcmCredentialArgs>? GcmCredential { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The MpnsCredential of the created NotificationHub
        /// </summary>
        [Input("mpnsCredential")]
        public Input<Inputs.MpnsCredentialArgs>? MpnsCredential { get; set; }

        /// <summary>
        /// The notification hub name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace name.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The RegistrationTtl of the created NotificationHub
        /// </summary>
        [Input("registrationTtl")]
        public Input<string>? RegistrationTtl { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku of the created namespace
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The WnsCredential of the created NotificationHub
        /// </summary>
        [Input("wnsCredential")]
        public Input<Inputs.WnsCredentialArgs>? WnsCredential { get; set; }

        public NotificationHubArgs()
        {
        }
    }
}
