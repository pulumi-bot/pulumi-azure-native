// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DesktopVirtualization.V20190924Preview
{
    public static class GetHostPool
    {
        public static Task<GetHostPoolResult> InvokeAsync(GetHostPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostPoolResult>("azurerm:desktopvirtualization/v20190924preview:getHostPool", args ?? new GetHostPoolArgs(), options.WithVersion());
    }


    public sealed class GetHostPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the host pool within the specified resource group
        /// </summary>
        [Input("hostPoolName", required: true)]
        public string HostPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHostPoolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHostPoolResult
    {
        /// <summary>
        /// List of applicationGroup links.
        /// </summary>
        public readonly ImmutableArray<string> ApplicationGroupReferences;
        /// <summary>
        /// Custom rdp property of HostPool.
        /// </summary>
        public readonly string? CustomRdpProperty;
        /// <summary>
        /// Description of HostPool.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Friendly name of HostPool.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// HostPool type for desktop.
        /// </summary>
        public readonly string HostPoolType;
        /// <summary>
        /// The type of the load balancer.
        /// </summary>
        public readonly string LoadBalancerType;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The max session limit of HostPool.
        /// </summary>
        public readonly int? MaxSessionLimit;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// PersonalDesktopAssignment type for HostPool.
        /// </summary>
        public readonly string? PersonalDesktopAssignmentType;
        /// <summary>
        /// The type of preferred application group type, default to Desktop Application Group
        /// </summary>
        public readonly string PreferredAppGroupType;
        /// <summary>
        /// The registration info of HostPool.
        /// </summary>
        public readonly Outputs.RegistrationInfoResponseResult? RegistrationInfo;
        /// <summary>
        /// The ring number of HostPool.
        /// </summary>
        public readonly int? Ring;
        /// <summary>
        /// Path to keyvault containing ssoContext secret.
        /// </summary>
        public readonly string? SsoContext;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Is validation environment.
        /// </summary>
        public readonly bool? ValidationEnvironment;
        /// <summary>
        /// VM template for sessionhosts configuration within hostpool.
        /// </summary>
        public readonly string? VmTemplate;

        [OutputConstructor]
        private GetHostPoolResult(
            ImmutableArray<string> applicationGroupReferences,

            string? customRdpProperty,

            string? description,

            string? friendlyName,

            string hostPoolType,

            string loadBalancerType,

            string location,

            int? maxSessionLimit,

            string name,

            string? personalDesktopAssignmentType,

            string preferredAppGroupType,

            Outputs.RegistrationInfoResponseResult? registrationInfo,

            int? ring,

            string? ssoContext,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? validationEnvironment,

            string? vmTemplate)
        {
            ApplicationGroupReferences = applicationGroupReferences;
            CustomRdpProperty = customRdpProperty;
            Description = description;
            FriendlyName = friendlyName;
            HostPoolType = hostPoolType;
            LoadBalancerType = loadBalancerType;
            Location = location;
            MaxSessionLimit = maxSessionLimit;
            Name = name;
            PersonalDesktopAssignmentType = personalDesktopAssignmentType;
            PreferredAppGroupType = preferredAppGroupType;
            RegistrationInfo = registrationInfo;
            Ring = ring;
            SsoContext = ssoContext;
            Tags = tags;
            Type = type;
            ValidationEnvironment = validationEnvironment;
            VmTemplate = vmTemplate;
        }
    }
}
