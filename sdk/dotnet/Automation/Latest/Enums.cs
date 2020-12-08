// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Automation.Latest
{
    /// <summary>
    /// Gets or sets the content source type.
    /// </summary>
    [EnumType]
    public readonly struct ContentSourceType : IEquatable<ContentSourceType>
    {
        private readonly string _value;

        private ContentSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContentSourceType EmbeddedContent { get; } = new ContentSourceType("embeddedContent");
        public static ContentSourceType Uri { get; } = new ContentSourceType("uri");

        public static bool operator ==(ContentSourceType left, ContentSourceType right) => left.Equals(right);
        public static bool operator !=(ContentSourceType left, ContentSourceType right) => !left.Equals(right);

        public static explicit operator string(ContentSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContentSourceType other && Equals(other);
        public bool Equals(ContentSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the type of the runbook.
    /// </summary>
    [EnumType]
    public readonly struct RunbookTypeEnum : IEquatable<RunbookTypeEnum>
    {
        private readonly string _value;

        private RunbookTypeEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RunbookTypeEnum Script { get; } = new RunbookTypeEnum("Script");
        public static RunbookTypeEnum Graph { get; } = new RunbookTypeEnum("Graph");
        public static RunbookTypeEnum PowerShellWorkflow { get; } = new RunbookTypeEnum("PowerShellWorkflow");
        public static RunbookTypeEnum PowerShell { get; } = new RunbookTypeEnum("PowerShell");
        public static RunbookTypeEnum GraphPowerShellWorkflow { get; } = new RunbookTypeEnum("GraphPowerShellWorkflow");
        public static RunbookTypeEnum GraphPowerShell { get; } = new RunbookTypeEnum("GraphPowerShell");

        public static bool operator ==(RunbookTypeEnum left, RunbookTypeEnum right) => left.Equals(right);
        public static bool operator !=(RunbookTypeEnum left, RunbookTypeEnum right) => !left.Equals(right);

        public static explicit operator string(RunbookTypeEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RunbookTypeEnum other && Equals(other);
        public bool Equals(RunbookTypeEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Day of the occurrence. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleDay : IEquatable<ScheduleDay>
    {
        private readonly string _value;

        private ScheduleDay(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleDay Monday { get; } = new ScheduleDay("Monday");
        public static ScheduleDay Tuesday { get; } = new ScheduleDay("Tuesday");
        public static ScheduleDay Wednesday { get; } = new ScheduleDay("Wednesday");
        public static ScheduleDay Thursday { get; } = new ScheduleDay("Thursday");
        public static ScheduleDay Friday { get; } = new ScheduleDay("Friday");
        public static ScheduleDay Saturday { get; } = new ScheduleDay("Saturday");
        public static ScheduleDay Sunday { get; } = new ScheduleDay("Sunday");

        public static bool operator ==(ScheduleDay left, ScheduleDay right) => left.Equals(right);
        public static bool operator !=(ScheduleDay left, ScheduleDay right) => !left.Equals(right);

        public static explicit operator string(ScheduleDay value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleDay other && Equals(other);
        public bool Equals(ScheduleDay other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the frequency of the schedule.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleFrequency : IEquatable<ScheduleFrequency>
    {
        private readonly string _value;

        private ScheduleFrequency(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleFrequency OneTime { get; } = new ScheduleFrequency("OneTime");
        public static ScheduleFrequency Day { get; } = new ScheduleFrequency("Day");
        public static ScheduleFrequency Hour { get; } = new ScheduleFrequency("Hour");
        public static ScheduleFrequency Week { get; } = new ScheduleFrequency("Week");
        public static ScheduleFrequency Month { get; } = new ScheduleFrequency("Month");
        /// <summary>
        /// The minimum allowed interval for Minute schedules is 15 minutes.
        /// </summary>
        public static ScheduleFrequency Minute { get; } = new ScheduleFrequency("Minute");

        public static bool operator ==(ScheduleFrequency left, ScheduleFrequency right) => left.Equals(right);
        public static bool operator !=(ScheduleFrequency left, ScheduleFrequency right) => !left.Equals(right);

        public static explicit operator string(ScheduleFrequency value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleFrequency other && Equals(other);
        public bool Equals(ScheduleFrequency other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the SKU name of the account.
    /// </summary>
    [EnumType]
    public readonly struct SkuNameEnum : IEquatable<SkuNameEnum>
    {
        private readonly string _value;

        private SkuNameEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuNameEnum Free { get; } = new SkuNameEnum("Free");
        public static SkuNameEnum Basic { get; } = new SkuNameEnum("Basic");

        public static bool operator ==(SkuNameEnum left, SkuNameEnum right) => left.Equals(right);
        public static bool operator !=(SkuNameEnum left, SkuNameEnum right) => !left.Equals(right);

        public static explicit operator string(SkuNameEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuNameEnum other && Equals(other);
        public bool Equals(SkuNameEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
