// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Devices.V20170821Preview
{
    /// <summary>
    /// Rights that this key has.
    /// </summary>
    [EnumType]
    public readonly struct AccessRightsDescription : IEquatable<AccessRightsDescription>
    {
        private readonly string _value;

        private AccessRightsDescription(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessRightsDescription ServiceConfig { get; } = new AccessRightsDescription("ServiceConfig");
        public static AccessRightsDescription EnrollmentRead { get; } = new AccessRightsDescription("EnrollmentRead");
        public static AccessRightsDescription EnrollmentWrite { get; } = new AccessRightsDescription("EnrollmentWrite");
        public static AccessRightsDescription DeviceConnect { get; } = new AccessRightsDescription("DeviceConnect");
        public static AccessRightsDescription RegistrationStatusRead { get; } = new AccessRightsDescription("RegistrationStatusRead");
        public static AccessRightsDescription RegistrationStatusWrite { get; } = new AccessRightsDescription("RegistrationStatusWrite");

        public static bool operator ==(AccessRightsDescription left, AccessRightsDescription right) => left.Equals(right);
        public static bool operator !=(AccessRightsDescription left, AccessRightsDescription right) => !left.Equals(right);

        public static explicit operator string(AccessRightsDescription value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessRightsDescription other && Equals(other);
        public bool Equals(AccessRightsDescription other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Allocation policy to be used by this provisioning service.
    /// </summary>
    [EnumType]
    public readonly struct AllocationPolicy : IEquatable<AllocationPolicy>
    {
        private readonly string _value;

        private AllocationPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AllocationPolicy Hashed { get; } = new AllocationPolicy("Hashed");
        public static AllocationPolicy GeoLatency { get; } = new AllocationPolicy("GeoLatency");
        public static AllocationPolicy Static { get; } = new AllocationPolicy("Static");

        public static bool operator ==(AllocationPolicy left, AllocationPolicy right) => left.Equals(right);
        public static bool operator !=(AllocationPolicy left, AllocationPolicy right) => !left.Equals(right);

        public static explicit operator string(AllocationPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AllocationPolicy other && Equals(other);
        public bool Equals(AllocationPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct IotDpsSku : IEquatable<IotDpsSku>
    {
        private readonly string _value;

        private IotDpsSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IotDpsSku S1 { get; } = new IotDpsSku("S1");

        public static bool operator ==(IotDpsSku left, IotDpsSku right) => left.Equals(right);
        public static bool operator !=(IotDpsSku left, IotDpsSku right) => !left.Equals(right);

        public static explicit operator string(IotDpsSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IotDpsSku other && Equals(other);
        public bool Equals(IotDpsSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Current state of the provisioning service.
    /// </summary>
    [EnumType]
    public readonly struct State : IEquatable<State>
    {
        private readonly string _value;

        private State(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static State Activating { get; } = new State("Activating");
        public static State Active { get; } = new State("Active");
        public static State Deleting { get; } = new State("Deleting");
        public static State Deleted { get; } = new State("Deleted");
        public static State ActivationFailed { get; } = new State("ActivationFailed");
        public static State DeletionFailed { get; } = new State("DeletionFailed");
        public static State Transitioning { get; } = new State("Transitioning");
        public static State Suspending { get; } = new State("Suspending");
        public static State Suspended { get; } = new State("Suspended");
        public static State Resuming { get; } = new State("Resuming");
        public static State FailingOver { get; } = new State("FailingOver");
        public static State FailoverFailed { get; } = new State("FailoverFailed");

        public static bool operator ==(State left, State right) => left.Equals(right);
        public static bool operator !=(State left, State right) => !left.Equals(right);

        public static explicit operator string(State value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is State other && Equals(other);
        public bool Equals(State other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
