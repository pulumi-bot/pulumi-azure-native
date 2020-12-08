// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.CostManagement.V20191001
{
    /// <summary>
    /// The type of the query.
    /// </summary>
    [EnumType]
    public readonly struct ExportType : IEquatable<ExportType>
    {
        private readonly string _value;

        private ExportType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ExportType Usage { get; } = new ExportType("Usage");

        public static bool operator ==(ExportType left, ExportType right) => left.Equals(right);
        public static bool operator !=(ExportType left, ExportType right) => !left.Equals(right);

        public static explicit operator string(ExportType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExportType other && Equals(other);
        public bool Equals(ExportType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The format of the export being delivered.
    /// </summary>
    [EnumType]
    public readonly struct FormatType : IEquatable<FormatType>
    {
        private readonly string _value;

        private FormatType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FormatType Csv { get; } = new FormatType("Csv");

        public static bool operator ==(FormatType left, FormatType right) => left.Equals(right);
        public static bool operator !=(FormatType left, FormatType right) => !left.Equals(right);

        public static explicit operator string(FormatType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FormatType other && Equals(other);
        public bool Equals(FormatType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the aggregation function to use.
    /// </summary>
    [EnumType]
    public readonly struct FunctionType : IEquatable<FunctionType>
    {
        private readonly string _value;

        private FunctionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FunctionType Sum { get; } = new FunctionType("Sum");

        public static bool operator ==(FunctionType left, FunctionType right) => left.Equals(right);
        public static bool operator !=(FunctionType left, FunctionType right) => !left.Equals(right);

        public static explicit operator string(FunctionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FunctionType other && Equals(other);
        public bool Equals(FunctionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The granularity of rows in the query.
    /// </summary>
    [EnumType]
    public readonly struct GranularityType : IEquatable<GranularityType>
    {
        private readonly string _value;

        private GranularityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GranularityType Daily { get; } = new GranularityType("Daily");
        public static GranularityType Hourly { get; } = new GranularityType("Hourly");

        public static bool operator ==(GranularityType left, GranularityType right) => left.Equals(right);
        public static bool operator !=(GranularityType left, GranularityType right) => !left.Equals(right);

        public static explicit operator string(GranularityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GranularityType other && Equals(other);
        public bool Equals(GranularityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The operator to use for comparison.
    /// </summary>
    [EnumType]
    public readonly struct OperatorType : IEquatable<OperatorType>
    {
        private readonly string _value;

        private OperatorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OperatorType In { get; } = new OperatorType("In");

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
    /// Has type of the column to group.
    /// </summary>
    [EnumType]
    public readonly struct QueryColumnType : IEquatable<QueryColumnType>
    {
        private readonly string _value;

        private QueryColumnType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static QueryColumnType Tag { get; } = new QueryColumnType("Tag");
        public static QueryColumnType Dimension { get; } = new QueryColumnType("Dimension");

        public static bool operator ==(QueryColumnType left, QueryColumnType right) => left.Equals(right);
        public static bool operator !=(QueryColumnType left, QueryColumnType right) => !left.Equals(right);

        public static explicit operator string(QueryColumnType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is QueryColumnType other && Equals(other);
        public bool Equals(QueryColumnType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The schedule recurrence.
    /// </summary>
    [EnumType]
    public readonly struct RecurrenceType : IEquatable<RecurrenceType>
    {
        private readonly string _value;

        private RecurrenceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RecurrenceType Daily { get; } = new RecurrenceType("Daily");
        public static RecurrenceType Weekly { get; } = new RecurrenceType("Weekly");
        public static RecurrenceType Monthly { get; } = new RecurrenceType("Monthly");
        public static RecurrenceType Annually { get; } = new RecurrenceType("Annually");

        public static bool operator ==(RecurrenceType left, RecurrenceType right) => left.Equals(right);
        public static bool operator !=(RecurrenceType left, RecurrenceType right) => !left.Equals(right);

        public static explicit operator string(RecurrenceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RecurrenceType other && Equals(other);
        public bool Equals(RecurrenceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The sorting direction
    /// </summary>
    [EnumType]
    public readonly struct SortDirection : IEquatable<SortDirection>
    {
        private readonly string _value;

        private SortDirection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SortDirection Ascending { get; } = new SortDirection("Ascending");
        public static SortDirection Descending { get; } = new SortDirection("Descending");

        public static bool operator ==(SortDirection left, SortDirection right) => left.Equals(right);
        public static bool operator !=(SortDirection left, SortDirection right) => !left.Equals(right);

        public static explicit operator string(SortDirection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SortDirection other && Equals(other);
        public bool Equals(SortDirection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the schedule. Whether active or not. If inactive, the export's scheduled execution is paused.
    /// </summary>
    [EnumType]
    public readonly struct StatusType : IEquatable<StatusType>
    {
        private readonly string _value;

        private StatusType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StatusType Active { get; } = new StatusType("Active");
        public static StatusType Inactive { get; } = new StatusType("Inactive");

        public static bool operator ==(StatusType left, StatusType right) => left.Equals(right);
        public static bool operator !=(StatusType left, StatusType right) => !left.Equals(right);

        public static explicit operator string(StatusType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StatusType other && Equals(other);
        public bool Equals(StatusType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The time frame for pulling data for the query. If custom, then a specific time period must be provided.
    /// </summary>
    [EnumType]
    public readonly struct TimeframeType : IEquatable<TimeframeType>
    {
        private readonly string _value;

        private TimeframeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TimeframeType WeekToDate { get; } = new TimeframeType("WeekToDate");
        public static TimeframeType MonthToDate { get; } = new TimeframeType("MonthToDate");
        public static TimeframeType YearToDate { get; } = new TimeframeType("YearToDate");
        public static TimeframeType TheLastWeek { get; } = new TimeframeType("TheLastWeek");
        public static TimeframeType TheLastMonth { get; } = new TimeframeType("TheLastMonth");
        public static TimeframeType TheLastYear { get; } = new TimeframeType("TheLastYear");
        public static TimeframeType Custom { get; } = new TimeframeType("Custom");

        public static bool operator ==(TimeframeType left, TimeframeType right) => left.Equals(right);
        public static bool operator !=(TimeframeType left, TimeframeType right) => !left.Equals(right);

        public static explicit operator string(TimeframeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TimeframeType other && Equals(other);
        public bool Equals(TimeframeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
