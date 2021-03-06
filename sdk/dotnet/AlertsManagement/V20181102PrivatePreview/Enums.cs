// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.AlertsManagement.V20181102PrivatePreview
{
    /// <summary>
    /// Indicates if the given action rule is enabled or disabled
    /// </summary>
    [EnumType]
    public readonly struct ActionRuleStatus : IEquatable<ActionRuleStatus>
    {
        private readonly string _value;

        private ActionRuleStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ActionRuleStatus Enabled { get; } = new ActionRuleStatus("enabled");
        public static ActionRuleStatus Disabled { get; } = new ActionRuleStatus("disabled");

        public static bool operator ==(ActionRuleStatus left, ActionRuleStatus right) => left.Equals(right);
        public static bool operator !=(ActionRuleStatus left, ActionRuleStatus right) => !left.Equals(right);

        public static explicit operator string(ActionRuleStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ActionRuleStatus other && Equals(other);
        public bool Equals(ActionRuleStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// type of target scope
    /// </summary>
    [EnumType]
    public readonly struct ScopeType : IEquatable<ScopeType>
    {
        private readonly string _value;

        private ScopeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScopeType ResourceGroup { get; } = new ScopeType("ResourceGroup");
        public static ScopeType Resource { get; } = new ScopeType("Resource");

        public static bool operator ==(ScopeType left, ScopeType right) => left.Equals(right);
        public static bool operator !=(ScopeType left, ScopeType right) => !left.Equals(right);

        public static explicit operator string(ScopeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScopeType other && Equals(other);
        public bool Equals(ScopeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies when the suppression should be applied
    /// </summary>
    [EnumType]
    public readonly struct SuppressionType : IEquatable<SuppressionType>
    {
        private readonly string _value;

        private SuppressionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SuppressionType Always { get; } = new SuppressionType("Always");
        public static SuppressionType Once { get; } = new SuppressionType("Once");
        public static SuppressionType Daily { get; } = new SuppressionType("Daily");
        public static SuppressionType Weekly { get; } = new SuppressionType("Weekly");
        public static SuppressionType Monthly { get; } = new SuppressionType("Monthly");

        public static bool operator ==(SuppressionType left, SuppressionType right) => left.Equals(right);
        public static bool operator !=(SuppressionType left, SuppressionType right) => !left.Equals(right);

        public static explicit operator string(SuppressionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SuppressionType other && Equals(other);
        public bool Equals(SuppressionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
