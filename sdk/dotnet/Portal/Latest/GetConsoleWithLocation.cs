// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.Latest
{
    public static class GetConsoleWithLocation
    {
        public static Task<GetConsoleWithLocationResult> InvokeAsync(GetConsoleWithLocationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConsoleWithLocationResult>("azurerm:portal/latest:getConsoleWithLocation", args ?? new GetConsoleWithLocationArgs(), options.WithVersion());
    }


    public sealed class GetConsoleWithLocationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the console
        /// </summary>
        [Input("consoleName", required: true)]
        public string ConsoleName { get; set; } = null!;

        /// <summary>
        /// The provider location
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        public GetConsoleWithLocationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConsoleWithLocationResult
    {
        /// <summary>
        /// Cloud shell console properties.
        /// </summary>
        public readonly Outputs.ConsolePropertiesResponseResult Properties;

        [OutputConstructor]
        private GetConsoleWithLocationResult(Outputs.ConsolePropertiesResponseResult properties)
        {
            Properties = properties;
        }
    }
}