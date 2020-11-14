// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforPostgreSQL.V20200214PrivatePreview
{
    /// <summary>
    /// Represents a server.
    /// </summary>
    public partial class Server : Pulumi.CustomResource
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        [Output("administratorLogin")]
        public Output<string?> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        [Output("administratorLoginPassword")]
        public Output<string?> AdministratorLoginPassword { get; private set; } = null!;

        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Status showing whether the data encryption is enabled with customer-managed keys.
        /// </summary>
        [Output("byokEnforcement")]
        public Output<string> ByokEnforcement { get; private set; } = null!;

        /// <summary>
        /// The mode to create a new PostgreSQL server.
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        [Output("delegatedSubnetArguments")]
        public Output<Outputs.ServerPropertiesResponseDelegatedSubnetArguments?> DelegatedSubnetArguments { get; private set; } = null!;

        /// <summary>
        /// The display name of a server.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        [Output("fullyQualifiedDomainName")]
        public Output<string> FullyQualifiedDomainName { get; private set; } = null!;

        /// <summary>
        /// stand by count value can be either enabled or disabled
        /// </summary>
        [Output("haEnabled")]
        public Output<string?> HaEnabled { get; private set; } = null!;

        /// <summary>
        /// A state of a HA server that is visible to user.
        /// </summary>
        [Output("haState")]
        public Output<string> HaState { get; private set; } = null!;

        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maintenance window of a server.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MaintenanceWindowResponse?> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Restore point creation time (ISO8601 format), specifying the time to restore from.
        /// </summary>
        [Output("pointInTimeUTC")]
        public Output<string?> PointInTimeUTC { get; private set; } = null!;

        /// <summary>
        /// public network access is enabled or not
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// The source PostgreSQL server name to restore from.
        /// </summary>
        [Output("sourceServerName")]
        public Output<string?> SourceServerName { get; private set; } = null!;

        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        [Output("standbyAvailabilityZone")]
        public Output<string> StandbyAvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// A state of a server that is visible to user.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        [Output("storageProfile")]
        public Output<Outputs.StorageProfileResponse?> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// PostgreSQL Server version.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:dbforpostgresql/v20200214privatepreview:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:dbforpostgresql/v20200214privatepreview:Server", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:dbforpostgresql/v20200214preview:Server"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Server(name, id, options);
        }
    }

    public sealed class ServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        [Input("administratorLoginPassword")]
        public Input<string>? AdministratorLoginPassword { get; set; }

        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The mode to create a new PostgreSQL server.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        [Input("delegatedSubnetArguments")]
        public Input<Inputs.ServerPropertiesDelegatedSubnetArgumentsArgs>? DelegatedSubnetArguments { get; set; }

        /// <summary>
        /// The display name of a server.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// stand by count value can be either enabled or disabled
        /// </summary>
        [Input("haEnabled")]
        public Input<string>? HaEnabled { get; set; }

        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Maintenance window of a server.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Restore point creation time (ISO8601 format), specifying the time to restore from.
        /// </summary>
        [Input("pointInTimeUTC")]
        public Input<string>? PointInTimeUTC { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// The source PostgreSQL server name to restore from.
        /// </summary>
        [Input("sourceServerName")]
        public Input<string>? SourceServerName { get; set; }

        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// PostgreSQL Server version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerArgs()
        {
        }
    }
}
