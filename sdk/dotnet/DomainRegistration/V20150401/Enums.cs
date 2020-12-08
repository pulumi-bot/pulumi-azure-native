// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.DomainRegistration.V20150401
{
    /// <summary>
    /// Target DNS type (would be used for migration)
    /// </summary>
    [EnumType]
    public readonly struct DnsType : IEquatable<DnsType>
    {
        private readonly string _value;

        private DnsType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DnsType AzureDns { get; } = new DnsType("AzureDns");
        public static DnsType DefaultDomainRegistrarDns { get; } = new DnsType("DefaultDomainRegistrarDns");

        public static bool operator ==(DnsType left, DnsType right) => left.Equals(right);
        public static bool operator !=(DnsType left, DnsType right) => !left.Equals(right);

        public static explicit operator string(DnsType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DnsType other && Equals(other);
        public bool Equals(DnsType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
