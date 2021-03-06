// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.AVS.V20200320
{
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
}
