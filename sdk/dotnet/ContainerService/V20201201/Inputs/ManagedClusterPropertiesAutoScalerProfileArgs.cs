// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201201.Inputs
{

    /// <summary>
    /// Parameters to be applied to the cluster-autoscaler when enabled
    /// </summary>
    public sealed class ManagedClusterPropertiesAutoScalerProfileArgs : Pulumi.ResourceArgs
    {
        [Input("balanceSimilarNodeGroups")]
        public Input<string>? BalanceSimilarNodeGroups { get; set; }

        [Input("expander")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerService.V20201201.Expander>? Expander { get; set; }

        [Input("maxEmptyBulkDelete")]
        public Input<string>? MaxEmptyBulkDelete { get; set; }

        [Input("maxGracefulTerminationSec")]
        public Input<string>? MaxGracefulTerminationSec { get; set; }

        [Input("maxNodeProvisionTime")]
        public Input<string>? MaxNodeProvisionTime { get; set; }

        [Input("maxTotalUnreadyPercentage")]
        public Input<string>? MaxTotalUnreadyPercentage { get; set; }

        [Input("newPodScaleUpDelay")]
        public Input<string>? NewPodScaleUpDelay { get; set; }

        [Input("okTotalUnreadyCount")]
        public Input<string>? OkTotalUnreadyCount { get; set; }

        [Input("scaleDownDelayAfterAdd")]
        public Input<string>? ScaleDownDelayAfterAdd { get; set; }

        [Input("scaleDownDelayAfterDelete")]
        public Input<string>? ScaleDownDelayAfterDelete { get; set; }

        [Input("scaleDownDelayAfterFailure")]
        public Input<string>? ScaleDownDelayAfterFailure { get; set; }

        [Input("scaleDownUnneededTime")]
        public Input<string>? ScaleDownUnneededTime { get; set; }

        [Input("scaleDownUnreadyTime")]
        public Input<string>? ScaleDownUnreadyTime { get; set; }

        [Input("scaleDownUtilizationThreshold")]
        public Input<string>? ScaleDownUtilizationThreshold { get; set; }

        [Input("scanInterval")]
        public Input<string>? ScanInterval { get; set; }

        [Input("skipNodesWithLocalStorage")]
        public Input<string>? SkipNodesWithLocalStorage { get; set; }

        [Input("skipNodesWithSystemPods")]
        public Input<string>? SkipNodesWithSystemPods { get; set; }

        public ManagedClusterPropertiesAutoScalerProfileArgs()
        {
        }
    }
}
