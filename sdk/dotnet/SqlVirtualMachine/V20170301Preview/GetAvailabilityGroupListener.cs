// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview
{
    public static class GetAvailabilityGroupListener
    {
        public static Task<GetAvailabilityGroupListenerResult> InvokeAsync(GetAvailabilityGroupListenerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityGroupListenerResult>("azurerm:sqlvirtualmachine/v20170301preview:getAvailabilityGroupListener", args ?? new GetAvailabilityGroupListenerArgs(), options.WithVersion());
    }


    public sealed class GetAvailabilityGroupListenerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the availability group listener.
        /// </summary>
        [Input("availabilityGroupListenerName", required: true)]
        public string AvailabilityGroupListenerName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine group.
        /// </summary>
        [Input("sqlVirtualMachineGroupName", required: true)]
        public string SqlVirtualMachineGroupName { get; set; } = null!;

        public GetAvailabilityGroupListenerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAvailabilityGroupListenerResult
    {
        /// <summary>
        /// Name of the availability group.
        /// </summary>
        public readonly string? AvailabilityGroupName;
        /// <summary>
        /// Create a default availability group if it does not exist.
        /// </summary>
        public readonly bool? CreateDefaultAvailabilityGroupIfNotExist;
        /// <summary>
        /// List of load balancer configurations for an availability group listener.
        /// </summary>
        public readonly ImmutableArray<Outputs.LoadBalancerConfigurationResponseResult> LoadBalancerConfigurations;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Listener port.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Provisioning state to track the async operation status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAvailabilityGroupListenerResult(
            string? availabilityGroupName,

            bool? createDefaultAvailabilityGroupIfNotExist,

            ImmutableArray<Outputs.LoadBalancerConfigurationResponseResult> loadBalancerConfigurations,

            string name,

            int? port,

            string provisioningState,

            string type)
        {
            AvailabilityGroupName = availabilityGroupName;
            CreateDefaultAvailabilityGroupIfNotExist = createDefaultAvailabilityGroupIfNotExist;
            LoadBalancerConfigurations = loadBalancerConfigurations;
            Name = name;
            Port = port;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
