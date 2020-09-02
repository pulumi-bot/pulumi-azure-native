// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AzureData.V20170301Preview
{
    public static class GetSqlServer
    {
        public static Task<GetSqlServerResult> InvokeAsync(GetSqlServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSqlServerResult>("azurerm:azuredata/v20170301preview:getSqlServer", args ?? new GetSqlServerArgs(), options.WithVersion());
    }


    public sealed class GetSqlServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL Server.
        /// </summary>
        [Input("sqlServerName", required: true)]
        public string SqlServerName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL Server registration.
        /// </summary>
        [Input("sqlServerRegistrationName", required: true)]
        public string SqlServerRegistrationName { get; set; } = null!;

        public GetSqlServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSqlServerResult
    {
        /// <summary>
        /// Cores of the Sql Server.
        /// </summary>
        public readonly int? Cores;
        /// <summary>
        /// Sql Server Edition.
        /// </summary>
        public readonly string? Edition;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Sql Server Json Property Bag.
        /// </summary>
        public readonly string? PropertyBag;
        /// <summary>
        /// ID for Parent Sql Server Registration.
        /// </summary>
        public readonly string? RegistrationID;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the Sql Server.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetSqlServerResult(
            int? cores,

            string? edition,

            string name,

            string? propertyBag,

            string? registrationID,

            string type,

            string? version)
        {
            Cores = cores;
            Edition = edition;
            Name = name;
            PropertyBag = propertyBag;
            RegistrationID = registrationID;
            Type = type;
            Version = version;
        }
    }
}
