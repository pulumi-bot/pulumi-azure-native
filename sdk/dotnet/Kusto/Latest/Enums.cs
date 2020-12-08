// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Kusto.Latest
{
    /// <summary>
    /// SKU name.
    /// </summary>
    [EnumType]
    public readonly struct AzureSkuName : IEquatable<AzureSkuName>
    {
        private readonly string _value;

        private AzureSkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AzureSkuName Standard_DS13_v2_1TB_PS { get; } = new AzureSkuName("Standard_DS13_v2+1TB_PS");
        public static AzureSkuName Standard_DS13_v2_2TB_PS { get; } = new AzureSkuName("Standard_DS13_v2+2TB_PS");
        public static AzureSkuName Standard_DS14_v2_3TB_PS { get; } = new AzureSkuName("Standard_DS14_v2+3TB_PS");
        public static AzureSkuName Standard_DS14_v2_4TB_PS { get; } = new AzureSkuName("Standard_DS14_v2+4TB_PS");
        public static AzureSkuName Standard_D13_v2 { get; } = new AzureSkuName("Standard_D13_v2");
        public static AzureSkuName Standard_D14_v2 { get; } = new AzureSkuName("Standard_D14_v2");
        public static AzureSkuName Standard_L8s { get; } = new AzureSkuName("Standard_L8s");
        public static AzureSkuName Standard_L16s { get; } = new AzureSkuName("Standard_L16s");
        public static AzureSkuName Standard_D11_v2 { get; } = new AzureSkuName("Standard_D11_v2");
        public static AzureSkuName Standard_D12_v2 { get; } = new AzureSkuName("Standard_D12_v2");
        public static AzureSkuName Standard_L4s { get; } = new AzureSkuName("Standard_L4s");
        public static AzureSkuName Dev_No_SLA_Standard_D11_v2 { get; } = new AzureSkuName("Dev(No SLA)_Standard_D11_v2");
        public static AzureSkuName Standard_E64i_v3 { get; } = new AzureSkuName("Standard_E64i_v3");
        public static AzureSkuName Standard_E2a_v4 { get; } = new AzureSkuName("Standard_E2a_v4");
        public static AzureSkuName Standard_E4a_v4 { get; } = new AzureSkuName("Standard_E4a_v4");
        public static AzureSkuName Standard_E8a_v4 { get; } = new AzureSkuName("Standard_E8a_v4");
        public static AzureSkuName Standard_E16a_v4 { get; } = new AzureSkuName("Standard_E16a_v4");
        public static AzureSkuName Standard_E8as_v4_1TB_PS { get; } = new AzureSkuName("Standard_E8as_v4+1TB_PS");
        public static AzureSkuName Standard_E8as_v4_2TB_PS { get; } = new AzureSkuName("Standard_E8as_v4+2TB_PS");
        public static AzureSkuName Standard_E16as_v4_3TB_PS { get; } = new AzureSkuName("Standard_E16as_v4+3TB_PS");
        public static AzureSkuName Standard_E16as_v4_4TB_PS { get; } = new AzureSkuName("Standard_E16as_v4+4TB_PS");
        public static AzureSkuName Dev_No_SLA_Standard_E2a_v4 { get; } = new AzureSkuName("Dev(No SLA)_Standard_E2a_v4");

        public static bool operator ==(AzureSkuName left, AzureSkuName right) => left.Equals(right);
        public static bool operator !=(AzureSkuName left, AzureSkuName right) => !left.Equals(right);

        public static explicit operator string(AzureSkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AzureSkuName other && Equals(other);
        public bool Equals(AzureSkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SKU tier.
    /// </summary>
    [EnumType]
    public readonly struct AzureSkuTier : IEquatable<AzureSkuTier>
    {
        private readonly string _value;

        private AzureSkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AzureSkuTier Basic { get; } = new AzureSkuTier("Basic");
        public static AzureSkuTier Standard { get; } = new AzureSkuTier("Standard");

        public static bool operator ==(AzureSkuTier left, AzureSkuTier right) => left.Equals(right);
        public static bool operator !=(AzureSkuTier left, AzureSkuTier right) => !left.Equals(right);

        public static explicit operator string(AzureSkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AzureSkuTier other && Equals(other);
        public bool Equals(AzureSkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Cluster principal role.
    /// </summary>
    [EnumType]
    public readonly struct ClusterPrincipalRole : IEquatable<ClusterPrincipalRole>
    {
        private readonly string _value;

        private ClusterPrincipalRole(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterPrincipalRole AllDatabasesAdmin { get; } = new ClusterPrincipalRole("AllDatabasesAdmin");
        public static ClusterPrincipalRole AllDatabasesViewer { get; } = new ClusterPrincipalRole("AllDatabasesViewer");

        public static bool operator ==(ClusterPrincipalRole left, ClusterPrincipalRole right) => left.Equals(right);
        public static bool operator !=(ClusterPrincipalRole left, ClusterPrincipalRole right) => !left.Equals(right);

        public static explicit operator string(ClusterPrincipalRole value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterPrincipalRole other && Equals(other);
        public bool Equals(ClusterPrincipalRole other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Kind of the endpoint for the data connection
    /// </summary>
    [EnumType]
    public readonly struct DataConnectionKind : IEquatable<DataConnectionKind>
    {
        private readonly string _value;

        private DataConnectionKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataConnectionKind EventHub { get; } = new DataConnectionKind("EventHub");
        public static DataConnectionKind EventGrid { get; } = new DataConnectionKind("EventGrid");
        public static DataConnectionKind IotHub { get; } = new DataConnectionKind("IotHub");

        public static bool operator ==(DataConnectionKind left, DataConnectionKind right) => left.Equals(right);
        public static bool operator !=(DataConnectionKind left, DataConnectionKind right) => !left.Equals(right);

        public static explicit operator string(DataConnectionKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataConnectionKind other && Equals(other);
        public bool Equals(DataConnectionKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Database principal role.
    /// </summary>
    [EnumType]
    public readonly struct DatabasePrincipalRole : IEquatable<DatabasePrincipalRole>
    {
        private readonly string _value;

        private DatabasePrincipalRole(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatabasePrincipalRole Admin { get; } = new DatabasePrincipalRole("Admin");
        public static DatabasePrincipalRole Ingestor { get; } = new DatabasePrincipalRole("Ingestor");
        public static DatabasePrincipalRole Monitor { get; } = new DatabasePrincipalRole("Monitor");
        public static DatabasePrincipalRole User { get; } = new DatabasePrincipalRole("User");
        public static DatabasePrincipalRole UnrestrictedViewers { get; } = new DatabasePrincipalRole("UnrestrictedViewers");
        public static DatabasePrincipalRole Viewer { get; } = new DatabasePrincipalRole("Viewer");

        public static bool operator ==(DatabasePrincipalRole left, DatabasePrincipalRole right) => left.Equals(right);
        public static bool operator !=(DatabasePrincipalRole left, DatabasePrincipalRole right) => !left.Equals(right);

        public static explicit operator string(DatabasePrincipalRole value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabasePrincipalRole other && Equals(other);
        public bool Equals(DatabasePrincipalRole other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default principals modification kind
    /// </summary>
    [EnumType]
    public readonly struct DefaultPrincipalsModificationKind : IEquatable<DefaultPrincipalsModificationKind>
    {
        private readonly string _value;

        private DefaultPrincipalsModificationKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DefaultPrincipalsModificationKind Union { get; } = new DefaultPrincipalsModificationKind("Union");
        public static DefaultPrincipalsModificationKind Replace { get; } = new DefaultPrincipalsModificationKind("Replace");
        public static DefaultPrincipalsModificationKind None { get; } = new DefaultPrincipalsModificationKind("None");

        public static bool operator ==(DefaultPrincipalsModificationKind left, DefaultPrincipalsModificationKind right) => left.Equals(right);
        public static bool operator !=(DefaultPrincipalsModificationKind left, DefaultPrincipalsModificationKind right) => !left.Equals(right);

        public static explicit operator string(DefaultPrincipalsModificationKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DefaultPrincipalsModificationKind other && Equals(other);
        public bool Equals(DefaultPrincipalsModificationKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The engine type
    /// </summary>
    [EnumType]
    public readonly struct EngineType : IEquatable<EngineType>
    {
        private readonly string _value;

        private EngineType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EngineType V2 { get; } = new EngineType("V2");
        public static EngineType V3 { get; } = new EngineType("V3");

        public static bool operator ==(EngineType left, EngineType right) => left.Equals(right);
        public static bool operator !=(EngineType left, EngineType right) => !left.Equals(right);

        public static explicit operator string(EngineType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EngineType other && Equals(other);
        public bool Equals(EngineType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove all identities.
    /// </summary>
    [EnumType]
    public readonly struct IdentityType : IEquatable<IdentityType>
    {
        private readonly string _value;

        private IdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityType None { get; } = new IdentityType("None");
        public static IdentityType SystemAssigned { get; } = new IdentityType("SystemAssigned");
        public static IdentityType UserAssigned { get; } = new IdentityType("UserAssigned");
        public static IdentityType SystemAssigned_UserAssigned { get; } = new IdentityType("SystemAssigned, UserAssigned");

        public static bool operator ==(IdentityType left, IdentityType right) => left.Equals(right);
        public static bool operator !=(IdentityType left, IdentityType right) => !left.Equals(right);

        public static explicit operator string(IdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityType other && Equals(other);
        public bool Equals(IdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Kind of the database
    /// </summary>
    [EnumType]
    public readonly struct Kind : IEquatable<Kind>
    {
        private readonly string _value;

        private Kind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Kind ReadWrite { get; } = new Kind("ReadWrite");
        public static Kind ReadOnlyFollowing { get; } = new Kind("ReadOnlyFollowing");

        public static bool operator ==(Kind left, Kind right) => left.Equals(right);
        public static bool operator !=(Kind left, Kind right) => !left.Equals(right);

        public static explicit operator string(Kind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Kind other && Equals(other);
        public bool Equals(Kind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Principal type.
    /// </summary>
    [EnumType]
    public readonly struct PrincipalType : IEquatable<PrincipalType>
    {
        private readonly string _value;

        private PrincipalType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrincipalType App { get; } = new PrincipalType("App");
        public static PrincipalType Group { get; } = new PrincipalType("Group");
        public static PrincipalType User { get; } = new PrincipalType("User");

        public static bool operator ==(PrincipalType left, PrincipalType right) => left.Equals(right);
        public static bool operator !=(PrincipalType left, PrincipalType right) => !left.Equals(right);

        public static explicit operator string(PrincipalType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrincipalType other && Equals(other);
        public bool Equals(PrincipalType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
