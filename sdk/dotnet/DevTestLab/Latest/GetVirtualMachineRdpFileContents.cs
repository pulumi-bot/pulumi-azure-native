// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Latest
{
    [Obsolete(@"The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:devtestlab:getVirtualMachineRdpFileContents'.")]
    public static class GetVirtualMachineRdpFileContents
    {
        /// <summary>
        /// Represents a .rdp file
        /// Latest API Version: 2018-09-15.
        /// </summary>
        public static Task<GetVirtualMachineRdpFileContentsResult> InvokeAsync(GetVirtualMachineRdpFileContentsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineRdpFileContentsResult>("azure-native:devtestlab/latest:getVirtualMachineRdpFileContents", args ?? new GetVirtualMachineRdpFileContentsArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineRdpFileContentsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineRdpFileContentsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineRdpFileContentsResult
    {
        /// <summary>
        /// The contents of the .rdp file
        /// </summary>
        public readonly string? Contents;

        [OutputConstructor]
        private GetVirtualMachineRdpFileContentsResult(string? contents)
        {
            Contents = contents;
        }
    }
}
