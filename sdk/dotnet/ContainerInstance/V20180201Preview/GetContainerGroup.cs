// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20180201Preview
{
    public static class GetContainerGroup
    {
        public static Task<GetContainerGroupResult> InvokeAsync(GetContainerGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContainerGroupResult>("azurerm:containerinstance/v20180201preview:getContainerGroup", args ?? new GetContainerGroupArgs(), options.WithVersion());
    }


    public sealed class GetContainerGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container group.
        /// </summary>
        [Input("containerGroupName", required: true)]
        public string ContainerGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetContainerGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContainerGroupResult
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
        public readonly Outputs.ContainerGroupResponseInstanceViewResult InstanceView;
        /// <summary>
        /// The IP address type of the container group.
        /// </summary>
        public readonly Outputs.IpAddressResponseResult? IpAddress;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
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
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The list of volumes that can be mounted by containers in this container group.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponseResult> Volumes;

        [OutputConstructor]
        private GetContainerGroupResult(
            ImmutableArray<Outputs.ContainerResponseResult> containers,

            ImmutableArray<Outputs.ImageRegistryCredentialResponseResult> imageRegistryCredentials,

            Outputs.ContainerGroupResponseInstanceViewResult instanceView,

            Outputs.IpAddressResponseResult? ipAddress,

            string? location,

            string name,

            string osType,

            string provisioningState,

            string? restartPolicy,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.VolumeResponseResult> volumes)
        {
            Containers = containers;
            ImageRegistryCredentials = imageRegistryCredentials;
            InstanceView = instanceView;
            IpAddress = ipAddress;
            Location = location;
            Name = name;
            OsType = osType;
            ProvisioningState = provisioningState;
            RestartPolicy = restartPolicy;
            Tags = tags;
            Type = type;
            Volumes = volumes;
        }
    }
}
