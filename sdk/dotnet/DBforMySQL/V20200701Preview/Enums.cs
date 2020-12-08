// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.DBforMySQL.V20200701Preview
{
    /// <summary>
    /// The mode to create a new MySQL server.
    /// </summary>
    [EnumType]
    public readonly struct CreateMode : IEquatable<CreateMode>
    {
        private readonly string _value;

        private CreateMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CreateMode Default { get; } = new CreateMode("Default");
        public static CreateMode PointInTimeRestore { get; } = new CreateMode("PointInTimeRestore");
        public static CreateMode Replica { get; } = new CreateMode("Replica");

        public static bool operator ==(CreateMode left, CreateMode right) => left.Equals(right);
        public static bool operator !=(CreateMode left, CreateMode right) => !left.Equals(right);

        public static explicit operator string(CreateMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CreateMode other && Equals(other);
        public bool Equals(CreateMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enable HA or not for a server.
    /// </summary>
    [EnumType]
    public readonly struct HaEnabledEnum : IEquatable<HaEnabledEnum>
    {
        private readonly string _value;

        private HaEnabledEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HaEnabledEnum Enabled { get; } = new HaEnabledEnum("Enabled");
        public static HaEnabledEnum Disabled { get; } = new HaEnabledEnum("Disabled");

        public static bool operator ==(HaEnabledEnum left, HaEnabledEnum right) => left.Equals(right);
        public static bool operator !=(HaEnabledEnum left, HaEnabledEnum right) => !left.Equals(right);

        public static explicit operator string(HaEnabledEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HaEnabledEnum other && Equals(other);
        public bool Equals(HaEnabledEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Status showing whether the server enabled infrastructure encryption.
    /// </summary>
    [EnumType]
    public readonly struct InfrastructureEncryptionEnum : IEquatable<InfrastructureEncryptionEnum>
    {
        private readonly string _value;

        private InfrastructureEncryptionEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfrastructureEncryptionEnum Enabled { get; } = new InfrastructureEncryptionEnum("Enabled");
        public static InfrastructureEncryptionEnum Disabled { get; } = new InfrastructureEncryptionEnum("Disabled");

        public static bool operator ==(InfrastructureEncryptionEnum left, InfrastructureEncryptionEnum right) => left.Equals(right);
        public static bool operator !=(InfrastructureEncryptionEnum left, InfrastructureEncryptionEnum right) => !left.Equals(right);

        public static explicit operator string(InfrastructureEncryptionEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfrastructureEncryptionEnum other && Equals(other);
        public bool Equals(InfrastructureEncryptionEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type.
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The key type like 'AzureKeyVault'.
    /// </summary>
    [EnumType]
    public readonly struct ServerKeyType : IEquatable<ServerKeyType>
    {
        private readonly string _value;

        private ServerKeyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServerKeyType AzureKeyVault { get; } = new ServerKeyType("AzureKeyVault");

        public static bool operator ==(ServerKeyType left, ServerKeyType right) => left.Equals(right);
        public static bool operator !=(ServerKeyType left, ServerKeyType right) => !left.Equals(right);

        public static explicit operator string(ServerKeyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServerKeyType other && Equals(other);
        public bool Equals(ServerKeyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Server version.
    /// </summary>
    [EnumType]
    public readonly struct ServerVersion : IEquatable<ServerVersion>
    {
        private readonly string _value;

        private ServerVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServerVersion _7 { get; } = new ServerVersion("5.7");

        public static bool operator ==(ServerVersion left, ServerVersion right) => left.Equals(right);
        public static bool operator !=(ServerVersion left, ServerVersion right) => !left.Equals(right);

        public static explicit operator string(ServerVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServerVersion other && Equals(other);
        public bool Equals(ServerVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The tier of the particular SKU, e.g. GeneralPurpose.
    /// </summary>
    [EnumType]
    public readonly struct SkuTier : IEquatable<SkuTier>
    {
        private readonly string _value;

        private SkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuTier Burstable { get; } = new SkuTier("Burstable");
        public static SkuTier GeneralPurpose { get; } = new SkuTier("GeneralPurpose");
        public static SkuTier MemoryOptimized { get; } = new SkuTier("MemoryOptimized");

        public static bool operator ==(SkuTier left, SkuTier right) => left.Equals(right);
        public static bool operator !=(SkuTier left, SkuTier right) => !left.Equals(right);

        public static explicit operator string(SkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuTier other && Equals(other);
        public bool Equals(SkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enable ssl enforcement or not when connect to server.
    /// </summary>
    [EnumType]
    public readonly struct SslEnforcementEnum : IEquatable<SslEnforcementEnum>
    {
        private readonly string _value;

        private SslEnforcementEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SslEnforcementEnum Enabled { get; } = new SslEnforcementEnum("Enabled");
        public static SslEnforcementEnum Disabled { get; } = new SslEnforcementEnum("Disabled");

        public static bool operator ==(SslEnforcementEnum left, SslEnforcementEnum right) => left.Equals(right);
        public static bool operator !=(SslEnforcementEnum left, SslEnforcementEnum right) => !left.Equals(right);

        public static explicit operator string(SslEnforcementEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SslEnforcementEnum other && Equals(other);
        public bool Equals(SslEnforcementEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enable Storage Auto Grow.
    /// </summary>
    [EnumType]
    public readonly struct StorageAutogrow : IEquatable<StorageAutogrow>
    {
        private readonly string _value;

        private StorageAutogrow(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StorageAutogrow Enabled { get; } = new StorageAutogrow("Enabled");
        public static StorageAutogrow Disabled { get; } = new StorageAutogrow("Disabled");

        public static bool operator ==(StorageAutogrow left, StorageAutogrow right) => left.Equals(right);
        public static bool operator !=(StorageAutogrow left, StorageAutogrow right) => !left.Equals(right);

        public static explicit operator string(StorageAutogrow value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageAutogrow other && Equals(other);
        public bool Equals(StorageAutogrow other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
