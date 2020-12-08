// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Insights.V20180617Preview
{
    /// <summary>
    /// The kind of workbook. Choices are user and shared.
    /// </summary>
    [EnumType]
    public readonly struct SharedTypeKind : IEquatable<SharedTypeKind>
    {
        private readonly string _value;

        private SharedTypeKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SharedTypeKind User { get; } = new SharedTypeKind("user");
        public static SharedTypeKind Shared { get; } = new SharedTypeKind("shared");

        public static bool operator ==(SharedTypeKind left, SharedTypeKind right) => left.Equals(right);
        public static bool operator !=(SharedTypeKind left, SharedTypeKind right) => !left.Equals(right);

        public static explicit operator string(SharedTypeKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SharedTypeKind other && Equals(other);
        public bool Equals(SharedTypeKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
