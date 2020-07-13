// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.Outputs
{

    [OutputType]
    public sealed class PoolPropertiesResponseResult
    {
        /// <summary>
        /// UUID v4 used to identify the Pool
        /// </summary>
        public readonly string PoolId;
        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The service level of the file system
        /// </summary>
        public readonly string ServiceLevel;
        /// <summary>
        /// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private PoolPropertiesResponseResult(
            string poolId,

            string provisioningState,

            string serviceLevel,

            int size)
        {
            PoolId = poolId;
            ProvisioningState = provisioningState;
            ServiceLevel = serviceLevel;
            Size = size;
        }
    }
}