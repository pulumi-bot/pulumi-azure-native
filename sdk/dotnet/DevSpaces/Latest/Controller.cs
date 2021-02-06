// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevSpaces.Latest
{
    /// <summary>
    /// Latest API Version: 2019-04-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:devspaces/latest:Controller")]
    public partial class Controller : Pulumi.CustomResource
    {
        /// <summary>
        /// DNS name for accessing DataPlane services
        /// </summary>
        [Output("dataPlaneFqdn")]
        public Output<string> DataPlaneFqdn { get; private set; } = null!;

        /// <summary>
        /// DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
        /// </summary>
        [Output("hostSuffix")]
        public Output<string> HostSuffix { get; private set; } = null!;

        /// <summary>
        /// Region where the Azure resource is located.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the Azure Dev Spaces Controller.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Model representing SKU for Azure Dev Spaces Controller.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// Tags for the Azure resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// DNS of the target container host's API server
        /// </summary>
        [Output("targetContainerHostApiServerFqdn")]
        public Output<string> TargetContainerHostApiServerFqdn { get; private set; } = null!;

        /// <summary>
        /// Credentials of the target container host (base64).
        /// </summary>
        [Output("targetContainerHostCredentialsBase64")]
        public Output<string> TargetContainerHostCredentialsBase64 { get; private set; } = null!;

        /// <summary>
        /// Resource ID of the target container host
        /// </summary>
        [Output("targetContainerHostResourceId")]
        public Output<string> TargetContainerHostResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Controller resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Controller(string name, ControllerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:devspaces/latest:Controller", name, args ?? new ControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Controller(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:devspaces/latest:Controller", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:devspaces/v20190401:Controller"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Controller resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Controller Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Controller(name, id, options);
        }
    }

    public sealed class ControllerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region where the Azure resource is located.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Model representing SKU for Azure Dev Spaces Controller.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags for the Azure resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Credentials of the target container host (base64).
        /// </summary>
        [Input("targetContainerHostCredentialsBase64", required: true)]
        public Input<string> TargetContainerHostCredentialsBase64 { get; set; } = null!;

        /// <summary>
        /// Resource ID of the target container host
        /// </summary>
        [Input("targetContainerHostResourceId", required: true)]
        public Input<string> TargetContainerHostResourceId { get; set; } = null!;

        public ControllerArgs()
        {
        }
    }
}
