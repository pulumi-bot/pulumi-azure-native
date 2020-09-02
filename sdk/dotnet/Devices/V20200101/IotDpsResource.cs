// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200101
{
    /// <summary>
    /// The description of the provisioning service.
    /// </summary>
    public partial class IotDpsResource : Pulumi.CustomResource
    {
        /// <summary>
        /// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Service specific properties for a provisioning service
        /// </summary>
        [Output("properties")]
        public Output<Outputs.IotDpsPropertiesDescriptionResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Sku info for a provisioning Service.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.IotDpsSkuInfoResponseResult> Sku { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IotDpsResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotDpsResource(string name, IotDpsResourceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devices/v20200101:IotDpsResource", name, args ?? new IotDpsResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotDpsResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devices/v20200101:IotDpsResource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:devices/latest:IotDpsResource"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20170821preview:IotDpsResource"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20171115:IotDpsResource"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20180122:IotDpsResource"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IotDpsResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotDpsResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IotDpsResource(name, id, options);
        }
    }

    public sealed class IotDpsResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Service specific properties for a provisioning service
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.IotDpsPropertiesDescriptionArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Name of provisioning service to create or update.
        /// </summary>
        [Input("provisioningServiceName", required: true)]
        public Input<string> ProvisioningServiceName { get; set; } = null!;

        /// <summary>
        /// Resource group identifier.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Sku info for a provisioning Service.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.IotDpsSkuInfoArgs> Sku { get; set; } = null!;

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

        public IotDpsResourceArgs()
        {
        }
    }
}
