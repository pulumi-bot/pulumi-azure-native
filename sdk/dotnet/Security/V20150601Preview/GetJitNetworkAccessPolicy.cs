// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20150601Preview
{
    public static class GetJitNetworkAccessPolicy
    {
        public static Task<GetJitNetworkAccessPolicyResult> InvokeAsync(GetJitNetworkAccessPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJitNetworkAccessPolicyResult>("azurerm:security/v20150601preview:getJitNetworkAccessPolicy", args ?? new GetJitNetworkAccessPolicyArgs(), options.WithVersion());
    }


    public sealed class GetJitNetworkAccessPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location where ASC stores the data of the subscription. can be retrieved from Get locations
        /// </summary>
        [Input("ascLocation", required: true)]
        public string AscLocation { get; set; } = null!;

        /// <summary>
        /// Name of a Just-in-Time access configuration policy.
        /// </summary>
        [Input("jitNetworkAccessPolicyName", required: true)]
        public string JitNetworkAccessPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetJitNetworkAccessPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJitNetworkAccessPolicyResult
    {
        /// <summary>
        /// Kind of the resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Location where the resource is stored
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets the provisioning state of the Just-in-Time policy.
        /// </summary>
        public readonly string ProvisioningState;
        public readonly ImmutableArray<Outputs.JitNetworkAccessRequestResponseResult> Requests;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Configurations for Microsoft.Compute/virtualMachines resource type.
        /// </summary>
        public readonly ImmutableArray<Outputs.JitNetworkAccessPolicyVirtualMachineResponseResult> VirtualMachines;

        [OutputConstructor]
        private GetJitNetworkAccessPolicyResult(
            string? kind,

            string location,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.JitNetworkAccessRequestResponseResult> requests,

            string type,

            ImmutableArray<Outputs.JitNetworkAccessPolicyVirtualMachineResponseResult> virtualMachines)
        {
            Kind = kind;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Requests = requests;
            Type = type;
            VirtualMachines = virtualMachines;
        }
    }
}
