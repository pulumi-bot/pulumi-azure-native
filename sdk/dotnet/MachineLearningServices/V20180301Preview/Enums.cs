// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.MachineLearningServices.V20180301Preview
{
    /// <summary>
    /// The type of compute
    /// </summary>
    [EnumType]
    public readonly struct ComputeType : IEquatable<ComputeType>
    {
        private readonly string _value;

        private ComputeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComputeType AKS { get; } = new ComputeType("AKS");
        public static ComputeType BatchAI { get; } = new ComputeType("BatchAI");
        public static ComputeType DataFactory { get; } = new ComputeType("DataFactory");
        public static ComputeType VirtualMachine { get; } = new ComputeType("VirtualMachine");
        public static ComputeType HDInsight { get; } = new ComputeType("HDInsight");

        public static bool operator ==(ComputeType left, ComputeType right) => left.Equals(right);
        public static bool operator !=(ComputeType left, ComputeType right) => !left.Equals(right);

        public static explicit operator string(ComputeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeType other && Equals(other);
        public bool Equals(ComputeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type.
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
