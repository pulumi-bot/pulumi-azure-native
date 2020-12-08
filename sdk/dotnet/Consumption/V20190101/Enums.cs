// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Consumption.V20190101
{
    /// <summary>
    /// The category of the budget, whether the budget tracks cost or usage.
    /// </summary>
    [EnumType]
    public readonly struct CategoryType : IEquatable<CategoryType>
    {
        private readonly string _value;

        private CategoryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CategoryType Cost { get; } = new CategoryType("Cost");
        public static CategoryType Usage { get; } = new CategoryType("Usage");

        public static bool operator ==(CategoryType left, CategoryType right) => left.Equals(right);
        public static bool operator !=(CategoryType left, CategoryType right) => !left.Equals(right);

        public static explicit operator string(CategoryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CategoryType other && Equals(other);
        public bool Equals(CategoryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The comparison operator.
    /// </summary>
    [EnumType]
    public readonly struct OperatorType : IEquatable<OperatorType>
    {
        private readonly string _value;

        private OperatorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OperatorType EqualTo { get; } = new OperatorType("EqualTo");
        public static OperatorType GreaterThan { get; } = new OperatorType("GreaterThan");
        public static OperatorType GreaterThanOrEqualTo { get; } = new OperatorType("GreaterThanOrEqualTo");

        public static bool operator ==(OperatorType left, OperatorType right) => left.Equals(right);
        public static bool operator !=(OperatorType left, OperatorType right) => !left.Equals(right);

        public static explicit operator string(OperatorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OperatorType other && Equals(other);
        public bool Equals(OperatorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
    /// </summary>
    [EnumType]
    public readonly struct TimeGrainType : IEquatable<TimeGrainType>
    {
        private readonly string _value;

        private TimeGrainType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TimeGrainType Monthly { get; } = new TimeGrainType("Monthly");
        public static TimeGrainType Quarterly { get; } = new TimeGrainType("Quarterly");
        public static TimeGrainType Annually { get; } = new TimeGrainType("Annually");
        public static TimeGrainType BillingMonth { get; } = new TimeGrainType("BillingMonth");
        public static TimeGrainType BillingQuarter { get; } = new TimeGrainType("BillingQuarter");
        public static TimeGrainType BillingAnnual { get; } = new TimeGrainType("BillingAnnual");

        public static bool operator ==(TimeGrainType left, TimeGrainType right) => left.Equals(right);
        public static bool operator !=(TimeGrainType left, TimeGrainType right) => !left.Equals(right);

        public static explicit operator string(TimeGrainType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TimeGrainType other && Equals(other);
        public bool Equals(TimeGrainType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
