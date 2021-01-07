// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Outputs
{

    [OutputType]
    public sealed class ResourceNavigationLinkResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Link to the external resource.
        /// </summary>
        public readonly string? Link;
        /// <summary>
        /// Resource type of the linked resource.
        /// </summary>
        public readonly string? LinkedResourceType;
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the resource navigation link resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ResourceNavigationLinkResponse(
            string etag,

            string id,

            string? link,

            string? linkedResourceType,

            string? name,

            string provisioningState,

            string type)
        {
            Etag = etag;
            Id = id;
            Link = link;
            LinkedResourceType = linkedResourceType;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
