// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class ServiceAssociationLinkPropertiesFormatResponseResult
    {
        /// <summary>
        /// If true, the resource can be deleted.
        /// </summary>
        public readonly bool? AllowDelete;
        /// <summary>
        /// Link to the external resource.
        /// </summary>
        public readonly string? Link;
        /// <summary>
        /// Resource type of the linked resource.
        /// </summary>
        public readonly string? LinkedResourceType;
        /// <summary>
        /// A list of locations.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// The provisioning state of the service association link resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private ServiceAssociationLinkPropertiesFormatResponseResult(
            bool? allowDelete,

            string? link,

            string? linkedResourceType,

            ImmutableArray<string> locations,

            string provisioningState)
        {
            AllowDelete = allowDelete;
            Link = link;
            LinkedResourceType = linkedResourceType;
            Locations = locations;
            ProvisioningState = provisioningState;
        }
    }
}