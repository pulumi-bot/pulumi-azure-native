// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Outputs
{

    [OutputType]
    public sealed class RouteFilterRuleResponse
    {
        /// <summary>
        /// The access type of the rule.
        /// </summary>
        public readonly string Access;
        /// <summary>
        /// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
        /// </summary>
        public readonly ImmutableArray<string> Communities;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the route filter rule resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The rule type of the rule.
        /// </summary>
        public readonly string RouteFilterRuleType;

        [OutputConstructor]
        private RouteFilterRuleResponse(
            string access,

            ImmutableArray<string> communities,

            string etag,

            string? id,

            string? location,

            string? name,

            string provisioningState,

            string routeFilterRuleType)
        {
            Access = access;
            Communities = communities;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            RouteFilterRuleType = routeFilterRuleType;
        }
    }
}
