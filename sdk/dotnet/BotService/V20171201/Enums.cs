// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.BotService.V20171201
{
    /// <summary>
    /// Required. Gets or sets the Kind of the resource.
    /// </summary>
    [EnumType]
    public readonly struct Kind : IEquatable<Kind>
    {
        private readonly string _value;

        private Kind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Kind Sdk { get; } = new Kind("sdk");
        public static Kind Designer { get; } = new Kind("designer");
        public static Kind Bot { get; } = new Kind("bot");
        public static Kind Function { get; } = new Kind("function");

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
    /// The sku name
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName F0 { get; } = new SkuName("F0");
        public static SkuName S1 { get; } = new SkuName("S1");

        public static bool operator ==(SkuName left, SkuName right) => left.Equals(right);
        public static bool operator !=(SkuName left, SkuName right) => !left.Equals(right);

        public static explicit operator string(SkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuName other && Equals(other);
        public bool Equals(SkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
