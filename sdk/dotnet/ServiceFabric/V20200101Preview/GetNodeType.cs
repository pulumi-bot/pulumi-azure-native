// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20200101Preview
{
    public static class GetNodeType
    {
        public static Task<GetNodeTypeResult> InvokeAsync(GetNodeTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodeTypeResult>("azurerm:servicefabric/v20200101preview:getNodeType", args ?? new GetNodeTypeArgs(), options.WithVersion());
    }


    public sealed class GetNodeTypeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the node type.
        /// </summary>
        [Input("nodeTypeName", required: true)]
        public string NodeTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNodeTypeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodeTypeResult
    {
        /// <summary>
        /// The range of ports from which cluster assigned port to Service Fabric applications.
        /// </summary>
        public readonly Outputs.EndpointRangeDescriptionResponseResult? ApplicationPorts;
        /// <summary>
        /// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Capacities;
        /// <summary>
        /// Disk size for each vm in the node type in GBs.
        /// </summary>
        public readonly int DataDiskSizeGB;
        /// <summary>
        /// The range of ephemeral ports that nodes in this node type should be configured with.
        /// </summary>
        public readonly Outputs.EndpointRangeDescriptionResponseResult? EphemeralPorts;
        /// <summary>
        /// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
        /// </summary>
        public readonly bool IsPrimary;
        /// <summary>
        /// Azure resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? PlacementProperties;
        /// <summary>
        /// The provisioning state of the managed cluster resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Set of extensions that should be installed onto the virtual machines.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMSSExtensionResponseResult> VmExtensions;
        /// <summary>
        /// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
        /// </summary>
        public readonly string? VmImageOffer;
        /// <summary>
        /// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
        /// </summary>
        public readonly string? VmImagePublisher;
        /// <summary>
        /// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
        /// </summary>
        public readonly string? VmImageSku;
        /// <summary>
        /// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
        /// </summary>
        public readonly string? VmImageVersion;
        /// <summary>
        /// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
        /// </summary>
        public readonly int VmInstanceCount;
        /// <summary>
        /// The secrets to install in the virtual machines.
        /// </summary>
        public readonly ImmutableArray<Outputs.VaultSecretGroupResponseResult> VmSecrets;
        /// <summary>
        /// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
        /// </summary>
        public readonly string? VmSize;

        [OutputConstructor]
        private GetNodeTypeResult(
            Outputs.EndpointRangeDescriptionResponseResult? applicationPorts,

            ImmutableDictionary<string, string>? capacities,

            int dataDiskSizeGB,

            Outputs.EndpointRangeDescriptionResponseResult? ephemeralPorts,

            bool isPrimary,

            string name,

            ImmutableDictionary<string, string>? placementProperties,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.VMSSExtensionResponseResult> vmExtensions,

            string? vmImageOffer,

            string? vmImagePublisher,

            string? vmImageSku,

            string? vmImageVersion,

            int vmInstanceCount,

            ImmutableArray<Outputs.VaultSecretGroupResponseResult> vmSecrets,

            string? vmSize)
        {
            ApplicationPorts = applicationPorts;
            Capacities = capacities;
            DataDiskSizeGB = dataDiskSizeGB;
            EphemeralPorts = ephemeralPorts;
            IsPrimary = isPrimary;
            Name = name;
            PlacementProperties = placementProperties;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            VmExtensions = vmExtensions;
            VmImageOffer = vmImageOffer;
            VmImagePublisher = vmImagePublisher;
            VmImageSku = vmImageSku;
            VmImageVersion = vmImageVersion;
            VmInstanceCount = vmInstanceCount;
            VmSecrets = vmSecrets;
            VmSize = vmSize;
        }
    }
}
