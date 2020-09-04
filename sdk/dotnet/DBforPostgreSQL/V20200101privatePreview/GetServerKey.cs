// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforPostgreSQL.V20200101privatePreview
{
    public static class GetServerKey
    {
        public static Task<GetServerKeyResult> InvokeAsync(GetServerKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerKeyResult>("azurerm:dbforpostgresql/v20200101privatepreview:getServerKey", args ?? new GetServerKeyArgs(), options.WithVersion());
    }


    public sealed class GetServerKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the PostgreSQL Server key to be retrieved.
        /// </summary>
        [Input("keyName", required: true)]
        public string KeyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerKeyResult
    {
        /// <summary>
        /// The key creation date.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// Kind of encryption protector. This is metadata used for the Azure portal experience.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The key type like 'AzureKeyVault'.
        /// </summary>
        public readonly string ServerKeyType;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URI of the key.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private GetServerKeyResult(
            string creationDate,

            string kind,

            string name,

            string serverKeyType,

            string type,

            string? uri)
        {
            CreationDate = creationDate;
            Kind = kind;
            Name = name;
            ServerKeyType = serverKeyType;
            Type = type;
            Uri = uri;
        }
    }
}
