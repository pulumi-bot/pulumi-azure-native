// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20171001Preview
{
    public static class GetManagedInstanceKey
    {
        public static Task<GetManagedInstanceKeyResult> InvokeAsync(GetManagedInstanceKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedInstanceKeyResult>("azurerm:sql/v20171001preview:getManagedInstanceKey", args ?? new GetManagedInstanceKeyArgs(), options.WithVersion());
    }


    public sealed class GetManagedInstanceKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed instance key to be retrieved.
        /// </summary>
        [Input("keyName", required: true)]
        public string KeyName { get; set; } = null!;

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public string ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagedInstanceKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagedInstanceKeyResult
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
        /// The key type like 'ServiceManaged', 'AzureKeyVault'.
        /// </summary>
        public readonly string ServerKeyType;
        /// <summary>
        /// Thumbprint of the key.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URI of the key. If the ServerKeyType is AzureKeyVault, then the URI is required.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private GetManagedInstanceKeyResult(
            string creationDate,

            string kind,

            string name,

            string serverKeyType,

            string thumbprint,

            string type,

            string? uri)
        {
            CreationDate = creationDate;
            Kind = kind;
            Name = name;
            ServerKeyType = serverKeyType;
            Thumbprint = thumbprint;
            Type = type;
            Uri = uri;
        }
    }
}
