// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Logic.V20150201Preview
{
    /// <summary>
    /// Gets or sets the type.
    /// </summary>
    [EnumType]
    public readonly struct ParameterType : IEquatable<ParameterType>
    {
        private readonly string _value;

        private ParameterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ParameterType NotSpecified { get; } = new ParameterType("NotSpecified");
        public static ParameterType String { get; } = new ParameterType("String");
        public static ParameterType SecureString { get; } = new ParameterType("SecureString");
        public static ParameterType Int { get; } = new ParameterType("Int");
        public static ParameterType Float { get; } = new ParameterType("Float");
        public static ParameterType Bool { get; } = new ParameterType("Bool");
        public static ParameterType Array { get; } = new ParameterType("Array");
        public static ParameterType Object { get; } = new ParameterType("Object");
        public static ParameterType SecureObject { get; } = new ParameterType("SecureObject");

        public static bool operator ==(ParameterType left, ParameterType right) => left.Equals(right);
        public static bool operator !=(ParameterType left, ParameterType right) => !left.Equals(right);

        public static explicit operator string(ParameterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ParameterType other && Equals(other);
        public bool Equals(ParameterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the name.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName NotSpecified { get; } = new SkuName("NotSpecified");
        public static SkuName Free { get; } = new SkuName("Free");
        public static SkuName Shared { get; } = new SkuName("Shared");
        public static SkuName Basic { get; } = new SkuName("Basic");
        public static SkuName Standard { get; } = new SkuName("Standard");
        public static SkuName Premium { get; } = new SkuName("Premium");

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

    /// <summary>
    /// Gets or sets the state.
    /// </summary>
    [EnumType]
    public readonly struct WorkflowState : IEquatable<WorkflowState>
    {
        private readonly string _value;

        private WorkflowState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkflowState NotSpecified { get; } = new WorkflowState("NotSpecified");
        public static WorkflowState Enabled { get; } = new WorkflowState("Enabled");
        public static WorkflowState Disabled { get; } = new WorkflowState("Disabled");
        public static WorkflowState Deleted { get; } = new WorkflowState("Deleted");
        public static WorkflowState Suspended { get; } = new WorkflowState("Suspended");

        public static bool operator ==(WorkflowState left, WorkflowState right) => left.Equals(right);
        public static bool operator !=(WorkflowState left, WorkflowState right) => !left.Equals(right);

        public static explicit operator string(WorkflowState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkflowState other && Equals(other);
        public bool Equals(WorkflowState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
