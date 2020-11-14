// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Kusto.V20200614
{
    public static class GetAttachedDatabaseConfiguration
    {
        public static Task<GetAttachedDatabaseConfigurationResult> InvokeAsync(GetAttachedDatabaseConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAttachedDatabaseConfigurationResult>("azure-nextgen:kusto/v20200614:getAttachedDatabaseConfiguration", args ?? new GetAttachedDatabaseConfigurationArgs(), options.WithVersion());
    }


    public sealed class GetAttachedDatabaseConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the attached database configuration.
        /// </summary>
        [Input("attachedDatabaseConfigurationName", required: true)]
        public string AttachedDatabaseConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAttachedDatabaseConfigurationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAttachedDatabaseConfigurationResult
    {
        /// <summary>
        /// The list of databases from the clusterResourceId which are currently attached to the cluster.
        /// </summary>
        public readonly ImmutableArray<string> AttachedDatabaseNames;
        /// <summary>
        /// The resource id of the cluster where the databases you would like to attach reside.
        /// </summary>
        public readonly string ClusterResourceId;
        /// <summary>
        /// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The default principals modification kind
        /// </summary>
        public readonly string DefaultPrincipalsModificationKind;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAttachedDatabaseConfigurationResult(
            ImmutableArray<string> attachedDatabaseNames,

            string clusterResourceId,

            string databaseName,

            string defaultPrincipalsModificationKind,

            string? location,

            string name,

            string provisioningState,

            string type)
        {
            AttachedDatabaseNames = attachedDatabaseNames;
            ClusterResourceId = clusterResourceId;
            DatabaseName = databaseName;
            DefaultPrincipalsModificationKind = defaultPrincipalsModificationKind;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
