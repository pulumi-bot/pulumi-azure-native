// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20200801Preview
{
    /// <summary>
    /// Describes an Azure Cognitive Search service and its current state.
    /// </summary>
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
        /// </summary>
        [Output("hostingMode")]
        public Output<string?> HostingMode { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network specific rules that determine how the Azure Cognitive Search service may be reached.
        /// </summary>
        [Output("networkRuleSet")]
        public Output<Outputs.NetworkRuleSetResponseResult?> NetworkRuleSet { get; private set; } = null!;

        /// <summary>
        /// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
        /// </summary>
        [Output("partitionCount")]
        public Output<int?> PartitionCount { get; private set; } = null!;

        /// <summary>
        /// The list of private endpoint connections to the Azure Cognitive Search service.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// The state of the last provisioning operation performed on the search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create search service. This is because the free service uses capacity that is already set up.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
        /// </summary>
        [Output("replicaCount")]
        public Output<int?> ReplicaCount { get; private set; } = null!;

        /// <summary>
        /// The list of shared private link resources managed by the Azure Cognitive Search service.
        /// </summary>
        [Output("sharedPrivateLinkResources")]
        public Output<ImmutableArray<Outputs.SharedPrivateLinkResourceResponseResult>> SharedPrivateLinkResources { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// The status of the search service. Possible values include: 'running': The search service is running and no provisioning operations are underway. 'provisioning': The search service is being provisioned or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This can occur when the underlying search units are not healthy. The search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service is disabled. In this state, the service will reject all API requests. 'error': The search service is in an error state. If your service is in the degraded, disabled, or error states, it means the Azure Cognitive Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The details of the search service status.
        /// </summary>
        [Output("statusDetails")]
        public Output<string> StatusDetails { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:search/v20200801preview:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:search/v20200801preview:Service", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:search/latest:Service"},
                    new Pulumi.Alias { Type = "azurerm:search/v20150819:Service"},
                    new Pulumi.Alias { Type = "azurerm:search/v20191001preview:Service"},
                    new Pulumi.Alias { Type = "azurerm:search/v20200313:Service"},
                    new Pulumi.Alias { Type = "azurerm:search/v20200801:Service"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
        /// </summary>
        [Input("hostingMode")]
        public Input<string>? HostingMode { get; set; }

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Network specific rules that determine how the Azure Cognitive Search service may be reached.
        /// </summary>
        [Input("networkRuleSet")]
        public Input<Inputs.NetworkRuleSetArgs>? NetworkRuleSet { get; set; }

        /// <summary>
        /// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
        /// </summary>
        [Input("publicNetworkAccess")]
        public Input<string>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://&lt;name&gt;.search.windows.net). You cannot change the service name after the service is created.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public Input<string> SearchServiceName { get; set; } = null!;

        /// <summary>
        /// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceArgs()
        {
        }
    }
}
