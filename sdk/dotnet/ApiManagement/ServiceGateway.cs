// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement
{
    /// <summary>
    /// Gateway details.
    /// </summary>
    public partial class ServiceGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gateway details.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.GatewayContractPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceGateway(string name, ServiceGatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement:ServiceGateway", name, args ?? new ServiceGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement:ServiceGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceGateway(name, id, options);
        }
    }

    public sealed class ServiceGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Gateway details.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.GatewayContractPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ServiceGatewayArgs()
        {
        }
    }
}
