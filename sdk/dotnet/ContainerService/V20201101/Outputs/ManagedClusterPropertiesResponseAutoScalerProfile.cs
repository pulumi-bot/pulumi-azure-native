// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201101.Outputs
{

    [OutputType]
    public sealed class ManagedClusterPropertiesResponseAutoScalerProfile
    {
        public readonly string? BalanceSimilarNodeGroups;
        public readonly string? Expander;
        public readonly string? MaxEmptyBulkDelete;
        public readonly string? MaxGracefulTerminationSec;
        public readonly string? MaxTotalUnreadyPercentage;
        public readonly string? NewPodScaleUpDelay;
        public readonly string? OkTotalUnreadyCount;
        public readonly string? ScaleDownDelayAfterAdd;
        public readonly string? ScaleDownDelayAfterDelete;
        public readonly string? ScaleDownDelayAfterFailure;
        public readonly string? ScaleDownUnneededTime;
        public readonly string? ScaleDownUnreadyTime;
        public readonly string? ScaleDownUtilizationThreshold;
        public readonly string? ScanInterval;
        public readonly string? SkipNodesWithLocalStorage;
        public readonly string? SkipNodesWithSystemPods;

        [OutputConstructor]
        private ManagedClusterPropertiesResponseAutoScalerProfile(
            string? balanceSimilarNodeGroups,

            string? expander,

            string? maxEmptyBulkDelete,

            string? maxGracefulTerminationSec,

            string? maxTotalUnreadyPercentage,

            string? newPodScaleUpDelay,

            string? okTotalUnreadyCount,

            string? scaleDownDelayAfterAdd,

            string? scaleDownDelayAfterDelete,

            string? scaleDownDelayAfterFailure,

            string? scaleDownUnneededTime,

            string? scaleDownUnreadyTime,

            string? scaleDownUtilizationThreshold,

            string? scanInterval,

            string? skipNodesWithLocalStorage,

            string? skipNodesWithSystemPods)
        {
            BalanceSimilarNodeGroups = balanceSimilarNodeGroups;
            Expander = expander;
            MaxEmptyBulkDelete = maxEmptyBulkDelete;
            MaxGracefulTerminationSec = maxGracefulTerminationSec;
            MaxTotalUnreadyPercentage = maxTotalUnreadyPercentage;
            NewPodScaleUpDelay = newPodScaleUpDelay;
            OkTotalUnreadyCount = okTotalUnreadyCount;
            ScaleDownDelayAfterAdd = scaleDownDelayAfterAdd;
            ScaleDownDelayAfterDelete = scaleDownDelayAfterDelete;
            ScaleDownDelayAfterFailure = scaleDownDelayAfterFailure;
            ScaleDownUnneededTime = scaleDownUnneededTime;
            ScaleDownUnreadyTime = scaleDownUnreadyTime;
            ScaleDownUtilizationThreshold = scaleDownUtilizationThreshold;
            ScanInterval = scanInterval;
            SkipNodesWithLocalStorage = skipNodesWithLocalStorage;
            SkipNodesWithSystemPods = skipNodesWithSystemPods;
        }
    }
}
