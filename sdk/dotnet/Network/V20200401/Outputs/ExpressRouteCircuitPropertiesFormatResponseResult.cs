// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Outputs
{

    [OutputType]
    public sealed class ExpressRouteCircuitPropertiesFormatResponseResult
    {
        /// <summary>
        /// Allow classic operations.
        /// </summary>
        public readonly bool? AllowClassicOperations;
        /// <summary>
        /// The list of authorizations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressRouteCircuitAuthorizationResponseResult> Authorizations;
        /// <summary>
        /// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
        /// </summary>
        public readonly double? BandwidthInGbps;
        /// <summary>
        /// The CircuitProvisioningState state of the resource.
        /// </summary>
        public readonly string? CircuitProvisioningState;
        /// <summary>
        /// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? ExpressRoutePort;
        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        public readonly string? GatewayManagerEtag;
        /// <summary>
        /// Flag denoting global reach status.
        /// </summary>
        public readonly bool? GlobalReachEnabled;
        /// <summary>
        /// The list of peerings.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponseResult> Peerings;
        /// <summary>
        /// The provisioning state of the express route circuit resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The ServiceKey.
        /// </summary>
        public readonly string? ServiceKey;
        /// <summary>
        /// The ServiceProviderNotes.
        /// </summary>
        public readonly string? ServiceProviderNotes;
        /// <summary>
        /// The ServiceProviderProperties.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitServiceProviderPropertiesResponseResult? ServiceProviderProperties;
        /// <summary>
        /// The ServiceProviderProvisioningState state of the resource.
        /// </summary>
        public readonly string? ServiceProviderProvisioningState;
        /// <summary>
        /// The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
        /// </summary>
        public readonly int Stag;

        [OutputConstructor]
        private ExpressRouteCircuitPropertiesFormatResponseResult(
            bool? allowClassicOperations,

            ImmutableArray<Outputs.ExpressRouteCircuitAuthorizationResponseResult> authorizations,

            double? bandwidthInGbps,

            string? circuitProvisioningState,

            Outputs.SubResourceResponseResult? expressRoutePort,

            string? gatewayManagerEtag,

            bool? globalReachEnabled,

            ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponseResult> peerings,

            string provisioningState,

            string? serviceKey,

            string? serviceProviderNotes,

            Outputs.ExpressRouteCircuitServiceProviderPropertiesResponseResult? serviceProviderProperties,

            string? serviceProviderProvisioningState,

            int stag)
        {
            AllowClassicOperations = allowClassicOperations;
            Authorizations = authorizations;
            BandwidthInGbps = bandwidthInGbps;
            CircuitProvisioningState = circuitProvisioningState;
            ExpressRoutePort = expressRoutePort;
            GatewayManagerEtag = gatewayManagerEtag;
            GlobalReachEnabled = globalReachEnabled;
            Peerings = peerings;
            ProvisioningState = provisioningState;
            ServiceKey = serviceKey;
            ServiceProviderNotes = serviceProviderNotes;
            ServiceProviderProperties = serviceProviderProperties;
            ServiceProviderProvisioningState = serviceProviderProvisioningState;
            Stag = stag;
        }
    }
}