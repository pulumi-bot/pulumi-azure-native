// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.DigitalTwins.V20201031
{
    /// <summary>
    /// The type of Digital Twins endpoint
    /// </summary>
    [EnumType]
    public readonly struct EndpointType : IEquatable<EndpointType>
    {
        private readonly string _value;

        private EndpointType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EndpointType EventHub { get; } = new EndpointType("EventHub");
        public static EndpointType EventGrid { get; } = new EndpointType("EventGrid");
        public static EndpointType ServiceBus { get; } = new EndpointType("ServiceBus");

        public static bool operator ==(EndpointType left, EndpointType right) => left.Equals(right);
        public static bool operator !=(EndpointType left, EndpointType right) => !left.Equals(right);

        public static explicit operator string(EndpointType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointType other && Equals(other);
        public bool Equals(EndpointType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
