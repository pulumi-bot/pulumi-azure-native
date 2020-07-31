// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20171001.Outputs
{

    [OutputType]
    public sealed class RedisLinkedServerPropertiesResponseResult
    {
        /// <summary>
        /// Fully qualified resourceId of the linked redis cache.
        /// </summary>
        public readonly string LinkedRedisCacheId;
        /// <summary>
        /// Location of the linked redis cache.
        /// </summary>
        public readonly string LinkedRedisCacheLocation;
        /// <summary>
        /// Terminal state of the link between primary and secondary redis cache.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Role of the linked server.
        /// </summary>
        public readonly string ServerRole;

        [OutputConstructor]
        private RedisLinkedServerPropertiesResponseResult(
            string linkedRedisCacheId,

            string linkedRedisCacheLocation,

            string provisioningState,

            string serverRole)
        {
            LinkedRedisCacheId = linkedRedisCacheId;
            LinkedRedisCacheLocation = linkedRedisCacheLocation;
            ProvisioningState = provisioningState;
            ServerRole = serverRole;
        }
    }
}