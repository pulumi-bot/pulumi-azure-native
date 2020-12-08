// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ServiceFabric.V20200101Preview
{
    /// <summary>
    /// the reference to the load balancer probe used by the load balancing rule.
    /// </summary>
    [EnumType]
    public readonly struct ProbeProtocol : IEquatable<ProbeProtocol>
    {
        private readonly string _value;

        private ProbeProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProbeProtocol Tcp { get; } = new ProbeProtocol("tcp");
        public static ProbeProtocol Http { get; } = new ProbeProtocol("http");
        public static ProbeProtocol Https { get; } = new ProbeProtocol("https");

        public static bool operator ==(ProbeProtocol left, ProbeProtocol right) => left.Equals(right);
        public static bool operator !=(ProbeProtocol left, ProbeProtocol right) => !left.Equals(right);

        public static explicit operator string(ProbeProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProbeProtocol other && Equals(other);
        public bool Equals(ProbeProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The reference to the transport protocol used by the load balancing rule.
    /// </summary>
    [EnumType]
    public readonly struct Protocol : IEquatable<Protocol>
    {
        private readonly string _value;

        private Protocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Protocol Tcp { get; } = new Protocol("tcp");
        public static Protocol Udp { get; } = new Protocol("udp");

        public static bool operator ==(Protocol left, Protocol right) => left.Equals(right);
        public static bool operator !=(Protocol left, Protocol right) => !left.Equals(right);

        public static explicit operator string(Protocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Protocol other && Equals(other);
        public bool Equals(Protocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
