// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201201
{
    public static class GetMaintenanceConfiguration
    {
        public static Task<GetMaintenanceConfigurationResult> InvokeAsync(GetMaintenanceConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMaintenanceConfigurationResult>("azure-nextgen:containerservice/v20201201:getMaintenanceConfiguration", args ?? new GetMaintenanceConfigurationArgs(), options.WithVersion());
    }


    public sealed class GetMaintenanceConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the maintenance configuration.
        /// </summary>
        [Input("configName", required: true)]
        public string ConfigName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetMaintenanceConfigurationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMaintenanceConfigurationResult
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Time slots on which upgrade is not allowed.
        /// </summary>
        public readonly ImmutableArray<Outputs.TimeSpanResponse> NotAllowedTime;
        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Weekday time slots allowed to upgrade.
        /// </summary>
        public readonly ImmutableArray<Outputs.TimeInWeekResponse> TimeInWeek;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMaintenanceConfigurationResult(
            string id,

            string name,

            ImmutableArray<Outputs.TimeSpanResponse> notAllowedTime,

            Outputs.SystemDataResponse systemData,

            ImmutableArray<Outputs.TimeInWeekResponse> timeInWeek,

            string type)
        {
            Id = id;
            Name = name;
            NotAllowedTime = notAllowedTime;
            SystemData = systemData;
            TimeInWeek = timeInWeek;
            Type = type;
        }
    }
}
