// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20190601Preview
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azurerm:sql/v20190601preview:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
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

        public GetServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// Administrator username for the server. Once created it cannot be changed.
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        public readonly string? AdministratorLoginPassword;
        /// <summary>
        /// The fully qualified domain name of the server.
        /// </summary>
        public readonly string FullyQualifiedDomainName;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponseResult? Identity;
        /// <summary>
        /// Kind of sql server. This is metadata used for the Azure portal experience.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
        /// </summary>
        public readonly string? MinimalTlsVersion;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of private endpoint connections on a server
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The state of the server.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The version of the server.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string? administratorLoginPassword,

            string fullyQualifiedDomainName,

            Outputs.ResourceIdentityResponseResult? identity,

            string kind,

            string location,

            string? minimalTlsVersion,

            string name,

            ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponseResult> privateEndpointConnections,

            string? publicNetworkAccess,

            string state,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            Identity = identity;
            Kind = kind;
            Location = location;
            MinimalTlsVersion = minimalTlsVersion;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            State = state;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
