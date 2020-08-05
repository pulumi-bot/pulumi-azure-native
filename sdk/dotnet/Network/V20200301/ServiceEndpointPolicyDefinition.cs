// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301
{
    /// <summary>
    /// Service Endpoint policy definitions.
    /// </summary>
    public partial class ServiceEndpointPolicyDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the service endpoint policy definition.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ServiceEndpointPolicyDefinitionPropertiesFormatResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointPolicyDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointPolicyDefinition(string name, ServiceEndpointPolicyDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:ServiceEndpointPolicyDefinition", name, args ?? new ServiceEndpointPolicyDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointPolicyDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:ServiceEndpointPolicyDefinition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointPolicyDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointPolicyDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointPolicyDefinition(name, id, options);
        }
    }

    public sealed class ServiceEndpointPolicyDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the service endpoint policy definition name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Service endpoint name.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// The name of the service endpoint policy.
        /// </summary>
        [Input("serviceEndpointPolicyName", required: true)]
        public Input<string> ServiceEndpointPolicyName { get; set; } = null!;

        [Input("serviceResources")]
        private InputList<string>? _serviceResources;

        /// <summary>
        /// A list of service resources.
        /// </summary>
        public InputList<string> ServiceResources
        {
            get => _serviceResources ?? (_serviceResources = new InputList<string>());
            set => _serviceResources = value;
        }

        public ServiceEndpointPolicyDefinitionArgs()
        {
        }
    }
}
