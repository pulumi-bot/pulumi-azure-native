// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple
{
    public static class GetManagerDeviceIscsiserver
    {
        public static Task<GetManagerDeviceIscsiserverResult> InvokeAsync(GetManagerDeviceIscsiserverArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagerDeviceIscsiserverResult>("azurerm:storsimple:getManagerDeviceIscsiserver", args ?? new GetManagerDeviceIscsiserverArgs(), options.WithVersion());
    }


    public sealed class GetManagerDeviceIscsiserverArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public string ManagerName { get; set; } = null!;

        /// <summary>
        /// The iSCSI server name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagerDeviceIscsiserverArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagerDeviceIscsiserverResult
    {
        /// <summary>
        /// The name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties.
        /// </summary>
        public readonly Outputs.ISCSIServerPropertiesResponseResult Properties;
        /// <summary>
        /// The type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagerDeviceIscsiserverResult(
            string name,

            Outputs.ISCSIServerPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}