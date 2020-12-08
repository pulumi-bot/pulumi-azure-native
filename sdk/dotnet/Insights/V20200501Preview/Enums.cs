// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Insights.V20200501Preview
{
    /// <summary>
    /// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
    /// </summary>
    [EnumType]
    public readonly struct AlertSeverity : IEquatable<AlertSeverity>
    {
        private readonly string _value;

        private AlertSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertSeverity Zero { get; } = new AlertSeverity("0");
        public static AlertSeverity One { get; } = new AlertSeverity("1");
        public static AlertSeverity Two { get; } = new AlertSeverity("2");
        public static AlertSeverity Three { get; } = new AlertSeverity("3");
        public static AlertSeverity Four { get; } = new AlertSeverity("4");

        public static bool operator ==(AlertSeverity left, AlertSeverity right) => left.Equals(right);
        public static bool operator !=(AlertSeverity left, AlertSeverity right) => !left.Equals(right);

        public static explicit operator string(AlertSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertSeverity other && Equals(other);
        public bool Equals(AlertSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The criteria operator.
    /// </summary>
    [EnumType]
    public readonly struct ConditionOperator : IEquatable<ConditionOperator>
    {
        private readonly string _value;

        private ConditionOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConditionOperator EqualsValue { get; } = new ConditionOperator("Equals");
        public static ConditionOperator GreaterThan { get; } = new ConditionOperator("GreaterThan");
        public static ConditionOperator GreaterThanOrEqual { get; } = new ConditionOperator("GreaterThanOrEqual");
        public static ConditionOperator LessThan { get; } = new ConditionOperator("LessThan");
        public static ConditionOperator LessThanOrEqual { get; } = new ConditionOperator("LessThanOrEqual");

        public static bool operator ==(ConditionOperator left, ConditionOperator right) => left.Equals(right);
        public static bool operator !=(ConditionOperator left, ConditionOperator right) => !left.Equals(right);

        public static explicit operator string(ConditionOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConditionOperator other && Equals(other);
        public bool Equals(ConditionOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Operator for dimension values
    /// </summary>
    [EnumType]
    public readonly struct DimensionOperator : IEquatable<DimensionOperator>
    {
        private readonly string _value;

        private DimensionOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DimensionOperator Include { get; } = new DimensionOperator("Include");
        public static DimensionOperator Exclude { get; } = new DimensionOperator("Exclude");

        public static bool operator ==(DimensionOperator left, DimensionOperator right) => left.Equals(right);
        public static bool operator !=(DimensionOperator left, DimensionOperator right) => !left.Equals(right);

        public static explicit operator string(DimensionOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DimensionOperator other && Equals(other);
        public bool Equals(DimensionOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Aggregation type
    /// </summary>
    [EnumType]
    public readonly struct TimeAggregation : IEquatable<TimeAggregation>
    {
        private readonly string _value;

        private TimeAggregation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TimeAggregation Count { get; } = new TimeAggregation("Count");
        public static TimeAggregation Average { get; } = new TimeAggregation("Average");
        public static TimeAggregation Minimum { get; } = new TimeAggregation("Minimum");
        public static TimeAggregation Maximum { get; } = new TimeAggregation("Maximum");
        public static TimeAggregation Total { get; } = new TimeAggregation("Total");

        public static bool operator ==(TimeAggregation left, TimeAggregation right) => left.Equals(right);
        public static bool operator !=(TimeAggregation left, TimeAggregation right) => !left.Equals(right);

        public static explicit operator string(TimeAggregation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TimeAggregation other && Equals(other);
        public bool Equals(TimeAggregation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
