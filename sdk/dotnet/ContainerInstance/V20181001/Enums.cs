// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ContainerInstance.V20181001
{
    /// <summary>
    /// Specifies if the IP is exposed to the public internet or private VNET.
    /// </summary>
    [EnumType]
    public readonly struct ContainerGroupIpAddressType : IEquatable<ContainerGroupIpAddressType>
    {
        private readonly string _value;

        private ContainerGroupIpAddressType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContainerGroupIpAddressType Public { get; } = new ContainerGroupIpAddressType("Public");
        public static ContainerGroupIpAddressType Private { get; } = new ContainerGroupIpAddressType("Private");

        public static bool operator ==(ContainerGroupIpAddressType left, ContainerGroupIpAddressType right) => left.Equals(right);
        public static bool operator !=(ContainerGroupIpAddressType left, ContainerGroupIpAddressType right) => !left.Equals(right);

        public static explicit operator string(ContainerGroupIpAddressType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerGroupIpAddressType other && Equals(other);
        public bool Equals(ContainerGroupIpAddressType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol associated with the port.
    /// </summary>
    [EnumType]
    public readonly struct ContainerGroupNetworkProtocol : IEquatable<ContainerGroupNetworkProtocol>
    {
        private readonly string _value;

        private ContainerGroupNetworkProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContainerGroupNetworkProtocol TCP { get; } = new ContainerGroupNetworkProtocol("TCP");
        public static ContainerGroupNetworkProtocol UDP { get; } = new ContainerGroupNetworkProtocol("UDP");

        public static bool operator ==(ContainerGroupNetworkProtocol left, ContainerGroupNetworkProtocol right) => left.Equals(right);
        public static bool operator !=(ContainerGroupNetworkProtocol left, ContainerGroupNetworkProtocol right) => !left.Equals(right);

        public static explicit operator string(ContainerGroupNetworkProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerGroupNetworkProtocol other && Equals(other);
        public bool Equals(ContainerGroupNetworkProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Restart policy for all containers within the container group. 
    /// - `Always` Always restart
    /// - `OnFailure` Restart on failure
    /// - `Never` Never restart
    /// </summary>
    [EnumType]
    public readonly struct ContainerGroupRestartPolicy : IEquatable<ContainerGroupRestartPolicy>
    {
        private readonly string _value;

        private ContainerGroupRestartPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContainerGroupRestartPolicy Always { get; } = new ContainerGroupRestartPolicy("Always");
        public static ContainerGroupRestartPolicy OnFailure { get; } = new ContainerGroupRestartPolicy("OnFailure");
        public static ContainerGroupRestartPolicy Never { get; } = new ContainerGroupRestartPolicy("Never");

        public static bool operator ==(ContainerGroupRestartPolicy left, ContainerGroupRestartPolicy right) => left.Equals(right);
        public static bool operator !=(ContainerGroupRestartPolicy left, ContainerGroupRestartPolicy right) => !left.Equals(right);

        public static explicit operator string(ContainerGroupRestartPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerGroupRestartPolicy other && Equals(other);
        public bool Equals(ContainerGroupRestartPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol associated with the port.
    /// </summary>
    [EnumType]
    public readonly struct ContainerNetworkProtocol : IEquatable<ContainerNetworkProtocol>
    {
        private readonly string _value;

        private ContainerNetworkProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContainerNetworkProtocol TCP { get; } = new ContainerNetworkProtocol("TCP");
        public static ContainerNetworkProtocol UDP { get; } = new ContainerNetworkProtocol("UDP");

        public static bool operator ==(ContainerNetworkProtocol left, ContainerNetworkProtocol right) => left.Equals(right);
        public static bool operator !=(ContainerNetworkProtocol left, ContainerNetworkProtocol right) => !left.Equals(right);

        public static explicit operator string(ContainerNetworkProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerNetworkProtocol other && Equals(other);
        public bool Equals(ContainerNetworkProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The SKU of the GPU resource.
    /// </summary>
    [EnumType]
    public readonly struct GpuSku : IEquatable<GpuSku>
    {
        private readonly string _value;

        private GpuSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GpuSku K80 { get; } = new GpuSku("K80");
        public static GpuSku P100 { get; } = new GpuSku("P100");
        public static GpuSku V100 { get; } = new GpuSku("V100");

        public static bool operator ==(GpuSku left, GpuSku right) => left.Equals(right);
        public static bool operator !=(GpuSku left, GpuSku right) => !left.Equals(right);

        public static explicit operator string(GpuSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GpuSku other && Equals(other);
        public bool Equals(GpuSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The log type to be used.
    /// </summary>
    [EnumType]
    public readonly struct LogAnalyticsLogType : IEquatable<LogAnalyticsLogType>
    {
        private readonly string _value;

        private LogAnalyticsLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LogAnalyticsLogType ContainerInsights { get; } = new LogAnalyticsLogType("ContainerInsights");
        public static LogAnalyticsLogType ContainerInstanceLogs { get; } = new LogAnalyticsLogType("ContainerInstanceLogs");

        public static bool operator ==(LogAnalyticsLogType left, LogAnalyticsLogType right) => left.Equals(right);
        public static bool operator !=(LogAnalyticsLogType left, LogAnalyticsLogType right) => !left.Equals(right);

        public static explicit operator string(LogAnalyticsLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LogAnalyticsLogType other && Equals(other);
        public bool Equals(LogAnalyticsLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The operating system type required by the containers in the container group.
    /// </summary>
    [EnumType]
    public readonly struct OperatingSystemTypes : IEquatable<OperatingSystemTypes>
    {
        private readonly string _value;

        private OperatingSystemTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OperatingSystemTypes Windows { get; } = new OperatingSystemTypes("Windows");
        public static OperatingSystemTypes Linux { get; } = new OperatingSystemTypes("Linux");

        public static bool operator ==(OperatingSystemTypes left, OperatingSystemTypes right) => left.Equals(right);
        public static bool operator !=(OperatingSystemTypes left, OperatingSystemTypes right) => !left.Equals(right);

        public static explicit operator string(OperatingSystemTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OperatingSystemTypes other && Equals(other);
        public bool Equals(OperatingSystemTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the container group.
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
        public static ResourceIdentityType UserAssigned { get; } = new ResourceIdentityType("UserAssigned");
        public static ResourceIdentityType SystemAssigned_UserAssigned { get; } = new ResourceIdentityType("SystemAssigned, UserAssigned");
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
