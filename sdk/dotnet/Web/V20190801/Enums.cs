// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Web.V20190801
{
    /// <summary>
    /// Action object.
    /// </summary>
    [EnumType]
    public readonly struct AccessControlEntryAction : IEquatable<AccessControlEntryAction>
    {
        private readonly string _value;

        private AccessControlEntryAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessControlEntryAction Permit { get; } = new AccessControlEntryAction("Permit");
        public static AccessControlEntryAction Deny { get; } = new AccessControlEntryAction("Deny");

        public static bool operator ==(AccessControlEntryAction left, AccessControlEntryAction right) => left.Equals(right);
        public static bool operator !=(AccessControlEntryAction left, AccessControlEntryAction right) => !left.Equals(right);

        public static explicit operator string(AccessControlEntryAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessControlEntryAction other && Equals(other);
        public bool Equals(AccessControlEntryAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Predefined action to be taken.
    /// </summary>
    [EnumType]
    public readonly struct AutoHealActionType : IEquatable<AutoHealActionType>
    {
        private readonly string _value;

        private AutoHealActionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AutoHealActionType Recycle { get; } = new AutoHealActionType("Recycle");
        public static AutoHealActionType LogEvent { get; } = new AutoHealActionType("LogEvent");
        public static AutoHealActionType CustomAction { get; } = new AutoHealActionType("CustomAction");

        public static bool operator ==(AutoHealActionType left, AutoHealActionType right) => left.Equals(right);
        public static bool operator !=(AutoHealActionType left, AutoHealActionType right) => !left.Equals(right);

        public static explicit operator string(AutoHealActionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutoHealActionType other && Equals(other);
        public bool Equals(AutoHealActionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Azure resource type.
    /// </summary>
    [EnumType]
    public readonly struct AzureResourceType : IEquatable<AzureResourceType>
    {
        private readonly string _value;

        private AzureResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AzureResourceType Website { get; } = new AzureResourceType("Website");
        public static AzureResourceType TrafficManager { get; } = new AzureResourceType("TrafficManager");

        public static bool operator ==(AzureResourceType left, AzureResourceType right) => left.Equals(right);
        public static bool operator !=(AzureResourceType left, AzureResourceType right) => !left.Equals(right);

        public static explicit operator string(AzureResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AzureResourceType other && Equals(other);
        public bool Equals(AzureResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Shared or dedicated app hosting.
    /// </summary>
    [EnumType]
    public readonly struct ComputeModeOptions : IEquatable<ComputeModeOptions>
    {
        private readonly string _value;

        private ComputeModeOptions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComputeModeOptions Shared { get; } = new ComputeModeOptions("Shared");
        public static ComputeModeOptions Dedicated { get; } = new ComputeModeOptions("Dedicated");
        public static ComputeModeOptions Dynamic { get; } = new ComputeModeOptions("Dynamic");

        public static bool operator ==(ComputeModeOptions left, ComputeModeOptions right) => left.Equals(right);
        public static bool operator !=(ComputeModeOptions left, ComputeModeOptions right) => !left.Equals(right);

        public static explicit operator string(ComputeModeOptions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeModeOptions other && Equals(other);
        public bool Equals(ComputeModeOptions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of database.
    /// </summary>
    [EnumType]
    public readonly struct ConnectionStringType : IEquatable<ConnectionStringType>
    {
        private readonly string _value;

        private ConnectionStringType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectionStringType MySql { get; } = new ConnectionStringType("MySql");
        public static ConnectionStringType SQLServer { get; } = new ConnectionStringType("SQLServer");
        public static ConnectionStringType SQLAzure { get; } = new ConnectionStringType("SQLAzure");
        public static ConnectionStringType Custom { get; } = new ConnectionStringType("Custom");
        public static ConnectionStringType NotificationHub { get; } = new ConnectionStringType("NotificationHub");
        public static ConnectionStringType ServiceBus { get; } = new ConnectionStringType("ServiceBus");
        public static ConnectionStringType EventHub { get; } = new ConnectionStringType("EventHub");
        public static ConnectionStringType ApiHub { get; } = new ConnectionStringType("ApiHub");
        public static ConnectionStringType DocDb { get; } = new ConnectionStringType("DocDb");
        public static ConnectionStringType RedisCache { get; } = new ConnectionStringType("RedisCache");
        public static ConnectionStringType PostgreSQL { get; } = new ConnectionStringType("PostgreSQL");

        public static bool operator ==(ConnectionStringType left, ConnectionStringType right) => left.Equals(right);
        public static bool operator !=(ConnectionStringType left, ConnectionStringType right) => !left.Equals(right);

        public static explicit operator string(ConnectionStringType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectionStringType other && Equals(other);
        public bool Equals(ConnectionStringType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Custom DNS record type.
    /// </summary>
    [EnumType]
    public readonly struct CustomHostNameDnsRecordType : IEquatable<CustomHostNameDnsRecordType>
    {
        private readonly string _value;

        private CustomHostNameDnsRecordType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CustomHostNameDnsRecordType CName { get; } = new CustomHostNameDnsRecordType("CName");
        public static CustomHostNameDnsRecordType A { get; } = new CustomHostNameDnsRecordType("A");

        public static bool operator ==(CustomHostNameDnsRecordType left, CustomHostNameDnsRecordType right) => left.Equals(right);
        public static bool operator !=(CustomHostNameDnsRecordType left, CustomHostNameDnsRecordType right) => !left.Equals(right);

        public static explicit operator string(CustomHostNameDnsRecordType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CustomHostNameDnsRecordType other && Equals(other);
        public bool Equals(CustomHostNameDnsRecordType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Database type (e.g. SqlAzure / MySql).
    /// </summary>
    [EnumType]
    public readonly struct DatabaseType : IEquatable<DatabaseType>
    {
        private readonly string _value;

        private DatabaseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatabaseType SqlAzure { get; } = new DatabaseType("SqlAzure");
        public static DatabaseType MySql { get; } = new DatabaseType("MySql");
        public static DatabaseType LocalMySql { get; } = new DatabaseType("LocalMySql");
        public static DatabaseType PostgreSql { get; } = new DatabaseType("PostgreSql");

        public static bool operator ==(DatabaseType left, DatabaseType right) => left.Equals(right);
        public static bool operator !=(DatabaseType left, DatabaseType right) => !left.Equals(right);

        public static explicit operator string(DatabaseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabaseType other && Equals(other);
        public bool Equals(DatabaseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The unit of time for how often the backup should be executed (e.g. for weekly backup, this should be set to Day and FrequencyInterval should be set to 7)
    /// </summary>
    [EnumType]
    public readonly struct FrequencyUnit : IEquatable<FrequencyUnit>
    {
        private readonly string _value;

        private FrequencyUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FrequencyUnit Day { get; } = new FrequencyUnit("Day");
        public static FrequencyUnit Hour { get; } = new FrequencyUnit("Hour");

        public static bool operator ==(FrequencyUnit left, FrequencyUnit right) => left.Equals(right);
        public static bool operator !=(FrequencyUnit left, FrequencyUnit right) => !left.Equals(right);

        public static explicit operator string(FrequencyUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FrequencyUnit other && Equals(other);
        public bool Equals(FrequencyUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// State of FTP / FTPS service
    /// </summary>
    [EnumType]
    public readonly struct FtpsState : IEquatable<FtpsState>
    {
        private readonly string _value;

        private FtpsState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FtpsState AllAllowed { get; } = new FtpsState("AllAllowed");
        public static FtpsState FtpsOnly { get; } = new FtpsState("FtpsOnly");
        public static FtpsState Disabled { get; } = new FtpsState("Disabled");

        public static bool operator ==(FtpsState left, FtpsState right) => left.Equals(right);
        public static bool operator !=(FtpsState left, FtpsState right) => !left.Equals(right);

        public static explicit operator string(FtpsState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FtpsState other && Equals(other);
        public bool Equals(FtpsState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Hostname type.
    /// </summary>
    [EnumType]
    public readonly struct HostNameType : IEquatable<HostNameType>
    {
        private readonly string _value;

        private HostNameType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostNameType Verified { get; } = new HostNameType("Verified");
        public static HostNameType Managed { get; } = new HostNameType("Managed");

        public static bool operator ==(HostNameType left, HostNameType right) => left.Equals(right);
        public static bool operator !=(HostNameType left, HostNameType right) => !left.Equals(right);

        public static explicit operator string(HostNameType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostNameType other && Equals(other);
        public bool Equals(HostNameType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether the hostname is a standard or repository hostname.
    /// </summary>
    [EnumType]
    public readonly struct HostType : IEquatable<HostType>
    {
        private readonly string _value;

        private HostType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostType Standard { get; } = new HostType("Standard");
        public static HostType Repository { get; } = new HostType("Repository");

        public static bool operator ==(HostType left, HostType right) => left.Equals(right);
        public static bool operator !=(HostType left, HostType right) => !left.Equals(right);

        public static explicit operator string(HostType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostType other && Equals(other);
        public bool Equals(HostType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
    /// </summary>
    [EnumType]
    public readonly struct InternalLoadBalancingMode : IEquatable<InternalLoadBalancingMode>
    {
        private readonly string _value;

        private InternalLoadBalancingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InternalLoadBalancingMode None { get; } = new InternalLoadBalancingMode("None");
        public static InternalLoadBalancingMode Web { get; } = new InternalLoadBalancingMode("Web");
        public static InternalLoadBalancingMode Publishing { get; } = new InternalLoadBalancingMode("Publishing");

        public static bool operator ==(InternalLoadBalancingMode left, InternalLoadBalancingMode right) => left.Equals(right);
        public static bool operator !=(InternalLoadBalancingMode left, InternalLoadBalancingMode right) => !left.Equals(right);

        public static explicit operator string(InternalLoadBalancingMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InternalLoadBalancingMode other && Equals(other);
        public bool Equals(InternalLoadBalancingMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines what this IP filter will be used for. This is to support IP filtering on proxies.
    /// </summary>
    [EnumType]
    public readonly struct IpFilterTag : IEquatable<IpFilterTag>
    {
        private readonly string _value;

        private IpFilterTag(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IpFilterTag Default { get; } = new IpFilterTag("Default");
        public static IpFilterTag XffProxy { get; } = new IpFilterTag("XffProxy");

        public static bool operator ==(IpFilterTag left, IpFilterTag right) => left.Equals(right);
        public static bool operator !=(IpFilterTag left, IpFilterTag right) => !left.Equals(right);

        public static explicit operator string(IpFilterTag value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IpFilterTag other && Equals(other);
        public bool Equals(IpFilterTag other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Managed pipeline mode.
    /// </summary>
    [EnumType]
    public readonly struct ManagedPipelineMode : IEquatable<ManagedPipelineMode>
    {
        private readonly string _value;

        private ManagedPipelineMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedPipelineMode Integrated { get; } = new ManagedPipelineMode("Integrated");
        public static ManagedPipelineMode Classic { get; } = new ManagedPipelineMode("Classic");

        public static bool operator ==(ManagedPipelineMode left, ManagedPipelineMode right) => left.Equals(right);
        public static bool operator !=(ManagedPipelineMode left, ManagedPipelineMode right) => !left.Equals(right);

        public static explicit operator string(ManagedPipelineMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedPipelineMode other && Equals(other);
        public bool Equals(ManagedPipelineMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of managed service identity.
    /// </summary>
    [EnumType]
    public readonly struct ManagedServiceIdentityType : IEquatable<ManagedServiceIdentityType>
    {
        private readonly string _value;

        private ManagedServiceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedServiceIdentityType SystemAssigned { get; } = new ManagedServiceIdentityType("SystemAssigned");
        public static ManagedServiceIdentityType UserAssigned { get; } = new ManagedServiceIdentityType("UserAssigned");
        public static ManagedServiceIdentityType SystemAssigned_UserAssigned { get; } = new ManagedServiceIdentityType("SystemAssigned, UserAssigned");
        public static ManagedServiceIdentityType None { get; } = new ManagedServiceIdentityType("None");

        public static bool operator ==(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => left.Equals(right);
        public static bool operator !=(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ManagedServiceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedServiceIdentityType other && Equals(other);
        public bool Equals(ManagedServiceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Public Certificate Location
    /// </summary>
    [EnumType]
    public readonly struct PublicCertificateLocation : IEquatable<PublicCertificateLocation>
    {
        private readonly string _value;

        private PublicCertificateLocation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PublicCertificateLocation CurrentUserMy { get; } = new PublicCertificateLocation("CurrentUserMy");
        public static PublicCertificateLocation LocalMachineMy { get; } = new PublicCertificateLocation("LocalMachineMy");
        public static PublicCertificateLocation Unknown { get; } = new PublicCertificateLocation("Unknown");

        public static bool operator ==(PublicCertificateLocation left, PublicCertificateLocation right) => left.Equals(right);
        public static bool operator !=(PublicCertificateLocation left, PublicCertificateLocation right) => !left.Equals(right);

        public static explicit operator string(PublicCertificateLocation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicCertificateLocation other && Equals(other);
        public bool Equals(PublicCertificateLocation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Site redundancy mode
    /// </summary>
    [EnumType]
    public readonly struct RedundancyMode : IEquatable<RedundancyMode>
    {
        private readonly string _value;

        private RedundancyMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RedundancyMode None { get; } = new RedundancyMode("None");
        public static RedundancyMode Manual { get; } = new RedundancyMode("Manual");
        public static RedundancyMode Failover { get; } = new RedundancyMode("Failover");
        public static RedundancyMode ActiveActive { get; } = new RedundancyMode("ActiveActive");
        public static RedundancyMode GeoRedundant { get; } = new RedundancyMode("GeoRedundant");

        public static bool operator ==(RedundancyMode left, RedundancyMode right) => left.Equals(right);
        public static bool operator !=(RedundancyMode left, RedundancyMode right) => !left.Equals(right);

        public static explicit operator string(RedundancyMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RedundancyMode other && Equals(other);
        public bool Equals(RedundancyMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of route this is:
    /// DEFAULT - By default, every app has routes to the local address ranges specified by RFC1918
    /// INHERITED - Routes inherited from the real Virtual Network routes
    /// STATIC - Static route set on the app only
    /// 
    /// These values will be used for syncing an app's routes with those from a Virtual Network.
    /// </summary>
    [EnumType]
    public readonly struct RouteType : IEquatable<RouteType>
    {
        private readonly string _value;

        private RouteType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RouteType DEFAULT { get; } = new RouteType("DEFAULT");
        public static RouteType INHERITED { get; } = new RouteType("INHERITED");
        public static RouteType STATIC { get; } = new RouteType("STATIC");

        public static bool operator ==(RouteType left, RouteType right) => left.Equals(right);
        public static bool operator !=(RouteType left, RouteType right) => !left.Equals(right);

        public static explicit operator string(RouteType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RouteType other && Equals(other);
        public bool Equals(RouteType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SCM type.
    /// </summary>
    [EnumType]
    public readonly struct ScmType : IEquatable<ScmType>
    {
        private readonly string _value;

        private ScmType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScmType None { get; } = new ScmType("None");
        public static ScmType Dropbox { get; } = new ScmType("Dropbox");
        public static ScmType Tfs { get; } = new ScmType("Tfs");
        public static ScmType LocalGit { get; } = new ScmType("LocalGit");
        public static ScmType GitHub { get; } = new ScmType("GitHub");
        public static ScmType CodePlexGit { get; } = new ScmType("CodePlexGit");
        public static ScmType CodePlexHg { get; } = new ScmType("CodePlexHg");
        public static ScmType BitbucketGit { get; } = new ScmType("BitbucketGit");
        public static ScmType BitbucketHg { get; } = new ScmType("BitbucketHg");
        public static ScmType ExternalGit { get; } = new ScmType("ExternalGit");
        public static ScmType ExternalHg { get; } = new ScmType("ExternalHg");
        public static ScmType OneDrive { get; } = new ScmType("OneDrive");
        public static ScmType VSO { get; } = new ScmType("VSO");
        public static ScmType VSTSRM { get; } = new ScmType("VSTSRM");

        public static bool operator ==(ScmType left, ScmType right) => left.Equals(right);
        public static bool operator !=(ScmType left, ScmType right) => !left.Equals(right);

        public static explicit operator string(ScmType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScmType other && Equals(other);
        public bool Equals(ScmType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Site load balancing.
    /// </summary>
    [EnumType]
    public readonly struct SiteLoadBalancing : IEquatable<SiteLoadBalancing>
    {
        private readonly string _value;

        private SiteLoadBalancing(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SiteLoadBalancing WeightedRoundRobin { get; } = new SiteLoadBalancing("WeightedRoundRobin");
        public static SiteLoadBalancing LeastRequests { get; } = new SiteLoadBalancing("LeastRequests");
        public static SiteLoadBalancing LeastResponseTime { get; } = new SiteLoadBalancing("LeastResponseTime");
        public static SiteLoadBalancing WeightedTotalTraffic { get; } = new SiteLoadBalancing("WeightedTotalTraffic");
        public static SiteLoadBalancing RequestHash { get; } = new SiteLoadBalancing("RequestHash");

        public static bool operator ==(SiteLoadBalancing left, SiteLoadBalancing right) => left.Equals(right);
        public static bool operator !=(SiteLoadBalancing left, SiteLoadBalancing right) => !left.Equals(right);

        public static explicit operator string(SiteLoadBalancing value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SiteLoadBalancing other && Equals(other);
        public bool Equals(SiteLoadBalancing other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SSL type.
    /// </summary>
    [EnumType]
    public readonly struct SslState : IEquatable<SslState>
    {
        private readonly string _value;

        private SslState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SslState Disabled { get; } = new SslState("Disabled");
        public static SslState SniEnabled { get; } = new SslState("SniEnabled");
        public static SslState IpBasedEnabled { get; } = new SslState("IpBasedEnabled");

        public static bool operator ==(SslState left, SslState right) => left.Equals(right);
        public static bool operator !=(SslState left, SslState right) => !left.Equals(right);

        public static explicit operator string(SslState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SslState other && Equals(other);
        public bool Equals(SslState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// MinTlsVersion: configures the minimum version of TLS required for SSL requests
    /// </summary>
    [EnumType]
    public readonly struct SupportedTlsVersions : IEquatable<SupportedTlsVersions>
    {
        private readonly string _value;

        private SupportedTlsVersions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SupportedTlsVersions _0 { get; } = new SupportedTlsVersions("1.0");
        public static SupportedTlsVersions _1 { get; } = new SupportedTlsVersions("1.1");
        public static SupportedTlsVersions _2 { get; } = new SupportedTlsVersions("1.2");

        public static bool operator ==(SupportedTlsVersions left, SupportedTlsVersions right) => left.Equals(right);
        public static bool operator !=(SupportedTlsVersions left, SupportedTlsVersions right) => !left.Equals(right);

        public static explicit operator string(SupportedTlsVersions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SupportedTlsVersions other && Equals(other);
        public bool Equals(SupportedTlsVersions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
