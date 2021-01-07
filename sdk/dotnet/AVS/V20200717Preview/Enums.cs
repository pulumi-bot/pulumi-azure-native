// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.AVS.V20200717Preview
{
    /// <summary>
    /// The type of private cloud addon
    /// </summary>
    [EnumType]
    public readonly struct AddonType : IEquatable<AddonType>
    {
        private readonly string _value;

        private AddonType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AddonType SRM { get; } = new AddonType("SRM");
        public static AddonType VR { get; } = new AddonType("VR");

        public static bool operator ==(AddonType left, AddonType right) => left.Equals(right);
        public static bool operator !=(AddonType left, AddonType right) => !left.Equals(right);

        public static explicit operator string(AddonType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddonType other && Equals(other);
        public bool Equals(AddonType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of DHCP: SERVER or RELAY.
    /// </summary>
    [EnumType]
    public readonly struct DhcpTypeEnum : IEquatable<DhcpTypeEnum>
    {
        private readonly string _value;

        private DhcpTypeEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DhcpTypeEnum SERVER_RELAY { get; } = new DhcpTypeEnum("SERVER, RELAY");

        public static bool operator ==(DhcpTypeEnum left, DhcpTypeEnum right) => left.Equals(right);
        public static bool operator !=(DhcpTypeEnum left, DhcpTypeEnum right) => !left.Equals(right);

        public static explicit operator string(DhcpTypeEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DhcpTypeEnum other && Equals(other);
        public bool Equals(DhcpTypeEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// DNS Service log level.
    /// </summary>
    [EnumType]
    public readonly struct DnsServiceLogLevelEnum : IEquatable<DnsServiceLogLevelEnum>
    {
        private readonly string _value;

        private DnsServiceLogLevelEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DnsServiceLogLevelEnum DEBUG { get; } = new DnsServiceLogLevelEnum("DEBUG");
        public static DnsServiceLogLevelEnum INFO { get; } = new DnsServiceLogLevelEnum("INFO");
        public static DnsServiceLogLevelEnum WARNING { get; } = new DnsServiceLogLevelEnum("WARNING");
        public static DnsServiceLogLevelEnum ERROR { get; } = new DnsServiceLogLevelEnum("ERROR");
        public static DnsServiceLogLevelEnum FATAL { get; } = new DnsServiceLogLevelEnum("FATAL");

        public static bool operator ==(DnsServiceLogLevelEnum left, DnsServiceLogLevelEnum right) => left.Equals(right);
        public static bool operator !=(DnsServiceLogLevelEnum left, DnsServiceLogLevelEnum right) => !left.Equals(right);

        public static explicit operator string(DnsServiceLogLevelEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DnsServiceLogLevelEnum other && Equals(other);
        public bool Equals(DnsServiceLogLevelEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connectivity to internet is enabled or disabled
    /// </summary>
    [EnumType]
    public readonly struct InternetEnum : IEquatable<InternetEnum>
    {
        private readonly string _value;

        private InternetEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InternetEnum Enabled { get; } = new InternetEnum("Enabled");
        public static InternetEnum Disabled { get; } = new InternetEnum("Disabled");

        public static bool operator ==(InternetEnum left, InternetEnum right) => left.Equals(right);
        public static bool operator !=(InternetEnum left, InternetEnum right) => !left.Equals(right);

        public static explicit operator string(InternetEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InternetEnum other && Equals(other);
        public bool Equals(InternetEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicate to rotate the NSX-T Manager password for the private cloud
    /// </summary>
    [EnumType]
    public readonly struct NsxtAdminRotateEnum : IEquatable<NsxtAdminRotateEnum>
    {
        private readonly string _value;

        private NsxtAdminRotateEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NsxtAdminRotateEnum OnetimeRotate { get; } = new NsxtAdminRotateEnum("OnetimeRotate");

        public static bool operator ==(NsxtAdminRotateEnum left, NsxtAdminRotateEnum right) => left.Equals(right);
        public static bool operator !=(NsxtAdminRotateEnum left, NsxtAdminRotateEnum right) => !left.Equals(right);

        public static explicit operator string(NsxtAdminRotateEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NsxtAdminRotateEnum other && Equals(other);
        public bool Equals(NsxtAdminRotateEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Direction of port mirroring profile.
    /// </summary>
    [EnumType]
    public readonly struct PortMirroringDirectionEnum : IEquatable<PortMirroringDirectionEnum>
    {
        private readonly string _value;

        private PortMirroringDirectionEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PortMirroringDirectionEnum INGRESS_EGRESS_BIDIRECTIONAL { get; } = new PortMirroringDirectionEnum("INGRESS, EGRESS, BIDIRECTIONAL");

        public static bool operator ==(PortMirroringDirectionEnum left, PortMirroringDirectionEnum right) => left.Equals(right);
        public static bool operator !=(PortMirroringDirectionEnum left, PortMirroringDirectionEnum right) => !left.Equals(right);

        public static explicit operator string(PortMirroringDirectionEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PortMirroringDirectionEnum other && Equals(other);
        public bool Equals(PortMirroringDirectionEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Protect LDAP communication using SSL certificate (LDAPS)
    /// </summary>
    [EnumType]
    public readonly struct SslEnum : IEquatable<SslEnum>
    {
        private readonly string _value;

        private SslEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SslEnum Enabled { get; } = new SslEnum("Enabled");
        public static SslEnum Disabled { get; } = new SslEnum("Disabled");

        public static bool operator ==(SslEnum left, SslEnum right) => left.Equals(right);
        public static bool operator !=(SslEnum left, SslEnum right) => !left.Equals(right);

        public static explicit operator string(SslEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SslEnum other && Equals(other);
        public bool Equals(SslEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicate to rotate the vCenter admin password for the private cloud
    /// </summary>
    [EnumType]
    public readonly struct VcsaAdminRotateEnum : IEquatable<VcsaAdminRotateEnum>
    {
        private readonly string _value;

        private VcsaAdminRotateEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VcsaAdminRotateEnum OnetimeRotate { get; } = new VcsaAdminRotateEnum("OnetimeRotate");

        public static bool operator ==(VcsaAdminRotateEnum left, VcsaAdminRotateEnum right) => left.Equals(right);
        public static bool operator !=(VcsaAdminRotateEnum left, VcsaAdminRotateEnum right) => !left.Equals(right);

        public static explicit operator string(VcsaAdminRotateEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VcsaAdminRotateEnum other && Equals(other);
        public bool Equals(VcsaAdminRotateEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
