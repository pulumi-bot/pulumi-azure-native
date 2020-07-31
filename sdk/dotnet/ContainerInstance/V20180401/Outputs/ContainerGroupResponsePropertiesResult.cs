// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20180401.Outputs
{

    [OutputType]
    public sealed class ContainerGroupResponsePropertiesResult
    {
        /// <summary>
        /// The containers within the container group.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerResponseResult> Containers;
        /// <summary>
        /// The image registry credentials by which the container group is created from.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageRegistryCredentialResponseResult> ImageRegistryCredentials;
        /// <summary>
        /// The instance view of the container group. Only valid in response.
        /// </summary>
        public readonly Outputs.ContainerGroupResponsePropertiesResult InstanceView;
        /// <summary>
        /// The IP address type of the container group.
        /// </summary>
        public readonly Outputs.IpAddressResponseResult? IpAddress;
        /// <summary>
        /// The operating system type required by the containers in the container group.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The provisioning state of the container group. This only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Restart policy for all containers within the container group. 
        /// - `Always` Always restart
        /// - `OnFailure` Restart on failure
        /// - `Never` Never restart
        /// </summary>
        public readonly string? RestartPolicy;
        /// <summary>
        /// The list of volumes that can be mounted by containers in this container group.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponseResult> Volumes;

        [OutputConstructor]
        private ContainerGroupResponsePropertiesResult(
            ImmutableArray<Outputs.ContainerResponseResult> containers,

            ImmutableArray<Outputs.ImageRegistryCredentialResponseResult> imageRegistryCredentials,

            Outputs.ContainerGroupResponsePropertiesResult instanceView,

            Outputs.IpAddressResponseResult? ipAddress,

            string osType,

            string provisioningState,

            string? restartPolicy,

            ImmutableArray<Outputs.VolumeResponseResult> volumes)
        {
            Containers = containers;
            ImageRegistryCredentials = imageRegistryCredentials;
            InstanceView = instanceView;
            IpAddress = ipAddress;
            OsType = osType;
            ProvisioningState = provisioningState;
            RestartPolicy = restartPolicy;
            Volumes = volumes;
        }
    }
}