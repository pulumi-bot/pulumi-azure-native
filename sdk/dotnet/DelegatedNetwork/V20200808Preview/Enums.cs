// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.DelegatedNetwork.V20200808Preview
{
    /// <summary>
    /// The kind of workbook. Choices are user and shared.
    /// </summary>
    [EnumType]
    public readonly struct OrchestratorKind : IEquatable<OrchestratorKind>
    {
        private readonly string _value;

        private OrchestratorKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OrchestratorKind Kubernetes { get; } = new OrchestratorKind("Kubernetes");

        public static bool operator ==(OrchestratorKind left, OrchestratorKind right) => left.Equals(right);
        public static bool operator !=(OrchestratorKind left, OrchestratorKind right) => !left.Equals(right);

        public static explicit operator string(OrchestratorKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OrchestratorKind other && Equals(other);
        public bool Equals(OrchestratorKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of identity used for orchestrator cluster. Type 'SystemAssigned' will use an implicitly created identity orchestrator clusters
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");
        public static ResourceIdentityType None { get; } = new ResourceIdentityType("None");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
