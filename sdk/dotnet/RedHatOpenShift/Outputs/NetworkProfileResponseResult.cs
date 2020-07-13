// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RedHatOpenShift.Outputs
{

    [OutputType]
    public sealed class NetworkProfileResponseResult
    {
        /// <summary>
        /// The CIDR used for OpenShift/Kubernetes Pods (immutable).
        /// </summary>
        public readonly string? PodCidr;
        /// <summary>
        /// The CIDR used for OpenShift/Kubernetes Services (immutable).
        /// </summary>
        public readonly string? ServiceCidr;

        [OutputConstructor]
        private NetworkProfileResponseResult(
            string? podCidr,

            string? serviceCidr)
        {
            PodCidr = podCidr;
            ServiceCidr = serviceCidr;
        }
    }
}