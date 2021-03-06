// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.NetApp.V20200201
{
    /// <summary>
    /// Indicates whether the local volume is the source or destination for the Volume Replication
    /// </summary>
    [EnumType]
    public readonly struct EndpointType : IEquatable<EndpointType>
    {
        private readonly string _value;

        private EndpointType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EndpointType Src { get; } = new EndpointType("src");
        public static EndpointType Dst { get; } = new EndpointType("dst");

        public static bool operator ==(EndpointType left, EndpointType right) => left.Equals(right);
        public static bool operator !=(EndpointType left, EndpointType right) => !left.Equals(right);

        public static explicit operator string(EndpointType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointType other && Equals(other);
        public bool Equals(EndpointType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The service level of the file system
    /// </summary>
    [EnumType]
    public readonly struct PoolServiceLevel : IEquatable<PoolServiceLevel>
    {
        private readonly string _value;

        private PoolServiceLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Standard service level
        /// </summary>
        public static PoolServiceLevel Standard { get; } = new PoolServiceLevel("Standard");
        /// <summary>
        /// Premium service level
        /// </summary>
        public static PoolServiceLevel Premium { get; } = new PoolServiceLevel("Premium");
        /// <summary>
        /// Ultra service level
        /// </summary>
        public static PoolServiceLevel Ultra { get; } = new PoolServiceLevel("Ultra");

        public static bool operator ==(PoolServiceLevel left, PoolServiceLevel right) => left.Equals(right);
        public static bool operator !=(PoolServiceLevel left, PoolServiceLevel right) => !left.Equals(right);

        public static explicit operator string(PoolServiceLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolServiceLevel other && Equals(other);
        public bool Equals(PoolServiceLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Schedule
    /// </summary>
    [EnumType]
    public readonly struct ReplicationSchedule : IEquatable<ReplicationSchedule>
    {
        private readonly string _value;

        private ReplicationSchedule(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReplicationSchedule ReplicationSchedule_10minutely { get; } = new ReplicationSchedule("_10minutely");
        public static ReplicationSchedule Hourly { get; } = new ReplicationSchedule("hourly");
        public static ReplicationSchedule Daily { get; } = new ReplicationSchedule("daily");

        public static bool operator ==(ReplicationSchedule left, ReplicationSchedule right) => left.Equals(right);
        public static bool operator !=(ReplicationSchedule left, ReplicationSchedule right) => !left.Equals(right);

        public static explicit operator string(ReplicationSchedule value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReplicationSchedule other && Equals(other);
        public bool Equals(ReplicationSchedule other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The service level of the file system
    /// </summary>
    [EnumType]
    public readonly struct VolumeServiceLevel : IEquatable<VolumeServiceLevel>
    {
        private readonly string _value;

        private VolumeServiceLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Standard service level
        /// </summary>
        public static VolumeServiceLevel Standard { get; } = new VolumeServiceLevel("Standard");
        /// <summary>
        /// Premium service level
        /// </summary>
        public static VolumeServiceLevel Premium { get; } = new VolumeServiceLevel("Premium");
        /// <summary>
        /// Ultra service level
        /// </summary>
        public static VolumeServiceLevel Ultra { get; } = new VolumeServiceLevel("Ultra");

        public static bool operator ==(VolumeServiceLevel left, VolumeServiceLevel right) => left.Equals(right);
        public static bool operator !=(VolumeServiceLevel left, VolumeServiceLevel right) => !left.Equals(right);

        public static explicit operator string(VolumeServiceLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeServiceLevel other && Equals(other);
        public bool Equals(VolumeServiceLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
