// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Management.V20180301
{
    /// <summary>
    /// The policy definition mode. Possible values are NotSpecified, Indexed, and All.
    /// </summary>
    [EnumType]
    public readonly struct PolicyMode : IEquatable<PolicyMode>
    {
        private readonly string _value;

        private PolicyMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyMode NotSpecified { get; } = new PolicyMode("NotSpecified");
        public static PolicyMode Indexed { get; } = new PolicyMode("Indexed");
        public static PolicyMode All { get; } = new PolicyMode("All");

        public static bool operator ==(PolicyMode left, PolicyMode right) => left.Equals(right);
        public static bool operator !=(PolicyMode left, PolicyMode right) => !left.Equals(right);

        public static explicit operator string(PolicyMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyMode other && Equals(other);
        public bool Equals(PolicyMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
    /// </summary>
    [EnumType]
    public readonly struct PolicyType : IEquatable<PolicyType>
    {
        private readonly string _value;

        private PolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyType NotSpecified { get; } = new PolicyType("NotSpecified");
        public static PolicyType BuiltIn { get; } = new PolicyType("BuiltIn");
        public static PolicyType Custom { get; } = new PolicyType("Custom");

        public static bool operator ==(PolicyType left, PolicyType right) => left.Equals(right);
        public static bool operator !=(PolicyType left, PolicyType right) => !left.Equals(right);

        public static explicit operator string(PolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyType other && Equals(other);
        public bool Equals(PolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
