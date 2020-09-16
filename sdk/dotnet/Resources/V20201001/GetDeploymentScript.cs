// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20201001
{
    public static class GetDeploymentScript
    {
        public static Task<GetDeploymentScriptResult> InvokeAsync(GetDeploymentScriptArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentScriptResult>("azure-nextgen:resources/v20201001:getDeploymentScript", args ?? new GetDeploymentScriptArgs(), options.WithVersion());
    }


    public sealed class GetDeploymentScriptArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment script.
        /// </summary>
        [Input("scriptName", required: true)]
        public string ScriptName { get; set; } = null!;

        public GetDeploymentScriptArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeploymentScriptResult
    {
        /// <summary>
        /// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// Type of the script.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The location of the ACI and the storage account for the deployment script.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDeploymentScriptResult(
            Outputs.ManagedServiceIdentityResponse? identity,

            string kind,

            string location,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Kind = kind;
            Location = location;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
