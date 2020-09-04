// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Inputs
{

    /// <summary>
    /// Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
    /// </summary>
    public sealed class NodeTypeDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of ports from which cluster assigned port to Service Fabric applications.
        /// </summary>
        [Input("applicationPorts")]
        public Input<Inputs.EndpointRangeDescriptionArgs>? ApplicationPorts { get; set; }

        [Input("capacities")]
        private InputMap<string>? _capacities;

        /// <summary>
        /// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        /// </summary>
        public InputMap<string> Capacities
        {
            get => _capacities ?? (_capacities = new InputMap<string>());
            set => _capacities = value;
        }

        /// <summary>
        /// The TCP cluster management endpoint port.
        /// </summary>
        [Input("clientConnectionEndpointPort", required: true)]
        public Input<int> ClientConnectionEndpointPort { get; set; } = null!;

        /// <summary>
        /// The durability level of the node type. Learn about [DurabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
        /// 
        ///   - Bronze - No privileges. This is the default.
        ///   - Silver - The infrastructure jobs can be paused for a duration of 10 minutes per UD.
        ///   - Gold - The infrastructure jobs can be paused for a duration of 2 hours per UD. Gold durability can be enabled only on full node VM skus like D15_V2, G5 etc.
        /// </summary>
        [Input("durabilityLevel")]
        public Input<string>? DurabilityLevel { get; set; }

        /// <summary>
        /// The range of ephemeral ports that nodes in this node type should be configured with.
        /// </summary>
        [Input("ephemeralPorts")]
        public Input<Inputs.EndpointRangeDescriptionArgs>? EphemeralPorts { get; set; }

        /// <summary>
        /// The HTTP cluster management endpoint port.
        /// </summary>
        [Input("httpGatewayEndpointPort", required: true)]
        public Input<int> HttpGatewayEndpointPort { get; set; } = null!;

        /// <summary>
        /// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
        /// </summary>
        [Input("isPrimary", required: true)]
        public Input<bool> IsPrimary { get; set; } = null!;

        /// <summary>
        /// The name of the node type.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("placementProperties")]
        private InputMap<string>? _placementProperties;

        /// <summary>
        /// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        /// </summary>
        public InputMap<string> PlacementProperties
        {
            get => _placementProperties ?? (_placementProperties = new InputMap<string>());
            set => _placementProperties = value;
        }

        /// <summary>
        /// The endpoint used by reverse proxy.
        /// </summary>
        [Input("reverseProxyEndpointPort")]
        public Input<int>? ReverseProxyEndpointPort { get; set; }

        /// <summary>
        /// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
        /// </summary>
        [Input("vmInstanceCount", required: true)]
        public Input<int> VmInstanceCount { get; set; } = null!;

        public NodeTypeDescriptionArgs()
        {
        }
    }
}
