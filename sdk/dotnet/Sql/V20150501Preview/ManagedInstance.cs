// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20150501Preview
{
    /// <summary>
    /// An Azure SQL managed instance.
    /// </summary>
    public partial class ManagedInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
        /// </summary>
        [Output("administratorLogin")]
        public Output<string?> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The administrator login password (required for managed instance creation).
        /// </summary>
        [Output("administratorLoginPassword")]
        public Output<string?> AdministratorLoginPassword { get; private set; } = null!;

        /// <summary>
        /// Collation of the managed instance.
        /// </summary>
        [Output("collation")]
        public Output<string?> Collation { get; private set; } = null!;

        /// <summary>
        /// The Dns Zone that the managed instance is in.
        /// </summary>
        [Output("dnsZone")]
        public Output<string> DnsZone { get; private set; } = null!;

        /// <summary>
        /// The resource id of another managed instance whose DNS zone this managed instance will share after creation.
        /// </summary>
        [Output("dnsZonePartner")]
        public Output<string?> DnsZonePartner { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name of the managed instance.
        /// </summary>
        [Output("fullyQualifiedDomainName")]
        public Output<string> FullyQualifiedDomainName { get; private set; } = null!;

        /// <summary>
        /// The Azure Active Directory identity of the managed instance.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ResourceIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Id of the instance pool this managed server belongs to.
        /// </summary>
        [Output("instancePoolId")]
        public Output<string?> InstancePoolId { get; private set; } = null!;

        /// <summary>
        /// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies maintenance configuration id to apply to this managed instance.
        /// </summary>
        [Output("maintenanceConfigurationId")]
        public Output<string?> MaintenanceConfigurationId { get; private set; } = null!;

        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: Regular instance creation.
        /// 
        /// Restore: Creates an instance by restoring a set of backups to specific point in time. RestorePointInTime and SourceManagedInstanceId must be specified.
        /// </summary>
        [Output("managedInstanceCreateMode")]
        public Output<string?> ManagedInstanceCreateMode { get; private set; } = null!;

        /// <summary>
        /// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
        /// </summary>
        [Output("minimalTlsVersion")]
        public Output<string?> MinimalTlsVersion { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Connection type used for connecting to the instance.
        /// </summary>
        [Output("proxyOverride")]
        public Output<string?> ProxyOverride { get; private set; } = null!;

        /// <summary>
        /// Whether or not the public data endpoint is enabled.
        /// </summary>
        [Output("publicDataEndpointEnabled")]
        public Output<bool?> PublicDataEndpointEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        /// </summary>
        [Output("restorePointInTime")]
        public Output<string?> RestorePointInTime { get; private set; } = null!;

        /// <summary>
        /// Managed instance SKU. Allowed values for sku.name: GP_Gen4, GP_Gen5, BC_Gen4, BC_Gen5
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the source managed instance associated with create operation of this instance.
        /// </summary>
        [Output("sourceManagedInstanceId")]
        public Output<string?> SourceManagedInstanceId { get; private set; } = null!;

        /// <summary>
        /// The state of the managed instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Storage size in GB. Minimum value: 32. Maximum value: 8192. Increments of 32 GB allowed only.
        /// </summary>
        [Output("storageSizeInGB")]
        public Output<int?> StorageSizeInGB { get; private set; } = null!;

        /// <summary>
        /// Subnet resource ID for the managed instance.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Id of the timezone. Allowed values are timezones supported by Windows.
        /// Windows keeps details on supported timezones, including the id, in registry under
        /// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
        /// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
        /// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
        /// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
        /// </summary>
        [Output("timezoneId")]
        public Output<string?> TimezoneId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
        /// </summary>
        [Output("vCores")]
        public Output<int?> VCores { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedInstance(string name, ManagedInstanceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20150501preview:ManagedInstance", name, args ?? new ManagedInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20150501preview:ManagedInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:sql/v20180601preview:ManagedInstance"},
                    new Pulumi.Alias { Type = "azurerm:sql/v20200202preview:ManagedInstance"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedInstance(name, id, options);
        }
    }

    public sealed class ManagedInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The administrator login password (required for managed instance creation).
        /// </summary>
        [Input("administratorLoginPassword")]
        public Input<string>? AdministratorLoginPassword { get; set; }

        /// <summary>
        /// Collation of the managed instance.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// The resource id of another managed instance whose DNS zone this managed instance will share after creation.
        /// </summary>
        [Input("dnsZonePartner")]
        public Input<string>? DnsZonePartner { get; set; }

        /// <summary>
        /// The Azure Active Directory identity of the managed instance.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Id of the instance pool this managed server belongs to.
        /// </summary>
        [Input("instancePoolId")]
        public Input<string>? InstancePoolId { get; set; }

        /// <summary>
        /// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies maintenance configuration id to apply to this managed instance.
        /// </summary>
        [Input("maintenanceConfigurationId")]
        public Input<string>? MaintenanceConfigurationId { get; set; }

        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: Regular instance creation.
        /// 
        /// Restore: Creates an instance by restoring a set of backups to specific point in time. RestorePointInTime and SourceManagedInstanceId must be specified.
        /// </summary>
        [Input("managedInstanceCreateMode")]
        public Input<string>? ManagedInstanceCreateMode { get; set; }

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public Input<string> ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
        /// </summary>
        [Input("minimalTlsVersion")]
        public Input<string>? MinimalTlsVersion { get; set; }

        /// <summary>
        /// Connection type used for connecting to the instance.
        /// </summary>
        [Input("proxyOverride")]
        public Input<string>? ProxyOverride { get; set; }

        /// <summary>
        /// Whether or not the public data endpoint is enabled.
        /// </summary>
        [Input("publicDataEndpointEnabled")]
        public Input<bool>? PublicDataEndpointEnabled { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// Managed instance SKU. Allowed values for sku.name: GP_Gen4, GP_Gen5, BC_Gen4, BC_Gen5
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// The resource identifier of the source managed instance associated with create operation of this instance.
        /// </summary>
        [Input("sourceManagedInstanceId")]
        public Input<string>? SourceManagedInstanceId { get; set; }

        /// <summary>
        /// Storage size in GB. Minimum value: 32. Maximum value: 8192. Increments of 32 GB allowed only.
        /// </summary>
        [Input("storageSizeInGB")]
        public Input<int>? StorageSizeInGB { get; set; }

        /// <summary>
        /// Subnet resource ID for the managed instance.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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
        /// Id of the timezone. Allowed values are timezones supported by Windows.
        /// Windows keeps details on supported timezones, including the id, in registry under
        /// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
        /// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
        /// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
        /// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
        /// </summary>
        [Input("timezoneId")]
        public Input<string>? TimezoneId { get; set; }

        /// <summary>
        /// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
        /// </summary>
        [Input("vCores")]
        public Input<int>? VCores { get; set; }

        public ManagedInstanceArgs()
        {
        }
    }
}
