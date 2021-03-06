// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Solutions.V20160901Preview
{
    /// <summary>
    /// The appliance artifact type.
    /// </summary>
    [EnumType]
    public readonly struct ApplianceArtifactType : IEquatable<ApplianceArtifactType>
    {
        private readonly string _value;

        private ApplianceArtifactType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplianceArtifactType Template { get; } = new ApplianceArtifactType("Template");
        public static ApplianceArtifactType Custom { get; } = new ApplianceArtifactType("Custom");

        public static bool operator ==(ApplianceArtifactType left, ApplianceArtifactType right) => left.Equals(right);
        public static bool operator !=(ApplianceArtifactType left, ApplianceArtifactType right) => !left.Equals(right);

        public static explicit operator string(ApplianceArtifactType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplianceArtifactType other && Equals(other);
        public bool Equals(ApplianceArtifactType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The appliance lock level.
    /// </summary>
    [EnumType]
    public readonly struct ApplianceLockLevel : IEquatable<ApplianceLockLevel>
    {
        private readonly string _value;

        private ApplianceLockLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplianceLockLevel CanNotDelete { get; } = new ApplianceLockLevel("CanNotDelete");
        public static ApplianceLockLevel ReadOnly { get; } = new ApplianceLockLevel("ReadOnly");
        public static ApplianceLockLevel None { get; } = new ApplianceLockLevel("None");

        public static bool operator ==(ApplianceLockLevel left, ApplianceLockLevel right) => left.Equals(right);
        public static bool operator !=(ApplianceLockLevel left, ApplianceLockLevel right) => !left.Equals(right);

        public static explicit operator string(ApplianceLockLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplianceLockLevel other && Equals(other);
        public bool Equals(ApplianceLockLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
